package dnsdb

import (
	"github.com/stretchr/testify/assert"

	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func Test_RRSetService_LookupName(t *testing.T) {
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
	_, _, err = c.RRSet.LookupName("*.farsightsecurity.com", nil)
	assert.NotNil(t, err)

	// Verify that it gets and parses a response correctly
	reportServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `{"count":51,"time_first":1372688083,"time_last":1374023864,"rrname":"farsightsecurity.com.","rrtype":"NS","bailiwick":"farsightsecurity.com.","rdata":["ns.lah1.vix.com.","ns1.isc-sns.net.","ns2.isc-sns.com.","ns3.isc-sns.info."]}
{"count":495241,"time_first":1374096380,"time_last":1468324876,"rrname":"farsightsecurity.com.","rrtype":"NS","bailiwick":"farsightsecurity.com.","rdata":["ns5.dnsmadeeasy.com.","ns6.dnsmadeeasy.com.","ns7.dnsmadeeasy.com."]}`)
	}))
	defer reportServer.Close()
	u, err = url.Parse(reportServer.URL)
	assert.Nil(t, err)
	c.BaseURL = u
	actual, _, err := c.RRSet.LookupName("*.farsightsecurity.com", &RRSetLookupNameOptions{
		RRType:    "NS",
		Bailiwick: "farsightsecurity.com",
	})
	assert.Nil(t, err)
	assert.Equal(t, []RRSet{
		RRSet{
			Count:     Uint64(51),
			Bailiwick: String("farsightsecurity.com."),
			TimeFirst: NewTimestamp(1372688083),
			TimeLast:  NewTimestamp(1374023864),
			RRName:    String("farsightsecurity.com."),
			RRType:    String("NS"),
			RData: []string{
				"ns.lah1.vix.com.",
				"ns1.isc-sns.net.",
				"ns2.isc-sns.com.",
				"ns3.isc-sns.info.",
			},
		},
		RRSet{
			Count:     Uint64(495241),
			Bailiwick: String("farsightsecurity.com."),
			TimeFirst: NewTimestamp(1374096380),
			TimeLast:  NewTimestamp(1468324876),
			RRName:    String("farsightsecurity.com."),
			RRType:    String("NS"),
			RData: []string{
				"ns5.dnsmadeeasy.com.",
				"ns6.dnsmadeeasy.com.",
				"ns7.dnsmadeeasy.com.",
			},
		},
	}, actual)
}
