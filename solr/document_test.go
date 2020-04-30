package solr

import (
	"crypto/md5"
	"fmt"
	"testing"
	"time"
)

func TestDocumentCreate(t *testing.T) {
	config := NewConfig().
		SetUrl("http://127.0.0.1:8981")

	client := NewClient(config)

	var documents []map[string]interface{}

	documents = append(documents, map[string]interface{}{
		"uuid": fmt.Sprintf("%x", md5.Sum([]byte(time.Now().String()))),
		"context": map[string]interface{}{
			"ip": "127.0.0.1",
		},
		"timestamp": "2020-04-27 16:43:57-0300",
	})

	response, err := client.Document.Update("identify-events", documents, DocumentParameters{
		Commit:       false,
	})
	if err != nil {
		t.Errorf("failed to create document %v", err)
	}

	if response.Header.Status != 0 {
		t.Errorf("failed to create document %v", err)
	}

	response, err = client.Document.Commit("identify-events")
	if err != nil {
		t.Errorf("failed to commit document %v", err)
	}

	if response.Header.Status != 0 {
		t.Errorf("failed to commit document %v", err)
	}
}

func TestDocumentCreateBulkWithCommit(t *testing.T) {
	config := NewConfig().
		SetUrl("http://127.0.0.1:8981")

	client := NewClient(config)

	var documents []map[string]interface{}
	for i:=0;i<10;i++ {
		documents = append(documents, map[string]interface{}{
			"uuid": fmt.Sprintf("%x", md5.Sum([]byte(time.Now().String()))),
			"context": map[string]interface{}{
				"ip": "127.0.0.1",
			},
			"timestamp": "2020-04-27 16:43:57-0300",
		})
	}

	response, err := client.Document.Update("identify-events", documents, DocumentParameters{
		Commit:       true,
	})
	if err != nil {
		t.Errorf("failed to create document %v", err)
	}

	if response.Header.Status != 0 {
		t.Errorf("failed to create document %v", err)
	}
}