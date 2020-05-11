package solr

import (
	"context"
	"crypto/md5"
	"fmt"
	"testing"
	"time"
)

func TestDocumentCreate(t *testing.T) {
	client := NewClient()

	response, err := client.Collection.Create(context.Background(), CollectionCreate{
		Name:                 "tests",
		RouterName:           "compositeId",
		NumShards:            1,
		ReplicationFactor: 	  1,
		CollectionConfigName: "_default",
		Async:                false,
	})
	if err != nil {
		t.Errorf("failed to create collection %v", err)
	}

	if response.ResponseHeader.Status != 0 {
		t.Errorf("failed to create collection %v", err)
	}

	var docs []Document
	docs = append(docs, map[string]interface{}{
		"uuid": fmt.Sprintf("%x", md5.Sum([]byte(time.Now().String()))),
		"context": map[string]interface{}{
			"ip": "127.0.0.1",
		},
		"timestamp": "2020-04-27 16:43:57-0300",
	})

	response, err = client.Document.UpdateMany(context.Background(), "tests", docs, &Parameters{
		Commit:       true,
	})
	if err != nil {
		t.Errorf("failed to create document and commit %v", err)
	}

	if response.ResponseHeader.Status != 0 {
		t.Errorf("failed to create document and commit %v", err)
	}
}

func TestDocumentCreateBulkWithCommit(t *testing.T) {
	client := NewClient()

	var docs []Document
	for i:=0;i<10;i++ {
		docs = append(docs, map[string]interface{}{
			"uuid": fmt.Sprintf("%x", md5.Sum([]byte(time.Now().String()))),
			"iteration": i,
			"context": map[string]interface{}{
				"ip": fmt.Sprintf("127.0.0.%v", i),
			},
			"timestamp": "2020-04-27 16:43:57-0300",
		})
	}

	response, err := client.Document.UpdateMany(context.Background(), "tests", docs, &Parameters{
		Commit:       true,
	})
	if err != nil {
		t.Errorf("failed to create documents and commit %v", err)
	}

	if response.ResponseHeader.Status != 0 {
		t.Errorf("failed to create documents and commit %v", err)
	}
}

func TestDocumentSelectAll(t *testing.T) {
	client := NewClient()

	var docs []Document
	for i:=0;i<10;i++ {
		docs = append(docs, map[string]interface{}{
			"uuid": fmt.Sprintf("%x", md5.Sum([]byte(time.Now().String()))),
			"iteration": i,
			"context": map[string]interface{}{
				"ip": fmt.Sprintf("127.0.0.%v", i),
			},
			"timestamp": "2020-04-27 16:43:57-0300",
		})
	}

	response, err := client.Document.UpdateMany(context.Background(), "tests", docs, &Parameters{
		Commit:       true,
	})
	if err != nil {
		t.Errorf("failed to add documents and commit for execute query %v", err)
	}

	response, err = client.Document.Select(context.Background(), "tests", "*:*")
	if err != nil {
		t.Errorf("failed to select all documents %v", err)
	}

	if response.ResponseHeader.Status != 0 {
		t.Errorf("failed to select all documents %v", err)
	}

}

func TestDocumentDelete(t *testing.T) {
	client := NewClient()

	id := fmt.Sprintf("%x", md5.Sum([]byte(time.Now().String())))

	var docs []Document
	docs = append(docs, map[string]interface{}{
		"id": fmt.Sprintf("%x", md5.Sum([]byte(time.Now().String()))),
		"context": map[string]interface{}{
			"ip": "127.0.0.1",
		},
		"timestamp": "2020-04-27 16:43:57-0300",
	})

	response, err := client.Document.UpdateMany(context.Background(), "tests", docs, &Parameters{
		Commit:       true,
	})
	if err != nil {
		t.Errorf("failed to create document and commit %v", err)
	}

	if response.ResponseHeader.Status != 0 {
		t.Errorf("failed to create document and commit %v", err)
	}

	d := Delete{
		Id: id,
	}

	response, err = client.Document.Delete(context.Background(), "tests", d, &Parameters{
		Commit:       true,
	})
	if err != nil {
		t.Errorf("failed to create document and commit %v", err)
	}

	if response.ResponseHeader.Status != 0 {
		t.Errorf("failed to create document and commit %v", err)
	}
}

func TestDocumentDeleteByQuery(t *testing.T) {
	client := NewClient()

	d := Delete{
		Query: "context.ip: 127.0.0.1",
	}

	response, err := client.Document.Delete(context.Background(), "tests", d, &Parameters{
		Commit:       true,
	})
	if err != nil {
		t.Errorf("failed to create document and commit %v", err)
	}

	if response.ResponseHeader.Status != 0 {
		t.Errorf("failed to create document and commit %v", err)
	}
}

func TestDocumentExtract(t *testing.T) {
	client := NewClient()

	_, err := client.Document.Extract(context.Background(), "tests", "../example/lorem-ipsum.pdf", &Parameters{
		Commit:       true,
	})
	if err != nil {
		t.Errorf("failed to extract pdf document %v", err)
	}

	//if response.ResponseHeader.Status != 0 {
	//	t.Errorf("failed to extract pdf document %v", err)
	//}
}

func TestDocumentDeleteAll(t *testing.T) {
	client := NewClient()

	d := Delete{
		Query: "*:*",
	}

	response, err := client.Document.Delete(context.Background(), "tests", d, &Parameters{
		Commit:       true,
	})
	if err != nil {
		t.Errorf("failed to create document and commit %v", err)
	}

	if response.ResponseHeader.Status != 0 {
		t.Errorf("failed to create document and commit %v", err)
	}

	response, err = client.Collection.Delete(context.Background(), CollectionDelete{
		Name:           "tests",
		Async:          false,
	})
	if err != nil {
		t.Errorf("failed to delete collection %v", err)
	}

	if response.ResponseHeader.Status != 0 {
		t.Errorf("failed to delete collection %v", err)
	}
}
