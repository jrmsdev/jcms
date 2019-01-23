// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jcms

import (
	"encoding/json"
	"net/http"
	"path"

	"github.com/jrmsdev/jcms"
	"github.com/jrmsdev/jcms/lib/internal/mime"
	"github.com/jrmsdev/jcms/lib/log"
)

type response struct {
	Version string `json:"jcms.version"`
}

func newResponse() *response {
	return &response{
		Version: jcms.Version(),
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	log.D("Handler %s", r.URL.Path)
	resp := newResponse()
	blob, err := json.MarshalIndent(&resp, "", "  ")
	if err != nil {
		log.E("%s", err)
		http.Error(w, "json error", http.StatusInternalServerError)
		return
	}
	setHeaders(w, r.URL.Path)
	if n, err := w.Write(blob); err != nil {
		log.E("%s", err)
	} else {
		log.Printf("sent: %s %d bytes", r.URL.Path, n)
	}
}

func setHeaders(w http.ResponseWriter, rp string) {
	log.D("setHeaders %s", rp)
	w.Header().Set("Content-Type", mime.TypeByExtension(path.Ext(rp)))
}
