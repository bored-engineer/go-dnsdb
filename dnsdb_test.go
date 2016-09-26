package dnsdb

// Imports
import (
	"github.com/stretchr/testify/assert"

	"net/http"
	"testing"
)

func Test_extractRateFromResponse(t *testing.T) {
	header := http.Header{}
	header.Add("X-RateLimit-Limit", "unlimited")
	header.Add("X-RateLimit-Remaining", "n/a")
	rate := extractRateFromResponse(&http.Response{
		Header: header,
	})
	assert.Equal(t, Rate{
		Limit:     -1,
		Remaining: -1,
	}, rate)
}
