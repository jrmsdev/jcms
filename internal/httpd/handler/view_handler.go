// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package handler

import (
	"io"
	"net/http"
	"path"
	"path/filepath"

	"github.com/jrmsdev/jcms/assets"
	"github.com/jrmsdev/jcms/internal/log"

	"github.com/gorilla/mux"
)

func setupView(r *mux.Router) {
	log.D("setupView")
	if r.Get("index") == nil {
		r.PathPrefix("/").Handler(http.StripPrefix("/",
			&viewServer{})).Name("index")
	}
}

type viewServer struct{}

func (s *viewServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rp := r.URL.String()
	fp := filepath.FromSlash(path.Join("view", rp))
	if rp != "index.html" {
		fp = filepath.Join(fp, "index.html")
	}
	log.D("ServeHTTP %s", rp)
	log.D("filepath %s", fp)
	blob, err := assets.ReadFile(fp)
	if err != nil {
		log.E("view handler %s: %s", fp, err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if n, err := io.WriteString(w, string(blob)); err != nil {
		log.E("view handler write %s: %s", fp, err)
	} else {
		log.Printf("sent: %s %d bytes", fp, n)
	}
}
