// Copyright (c) 2024-2025 Alexej Disterhoft
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
//
// SPDX-License-Identifier: MIT

package parse

import (
	"encoding/json"
	"time"

	mapr "github.com/nobbs/mapr-ticket-parser/internal/ezmeral.hpe.com/datafab/fs/proto"

	"k8s.io/apimachinery/pkg/util/duration"
)

// Custom JSON marshalling to make the output more human readable.
type humanTicketAndKey mapr.TicketAndKey

// prettyJSON returns a byte slice containing the JSON representation of the ticket. The JSON
// representation is human readable, ie. timestamps are converted from their unix epoch format to
// RFC3339 and durations are converted to human readable format.
func (t *MaprTicket) prettyJSON() ([]byte, error) {
	return json.MarshalIndent(&struct {
		Cluster      string             `json:"cluster"`
		TicketAndKey *humanTicketAndKey `json:"ticket"`
	}{
		Cluster:      t.Cluster,
		TicketAndKey: (*humanTicketAndKey)(t.TicketAndKey),
	}, "", "  ")
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
