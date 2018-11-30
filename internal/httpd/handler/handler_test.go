// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package handler_test

import (
	"testing"

	//~ "github.com/jrmsdev/jcms/internal/_t/check"
	"github.com/jrmsdev/jcms/internal/_t/test"
)

func TestMain(m *testing.M) {
	test.Main(m, "")
}

func TestStatic(t *testing.T) {
	c := test.Client()
	r, err := c.Get("/static/test.txt")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(r)
}
