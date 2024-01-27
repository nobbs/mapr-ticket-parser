// Copyright (c) 2024 Alexej Disterhoft
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
//
// SPX-License-Identifier: MIT

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
	t, err := parse.Unmarshal(blob)
	if err != nil {
		log.Fatal(err)
	}

	// Print the ticket information, hiding the sensitive fields
	fmt.Println(t.Mask())
}
