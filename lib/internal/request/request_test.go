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
	"": reqt{path: "/"},
	".": reqt{path: "/"},
	"..": reqt{path: "/"},
	"/t0": reqt{path: "/t0"},
	"//t1": reqt{path: "/t1"},
	"./t2": reqt{path: "/t2"},
	"../t3": reqt{path: "/t3"},
	"/t4": reqt{path: "/t4"},
	"/t/../../../t5": reqt{path: "/t5"},
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
