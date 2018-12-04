// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

//go:generate python lib_generate.py

package handler

import (
	"encoding/base64"
	"io"
	"net/http"

	"github.com/jrmsdev/jcms/internal/log"

	"github.com/gorilla/mux"
)

func setupLib(r *mux.Router) {
	log.D("setupLib")
	r.PathPrefix("/_/lib/").Handler(http.StripPrefix("/_/", &libServer{}))
}

type libServer struct{}

func (s *libServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fn := r.URL.String()
	log.D("ServeHTTP %s", fn)
	encBody, found := libFiles[fn]
	if found {
		body, err := base64.StdEncoding.DecodeString(encBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		io.WriteString(w, string(body))
	} else {
		http.Error(w, "file not found", http.StatusNotFound)
	}
}
