// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jcms

import (
	"testing"

	"github.com/jrmsdev/jcms/internal/_t/check"
)

func TestNewWebapp(t *testing.T) {
	wapp := NewWebapp("testing")
	if check.NotEqual(t, wapp.Name, "testing", "webapp Name") {
		t.FailNow()
	}
}
