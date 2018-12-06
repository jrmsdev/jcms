// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"net/http"

	"github.com/jrmsdev/jcms/internal/httpd/handler"
	"github.com/jrmsdev/jcms/internal/log"

	"github.com/gorilla/mux"
)

func libHandlerSetup(r *mux.Router) {
	log.D("libHandlerSetup")
	r.PathPrefix("/_lib/").Handler(http.StripPrefix("/_lib/",
		handler.NewFileServer("_lib_devel"))).Name("_lib")
}
