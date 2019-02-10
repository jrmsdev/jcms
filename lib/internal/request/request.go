// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package request

import (
	"net/http"
	"path"

	"github.com/jrmsdev/jcms/lib/log"
)

type Request struct {
	*http.Request
}

func New(r *http.Request) *Request {
	log.D("new '%s'", r.URL.Path)
	r.URL.Path = path.Clean(r.URL.Path)
	if r.URL.Path == "." {
		r.URL.Path = "/"
	}
	return &Request{r}
}

func (r *Request) Path() string {
	log.D("path %s", r.URL.Path)
	return r.URL.Path
}
