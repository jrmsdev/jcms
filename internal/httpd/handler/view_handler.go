// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package handler

import (
	"net/http"

	"github.com/jrmsdev/jcms/internal/log"

	"github.com/gorilla/mux"
)

func setupView(r *mux.Router) {
	log.D("setupView")
	if r.Get("index") == nil {
		r.PathPrefix("/").Handler(http.StripPrefix("/",
			newFileServer("html", "view"))).Name("index")
	}
}
