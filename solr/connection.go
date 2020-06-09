package solr

import (
	"net/url"
	"sync"
)

type ConnectionPool interface {
	Next() (*Connection, error)
	URLs() []*url.URL
}

type Connection struct {
	sync.Mutex
	URL *url.URL
}

type singleConnectionPool struct {
	connection *Connection
}

type statusConnectionPool struct {
	sync.Mutex
	live []*Connection
	dead []*Connection
}
