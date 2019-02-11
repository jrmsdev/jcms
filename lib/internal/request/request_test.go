// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package request

import (
	"path"
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
	"/t6.html":       {path: "/t6.html"},
	"/t7/t.json":     {path: "/t7/t.json"},
	"/t8/a/b/c":      {path: "/t8/a/b/c"},
}

func TestRequest(t *testing.T) {
	for uri, x := range rt {
		//~ t.Log(uri, x)
		req := New(http.Request(uri))
		p := req.Path()
		if check.NotEqual(t, p, x.path, uri+" request path") {
			t.Fail()
		}
		fn := p
		if path.Ext(p) == "" {
			fn = path.Join(p, "index.html")
		}
		if check.NotEqual(t, req.Filename(), fn, uri+" request filename") {
			t.Fail()
		}
	}
}
