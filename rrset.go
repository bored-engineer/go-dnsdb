package dnsdb

// Imports
import (
	"encoding/json"
	"io"
)

// RRSet as described at https://api.dnsdb.info/#rrest-results
type RRSet struct {
	Count         *uint64    `json:"count"`
	Bailiwick     *string    `json:"bailiwick"`
	TimeFirst     *Timestamp `json:"time_first"`
	TimeLast      *Timestamp `json:"time_last"`
	ZoneTimeFirst *Timestamp `json:"zone_time_first"`
	ZoneTimeLast  *Timestamp `json:"zone_time_last"`
	RRName        *string    `json:"rrname"`
	RRType        *string    `json:"rrtype"`
	RData         []string   `json:"rdata"`
}

// RRSetService communicates with the rrset related methods of the DNSDB API.
type RRSetService service

// RRSetLookupNameOptions specifies the optional parameters to the RRSetService.LookupName method.
type RRSetLookupNameOptions struct {
	RRType    string
	Bailiwick string

	LookupOptions
}

// decodeRRSet is a helper function for json streams
func decodeRRSet(reader io.Reader) ([]RRSet, error) {
	var result []RRSet
	dec := json.NewDecoder(reader)
	for {
		var r RRSet
		if err := dec.Decode(&r); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		result = append(result, r)
	}
	return result, nil
}

// LookupName fetches all matching records for the given owner name
func (s *RRSetService) LookupName(ownerName string, opt *RRSetLookupNameOptions) ([]RRSet, *Response, error) {
	path := "lookup/rrset/name/" + ownerName
	var lookupOpt LookupOptions
	if opt != nil {
		lookupOpt = opt.LookupOptions
		if opt.RRType != "" {
			path = path + "/" + opt.RRType
			if opt.Bailiwick != "" {
				path = path + "/" + opt.Bailiwick
			}
		}
	}
	req, err := s.client.NewLookupRequest("GET", path, lookupOpt)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, resp, err
	}

	result, err := decodeRRSet(resp.Body)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, err
}
