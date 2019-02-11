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
		fn = path.Join(p, "index.html")
	}
	return &Request{p, fn}
}

func (r *Request) Path() string {
	log.D("path %s", r.path)
	return r.path
}

func (r *Request) Filename() string {
	log.D("filename %s", r.filename)
	return r.filename
}
