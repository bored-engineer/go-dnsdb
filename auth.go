package dnsdb

// Imports
import (
	"net/http"
)

// APIKeyTransport is a http.RoundTripper that adds the APIKey to each request.
type APIKeyTransport struct {
	APIKey string

	// Transport to use when making requests, defaults to http.DefaultTransport if nil.
	Transport http.RoundTripper
}

// RoundTrip implements the RoundTripper interface.
func (t *APIKeyTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Clone the request (don't leak the key)
	reqClone := new(http.Request)
	*reqClone = *req
	reqClone.Header = make(http.Header, len(req.Header))
	for idx, header := range req.Header {
		reqClone.Header[idx] = append([]string(nil), header...)
	}
	reqClone.Header.Add("X-API-Key", t.APIKey)
	if t.Transport == nil {
		return http.DefaultTransport.RoundTrip(reqClone)
	}
	return t.Transport.RoundTrip(reqClone)
}

// Client returns a *http.Client that makes APIKey authenticated requests
func (t *APIKeyTransport) Client() *http.Client {
	return &http.Client{
		Transport: t,
	}
}
