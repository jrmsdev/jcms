// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package template

import (
	"bytes"
	"testing"

	"github.com/jrmsdev/jcms/_t/check"
)

var tcfg = `{
	"default": "main.html",
	"templates": {
		"/": "index.html"
	}
}`

func init() {
	cfg = new(Config)
	cfgLoad(cfg, []byte(tcfg))
}

func TestConfig(t *testing.T) {
	//~ t.Log(cfg)
	if check.NotEqual(t, cfg.Default, "main.html", "default template") {
		t.Fail()
	}
	if check.NotEqual(t, cfg.Templates["/"], "index.html", "/ template") {
		t.Fail()
	}
}

type tpltest struct {
	path string
	src  string
	rst  string
}

var tt = []tpltest{
	{"/", "testing", "testing"},
}

func TestTemplate(t *testing.T) {
	for _, x := range tt {
		//~ t.Log(x)
		src := bytes.NewBufferString(x.src)
		dst := new(bytes.Buffer)
		err := Parse(dst, src, x.path)
		if err != nil {
			t.Log(err)
			t.FailNow()
		}
		if check.NotEqual(t, dst.String(), x.rst, x.path) {
			t.Fail()
		}
	}
}
