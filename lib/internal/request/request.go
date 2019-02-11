// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package request

import (
	"net/http"
	"path"

	"github.com/jrmsdev/jcms/lib/log"
)

type Request struct {
	path     string
	filename string
}

func New(r *http.Request) *Request {
	log.D("new '%s'", r.URL.Path)
	p := path.Clean(r.URL.Path)
	if p == "." {
		p = "/"
	}
	fn := p
	if path.Ext(p) == "" {
		if p == "/" {
			fn = "index.html"
		} else {
			fn = path.Join(p, "index.html")
		}
	}
	log.D("path %s", p)
	log.D("filename %s", fn)
	return &Request{p, fn}
}

func (r *Request) Path() string {
	return r.path
}

func (r *Request) Filename() string {
	return r.filename
}
