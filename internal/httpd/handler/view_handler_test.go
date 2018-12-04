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
	r.Status(404)
}

func TestViewIndex(t *testing.T) {
	c := test.Client(t)
	r := c.Get("/index.html")
	r.Status(200)
	r.BodyChecksumMatch("testdata/assets/testing/view/index.html")
}

func TestViewIndexSlash(t *testing.T) {
	c := test.Client(t)
	r := c.Get("/")
	r.Status(200)
	r.BodyChecksumMatch("testdata/assets/testing/view/index.html")
}

func TestViewIndexEmpty(t *testing.T) {
	c := test.Client(t)
	r := c.Get("")
	r.Status(200)
	r.BodyChecksumMatch("testdata/assets/testing/view/index.html")
}

func TestViewNoSlash(t *testing.T) {
	c := test.Client(t)
	r := c.Get("/testing")
	r.Status(200)
	r.BodyChecksumMatch("testdata/assets/testing/view/testing/index.html")
}
