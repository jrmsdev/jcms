// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package handler_test

import (
	"testing"

	"github.com/jrmsdev/jcms/internal/_t/test"
)

func TestStatic(t *testing.T) {
	c := test.Client(t)
	r := c.Get("/static/test.txt")
	r.Check(200, "text/plain")
	r.StatusInfo("200 OK")
	r.Body("testing")
}

func TestStaticNotFound(t *testing.T) {
	c := test.Client(t)
	r := c.Get("/static/notfound.txt")
	r.Check(404, "text/plain")
	r.StatusInfo("404 Not Found")
	r.Body("static/notfound.txt: not found")
}

func TestStaticGetDir(t *testing.T) {
	c := test.Client(t)
	r := c.Get("/static/testdir/")
	r.Check(404, "text/plain")
	r.Body("static/testdir: not found")
}
