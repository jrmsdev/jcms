// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package client

import (
	"net/http"
	"time"

	"github.com/jrmsdev/jcms/internal/log"
)

type Client struct {
	uri string
	cli *http.Client
}

func New(uri string) *Client {
	log.D("New")
	return &Client{uri, &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    10 * time.Second,
			DisableCompression: true,
		},
		CheckRedirect: checkRedirect,
	}}
}

func checkRedirect(req *http.Request, via []*http.Request) error {
	log.D("checkRedirect %s", req.URL.String())
	return http.ErrUseLastResponse
}

func (c *Client) url(path string) string {
	return c.uri + path
}

func (c *Client) Get(path string) (*http.Response, error) {
	log.D("Get %s", path)
	return c.cli.Get(c.url(path))
}
