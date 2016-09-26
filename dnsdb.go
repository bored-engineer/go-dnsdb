package dnsdb

// Imports
import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

const (
	baseURL   = "https://api.dnsdb.info/"
	userAgent = "go-dnsdb"
)

// A Client manages communication with the DNSDB API.
type Client struct {
	client *http.Client // HTTP client used to communicate with the API.

	// Base URL for API requests. Defaults to the public DNSDB API. BaseURL should always be specified with a trailing slash.
	BaseURL *url.URL

	// User agent used when communicating with the DNSDB API.
	UserAgent string

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Service are used for communication with the different parts of the DNSDB API.
	RRSet *RRSetService
	RData *RDataService
}

type service struct {
	client *Client
}

// NewClient returns a new DNSDB API client. It will fallback to http.DefaultClient if no client is provided
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(baseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent}
	c.common.client = c
	c.RRSet = (*RRSetService)(&c.common)
	c.RData = (*RDataService)(&c.common)

	return c
}

// NewRequest creates an http.Request object or returns and error. A relative URL can be provided in urlStr
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, c.BaseURL.ResolveReference(rel).String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", c.UserAgent)

	return req, nil
}

// Rate represents the current rate limit
type Rate struct {
	Limit     int       `json: "reset"`
	Remaining int       `json: "limit"`
	Reset     Timestamp `json: "remaining"`
}

// Extracts a rate from the response
func extractRateFromResponse(resp *http.Response) (r Rate) {
	limit := resp.Header.Get("X-RateLimit-Limit")
	if limit == "unlimited" {
		r.Limit = -1
	} else {
		(r.Limit), _ = strconv.Atoi(limit)
	}
	remaining := resp.Header.Get("X-RateLimit-Remaining")
	if remaining == "n/a" {
		r.Remaining = -1
	} else {
		(r.Remaining), _ = strconv.Atoi(remaining)
	}
	r.Reset.UnmarshalJSON([]byte(resp.Header.Get("X-RateLimit-Reset")))
	return
}

// Response is a DNSDB API response. We wrap the standard http.Response and provided the Rate object from the response
type Response struct {
	*http.Response

	Rate
}

// Do sends the provided http.Request and returns the response from DNSDB.
func (c *Client) Do(req *http.Request) (*Response, error) {
	// Actually do the request
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	// Make a response instance using the response
	response := &Response{
		Response: resp,
		Rate:     extractRateFromResponse(resp),
	}

	// If API returned an error, return the response and err back to user to inspect
	if resp.StatusCode != 200 {
		return response, fmt.Errorf("A unexpected status code (%d) was returned", resp.StatusCode)
	}

	// Return success
	return response, nil
}

// String allocates a new string value to store and returns a pointer to it.
func String(string string) *string {
	return &string
}

// UInt64 allocates a new uint64 value to store and returns a pointer to it.
func Uint64(uint64 uint64) *uint64 {
	return &uint64
}
