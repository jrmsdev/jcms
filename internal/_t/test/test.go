// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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

//
// test main
//

func Main(m *testing.M, name string) {
	if wapp != nil {
		panic("wapp is not nil")
	}
	wapp = webapp.New(newConfig(name))
	serverURI = wapp.Start()
	go func() {
		wapp.Serve()
		wapp.Stop()
	}()
	cli = wapp.Client()
	os.Exit(m.Run())
}

// config

func newConfig(name string) *config.Config {
	srcdir := filepath.Join(os.Getenv("GOPATH"), "src")
	cfg := config.New(name)
	cfg.Basedir = filepath.Join(srcdir,
		"github.com", "jrmsdev", "jcms", "testdata", "basedir")
	return cfg
}

// webapp

func Webapp() *webapp.Webapp {
	return wapp
}

//
// test response
//

type TestResponse struct {
	t    *testing.T
	orig *http.Response
}

func newResponse(t *testing.T, r *http.Response) *TestResponse {
	return &TestResponse{t, r}
}

func (r *TestResponse) String() string {
	return fmt.Sprintf("%s", r.orig.Status)
}

func (r *TestResponse) Status(expect int) {
	r.t.Helper()
	if check.NotEqual(r.t, r.orig.StatusCode, expect, "response status") {
		r.t.FailNow()
	}
}

func (r *TestResponse) StatusInfo(expect string) {
	r.t.Helper()
	if check.NotEqual(r.t, r.orig.Status, expect, "response status info") {
		r.t.FailNow()
	}
}

func (r *TestResponse) getBody() string {
	r.t.Helper()
	body, err := ioutil.ReadAll(r.orig.Body)
	if err != nil {
		r.t.Fatal(err)
	}
	r.orig.Body.Close()
	return strings.TrimSpace(string(body))
}

func (r *TestResponse) Body(expect string) {
	r.t.Helper()
	if check.NotEqual(r.t, r.getBody(), expect, "response body") {
		r.t.FailNow()
	}
}

func (r *TestResponse) BodyMatch(pat string) {
	r.t.Helper()
	body := r.getBody()
	m, err := regexp.MatchString(pat, body)
	if err != nil {
		r.t.Fatal(err)
	}
	if !m {
		r.t.Fatalf("'%s' not match body '%s'", pat, body)
	}
}

//
// test client
//

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
