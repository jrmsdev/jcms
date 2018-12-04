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
	if r.Get("_lib") == nil {
		r.PathPrefix("/_lib/").Handler(http.StripPrefix("/_",
			&libServer{})).Name("_lib")
	}
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
		if n, err := io.WriteString(w, string(body)); err != nil {
			log.E("lib handler write %s: %s", fn, err)
		} else {
			log.Printf("sent: %s %d bytes", fn, n)
		}
	} else {
		http.Error(w, "file not found", http.StatusNotFound)
	}
}
