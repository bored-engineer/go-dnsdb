go-dnsdb [![GoDoc][doc-img]][doc] [![Build Status][ci-img]][ci] [![Coverage Status][cov-img]][cov]
======
An (Unofficial) Golang client for [Farsight Security's](https://www.farsightsecurity.com) [DNSDB](https://www.dnsdb.info/) (https://api.dnsdb.info/)

# Usage
```go
import "github.com/bored-engineer/go-dnsdb"
```

To list all RRSet records matching `www.farsightsecurity.com`:
```go
records, _, err := client.RRSet.LookupName("www.farsightsecurity.com", nil)
if err != nil {
	panic(err)
}
for _, record := range records {
	fmt.Println("%s: %v", *record.RRName, record.RData)
}
```

## Authentication
The `dnsdb` library does not directly handle authentication. Instead, when creating a new client, you can pass a `http.Client` that handles authentication for you. It does provide a `APIKeyTransport` structure when using API Key authentication. It is used like this:
```go
tp := dnsdb.APIKeyTransport{
	APIKey: "d41d8cd98f00b204e9800998ecf8427e",
}

client := dnsdb.NewClient(tp.Client())
```

[doc-img]: https://godoc.org/github.com/bored-engineer/go-dnsdb?status.svg
[doc]: https://godoc.org/github.com/bored-engineer/go-dnsdb
[ci-img]: https://travis-ci.org/bored-engineer/go-dnsdb.svg?branch=master
[ci]: https://travis-ci.org/bored-engineer/go-dnsdb
[cov-img]: https://coveralls.io/repos/github/bored-engineer/go-dnsdb/badge.svg?branch=master&
[cov]: https://coveralls.io/github/bored-engineer/go-dnsdb?branch=master
