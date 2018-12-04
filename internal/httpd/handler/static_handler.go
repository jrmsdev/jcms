// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package handler

import (
	"net/http"

	"github.com/jrmsdev/jcms/internal/log"

	"github.com/gorilla/mux"
)

func setupStatic(r *mux.Router) {
	log.D("setupStatic")
	if r.Get("static") == nil {
		r.PathPrefix("/static/").Handler(http.StripPrefix("/static",
			newFileServer("static"))).Name("static")
	}
}
