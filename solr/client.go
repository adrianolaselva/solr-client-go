package solr

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/google/go-querystring/query"
	"net/http"
	"net/url"
)

const (
	DefaultHost = "http://127.0.0.1:8983"
	DefaultContentType = "application/json"
)

type Client struct {
	client 				*http.Client
	baseURL 			*url.URL
	Document 			DocumentAPI
	Collection 			CollectionAPI
	Config 				ConfigAPI
	onRequestCompleted 	RequestCompletionCallback
	username			string
	password			string
}

type RequestCompletionCallback func(*http.Request, *http.Response)

func NewClient() Client {

	httpClient := http.DefaultClient
	baseURL, _ := url.Parse(DefaultHost)
	
	client := Client{
		client:             httpClient,
		baseURL:            baseURL,
	}

	client.Initialize()

	return client
}

func (c *Client) Initialize() {
	document := DocumentAPI{client: c}
	c.Document = document
	collection := CollectionAPI{client: c}
	c.Collection = collection
	config := ConfigAPI{client: c}
	c.Config = config
}

func (c *Client) SetHttpClient(httpClient *http.Client) *Client {
	c.client = httpClient
	c.Initialize()
	return c
}

func (c *Client) SetBasicAuth(username string, password string) *Client {
	c.username = username
	c.password = password
	c.Initialize()
	return c
}

func (c *Client) SetBaseURL(baseURL string) *Client{
	c.baseURL, _ = url.Parse(baseURL)
	c.Initialize()
	return c
}

func (c *Client) NewRequest(ctx context.Context, method, urlStr string, body interface{}, queryStrings interface{}, headers *map[string]string) (*http.Request, error) {
	u, err := c.baseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if body != nil {
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	params, _ := query.Values(queryStrings)
	u.RawQuery = params.Encode()

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", DefaultContentType)
	req.Header.Add("Accept", DefaultContentType)

	if headers != nil {
		for key, value := range *headers {
			req.Header.Set(key, value)
		}
	}

	if c.username != "" && c.password != "" {
		req.SetBasicAuth(c.username, c.password)
	}

	return req, nil
}

func (c *Client) NewRequestUpload(ctx context.Context, method, urlStr string, body interface{}, queryStrings interface{}) (*http.Request, error) {
	u, err := c.baseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if body != nil {
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	params, _ := query.Values(queryStrings)
	u.RawQuery = params.Encode()

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", DefaultContentType)
	req.Header.Add("Accept", DefaultContentType)

	if c.username != "" && c.password != "" {
		req.SetBasicAuth(c.username, c.password)
	}

	return req, nil
}

func (c *Client) Do(ctx context.Context, req *http.Request) (*Response, error) {
	req = req.WithContext(ctx)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, resp)
	}

	defer func() {
		if rerr := resp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	response := Response{HttpResponse: resp}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) OnRequestCompleted(rc RequestCompletionCallback) {
	c.onRequestCompleted = rc
}