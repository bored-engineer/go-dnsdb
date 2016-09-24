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
