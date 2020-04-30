package solr

import (
	"net/http"
	"net/url"
)

type Config struct {
	HTTPClient 	*http.Client
	Url 		*url.URL
	MaxRetries 	*int
	Logger 		*Logger
	Timeout 	*int
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) SetUrl(host string) *Config {
	c.Url, _ = url.Parse(host)
	return c
}

func (c *Config) WithMaxRetries(max int) *Config {
	c.MaxRetries = &max
	return c
}

func (c *Config) WithLogger(logger Logger) *Config {
	c.Logger = &logger
	return c
}

func (c *Config) SetHost(host string) *Config {
	c.Host = &host
	return c
}

func (c *Config) SetTimeout(timeout int) *Config {
	c.Timeout = &timeout
	return c
}
