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
	DefaultHost = "http://127.0.0.1:8981"
	DefaultContentType = "application/json"
)

type Client struct {
	client 			*http.Client
	baseURL 		*url.URL
	Document 		DocumentAPI
	Collection 		CollectionAPI
	onRequestCompleted RequestCompletionCallback
}

type RequestCompletionCallback func(*http.Request, *http.Response)

func NewClient() Client {

	httpClient := http.DefaultClient
	baseURL, _ := url.Parse(DefaultHost)
	
	client := Client{
		client:             httpClient,
		baseURL:            baseURL,
	}

	initialize(&client)

	return client
}

func initialize(client *Client) {
	document := DocumentAPI{client: client}
	client.Document = document
	collection := CollectionAPI{client: client}
	client.Collection = collection
}

func (c *Client) SetHttpClient(httpClient *http.Client) *Client {
	c.client = httpClient
	initialize(c)
	return c
}

func (c *Client) SetBaseURL(baseURL string) *Client{
	c.baseURL, _ = url.Parse(baseURL)
	initialize(c)
	return c
}

func (c *Client) NewRequest(ctx context.Context, method, urlStr string, body interface{}, queryStrings interface{}) (*http.Request, error) {
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