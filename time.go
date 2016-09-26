package dnsdb

// Imports
import (
	"strconv"
	"time"
)

// Timestamp represents a time generated from a JSON string
type Timestamp struct {
	time.Time
}

// Make a new Timestamp from a unix timestamp
func NewTimestamp(secs int64) *Timestamp {
	return &Timestamp{
		Time: time.Unix(secs, 0),
	}
}

// UnmarshalJSON helps unmarshal UNIX dates in JSON
func (t *Timestamp) UnmarshalJSON(data []byte) error {
	var err error
	i, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	(*t).Time = time.Unix(i, 0)
	return nil
}
