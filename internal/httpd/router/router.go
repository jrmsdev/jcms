// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package router

import (
	"github.com/jrmsdev/jcms/internal/log"

	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	log.D("Init")
	r := mux.NewRouter()
	return r.Host("127.0.0.1").
		Methods("GET").
		Schemes("http").
		Subrouter().
		StrictSlash(true)
}
