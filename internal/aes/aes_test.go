package aes_test

import (
	"testing"

	ticket "github.com/nobbs/mapr-ticket-parser/pkg/parse"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecodeAndReencode(t *testing.T) {
	originalTicketStr := []byte("demo.mapr.com WWup1XJ7wLYtrXA7vydwiorHLq2iqMvdLYa1qOCDqWE+qDtbAm0zYT1LZAqJ6NXNbgXXS+hP9vbcYTv69Fe+F3hJtBSPs6YHnd/b0T9hX7AyejyIJ/4RJ4ixjVpcl2Y9OMYKR28KjS7rV5sczOVLBf4kIVHa9PG0mRUopCOOK/HPsuaRSJRuvv1U5K9xUR+nND0cq6eN02PGVCsjERecNxtPgeHIdts3wDgld/JE2jWcWptNGwozmwqSAUym7VPNZNtjm67s+wXPSrjxvH69WcJmmpjJHpA=")

	// decode the ticket
	originalTicket, err := ticket.Unmarshal(originalTicketStr)
	require.NoError(t, err)

	// re-encode the ticket
	rebuildTicketStr, err := ticket.Marshal(originalTicket)
	require.NoError(t, err)

	// we can't simply compare the two tickets because the encrypted ticket
	// uses a different IV each time it is encrypted - so we instead decode
	// the re-encoded ticket and compare the fields of the ticketAndKey

	// assert that the re-encoded ticket can be decoded
	rebuildTicket, err := ticket.Unmarshal(rebuildTicketStr)
	require.NoError(t, err)

	// assert that the decoded ticket is the same as the original
	assert.Equal(t, originalTicket.Host, rebuildTicket.Host)
	assert.Equal(t, originalTicket.TicketAndKey.UserCreds.Gids, rebuildTicket.TicketAndKey.UserCreds.Gids)
	assert.Equal(t, originalTicket.TicketAndKey.UserCreds.Uid, rebuildTicket.TicketAndKey.UserCreds.Uid)
	assert.Equal(t, originalTicket.TicketAndKey.UserCreds.UserName, rebuildTicket.TicketAndKey.UserCreds.UserName)
	assert.Equal(t, originalTicket.TicketAndKey.ExpiryTime, rebuildTicket.TicketAndKey.ExpiryTime)
	assert.Equal(t, originalTicket.TicketAndKey.CreationTimeSec, rebuildTicket.TicketAndKey.CreationTimeSec)
	assert.Equal(t, originalTicket.TicketAndKey.MaxRenewalDurationSec, rebuildTicket.TicketAndKey.MaxRenewalDurationSec)
	assert.Equal(t, originalTicket.TicketAndKey.CanUserImpersonate, rebuildTicket.TicketAndKey.CanUserImpersonate)
	assert.Equal(t, originalTicket.TicketAndKey.IsExternal, rebuildTicket.TicketAndKey.IsExternal)
}
