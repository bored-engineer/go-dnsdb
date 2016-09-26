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

func Test_RDataService_LookupName(t *testing.T) {
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
	_, _, err = c.RData.LookupName("ns5.dnsmadeeasy.com", nil)
	assert.NotNil(t, err)

	// Verify that it gets and parses a response correctly
	reportServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `{"count":45644,"time_first":1372706073,"time_last":1468330740,"rrname":"fsi.io.","rrtype":"MX","rdata":"10 hq.fsi.io."}
{"count":19304,"time_first":1374098929,"time_last":1468333042,"rrname":"farsightsecurity.com.","rrtype":"MX","rdata":"10 hq.fsi.io."}`)
	}))
	defer reportServer.Close()
	u, err = url.Parse(reportServer.URL)
	assert.Nil(t, err)
	c.BaseURL = u
	actual, _, err := c.RData.LookupName("hq.fsi.io", &RDataLookupNameOptions{
		RRType: "MX",
	})
	assert.Nil(t, err)
	assert.Equal(t, []RData{
		RData{
			Count:         Uint64(45644),
			ZoneTimeFirst: NewTimestamp(1372706073),
			ZoneTimeLast:  NewTimestamp(1468330740),
			RRName:        String("fsi.io."),
			RRType:        String("MX"),
			RData:         String("10 hq.fsi.io."),
		},
		RData{
			Count:     Uint64(19304),
			TimeFirst: NewTimestamp(1374098929),
			TimeLast:  NewTimestamp(1468333042),
			RRName:    String("farsightsecurity.com."),
			RRType:    String("MX"),
			RData:     String("10 hq.fsi.io."),
		},
	}, actual)
}

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

func Test_RDataService_LookupIPNet(t *testing.T) {
	// Setup a client
	c := NewClient(nil)
	_, ipnet, _ := net.ParseCIDR("104.244.13.104/29")

	// Verify that an error response fails
	errorServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Oh No", 500)
	}))
	defer errorServer.Close()
	u, err := url.Parse(errorServer.URL)
	assert.Nil(t, err)
	c.BaseURL = u
	_, _, err = c.RData.LookupIPNet(*ipnet, nil)
	assert.NotNil(t, err)

	// Verify that it gets and parses a response correctly
	reportServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `{"count":24,"time_first":1433550785,"time_last":1468312116,"rrname":"www.farsighsecurity.com.","rrtype":"A","rdata":"104.244.13.104"}
{"count":9429,"time_first":1427897872,"time_last":1468333042,"rrname":"farsightsecurity.com.","rrtype":"A","rdata":"104.244.13.104"}
`)
	}))
	defer reportServer.Close()
	u, err = url.Parse(reportServer.URL)
	assert.Nil(t, err)
	c.BaseURL = u
	actual, _, err := c.RData.LookupIPNet(*ipnet, nil)
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
