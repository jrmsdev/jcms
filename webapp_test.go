// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jcms

import (
	"testing"

	"github.com/jrmsdev/jcms/internal/_t/check"
	"github.com/jrmsdev/jcms/internal/_t/test"
)

func TestNewWebapp(t *testing.T) {
	wapp := NewWebapp(test.Config())
	if check.NotEqual(t, wapp.Name(), "testing", "webapp Name") {
		t.FailNow()
	}
}
