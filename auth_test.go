package dnsdb

// Imports
import (
	"github.com/stretchr/testify/assert"

	"net/http"
	"testing"
)

// Make a test struct for testing
type APIKeyTransportTest struct {
	Request *http.Request
}

func (t *APIKeyTransportTest) RoundTrip(req *http.Request) (*http.Response, error) {
	t.Request = req
	return nil, nil
}

func Test_APIKeyTransport(t *testing.T) {
	auth := APIKeyTransport{
		APIKey: "d41d8cd98f00b204e9800998ecf8427e",
	}
	// Verify it returns a correct client
	client := auth.Client()
	assert.NotNil(t, client)
	assert.Equal(t, client.Transport, &auth)
	// Send a request with the default client
	_, err := auth.RoundTrip(&http.Request{})
	assert.NotNil(t, err)
	// Setup a test transport and send a request
	testTransport := &APIKeyTransportTest{}
	auth.Transport = testTransport
	auth.RoundTrip(&http.Request{
		Header: http.Header{
			"X-Test": []string{"Header"},
		},
	})
	assert.NotNil(t, testTransport.Request)
	assert.Equal(t, testTransport.Request.Header.Get("X-API-Key"), "d41d8cd98f00b204e9800998ecf8427e")
}
