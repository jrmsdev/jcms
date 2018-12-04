// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package handler

import (
	"encoding/base64"
	"mime"
	"net/http"
	"path"
	"strings"

	"github.com/jrmsdev/jcms/assets"
	"github.com/jrmsdev/jcms/internal/log"
	"github.com/jrmsdev/jcms/webapp/config"

	"github.com/gorilla/mux"
)

func Setup(r *mux.Router, cfg *config.Config) {
	log.D("Setup")
	handlersSetup(r, cfg)
	if cfg.StaticEnable {
		setupStatic(r)
	}
	setupLib(r)
	setupView(r)
}

func handlersSetup(r *mux.Router, cfg *config.Config) {
	for n, f := range cfg.HandlerSetup {
		log.D("handlerSetup %s", n)
		f(r)
	}
}

// struct to serve static files

type fileServer struct {
	typ string
}

func newFileServer(typ string) *fileServer {
	return &fileServer{typ}
}

func (s *fileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		body []byte
		err  error
	)
	fp := path.Join(s.typ, r.URL.Path)
	log.D("ServeHTTP %s", fp)
	if s.typ == "view" {
		if path.Ext(fp) != ".html" {
			fp = path.Join(fp, "index.html")
		}
	} else {
		if strings.HasSuffix(fp, ".html") {
			log.D("denied .html access")
			http.Error(w, fp+": invalid request",
				http.StatusBadRequest)
			return
		}
	}
	if s.typ == "_lib" {
		encBody, found := libFiles[fp]
		if found {
			body, err = base64.StdEncoding.DecodeString(encBody)
			if err != nil {
				log.E("lib file serve %s: %s", fp, err)
				http.Error(w, fp+": "+err.Error(),
					http.StatusInternalServerError)
				return
			}
		} else {
			log.E("%s: not found", fp)
			http.Error(w, fp+": not found", http.StatusNotFound)
			return
		}
	} else {
		body, err = assets.ReadFile(fp)
		if err != nil {
			log.E("file serve %s: %s", fp, err)
			http.Error(w, fp+": not found", http.StatusNotFound)
			return
		}
	}
	s.setHeaders(w, fp)
	if n, err := w.Write(body); err != nil {
		log.E("file serve write %s: %s", fp, err)
	} else {
		log.Printf("sent: %s %d bytes", fp, n)
	}
}

func (s *fileServer) setHeaders(w http.ResponseWriter, fp string) {
	log.D("file server setHeaders %s", fp)
	h := w.Header()
	typ := mime.TypeByExtension(path.Ext(fp))
	if typ == "" {
		h.Set("Content-Type", "application/octet-stream")
	} else {
		h.Set("Content-Type", typ)
	}
}
