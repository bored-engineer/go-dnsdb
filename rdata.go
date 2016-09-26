package dnsdb

// Imports
import (
	"encoding/hex"
	"encoding/json"
	"io"
	"net"
	"strings"
)

// RData as described at https://api.dnsdb.info/#rdata-lookups
type RData struct {
	Count         *uint64    `json:"count"`
	TimeFirst     *Timestamp `json:"time_first"`
	TimeLast      *Timestamp `json:"time_last"`
	ZoneTimeFirst *Timestamp `json:"zone_time_first"`
	ZoneTimeLast  *Timestamp `json:"zone_time_last"`
	RRName        *string    `json:"rrname"`
	RRType        *string    `json:"rrtype"`
	RData         *string    `json:"rdata"`
}

// decodeRData is a helper function for json streams
func decodeRData(reader io.Reader) ([]RData, error) {
	var result []RData
	dec := json.NewDecoder(reader)
	for {
		var r RData
		if err := dec.Decode(&r); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		result = append(result, r)
	}
	return result, nil
}

// RDataService communicates with the rdata related methods of the DNSDB API.
type RDataService service

// RDataLookupNameOptions specifies the optional parameters to the RDataService.LookupName method.
type RDataLookupNameOptions struct {
	RRType string

	LookupOptions
}

// LookupName fetches all matching records for the provided name
func (s *RDataService) LookupName(name string, opt *RDataLookupNameOptions) ([]RData, *Response, error) {
	path := "lookup/rdata/name/" + name
	var lookupOpt LookupOptions
	if opt != nil {
		lookupOpt = opt.LookupOptions
		if opt.RRType != "" {
			path = path + "/" + opt.RRType
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

	result, err := decodeRData(resp.Body)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, err
}

// RDataLookupIPOptions specifies the optional parameters to the RDataService.LookupIP method.
type RDataLookupIPOptions struct {
	RRType string

	LookupOptions
}

// LookupIP fetches all matching records for the provided IP
func (s *RDataService) LookupIP(ip net.IP, opt *RDataLookupIPOptions) ([]RData, *Response, error) {
	if opt == nil {

	}
	path := "lookup/rdata/ip/" + ip.String()
	var lookupOpt LookupOptions
	if opt != nil {
		lookupOpt = opt.LookupOptions
		if opt.RRType != "" {
			path = path + "/" + opt.RRType
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

	result, err := decodeRData(resp.Body)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, err
}

// RDataLookupIPNetOptions specifies the optional parameters to the RDataService.LookupIPNet method.
type RDataLookupIPNetOptions struct {
	RRType string

	LookupOptions
}

// LookupIPNet fetches all matching records for the provided IPNet
func (s *RDataService) LookupIPNet(ipnet net.IPNet, opt *RDataLookupIPNetOptions) ([]RData, *Response, error) {
	path := "lookup/rdata/ip/" + strings.Replace(ipnet.String(), "/", ",", 1)
	var lookupOpt LookupOptions
	if opt != nil {
		lookupOpt = opt.LookupOptions
		if opt.RRType != "" {
			path = path + "/" + opt.RRType
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

	result, err := decodeRData(resp.Body)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, err
}

// RDataLookupRawOptions specifies the optional parameters to the RDataService.LookupRaw method.
type RDataLookupRawOptions struct {
	RRType string

	LookupOptions
}

// LookupRaw fetches all matching records for the provided raw bytes and optional RRType (set to "")
func (s *RDataService) LookupRaw(raw []byte, opt *RDataLookupRawOptions) ([]RData, *Response, error) {
	path := "lookup/rdata/raw/" + hex.EncodeToString(raw)
	var lookupOpt LookupOptions
	if opt != nil {
		lookupOpt = opt.LookupOptions
		if opt.RRType != "" {
			path = path + "/" + opt.RRType
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

	result, err := decodeRData(resp.Body)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, err
}
