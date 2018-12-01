// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package test

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/jrmsdev/jcms/internal/_t/check"
	"github.com/jrmsdev/jcms/webapp"
	"github.com/jrmsdev/jcms/webapp/client"
	"github.com/jrmsdev/jcms/webapp/config"
)

var (
	wapp      *webapp.Webapp
	cli       *client.Client
	serverURI string
)

func Main(m *testing.M, name string) {
	if wapp != nil {
		panic("wapp is not nil")
	}
	wapp = webapp.New(newConfig(name))
	serverURI = wapp.Start()
	defer wapp.Stop()
	go func() {
		wapp.Serve()
	}()
	cli = wapp.Client()
	rc := m.Run()
	os.Exit(rc)
}

func newConfig(name string) *config.Config {
	cfg := config.New(name)
	srcdir := filepath.Join(os.Getenv("GOPATH"), "src")
	if srcdir == "" {
		panic("no GOPATH")
	}
	cfg.Basedir = filepath.Join(srcdir,
		"github.com", "jrmsdev", "jcms", "testdata")
	return cfg
}

func Webapp() *webapp.Webapp {
	return wapp
}

type TestResponse struct {
	t *testing.T
	orig *http.Response
}

func newResponse(t *testing.T, r *http.Response) *TestResponse {
	return &TestResponse{t, r}
}

func (r *TestResponse) Status(expect int) {
	if check.NotEqual(r.t, r.orig.StatusCode, expect, "response status") {
		r.t.FailNow()
	}
}

func (r *TestResponse) StatusInfo(expect string) {
	if check.NotEqual(r.t, r.orig.Status, expect, "response status info") {
		r.t.FailNow()
	}
}

func (r *TestResponse) Body(expect string) {
	body, err := ioutil.ReadAll(r.orig.Body)
	if err != nil {
		r.t.Fatal(err)
	}
	r.orig.Body.Close()
	if check.NotEqual(r.t, string(body), expect, "response body") {
		r.t.FailNow()
	}
}

type TestClient struct {
	t *testing.T
	cli *client.Client
}

func Client(t *testing.T) *TestClient {
	return &TestClient{t, cli}
}

func (c *TestClient) Get(p string) *TestResponse {
	resp, err := c.cli.Get(p)
	if err != nil {
		c.t.Fatal(err)
	}
	return newResponse(c.t, resp)
}
