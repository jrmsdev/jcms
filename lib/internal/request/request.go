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
	r.URL.Path = path.Clean(r.URL.Path)
	if r.URL.Path == "" {
		r.URL.Path = "index.html"
	} else if r.URL.Path == "." {
		r.URL.Path = "/"
	}
	log.D("new %s", r.URL.Path)
	return &Request{r}
}

func (r *Request) Path() string {
	return r.URL.Path
}
