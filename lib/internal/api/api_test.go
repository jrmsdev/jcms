// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package api

import (
	gohttp "net/http"
	"testing"

	"github.com/jrmsdev/jcms/_t/check"
	"github.com/jrmsdev/jcms/_t/http"
)

func TestAPI(t *testing.T) {
	r, w := http.GET("/_/jcms.json")
	s := newServer()
	s.ServeHTTP(w, r)
	res := w.Result()
	if check.NotEqual(t, res.StatusCode, gohttp.StatusOK, "response status") {
		t.Fail()
	}
}

func TestNotFound(t *testing.T) {
	r, w := http.GET("/_/noapi")
	s := newServer()
	s.ServeHTTP(w, r)
	res := w.Result()
	if check.NotEqual(t, res.StatusCode, gohttp.StatusNotFound, "response status") {
		t.Fail()
	}
}
