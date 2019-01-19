// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package handler

import (
	"net/http"

	"github.com/jrmsdev/jcms/internal/errors"
	"github.com/jrmsdev/jcms/internal/log"
)

type errorServer struct {
	err errors.Error
}

func ServeError(err errors.Error) *errorServer {
	return &errorServer{err}
}

func (s *errorServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.D("ServeError %s", r.URL.Path)
	log.E("%s: %s", r.URL.Path, s.err.Error())
	s.err.WriteResponse(w)
}
