// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"net/http"
	"net/http/pprof"

	"github.com/jrmsdev/jcms/internal/log"

	"github.com/gorilla/mux"
)

func pprofSetup(r *mux.Router) {
	log.D("pprofSetup")
	r.PathPrefix("/debug/pprof/").Handler(&pprofServer{}).Name("_pprof")
}

type pprofServer struct{}

func (s *pprofServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pprof.Index(w, r)
}
