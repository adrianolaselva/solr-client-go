package solr

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type Parameters struct {
	CommitWithin 	int 						`url:"commitWithin,omitempty"`
	Commit 			bool 						`url:"commit,omitempty"`
	Query 			string 						`url:"q,omitempty"`
	Delete 			interface{} 				`url:"delete,omitempty"`
	LiteralId 		string 						`url:"literal.id,omitempty"`
}

type Delete struct {
	Id 				string 						`json:"id,omitempty"`
	Query 			string 						`json:"query,omitempty"`
}

type Document map[string]interface{}

type DocumentAPI struct {
	client *Client
}

// SELECT: Select documents
func (d *DocumentAPI) Select(ctx context.Context, collection string, query string) (*Response, error) {
	path := fmt.Sprintf("/solr/%s/select", collection)

	req, err := d.client.NewRequest(ctx, http.MethodGet, path, nil, &Parameters{
		Query:	query,
	}, nil)
	if err != nil {
		return nil, err
	}

	response, err := d.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, err
}

// UPDATE: Update/Insert document
func (d *DocumentAPI) Update(ctx context.Context, collection string, doc Document, params *Parameters) (*Response, error) {

	path := fmt.Sprintf("/api/collections/%s/update/json", collection)

	req, err := d.client.NewRequest(ctx, http.MethodPost, path, doc, params, nil)
	if err != nil {
		return nil, err
	}

	response, err := d.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, err
}

// UPDATE MANY: Update/Insert documents
func (d *DocumentAPI) UpdateMany(ctx context.Context, collection string, docs []Document, params *Parameters) (*Response, error) {

	path := fmt.Sprintf("/api/collections/%s/update/json", collection)

	req, err := d.client.NewRequest(ctx, http.MethodPost, path, docs, params, nil)
	if err != nil {
		return nil, err
	}

	response, err := d.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, err
}

// COMMIT: Commit documents
func (d *DocumentAPI) Commit(ctx context.Context, collection string) (*Response, error) {

	path := fmt.Sprintf("/api/collections/%s/update/json", collection)

	req, err := d.client.NewRequest(ctx, http.MethodPost, path, nil, Parameters{
		Commit:       true,
	}, nil)
	if err != nil {
		return nil, err
	}

	response, err := d.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, err
}

// DELETE: Delete document
func (d *DocumentAPI) Delete(ctx context.Context, collection string, delete Delete, params *Parameters) (*Response, error) {

	path := fmt.Sprintf("/solr/%s/update", collection)

	req, err := d.client.NewRequest(ctx, http.MethodPost, path, map[string]interface{}{
		"delete": delete,
	}, params, nil)
	if err != nil {
		return nil, err
	}

	response, err := d.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, err
}

// EXTRACT: Uploading Data with Solr Cell using Apache Tika
func (d *DocumentAPI) Extract(ctx context.Context, collection string, filename string, params *Parameters) (*Response, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("document", filepath.Base(file.Name()))
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("/api/collections/%s/update/extract", collection)

	req, err := d.client.NewRequest(ctx, http.MethodPost, path, body, params, &map[string]string{
		"Content-Type": "application/octet-stream",
	})
	if err != nil {
		return nil, err
	}

	response, err := d.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, err
}