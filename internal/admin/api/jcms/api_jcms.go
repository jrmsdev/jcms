// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jcms

import (
	"encoding/json"
	"net/http"

	"github.com/jrmsdev/jcms"
	"github.com/jrmsdev/jcms/internal/log"
)

type response struct {
	Version string `json:"version"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	log.D("Handle %s", r.URL.Path)
	resp := &response{
		Version: jcms.Version(),
	}
	blob, err := json.MarshalIndent(&resp, "", "  ")
	if err != nil {
		log.E("%s", err)
		http.Error(w, "json error", http.StatusInternalServerError)
		return
	}
	if n, err := w.Write(blob); err != nil {
		log.E("%s", err)
	} else {
		log.Printf("sent: %s %d bytes", r.URL.Path, n)
	}
}
