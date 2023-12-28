package parse

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/nobbs/mapr-ticket-parser/internal/aes"
	mapr "github.com/nobbs/mapr-ticket-parser/internal/ezmeral.hpe.com/datafab/fs/proto"
	"google.golang.org/protobuf/proto"
)

const errInvalidTicket = "invalid mapr ticket"

type MaprTicket struct {
	Host               string `json:"host"`
	*mapr.TicketAndKey `json:"ticketAndKey"`
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

// Unmarshal takes a byte slice containing an encoded MapR ticket string representation of a ticket
// and returns a MaprTicket object.
func Unmarshal(in []byte) (*MaprTicket, error) {
	// split the input into the two parts (host and ticket blob)
	parts := strings.Split(string(in), " ")
	if len(parts) != 2 {
		return nil, fmt.Errorf("%s: %s", errInvalidTicket, "cannot split ticket into host and ticket")
	}

	host := parts[0]
	blob := parts[1]

	// base64 decode the ticket
	ticketEncrypted, err := base64.StdEncoding.DecodeString(blob)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errInvalidTicket, err)
	}

	// decrypt the ticket
	decryptedTicket, err := aes.Decrypt(ticketEncrypted)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errInvalidTicket, err)
	}

	// unmarshal the ticket
	ticketAndKey, err := unmarshal(decryptedTicket)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errInvalidTicket, err)
	}

	return &MaprTicket{
		Host:         host,
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
	ticketEncrypted, err := aes.Encrypt(ticketDecrypted)
	if err != nil {
		return nil, err
	}

	// base64 encode the ticket
	ticketEncoded := base64.StdEncoding.EncodeToString(ticketEncrypted)

	// join the host and the ticket
	return []byte(fmt.Sprintf("%s %s", in.Host, ticketEncoded)), nil
}

// WriteJSON writes the ticket to the writer as a pretty-printed json object.
func (t *MaprTicket) WriteJSON(w io.Writer) error {
	// json encoding
	jsonBytes, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err
	}

	// write the json to the writer
	if _, err = w.Write(jsonBytes); err != nil {
		return err
	}

	// print a newline at the end
	if _, err = w.Write([]byte("\n")); err != nil {
		return err
	}

	return nil
}

// Redact returns a new MaprTicket object with the properties UserKey and EncryptedTicket removed.
func (t *MaprTicket) Redact() *MaprTicket {
	ticketAndKey := proto.Clone(t.TicketAndKey).(*mapr.TicketAndKey)
	ticketAndKey.UserKey = nil
	ticketAndKey.EncryptedTicket = nil

	return &MaprTicket{
		Host:         t.Host,
		TicketAndKey: ticketAndKey,
	}
}
