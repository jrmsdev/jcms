// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jrmsdev/jcms/lib/internal/api/jcms"
	"github.com/jrmsdev/jcms/lib/internal/error/handler"
	"github.com/jrmsdev/jcms/lib/internal/request"
	"github.com/jrmsdev/jcms/lib/log"
)

var hreg = map[string]http.HandlerFunc{
	"_/jcms.json": jcms.Handler,
}

func Setup(r *mux.Router) {
	log.D("setup")
	r.PathPrefix("/_/").Handler(newServer())
}

type apisvr struct{}

func newServer() http.Handler {
	return http.StripPrefix("/", &apisvr{})
}

func (s *apisvr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := request.New(r)
	rp := req.Path()
	log.D("serve '%s'", rp)
	f, ok := hreg[rp]
	if !ok {
		log.E("api '%s' not found", rp)
		handler.Error(w, "not found", http.StatusNotFound)
		return
	}
	f(w, r)
}
