[![Build Status](https://travis-ci.org/adrianolaselva/solr-client-go.svg?branch=master)](https://travis-ci.org/adrianolaselva/solr-client-go)
[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/adrianolaselva/solr-client-go/badges/quality-score.png?b=master)](https://scrutinizer-ci.com/g/adrianolaselva/solr-client-go/?branch=master)
[![GoDoc](https://godoc.org/github.com/adrianolaselva/solr-client-go?status.svg)](https://pkg.go.dev/github.com/adrianolaselva/solr-client-go)
![license](http://img.shields.io/badge/license-Apache%20v2-blue.svg)

Apache Solr client Go
=======

This is a Go client library to access Apache Solr APIs to enable the use of operations from document manipulation to index management.

## Install

```sh
go get github.com/adrianolaselva/solr-client-go/solr
```

## Usage

```go
import "github.com/adrianolaselva/solr-client-go/solr"
```

Create a new Solr customer to consume the endpoints.

## Examples

Instantiate solr client:

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
    Query: "context.ip: 127.0.0.1",
}

response, err = client.Document.Delete(context.Background(), "identify-events", d, &Parameters{
    Commit:       true,
})
```

Create new collection:

```go
response, err := client.Collection.Create(context.Background(), CollectionCreate{
    Name:                 "collection-test",
    RouterName:           "compositeId",
    NumShards:            1,
    ReplicationFactor: 	  1,
    CollectionConfigName: "_default",
    Async:                false,
})
```

Reload Collection:

```go
response, err := client.Collection.Reload(context.Background(), CollectionReload{
    Name:           "collection-test",
    Async:          false,
})
```

Modify Collection:

```go
response, err := client.Collection.Modify(context.Background(), CollectionModifyCollection{
    Collection: "collection-test",
    MaxShardsPerNode:      1,
})
```


List Collections:

```go
response, err := client.Collection.List(context.Background())
```

Migrate collection

```go
response, err = client.Collection.Migrate(context.Background(), CollectionMigrate{
    Collection:       "collection-test",
    TargetCollection: "collection-test-migrate",
    SplitKey:         "a!",
    ForwardTimeout:   100000,
    Async:            false,
})
```

Collection backup:

```go
response, err := client.Collection.Backup(context.Background(), CollectionBackup{
	Collection:     "collection-test",
	Name:           backupFilePath,
	Location:       "/tmp/",
	Async:          false,
})
```

Collection restore:

```go
response, err = client.Collection.Restore(context.Background(), CollectionRestore{
	Collection:           backupFilePath,
	Name:                 backupFilePath,
	Location:       	  "/tmp/",
	Async:                false,
	ReplicationFactor:    1,
})
```

Create collection configuration:

```go
response, err := client.Config.Upload(context.Background(), "../example/configs/identify-events.zip", "identify-events")
```

>Obs: example collection settings in the `./example/configs/` directory.

Create a new configuration using another as a base:

```go
response, err := client.Config.Create(context.Background(), CreateConfig{
    Create: Config{
        Name:                   "identify-events.CREATE",
        BaseConfigSet:          "identify-events",
    },
})
```

Remove collection configuration:

```go
response, err := client.Config.Delete(context.Background(), "identify-events")
```

Atomic Updates:

```go
var docs []Document
docs = append(docs, map[string]interface{}{
    "id": id,
    "author_s": map[string]interface{}{
        "set": "Teste 2",
    },
    "copies_i": map[string]interface{}{
        "inc": 5,
    },
    "cat_ss": map[string]interface{}{
        "add": time.Now().Format(time.RFC3339),
    },
})

response, err = client.Document.AtomicUpdateMany(context.Background(), "identify-events", docs, &Parameters{
    Commit:       true,
})
```

>Obs: Solr supports several modifiers that atomically update values of a document. This allows updating only specific fields, which can help speed indexing processes in an environment where speed of index additions is critical to the application.

>To use atomic updates, add a modifier to the field that needs to be updated. The content can be updated, added to, or incrementally increased if a number.

>[Atomic update](https://lucene.apache.org/solr/guide/6_6/updating-parts-of-documents.html#UpdatingPartsofDocuments-AtomicUpdates)

## Versioning

Each version of the client is tagged and the version is updated accordingly.

To see the list of past versions, run `git tag`.

## License
MIT
