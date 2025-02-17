// Copyright (c) 2024-2025 Alexej Disterhoft
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
//
// SPDX-License-Identifier: MIT

package parse_test

import (
	"fmt"
	"log"
	"os"

	. "github.com/nobbs/mapr-ticket-parser/pkg/parse"

	"k8s.io/utils/ptr"
)

// Set the timezone to Europe/Berlin for the examples to work as expected in CI.
func init() {
	os.Setenv("TZ", "Europe/Berlin")
}

// This example shows a complete roundtrip of a MapR ticket. It takes a base64 encoded ticket,
// unmarshals it, marshals it again and prints the re-encrypted ticket. The re-encrypted ticket
// will be different from the original ticket even though the ticket was not changed, because
// the encryption uses a random nonce.
func Example() {
	// input is a base64 encoded ticket
	blob := []byte(`demo.mapr.com cj1FDarNNKh7f+hL5ho1m32RzYyHPKuGIPJzE/CkUqEfcTGEP4YJuFlTsBmHuifI5
LvNob/Y4xmDsrz9OxrBnhly/0g9xAs5ApZWNY8Rcab8q70IBYIbpu7xsBBTAiVRyLJkAtGFXNn104BB0AsS55GbQFUN9NAiWL
zZY3/X1ITfGfDEGaYbWWTb1LGx6C0Jjgnr7TzXv1GqwiASbcUQCXOx4inguwMneYt9KhOp89smw6GBKP064DfIMHHR6lgv0Xh
BP6d9FVJ1QWKvcccvi2F3LReBtqA=`)

	// Parse the ticket
	t, err := Unmarshal(blob)
	if err != nil {
		log.Fatal(err)
	}

	// Marshal the ticket again, this will encrypt the ticket again with a random nonce, so even if
	// the ticket was encrypted before, it will be different now.
	b, err := Marshal(t)
	if err != nil {
		log.Fatal(err)
	}

	// Print the re-encrypted ticket
	fmt.Println(string(b))
}

// This example shows how to unmarshal a MapR ticket given a base64 encoded ticket.
func ExampleUnmarshal() {
	// input is a base64 encoded ticket
	blob := []byte(`demo.mapr.com cj1FDarNNKh7f+hL5ho1m32RzYyHPKuGIPJzE/CkUqEfcTGEP4YJuFlTsBmHuifI5
LvNob/Y4xmDsrz9OxrBnhly/0g9xAs5ApZWNY8Rcab8q70IBYIbpu7xsBBTAiVRyLJkAtGFXNn104BB0AsS55GbQFUN9NAiWL
zZY3/X1ITfGfDEGaYbWWTb1LGx6C0Jjgnr7TzXv1GqwiASbcUQCXOx4inguwMneYt9KhOp89smw6GBKP064DfIMHHR6lgv0Xh
BP6d9FVJ1QWKvcccvi2F3LReBtqA=`)

	// Parse the ticket
	t, err := Unmarshal(blob)
	if err != nil {
		log.Fatal(err)
	}

	// Print the ticket information as formatted JSON
	fmt.Println(t.PrettyString())
	// Output:
	// {
	//   "cluster": "demo.mapr.com",
	//   "ticket": {
	//     "expiryTime": "2019-02-19T13:13:49+01:00",
	//     "creationTimeSec": "2019-02-05T13:13:49+01:00",
	//     "maxRenewalDurationSec": "30d",
	//     "encryptedTicket": "AggBH+N6bCF5TEMUkaeHo7IHYUawyrf4ncz4JdGj7uupD5ll4vY2ddibb/4rBU1hob2aMhZIwdJakKgoBxJi4fWkyXVRCj/rihDWeBoInszw1Ni6ovwqO7Q3s8kvn7QIV+8yJlilCCNM1DZwVw==",
	//     "userKey": {
	//       "key": "KPNSEc96euKFcp5DLc0gbfuLrCDolaSFIQBJzW8YSYY="
	//     },
	//     "userCreds": {
	//       "uid": 5000,
	//       "gids": [
	//         5000,
	//         1000
	//       ],
	//       "userName": "mapr"
	//     },
	//     "canUserImpersonate": true
	//   }
	// }
}

