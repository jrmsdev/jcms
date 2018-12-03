// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"net/http"

	"github.com/jrmsdev/jcms/internal/log"

	"github.com/gorilla/mux"
)

func libHandlerSetup(r *mux.Router) {
	log.D("libHandlerSetup")
	r.PathPrefix("/_/lib/").Handler(http.StripPrefix("/_/lib/",
		http.FileServer(http.Dir("internal/httpd/handler/lib"))))
}
