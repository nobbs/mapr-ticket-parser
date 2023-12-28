package parse_test

import (
	"testing"

	mapr "github.com/nobbs/mapr-ticket-parser/internal/ezmeral.hpe.com/datafab/fs/proto"
	"github.com/nobbs/mapr-ticket-parser/pkg/parse"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/utils/ptr"
)

func TestFromBytes(t *testing.T) {
	tests := []struct {
		name    string
		args    []byte
		want    *parse.MaprTicket
		wantErr bool
	}{
		{
			name: "valid ticket",
			args: []byte("demo.mapr.com +Cze+qwYCbAXGbz56OO7UF+lGqL3WPXrNkO1SLawEEDmSbgNl019xBeBY3kvh+R13iz/mCnwpzsLQw4Y5jEnv5GtuIWbeoC95ha8VKwX8MKcE6Kn9nZ2AF0QminkHwNVBx6TDriGZffyJCfZzivBwBSdKoQEWhBOPFCIMAi7w2zV/SX5Ut7u4qIKvEpr0JHV7sLMWYLhYncM6CKMd7iECGvECsBvEZRVj+dpbEY0BaRN/W54/7wNWaSVELUF6JWHQ8dmsqty4cZlI0/MV10HZzIbl9sMLFQ="),
			want: &parse.MaprTicket{
				Host: "demo.mapr.com",
				TicketAndKey: &mapr.TicketAndKey{
					UserCreds: &mapr.CredentialsMsg{
						UserName: ptr.To[string]("mapr"),
						Uid:      ptr.To[uint32](5000),
						Gids:     []uint32{5000, 0, 5001},
					},
					MaxRenewalDurationSec: ptr.To[uint64](0),
					ExpiryTime:            ptr.To[uint64](922337203685477),
					CreationTimeSec:       ptr.To[uint64](1522852297),
				}},
			wantErr: false,
		},
		{
			name: "another valid ticket",
			args: []byte("demo.mapr.com cj1FDarNNKh7f+hL5ho1m32RzYyHPKuGIPJzE/CkUqEfcTGEP4YJuFlTsBmHuifI5LvNob/Y4xmDsrz9OxrBnhly/0g9xAs5ApZWNY8Rcab8q70IBYIbpu7xsBBTAiVRyLJkAtGFXNn104BB0AsS55GbQFUN9NAiWLzZY3/X1ITfGfDEGaYbWWTb1LGx6C0Jjgnr7TzXv1GqwiASbcUQCXOx4inguwMneYt9KhOp89smw6GBKP064DfIMHHR6lgv0XhBP6d9FVJ1QWKvcccvi2F3LReBtqA="),
			want: &parse.MaprTicket{
				Host: "demo.mapr.com",
				TicketAndKey: &mapr.TicketAndKey{
					UserCreds: &mapr.CredentialsMsg{
						UserName: ptr.To[string]("mapr"),
						Uid:      ptr.To[uint32](5000),
						Gids:     []uint32{5000, 1000},
					},
					MaxRenewalDurationSec: ptr.To[uint64](2592000),
					ExpiryTime:            ptr.To[uint64](1550578429),
					CreationTimeSec:       ptr.To[uint64](1549368829),
					CanUserImpersonate:    ptr.To[bool](true),
				}},
			wantErr: false,
		},
		{
			name: "and another valid ticket",
			args: []byte("demo.mapr.com IGem6fUksZ1pd4iut978SKElS4ktecRsAkrl+qwPYc7xhfMg4wkwALKDmFmpc8Xvrm1L9Et0jVBoyhCWMDCjhToZ8b6FsfCn8wdCOB0MWm9CRobGv7MDsoEO2TQ5Bnh8i/VfuthKFxd3Om9iZPVCI4I1S9h4p/77Al1GzTGcfFFf1g9fq1HXftT9TEDyLdABIyATJbzv8zD10IDT8P1f8nxl7lgT/7ZhGz7N24vSz6jBxHE7oHmvHzjW22xJwt7TJgvrP21boH9HTsTPiKZOpQMZ4zFo6JA4aNVlQQ0="),
			want: &parse.MaprTicket{
				Host: "demo.mapr.com",
				TicketAndKey: &mapr.TicketAndKey{
					UserCreds: &mapr.CredentialsMsg{
						UserName: ptr.To[string]("mapr"),
						Uid:      ptr.To[uint32](5000),
						Gids:     []uint32{5000, 5003, 0},
					},
					MaxRenewalDurationSec: ptr.To[uint64](2592000),
					ExpiryTime:            ptr.To[uint64](1619735566),
					CreationTimeSec:       ptr.To[uint64](1618525966),
					CanUserImpersonate:    ptr.To[bool](true),
					IsExternal:            ptr.To[bool](true),
				}},
			wantErr: false,
		},
		{
			name:    "invalid ticket",
			args:    []byte("demo.mapr.com IGem6fUksZ1pd4iut978SKElS4ktecRsAkrl+qwPYc7xhfMg4wkwALKDmFmpc8Xvrm1L9Et0jVBoyhCWMDCjhToZ8b6FsfCn8wdCOB0MWm9CRobGv7MDsoEO2TQ5Bnh8i/VfuthKFxd3Om9iZPVCI4I1S9h4p/77Al1GzTGcfFFf1g9fq1HXftT9TEDyLdABIyATJbzv8zD10IDT8P1f8nxl7lgT/7ZhGz7N24vSz6jBxHE7oHmvHzjW22xJwt7TJgvrP21boH9HTsTPiKZOpQMZ4zFo6JA4aNVl"),
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := parse.Unmarshal(tt.args)

			// assert error
			if tt.wantErr {
				require.Error(t, gotErr)
				return
			}

			// manually assert fields of the ticketAndKey
			assert.Equal(t, tt.want.Host, got.Host)
			assert.Equal(t, tt.want.TicketAndKey.UserCreds, got.TicketAndKey.UserCreds)
			assert.Equal(t, tt.want.TicketAndKey.ExpiryTime, got.TicketAndKey.ExpiryTime)
			assert.Equal(t, tt.want.TicketAndKey.CreationTimeSec, got.TicketAndKey.CreationTimeSec)
			assert.Equal(t, tt.want.TicketAndKey.MaxRenewalDurationSec, got.TicketAndKey.MaxRenewalDurationSec)
			assert.Equal(t, tt.want.TicketAndKey.CanUserImpersonate, got.TicketAndKey.CanUserImpersonate)
			assert.Equal(t, tt.want.TicketAndKey.IsExternal, got.TicketAndKey.IsExternal)
		})
	}
}
