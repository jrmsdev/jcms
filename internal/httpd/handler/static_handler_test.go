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
	r.Status(200)
	r.StatusInfo("200 OK")
	r.Body("testing")
}

func TestStaticNotFound(t *testing.T) {
	c := test.Client(t)
	r := c.Get("/static/notfound.txt")
	r.Status(404)
	r.StatusInfo("404 Not Found")
	r.Body("/static/notfound.txt: not found")
}

func TestStaticGetDir(t *testing.T) {
	c := test.Client(t)
	r := c.Get("/static/testdir/")
	r.Status(404)
	r.StatusInfo("404 Not Found")
	r.Body("/static/testdir: not found")
}
