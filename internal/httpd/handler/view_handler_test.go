// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package handler_test

import (
	"testing"

	"github.com/jrmsdev/jcms/internal/_t/test"
)

func TestViewNotFound(t *testing.T) {
	c := test.Client(t)
	r := c.Get("/notfound/")
	r.Check(404, "text/plain")
}

func TestViewIndex(t *testing.T) {
	c := test.Client(t)
	r := c.Get("/index.html")
	r.Check(200, "text/html")
	r.BodyChecksumMatch("testdata/assets/testing/view/index.html")
}

func TestViewIndexSlash(t *testing.T) {
	c := test.Client(t)
	r := c.Get("/")
	r.Check(200, "text/html")
	r.BodyChecksumMatch("testdata/assets/testing/view/index.html")
}

func TestViewIndexEmpty(t *testing.T) {
	c := test.Client(t)
	r := c.Get("")
	r.Check(200, "text/html")
	r.BodyChecksumMatch("testdata/assets/testing/view/index.html")
}

func TestViewNoSlash(t *testing.T) {
	c := test.Client(t)
	r := c.Get("/testing")
	r.Check(200, "text/html")
	r.BodyChecksumMatch("testdata/assets/testing/view/testing/index.html")
}

func TestViewStaticRedirect(t *testing.T) {
	c := test.Client(t)
	r := c.Get("/static.ext")
	r.Check(301, "text/html")
	r.Header("location", "/static/static.ext")
}
