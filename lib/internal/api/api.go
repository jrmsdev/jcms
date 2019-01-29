// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jrmsdev/jcms/lib/internal/api/jcms"
	"github.com/jrmsdev/jcms/lib/internal/error/handler"
	"github.com/jrmsdev/jcms/lib/log"
)

var rreg = map[string]http.HandlerFunc{
	"_/jcms.json": jcms.Handler,
}

type apisvr struct{}

func Setup(r *mux.Router) {
	log.D("setup")
	r.PathPrefix("/_/").Handler(http.StripPrefix("/", &apisvr{}))
}

func (s *apisvr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rp := r.URL.Path
	log.D("serve '%s'", rp)
	f, ok := rreg[rp]
	if !ok {
		log.E("api '%s' not found", rp)
		handler.Error(w, "not found", http.StatusNotFound)
		return
	}
	f(w, r)
}
