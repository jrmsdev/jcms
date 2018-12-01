// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package test

import (
	"testing"

	"github.com/jrmsdev/jcms/webapp/client"
)

type TestClient struct {
	t   *testing.T
	cli *client.Client
}

func Client(t *testing.T) *TestClient {
	return &TestClient{t, cli}
}

func (c *TestClient) Get(p string) *TestResponse {
	c.t.Helper()
	resp, err := c.cli.Get(p)
	if err != nil {
		c.t.Fatal(err)
	}
	return newResponse(c.t, resp)
}
