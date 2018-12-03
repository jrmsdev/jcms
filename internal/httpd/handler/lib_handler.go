// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package handler

import (
	"io"
	"net/http"

	"github.com/jrmsdev/jcms/internal/log"

	"github.com/gorilla/mux"
)

func setupLib(r *mux.Router) {
	log.D("setupLib")
	r.PathPrefix("/_/lib/").Handler(http.StripPrefix("/_/", &libServer{}))
}

type libServer struct {
}

func (s *libServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fn := r.URL.String()
	log.D("ServeHTTP %s", fn)
	body, found := libFiles[fn]
	if found {
		io.WriteString(w, body)
	} else {
		http.Error(w, "file not found", http.StatusNotFound)
	}
}

var libFiles = map[string]string{
	"lib/w3.js": "testing",
}
