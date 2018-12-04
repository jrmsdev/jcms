// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package handler_test

import (
	"testing"

	"github.com/jrmsdev/jcms/internal/_t/check"
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
	if check.NotFileChecksum(t, r.ReadBody(),
		"testdata/assets/testing/view/index.html") {
		t.FailNow()
	}
}

func TestViewIndexSlash(t *testing.T) {
	c := test.Client(t)
	r := c.Get("/")
	r.Status(200)
	if check.NotFileChecksum(t, r.ReadBody(),
		"testdata/assets/testing/view/index.html") {
		t.FailNow()
	}
}

func TestViewIndexEmpty(t *testing.T) {
	c := test.Client(t)
	r := c.Get("")
	r.Status(200)
	if check.NotFileChecksum(t, r.ReadBody(),
		"testdata/assets/testing/view/index.html") {
		t.FailNow()
	}
}
