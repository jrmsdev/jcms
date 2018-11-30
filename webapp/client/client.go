// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package client

import (
	"net/http"
	"time"
)

type Client struct {
	uri string
	cli *http.Client
}

func New(uri string) *Client {
	return &Client{uri, &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    10 * time.Second,
			DisableCompression: true,
		},
	}}
}

func (c *Client) url(path string) string {
	return c.uri + path
}

func (c *Client) Get(path string) (*http.Response, error) {
	return c.cli.Get(c.url(path))
}
