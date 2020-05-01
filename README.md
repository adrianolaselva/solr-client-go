[![Build Status](https://travis-ci.org/adrianolaselva/solr-client-go.svg?branch=master)](https://travis-ci.org/adrianolaselva/solr-client-go)
[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/adrianolaselva/solr-client-go/badges/quality-score.png?b=master)](https://scrutinizer-ci.com/g/adrianolaselva/solr-client-go/?branch=master)
[![GoDoc](https://godoc.org/github.com/adrianolaselva/solr-client-go?status.svg)](https://godoc.org/github.com/adrianolaselva/solr-client-go)
![license](http://img.shields.io/badge/license-Apache%20v2-blue.svg)

solr-client-go
=======

This is a Go client library to access Apache Solr APIs to enable the use of operations from document manipulation to index management.

## Install

go get github.com/adrianolaselva/solr-client-go/solr

## Usage

```go
import "github.com/adrianolaselva/solr-client-go/solr"
```

Create a new Solr customer to consume the endpoints.

## Examples

Instantiate client:

```go
client := solr.NewClient()
```

Find Documents:

```go
response, err := client.Document.Select(context.Background(), "identify-events", "*:*")
if err != nil {
    log.Info(err)
}
```

Create Document(s):

```go
var docs []solr.Document
docs = append(docs, map[string]interface{}{
    "uuid": fmt.Sprintf("%x", md5.Sum([]byte(time.Now().String()))),
    "context": map[string]interface{}{
        "ip": "127.0.0.1",
    },
    "timestamp": "2020-04-27 16:43:57-0300",
})

response, err := client.Document.Update(context.Background(), "identify-events", docs, &solr.Parameters{
    Commit:       true,
})
```

Delete by ID:

```go
d := Delete{
    Id: id,
}

response, err = client.Document.Delete(context.Background(), "identify-events", d, &Parameters{
    Commit:       true,
})
```

Delete by Query:

```go
d := Delete{
    d := Delete{
        Query: "context.ip: 127.0.0.1",
    }
}

response, err = client.Document.Delete(context.Background(), "identify-events", d, &Parameters{
    Commit:       true,
})
```

## Versioning

Each version of the client is tagged and the version is updated accordingly.

To see the list of past versions, run `git tag`.

## License
MIT