# mapr-ticket-parser

This module provides a simple way to parse MapR tickets and extract information from their content without having to use the `maprlogin` command line tool. A MapR ticket is an authentication token that is used to authenticate a user to a MapR cluster.

The module is based on [this blog post of mine](https://nobbs.dev/posts/reverse-engineering-mapr-ticket-format/) which describes the process of reverse engineering the MapR ticket format. A summary including the code snippets used during the reverse engineering process can be found in the [hack](./hack) directory.

## Disclaimer

_The only "proprietary" code found in this repository is the Protobuf definition of the ticket format which was reconstructed as described in the blog post. No code was used or in any way copied from the MapR source code._

## Installation

To install the module, use the following command:

<!-- x-release-please-start-version -->

```bash
go get github.com/nobbs/mapr-ticket-parser@v0.1.2
```

<!-- x-release-please-end -->

## Usage

The [following example](./examples/main.go) shows how to read a MapR ticket from a file called `ticket` and print the ticket information to the console in a human-readable format (JSON) - `Mask()` is used to hide the sensitive fields such as the encrypted ticket and user key.

```go
package main

import (
  "fmt"
  "log"
  "os"

  "github.com/nobbs/mapr-ticket-parser/pkg/parse"
)

func main() {
  // Read the blob from a file
  blob, err := os.ReadFile("ticket")
  if err != nil {
    log.Fatal(err)
  }

  // Parse the ticket
  ticket, err := parse.Unmarshal(blob)
  if err != nil {
    log.Fatal(err)
  }

  // Print the ticket information, hiding the sensitive fields
  fmt.Println(ticket.Mask().PrettyString())
}
```

The above example will print the following output when used on the following example ticket:

```console
$ cat ticket
demo.mapr.com +Cze+qwYCbAXGbz56OO7UF+lGqL3WPXrNkO1SLawEEDmSbgNl019xBeBY3kvh+R13iz/mCnwpzsLQw4Y5jEnv5GtuIWbeoC95ha8VKwX8MKcE6Kn9nZ2AF0QminkHwNVBx6TDriGZffyJCfZzivBwBSdKoQEWhBOPFCIMAi7w2zV/SX5Ut7u4qIKvEpr0JHV7sLMWYLhYncM6CKMd7iECGvECsBvEZRVj+dpbEY0BaRN/W54/7wNWaSVELUF6JWHQ8dmsqty4cZlI0/MV10HZzIbl9sMLFQ=

# put the above example code in a file called example.go and run it
$ go run ./example.go
{
  "cluster": "demo.mapr.com",
  "ticket": {
    "userCreds": {
      "uid": 5000,
      "gids": [
        5000,
        0,
        5001
      ],
      "userName": "mapr"
    },
    "expiryTime": 922337203685477,
    "creationTimeSec": 1522852297,
    "maxRenewalDurationSec": 0
  }
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
