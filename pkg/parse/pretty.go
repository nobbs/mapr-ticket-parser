package parse

import (
	"encoding/json"
	"time"

	mapr "github.com/nobbs/mapr-ticket-parser/internal/ezmeral.hpe.com/datafab/fs/proto"

	"k8s.io/apimachinery/pkg/util/duration"
)

// humanMaprTicket is a struct representing a MapR ticket, with custom JSON marshalling to make the
// output more human readable.
type humanMaprTicket struct {
	Cluster            string `json:"cluster"`
	*humanTicketAndKey `json:"ticket"`
}

type humanTicketAndKey mapr.TicketAndKey

// Custom JSON marshalling to make the output more human readable.
func (mt *MaprTicket) MarshalJSON() ([]byte, error) {
	return json.Marshal(&humanMaprTicket{
		Cluster:           mt.Cluster,
		humanTicketAndKey: (*humanTicketAndKey)(mt.TicketAndKey),
	})
}

// Custom JSON marshalling to make the output more human readable.
func (ht *humanTicketAndKey) MarshalJSON() ([]byte, error) {
	type Alias humanTicketAndKey
	return json.Marshal(&struct {
		ExpiryTime            string `json:"expiryTime,omitempty"`
		CreationTimeSec       string `json:"creationTimeSec,omitempty"`
		MaxRenewalDurationSec string `json:"maxRenewalDurationSec,omitempty"`
		LastRenewalTime       string `json:"lastRenewalTime,omitempty"`
		*Alias
	}{
		ExpiryTime:            unixTimeToRFC3339(ht.ExpiryTime),
		CreationTimeSec:       unixTimeToRFC3339(ht.CreationTimeSec),
		LastRenewalTime:       unixTimeToRFC3339(ht.LastRenewalTime),
		MaxRenewalDurationSec: secondsToHumanDuration(ht.MaxRenewalDurationSec),
		Alias:                 (*Alias)(ht),
	})
}

// unixTimeToRFC3339 takes a pointer to a uint64 and returns a string containing the RFC3339
func unixTimeToRFC3339(unixTime *uint64) string {
	if unixTime == nil {
		return ""
	}

	return time.Unix(int64(*unixTime), 0).Format(time.RFC3339)
}

// secondsToHumanDuration takes a pointer to a uint64 and returns a string containing the duration
// in human readable format.
func secondsToHumanDuration(seconds *uint64) string {
	if seconds == nil {
		return ""
	}

	return duration.HumanDuration(time.Duration(*seconds) * time.Second)
}
