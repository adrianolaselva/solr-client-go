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

const (
	ActionUpload	=	"UPLOAD"
)

type Config struct {
	Name 			string 		`url:"name,omitempty"`
}

type ConfigParameter struct {
	Action 			string 		`url:"action,omitempty"`
	Name 			string 		`url:"name,omitempty"`
	OmitHeader 		bool 		`url:"omitHeader,omitempty"`
}

type ConfigAPI struct {
	client *Client
}

type CreateConfig struct {
	Name 					string	`json:"name,omitempty"`
	BaseConfigSet 			string	`json:"baseConfigSet,omitempty"`
	ConfigSetPropImmutable 	bool	`json:"configSetProp.immutable,omitempty"`
}

func (c *ConfigAPI) List(ctx context.Context) (*Response, error) {
	req, err := c.client.NewRequest(ctx, http.MethodGet, "/api/cluster/configs", nil, &ConfigParameter{
		OmitHeader: true,
	}, nil)
	if err != nil {
		return nil, err
	}

	response, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, err
}

func (c *ConfigAPI) Upload(ctx context.Context, filename string, name string) (*Response, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("application/octet-stream", filepath.Base(file.Name()))
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

	req, err := c.client.NewRequest(ctx, http.MethodPost, "/solr/admin/configs", body, &ConfigParameter{
		Action: ActionUpload,
		Name:   name,
	}, &map[string]string{
		"Content-Type": "application/octet-stream",
	})
	if err != nil {
		return nil, err
	}

	response, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, err
}

func (c *ConfigAPI) Create(ctx context.Context, config CreateConfig) (*Response, error) {
	req, err := c.client.NewRequest(ctx, http.MethodPost, "/api/cluster/configs", config, &ConfigParameter{
		OmitHeader: true,
	}, nil)
	if err != nil {
		return nil, err
	}

	response, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, err
}

func (c *ConfigAPI) Delete(ctx context.Context, name string) (*Response, error) {
	path := fmt.Sprintf("/api/cluster/configs/%s", name)

	req, err := c.client.NewRequest(ctx, http.MethodDelete, path, nil, &ConfigParameter{
		OmitHeader: true,
	}, nil)
	if err != nil {
		return nil, err
	}

	response, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, err
}