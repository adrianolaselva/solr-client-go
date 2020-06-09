package solr

import (
	"context"
	"fmt"
	"net/http"
)

const (
	ActionUpload = "UPLOAD"
)

type Config struct {
	Name                   string `json:"name,omitempty"url:"name,omitempty"`
	BaseConfigSet          string `json:"baseConfigSet,omitempty"`
	ConfigSetPropImmutable bool   `json:"configSetProp.immutable,omitempty"`
}

type ConfigParameter struct {
	Action     string `url:"action,omitempty"`
	Name       string `url:"name,omitempty"`
	OmitHeader bool   `url:"omitHeader,omitempty"`
}

type ConfigAPI struct {
	client *Client
}

type CreateConfig struct {
	Create Config `json:"create,omitempty"`
}

// LIST: List a Configset
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

// UPLOAD: Upload a Configset
func (c *ConfigAPI) Upload(ctx context.Context, filename string, name string) (*Response, error) {
	response, err := c.client.NewUpload(ctx, "/solr/admin/configs", filename, &ConfigParameter{
		Action: ActionUpload,
		Name:   name,
	})
	if err != nil {
		return nil, err
	}

	return response, err
}

// CREATE: Create a Configset
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

// DELETE: Create a Configset
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
