// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package request

import (
	"testing"

	"github.com/jrmsdev/jcms/_t/check"
	"github.com/jrmsdev/jcms/_t/http"
)

type reqt struct {
	path string
}

var rt = map[string]reqt{
	"":               {path: "/"},
	".":              {path: "/"},
	"..":             {path: "/"},
	"/t0":            {path: "/t0"},
	"//t1":           {path: "/t1"},
	"./t2":           {path: "/t2"},
	"../t3":          {path: "/t3"},
	"/t4":            {path: "/t4"},
	"/t/../../../t5": {path: "/t5"},
}

func TestRequest(t *testing.T) {
	for uri, x := range rt {
		//~ t.Log(uri, x)
		req := New(http.Request(uri))
		if x.path != "" {
			if check.NotEqual(t, req.Path(), x.path, "request path") {
				t.Fail()
			}
		}
	}
}
