// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package response

import (
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
	{"/",           "content-type", "application/octet-stream"},
	{"/index.html", "content-type", "text/html; charset=utf-8"},
	{"/jcms.json",  "content-type", "application/json"},
}

func TestHeaders(t *testing.T) {
	w := http.Writer()
	for _, x := range ht {
		setHeaders(w, x.path)
		if check.NotEqual(t, w.Header().Get(x.key), x.expect, x.path + " " + x.key) {
			t.Fail()
		}
	}
}
