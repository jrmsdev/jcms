// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package handler_test

import (
	"testing"

	"github.com/jrmsdev/jcms/internal/_t/test"
)

func TestLibNotFound(t *testing.T) {
	c := test.Client(t)
	r := c.Get("/_lib/notfound.js")
	r.Check(404, "text/plain")
	r.Body("/_lib/notfound.js: not found")
}

func TestLibW3JS(t *testing.T) {
	c := test.Client(t)
	r := c.Get("/_lib/w3.js")
	r.Check(200, "application/javascript")
	r.BodyChecksumMatch("lib/w3.js")
}

func TestLibW3CSS(t *testing.T) {
	c := test.Client(t)
	r := c.Get("/_lib/w3.css")
	r.Check(200, "text/css")
	r.BodyChecksumMatch("lib/w3.css")
}