// This example shows how to unmarshal a MapR ticket given a base64 encoded ticket and then
// remove sensitive information from the ticket, ie. the encrypted ticket and the user key.
func ExampleMaprTicket_Mask() {
	// input is a base64 encoded ticket
	blob := []byte(`demo.mapr.com cj1FDarNNKh7f+hL5ho1m32RzYyHPKuGIPJzE/CkUqEfcTGEP4YJuFlTsBmHuifI5
LvNob/Y4xmDsrz9OxrBnhly/0g9xAs5ApZWNY8Rcab8q70IBYIbpu7xsBBTAiVRyLJkAtGFXNn104BB0AsS55GbQFUN9NAiWL
zZY3/X1ITfGfDEGaYbWWTb1LGx6C0Jjgnr7TzXv1GqwiASbcUQCXOx4inguwMneYt9KhOp89smw6GBKP064DfIMHHR6lgv0Xh
BP6d9FVJ1QWKvcccvi2F3LReBtqA=`)

	// Parse the ticket
	t, err := Unmarshal(blob)
	if err != nil {
		log.Fatal(err)
	}

	// Mask the ticket, removing sensitive information
	t = t.Mask()

	// Print the ticket information as formatted JSON
	fmt.Println(t.PrettyString())
	// Output:
	// {
	//   "cluster": "demo.mapr.com",
	//   "ticket": {
	//     "expiryTime": "2019-02-19T13:13:49+01:00",
	//     "creationTimeSec": "2019-02-05T13:13:49+01:00",
	//     "maxRenewalDurationSec": "30d",
	//     "userCreds": {
	//       "uid": 5000,
	//       "gids": [
	//         5000,
	//         1000
	//       ],
	//       "userName": "mapr"
	//     },
	//     "canUserImpersonate": true
	//   }
	// }
}

func ExampleMaprTicket_String() {

	// create a new MaprTicket object with some values
	t := NewMaprTicket()
	t.Cluster = "demo.mapr.com"
	t.UserCreds.Uid = ptr.To[uint32](1000)
	t.UserCreds.UserName = ptr.To[string]("mapr")
	t.CreationTimeSec = ptr.To[uint64](1549385629) // 2019-02-05T17:53:49+01:00
	t.ExpiryTime = ptr.To[uint64](1549576429)      // 2019-02-07T22:53:49+01:00

	// Print the ticket information as formatted JSON
	fmt.Println(t.String())
	// Output:
	// {"cluster":"demo.mapr.com","ticket":{"userKey":{},"userCreds":{"uid":1000,"userName":"mapr","capabilities":{}},"expiryTime":1549576429,"creationTimeSec":1549385629}}
}

func ExampleMaprTicket_PrettyString() {
	// create a new MaprTicket object with some values
	t := NewMaprTicket()
	t.Cluster = "demo.mapr.com"
	t.UserCreds.Uid = ptr.To[uint32](1000)
	t.UserCreds.UserName = ptr.To[string]("mapr")
	t.CreationTimeSec = ptr.To[uint64](1549385629) // 2019-02-05T17:53:49+01:00
	t.ExpiryTime = ptr.To[uint64](1549576429)      // 2019-02-07T22:53:49+01:00

	// Print the ticket information as formatted JSON
	fmt.Println(t.PrettyString())
	// Output:
	// {
	//   "cluster": "demo.mapr.com",
	//   "ticket": {
	//     "expiryTime": "2019-02-07T22:53:49+01:00",
	//     "creationTimeSec": "2019-02-05T17:53:49+01:00",
	//     "userKey": {},
	//     "userCreds": {
	//       "uid": 1000,
	//       "userName": "mapr",
	//       "capabilities": {}
	//     }
	//   }
	// }
}
