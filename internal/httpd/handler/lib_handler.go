// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

//go:generate python lib_generate.py

package handler

import (
	"net/http"

	"github.com/jrmsdev/jcms/internal/log"

	"github.com/gorilla/mux"
)

func setupLib(r *mux.Router) {
	log.D("setupLib")
	if r.Get("_lib") == nil {
		r.PathPrefix("/_lib/").Handler(http.StripPrefix("/_lib",
			newFileServer("_lib"))).Name("_lib")
	}
}
