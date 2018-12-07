// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package assets_test

import (
	"testing"

	"github.com/jrmsdev/jcms/internal/_t/test"
)

func TestMain(m *testing.M) {
	test.Main(m, "testing")
}

func TestReadFile(t *testing.T) {
	c := test.Client(t)
	r := c.Get("/static/test.txt")
	r.Status(200)
	r.Body("testing")
}

func TestFileNotFound(t *testing.T) {
	c := test.Client(t)
	r := c.Get("/static/notfound.txt")
	r.Status(404)
	r.Body("FileNotFound /static/notfound.txt")
}
