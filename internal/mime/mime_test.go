// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package mime

import (
	"testing"

	"github.com/jrmsdev/jcms/internal/_t/check"
)

func TestExt(t *testing.T) {
	x := TypeByExtension(".txt")
	if check.NotEqual(t, x, "text/plain; charset=utf-8", ".txt type") {
		t.FailNow()
	}
}

func TestDefault(t *testing.T) {
	x := TypeByExtension(".unknown")
	if check.NotEqual(t, x, "application/octet-stream", "default type") {
		t.FailNow()
	}
}

func TestJSCharset(t *testing.T) {
	x := TypeByExtension(".js")
	if check.NotEqual(t, x, "application/javascript; charset=utf-8",
		".js type/charset") {
		t.FailNow()
	}
}
