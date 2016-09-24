package dnsdb

// Imports
import (
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

// LookupOptions defines the optional parameters to all lookup calls
type LookupOptions struct {
	Limit           int64     `url:"limit,omitempty"`
	TimeFirstBefore time.Time `url:"time_first_before,omitempty"`
	TimeFirstAfter  time.Time `url:"time_first_after,omitempty"`
	TimeLastBefore  time.Time `url:"time_last_before,omitempty"`
	TimeLastAfter   time.Time `url:"time_last_after,omitempty"`
}

// NewLookupRequest is a convienience function that extends NewRequest for Lookup methods
func (c *Client) NewLookupRequest(method, urlStr string, opt LookupOptions) (*http.Request, error) {
	qs, err := query.Values(opt)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest(method, urlStr+"?"+qs.Encode(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")

	return req, nil
}
