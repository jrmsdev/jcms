// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// +build jcmsadmin

package handler

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jrmsdev/jcms/lib/log"
)

func init () {
	adminSetup = func(r *mux.Router) {
		log.D("setup admin file server")
		htmldir = "./webapp/admin"
		r.PathPrefix("/inc/").Handler(http.StripPrefix("/inc/",
			newFileServer("./webapp/html/inc")))
	}
}
