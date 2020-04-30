package solr

import (
	"fmt"
	"github.com/google/go-querystring/query"
	"net/http"
	"net/url"
	"time"
	"github.com/pkg/errors"
)

type Config struct {
	http 		http.Client
	url 		*url.URL
	prefix		string
}

func NewConfig() *Config {
	return &Config{http: http.Client{
		Timeout:       20 * time.Second,
	}, url: &url.URL{
		Scheme:     "http",
		Host:       "127.0.0.1:8981",
	}}
}

func (c *Config) SetUrl(host string) *Config {
	c.url, _ = url.Parse(host)
	return c
}

func (c *Config) SetHttpScheme(httpScheme string) *Config {
	c.url.Scheme = httpScheme
	return c
}

func (c *Config) SetPrefix(prefix string) *Config {
	c.prefix = prefix
	return c
}

func (c *Config) SetTimeout(timeout int) *Config {
	c.http = http.Client{
		Timeout:	time.Duration(timeout),
	}
	return c
}

func (c *Config) getUrlWithQueryStrings(endpoint string, params interface{}) (string, error) {
	if c.url != nil {
		endpoint, err := c.url.Parse(fmt.Sprintf("%s%s", c.prefix, endpoint))
		if err != nil {
			return "", err
		}

		params, err := query.Values(params)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("%s?%s", endpoint.String(), params.Encode()), nil
	}

	return "", errors.Errorf("failed to build request url")
}

func (c *Config) getUrl(endpoint string) (string, error) {
	if c.url != nil {
		endpoint, err := c.url.Parse(fmt.Sprintf("%s%s", c.prefix, endpoint))
		if err != nil {
			return "", err
		}

		return endpoint.String(), nil
	}

	return "", errors.Errorf("failed to build request url")
}

