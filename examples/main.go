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
