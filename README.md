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
For furthur usage see the [GoDocs][doc].

## Authentication
The `dnsdb` library does not directly handle authentication. Instead, when creating a new client, you can pass a `http.Client` that handles authentication for you. It does provide a `APIKeyTransport` structure when using API Key authentication. It is used like this:
```go
tp := dnsdb.APIKeyTransport{
    APIKey: "<<your API key here>>",
}

client := dnsdb.NewClient(tp.Client())
```


Assemblying those two fragments of code into a complete GoLang program:
```go
package main

import "github.com/bored-engineer/go-dnsdb"
import "fmt"

func main() {
    fmt.Println("starting dnsdb api test...")

    tp := dnsdb.APIKeyTransport{
        APIKey: "<<your API key here>>",
    }

    client := dnsdb.NewClient(tp.Client())

    records, _, err := client.RRSet.LookupName("www.farsightsecurity.com", nil)
    if err != nil {
        panic(err)
    }

    for _, record := range records {
        fmt.Println("%s: %v", *record.RRName, record.RData)
    }
}
```

[doc-img]: https://godoc.org/github.com/bored-engineer/go-dnsdb?status.svg
[doc]: https://godoc.org/github.com/bored-engineer/go-dnsdb
[ci-img]: https://travis-ci.org/bored-engineer/go-dnsdb.svg?branch=master
[ci]: https://travis-ci.org/bored-engineer/go-dnsdb
[cov-img]: https://coveralls.io/repos/github/bored-engineer/go-dnsdb/badge.svg?branch=master&
[cov]: https://coveralls.io/github/bored-engineer/go-dnsdb?branch=master
