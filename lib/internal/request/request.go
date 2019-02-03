// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package request

import (
	"net/http"
	"path"
)

type Request struct {
	*http.Request
}

func New(r *http.Request) *Request {
	r.URL.Path = path.Clean(r.URL.Path)
	return &Request{r}
}

func (r *Request) Path() string {
	if r.URL.Path == "" {
		r.URL.Path = "index.html"
	}
	return r.URL.Path
}
