# Reverse-Engineering the MapR Ticket format

A detailed write-up of the reverse-engineering process can be found in [this blog post](https://nobbs.dev/posts/reverse-engineering-mapr-ticket-format). This README is only a short usage guide to reproduce the results.

**Disclaimer:** This repository is for educational purposes only and does not contain any proprietary MapR code. The reconstructed protobuf schema is truncated and does not contain all fields of the original schema.

## Usage

The `hack` directory contains a set of scripts to reverse-engineer the MapR Ticket format. The scripts are mostly self-contained.

### Sample Tickets

The following sample tickets can be found on the internet and were used to test the implementation:

```
demo.mapr.com +Cze+qwYCbAXGbz56OO7UF+lGqL3WPXrNkO1SLawEEDmSbgNl019xBeBY3kvh+R13iz/mCnwpzsLQw4Y5jEnv5GtuIWbeoC95ha8VKwX8MKcE6Kn9nZ2AF0QminkHwNVBx6TDriGZffyJCfZzivBwBSdKoQEWhBOPFCIMAi7w2zV/SX5Ut7u4qIKvEpr0JHV7sLMWYLhYncM6CKMd7iECGvECsBvEZRVj+dpbEY0BaRN/W54/7wNWaSVELUF6JWHQ8dmsqty4cZlI0/MV10HZzIbl9sMLFQ=
demo.mapr.com cj1FDarNNKh7f+hL5ho1m32RzYyHPKuGIPJzE/CkUqEfcTGEP4YJuFlTsBmHuifI5LvNob/Y4xmDsrz9OxrBnhly/0g9xAs5ApZWNY8Rcab8q70IBYIbpu7xsBBTAiVRyLJkAtGFXNn104BB0AsS55GbQFUN9NAiWLzZY3/X1ITfGfDEGaYbWWTb1LGx6C0Jjgnr7TzXv1GqwiASbcUQCXOx4inguwMneYt9KhOp89smw6GBKP064DfIMHHR6lgv0XhBP6d9FVJ1QWKvcccvi2F3LReBtqA=
demo.mapr.com IGem6fUksZ1pd4iut978SKElS4ktecRsAkrl+qwPYc7xhfMg4wkwALKDmFmpc8Xvrm1L9Et0jVBoyhCWMDCjhToZ8b6FsfCn8wdCOB0MWm9CRobGv7MDsoEO2TQ5Bnh8i/VfuthKFxd3Om9iZPVCI4I1S9h4p/77Al1GzTGcfFFf1g9fq1HXftT9TEDyLdABIyATJbzv8zD10IDT8P1f8nxl7lgT/7ZhGz7N24vSz6jBxHE7oHmvHzjW22xJwt7TJgvrP21boH9HTsTPiKZOpQMZ4zFo6JA4aNVlQQ0=
```

### Prerequisites

The following tools are required to run the scripts in full:

- Python 3.8+
- [Poetry](https://python-poetry.org/), a Python dependency manager
- [Protobuf](https://protobuf.dev/), a protocol buffer library

### Installing Dependencies

To install the Python dependencies, run the following command:

```bash
poetry install
```

### Reconstructing the Protobuf Schema

First, we need to compile the `fds2proto.cpp` file to a binary. The following command assumes that you have installed the [Homebrew](https://brew.sh/) package manager on macOS and that you have installed the `protobuf` package via Homebrew.

```bash
export CPATH=/opt/homebrew/include
export LIBRARY_PATH=/opt/homebrew/lib
export LD_LIBRARY_PATH=/opt/homebrew/lib:$LD_LIBRARY_PATH
g++ -o fds2proto fds2proto.cpp -std=c++17 -lprotobuf -pthread
```

To reconstruct the protobuf schema, run the following command:

```bash
poetry run python rebuild_textproto.py < Security.java | ./fds2proto | tee security.proto
```

### Decoding a Ticket

To decode a ticket, we need to compile the protobuf schema to Python code. The following command assumes that you have installed the `protoc` compiler via Homebrew - which is the case if you have installed the `protobuf` package.

```bash
protoc --proto_path=./ --pyi_out=./ --python_out=./ security.proto
```

To decode a ticket, run the following command:

```bash
$ poetry run python parse.py <<<"demo.mapr.com +Cze+qwYCbAXGbz56OO7UF+lGqL3WPXrNkO1SLawEEDmSbgNl019xBeBY3kvh+R13iz/mCnwpzsLQw4Y5jEnv5GtuIWbeoC95ha8VKwX8MKcE6Kn9nZ2AF0QminkHwNVBx6TDriGZffyJCfZzivBwBSdKoQEWhBOPFCIMAi7w2zV/SX5Ut7u4qIKvEpr0JHV7sLMWYLhYncM6CKMd7iECGvECsBvEZRVj+dpbEY0BaRN/W54/7wNWaSVELUF6JWHQ8dmsqty4cZlI0/MV10HZzIbl9sMLFQ="

encryptedTicket: "\002\010\001zwP\014rA\244\0374e\232f\245\327\277\33061}\270\004<\340c\005\224\340,\331\350\363\216)\320\216X\344\346{>\365\035m\304\265\037\013\306B3\226\236\305\367\240\204\267\027\246k\030\267\215[\022\260AoU\265\001\273\021\303\231\024[G\010S#ru\224\335\024\206\243\247Q\237\301\'u\363\2331\220[\\\004\030&K9C"
userKey {
  key: "\267\211wq!N\016\370\247!&\024\257\240\267\367\211\315\367\r\340n\017\230\262\225x|\213\276\276D"
}
userCreds {
  uid: 5000
  gids: 5000
  gids: 0
  gids: 5001
  userName: "mapr"
}
expiryTime: 922337203685477
creationTimeSec: 1522852297
maxRenewalDurationSec: 0
```
