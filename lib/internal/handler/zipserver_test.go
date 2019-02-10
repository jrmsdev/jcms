// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package handler

import (
	gohttp "net/http"
	"testing"

	"github.com/jrmsdev/jcms/_t/check"
	"github.com/jrmsdev/jcms/_t/http"
)

func TestZipHeaders(t *testing.T) {
	zs := &zipServer{}
	w := http.Writer()
	for _, x := range ht {
		zs.setHeaders(w, x.path)
		if check.NotEqual(t, w.Header().Get(x.key), x.expect, x.path+" "+x.key) {
			t.Fail()
		}
	}
}

func TestZipServer(t *testing.T) {
	zs := newZipServer()
	testServer(t, gohttp.StripPrefix("/", zs))
}
