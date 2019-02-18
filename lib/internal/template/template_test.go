// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package template

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/jrmsdev/jcms/_t/check"
	"github.com/jrmsdev/jcms/lib/internal/asset"
)

type tpltest struct {
	path string
	src  string
	rst  string
}

var tt = []tpltest{
	{"/", "testing", "testing"},
	//~ {"/notpl", "testing", "testing"},
	{"/test", `{{define "testdata"}}testing{{end}}`, "testing\n"},
}

func init() {
	asset.InitTest()
	Setup()
}

func TestConfig(t *testing.T) {
	//~ t.Log(cfg)
	if check.NotEqual(t, cfg.Default, "main", "default template") {
		t.Fail()
	}
	if check.NotEqual(t, cfg.Templates["/"], "index", "/ template") {
		t.Fail()
	}
	if check.NotEqual(t, cfg.Get("nopath"), "main", "get default template") {
		t.Fail()
	}
	if check.NotEqual(t, cfg.Get("/"), "index", "get / template") {
		t.Fail()
	}
}

func TestTemplate(t *testing.T) {
	for _, x := range tt {
		//~ t.Log(x)
		src := bytes.NewBufferString(x.src)
		dst := new(bytes.Buffer)
		tname := cfg.Get(x.path)
		if tname != "" {
			fn := filepath.Join("testdata", "wapp", "tpl", tname+".html")
			fh, err := os.Open(fn)
			if err != nil {
				t.Log(x.path, err)
				t.FailNow()
			}
			err = Parse(dst, src, fh)
			if err != nil {
				t.Log(err)
				t.FailNow()
			}
			if check.NotEqual(t, dst.String(), x.rst, x.path) {
				t.Fail()
			}
		}
	}
}
