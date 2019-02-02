// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package handler

import (
	"path/filepath"
	"testing"

	"github.com/jrmsdev/jcms/_t/check"
	"github.com/jrmsdev/jcms/_t/http"
)

func TestHeaders(t *testing.T) {
	fs := &fileServer{}
	w := http.Writer()
	for _, x := range ht {
		fs.setHeaders(w, x.path)
		if check.NotEqual(t, w.Header().Get(x.key), x.expect, x.path+" "+x.key) {
			t.Fail()
		}
	}
}

func TestFileServer(t *testing.T) {
	fs := newFileServer(filepath.FromSlash("./testdata"))
	if check.NotEqual(t, filepath.ToSlash(fs.dir), "./testdata", "file server dir") {
		t.FailNow()
	}
	for _, x := range st {
		r, w := http.GET(x.path)
		fs.ServeHTTP(w, r)
		res := w.Result()
		if check.NotEqual(t, res.StatusCode, x.status, "response status") {
			t.FailNow()
		}
	}
}
