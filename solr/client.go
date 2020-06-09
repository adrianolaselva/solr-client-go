package solr

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/google/go-querystring/query"
)

const (
	DefaultHost        = "http://127.0.0.1:8983"
	DefaultContentType = "application/json"
)

type Client struct {
	client             *http.Client
	baseURL            *url.URL
	Document           DocumentAPI
	Collection         CollectionAPI
	Config             ConfigAPI
	onRequestCompleted RequestCompletionCallback
	username           string
	password           string
}

type RequestCompletionCallback func(*http.Request, *http.Response)

// NEW CLIENT: New Client Instance
func NewClient() Client {
	httpClient := http.DefaultClient
	baseURL, _ := url.Parse(DefaultHost)

	client := Client{
		client:  httpClient,
		baseURL: baseURL,
	}

	client.Initialize()

	return client
}

// INITIALIZE: Initialize Instances
func (c *Client) Initialize() {
	document := DocumentAPI{client: c}
	c.Document = document
	collection := CollectionAPI{client: c}
	c.Collection = collection
	config := ConfigAPI{client: c}
	c.Config = config
}

// SET HTTP CLIENT: Set HTTP Client Instance
func (c *Client) SetHttpClient(httpClient *http.Client) *Client {
	c.client = httpClient
	c.Initialize()
	return c
}

// SET BASIC AUTH: Add Credentials for use Basic Authentication
func (c *Client) SetBasicAuth(username string, password string) *Client {
	c.username = username
	c.password = password
	c.Initialize()
	return c
}

// SET BASE URL: Ser Base URL
func (c *Client) SetBaseURL(baseURL string) *Client {
	c.baseURL, _ = url.Parse(baseURL)
	c.Initialize()
	return c
}

// NEW UPLOAD: New Request Upload
func (c *Client) NewUpload(ctx context.Context, urlStr string, filepath string, queryStrings interface{}) (*Response, error) {
	u, err := c.baseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	params, _ := query.Values(queryStrings)
	u.RawQuery = params.Encode()

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	resp, err := http.Post(u.String(), "application/octet-stream", file)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := Response{HttpResponse: resp}

	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// NEW REQUEST: New Request
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

// NEW REQUEST UPLOAD: New Request Upload
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

// DO: Response Handle
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

	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := Response{HttpResponse: resp}

	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// ON REQUEST COMPLETED: On Request Completed Handle
func (c *Client) OnRequestCompleted(rc RequestCompletionCallback) {
	c.onRequestCompleted = rc
}
