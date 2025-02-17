// Copyright (c) 2024-2025 Alexej Disterhoft
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
//
// SPDX-License-Identifier: MIT

// Package parse provides functions to parse MapR ticket strings to obtain the information contained
// in the ticket.
//
// A MapR ticket is a string containing two parts, separated by a space. The first part is the
// cluster the ticket is for, the second part is the ticket itself. The ticket is base64 encoded
// and encrypted using AES-256-GCM. This package provides functions to marshal and unmarshal
// MapR tickets from their string representation into a struct and back.
package parse

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"google.golang.org/protobuf/proto"

	"github.com/nobbs/mapr-ticket-parser/internal/aes"
	mapr "github.com/nobbs/mapr-ticket-parser/internal/ezmeral.hpe.com/datafab/fs/proto"
)

// ErrInvalidTicket is the error message returned when a ticket is invalid and cannot be parsed.
const ErrInvalidTicket = "invalid mapr ticket"

// MaprTicket is a struct representing a MapR ticket.
type MaprTicket struct {
	// the cluster the ticket is for
	Cluster string `json:"cluster"`
	// the ticket and key
	*mapr.TicketAndKey `json:"ticket"`
}

// NewMaprTicket returns a new empty MaprTicket object, initializing fields of internal types.
func NewMaprTicket() *MaprTicket {
	return &MaprTicket{
		TicketAndKey: &mapr.TicketAndKey{
			UserCreds: &mapr.CredentialsMsg{
				Capabilities: &mapr.Capabilities{},
			},
			UserKey: &mapr.Key{},
		},
	}
}

// Unmarshal takes a byte slice containing an encoded MapR ticket string representation of a ticket
// and returns a MaprTicket object.
func Unmarshal(in []byte) (*MaprTicket, error) {
	// split the input into the two parts (cluster and ticket blob)
	parts := strings.Split(string(in), " ")
	if len(parts) != 2 {
		return nil, fmt.Errorf("%s: %s", ErrInvalidTicket, "cannot split ticket into cluster and ticket")
	}

	cluster := parts[0]
	blob := parts[1]

	// base64 decode the ticket
	ticketEncrypted, err := base64.StdEncoding.DecodeString(blob)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ErrInvalidTicket, err)
	}

	// decrypt the ticket
	aes, err := aes.New()
	if err != nil {
		return nil, err
	}

	decryptedTicket, err := aes.Decrypt(ticketEncrypted)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ErrInvalidTicket, err)
	}

	// unmarshal the ticket
	ticketAndKey, err := unmarshal(decryptedTicket)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ErrInvalidTicket, err)
	}

	return &MaprTicket{
		Cluster:      cluster,
		TicketAndKey: ticketAndKey,
	}, nil
}

// Marshal takes a Ticket object and returns a byte slice containing an encoded MapR ticket string
// representation of the ticket as a byte slice.
func Marshal(in *MaprTicket) ([]byte, error) {
	// marshal the ticket
	ticketDecrypted, err := marshal(in.TicketAndKey)
	if err != nil {
		return nil, err
	}

	// encrypt the ticket
	aes, err := aes.New()
	if err != nil {
		return nil, err
	}

	ticketEncrypted, err := aes.Encrypt(ticketDecrypted)
	if err != nil {
		return nil, err
	}

	// base64 encode the ticket
	ticketEncoded := base64.StdEncoding.EncodeToString(ticketEncrypted)

	// join the cluster and the ticket
	return []byte(fmt.Sprintf("%s %s", in.Cluster, ticketEncoded)), nil
}

// String returns a string representation of the ticket.
func (t *MaprTicket) String() string {
	// encode to pretty-printed json
	jsonBytes, err := json.Marshal(t)
	if err != nil {
		return ""
	}

	return string(jsonBytes)
}

// PrettyString returns a pretty-printed string representation of the ticket, converting timestamps
// to RFC3339 and durations to human readable format.
func (t *MaprTicket) PrettyString() string {
	// encode to pretty-printed json
	jsonBytes, err := t.prettyJSON()
	if err != nil {
		return ""
	}

	return string(jsonBytes)
}

// Mask returns a new MaprTicket object with the properties UserKey and EncryptedTicket removed.
func (t *MaprTicket) Mask() *MaprTicket {
	ticketAndKey := proto.Clone(t.TicketAndKey).(*mapr.TicketAndKey)
	ticketAndKey.UserKey = nil
	ticketAndKey.EncryptedTicket = nil

	return &MaprTicket{
		Cluster:      t.Cluster,
		TicketAndKey: ticketAndKey,
	}
}

// Equal returns true if the two tickets are equal, false otherwise. This comparison is done by
// using the proto.Equal function provided by the protobuf library.
func Equal(x, y *MaprTicket) bool {
	if x.Cluster != y.Cluster {
		return false
	}

	return proto.Equal(x.TicketAndKey, y.TicketAndKey)
}

// unmarshal takes a byte slice containing a decrypted ticket and returns a TicketAndKey object.
func unmarshal(ticket []byte) (*mapr.TicketAndKey, error) {
	ticketAndKey := &mapr.TicketAndKey{}
	if err := proto.Unmarshal(ticket, ticketAndKey); err != nil {
		return nil, err
	}

	return ticketAndKey, nil
}

// marshal takes a Ticket object and returns a byte slice containing a decrypted ticket.
func marshal(t *mapr.TicketAndKey) ([]byte, error) {
	return proto.Marshal(t)
}
