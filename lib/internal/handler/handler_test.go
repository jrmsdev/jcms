// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package handler

import (
	"io/ioutil"
	gohttp "net/http"
	"testing"

	"github.com/jrmsdev/jcms/_t/check"
	"github.com/jrmsdev/jcms/_t/http"
)

type headerTest struct {
	path   string
	key    string
	expect string
}

var ht = []headerTest{
	{"/", "content-type", "application/octet-stream"},
	{"/t.html", "content-type", "text/html; charset=utf-8"},
	{"/t.json", "content-type", "application/json"},
	{"/t.js", "content-type", "application/javascript; charset=utf-8"},
	{"/t.css", "content-type", "text/css; charset=utf-8"},
	{"/t.ico", "content-type", "image/vnd.microsoft.icon"},
}

type serverTest struct {
	path   string
	status int
	body   string
}

var st = []serverTest{
	{"/", gohttp.StatusNotFound, ""},
	{"/test.txt", gohttp.StatusOK, "testing\n"},
	{"/nofile.txt", gohttp.StatusNotFound, ""},
}

func testServer(t *testing.T, s gohttp.Handler) {
	t.Helper()
	for _, x := range st {
		r, w := http.GET(x.path)
		s.ServeHTTP(w, r)
		res := w.Result()
		// status
		if check.NotEqual(t, res.StatusCode, x.status, x.path+" response status") {
			t.FailNow()
		}
		// body
		if x.body != "" {
			if blob, err := ioutil.ReadAll(res.Body); err != nil {
				t.Log(err)
				t.FailNow()
			} else {
				if check.NotEqual(t, string(blob), x.body, "response body") {
					t.Fail()
				}
			}
		}
	}
}
