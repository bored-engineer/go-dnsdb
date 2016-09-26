package dnsdb

import (
	"github.com/stretchr/testify/assert"

	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func Test_RDataService_LookupIP(t *testing.T) {
	// Setup a client
	c := NewClient(nil)

	// Verify that an error response fails
	errorServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Oh No", 500)
	}))
	defer errorServer.Close()
	u, err := url.Parse(errorServer.URL)
	assert.Nil(t, err)
	c.BaseURL = u
	_, _, err = c.RData.LookupIP(net.ParseIP("104.244.13.104"), nil)
	assert.NotNil(t, err)

	// Verify that it gets and parses a response correctly
	reportServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `{"count":24,"time_first":1433550785,"time_last":1468312116,"rrname":"www.farsighsecurity.com.","rrtype":"A","rdata":"104.244.13.104"}
{"count":9429,"time_first":1427897872,"time_last":1468333042,"rrname":"farsightsecurity.com.","rrtype":"A","rdata":"104.244.13.104"}`)
	}))
	defer reportServer.Close()
	u, err = url.Parse(reportServer.URL)
	assert.Nil(t, err)
	c.BaseURL = u
	actual, _, err := c.RData.LookupIP(net.ParseIP("104.244.13.104"), nil)
	assert.Nil(t, err)
	assert.Equal(t, []RData{
		RData{
			Count:     Uint64(24),
			TimeFirst: NewTimestamp(1433550785),
			TimeLast:  NewTimestamp(1468312116),
			RRName:    String("www.farsighsecurity.com."),
			RRType:    String("A"),
			RData:     String("104.244.13.104"),
		},
		RData{
			Count:     Uint64(9429),
			TimeFirst: NewTimestamp(1427897872),
			TimeLast:  NewTimestamp(1468333042),
			RRName:    String("farsightsecurity.com."),
			RRType:    String("A"),
			RData:     String("104.244.13.104"),
		},
	}, actual)
}
