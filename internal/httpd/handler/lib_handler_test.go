// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package handler_test

import (
	"testing"

	"github.com/jrmsdev/jcms/internal/_t/test"
)

func TestLibNotFound(t *testing.T) {
	c := test.Client(t)
	r := c.Get("/_/lib/notfound.js")
	r.Status(404)
}

func TestW3JS(t *testing.T) {
	c := test.Client(t)
	r := c.Get("/_/lib/w3.js")
	r.Status(200)
	//~ r.Body("testing")
}

func TestW3CSS(t *testing.T) {
	c := test.Client(t)
	r := c.Get("/_/lib/w3.css")
	r.Status(200)
	//~ r.Body("testing")
}
