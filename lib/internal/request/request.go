// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package request

import (
	"net/http"
)

type Request struct {
	*http.Request
}

func New(r *http.Request) *Request {
	return &Request{r}
}

func (r *Request) Path() string {
	if r.URL.Path == "" {
		r.URL.Path = "index.html"
	}
	return r.URL.Path
}
