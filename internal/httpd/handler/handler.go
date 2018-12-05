// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package handler

import (
	"encoding/base64"
	"net/http"
	"path"
	"strings"

	"github.com/jrmsdev/jcms/assets"
	"github.com/jrmsdev/jcms/internal/log"
	"github.com/jrmsdev/jcms/internal/mime"
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
	rp := path.Join("/", fp)
	// pre checks
	if s.typ == "view" {
		rp = path.Join("/", r.URL.Path)
		if x := path.Ext(fp); x == "" {
			fp = path.Join(fp, "index.html")
		} else if x != ".html" {
			log.Printf("view redirect static: %s", fp)
			http.Redirect(w, r, path.Join("/", "static", rp),
				http.StatusMovedPermanently)
			return
		}
	} else {
		if strings.HasSuffix(fp, ".html") {
			log.D("denied .html access")
			http.Error(w, rp+": invalid request",
				http.StatusBadRequest)
			return
		}
	}
	// get file content (body)
	if s.typ == "_lib" {
		// _lib files
		encBody, found := libFiles[fp]
		if found {
			body, err = base64.StdEncoding.DecodeString(encBody)
			if err != nil {
				log.E("lib file serve %s: %s", fp, err)
				http.Error(w, rp+": base64 error",
					http.StatusInternalServerError)
				return
			}
		} else {
			log.E("%s: not found", fp)
			http.Error(w, rp+": not found", http.StatusNotFound)
			return
		}
	} else {
		// asset files (static and view)
		body, err = assets.ReadFile(fp)
		if err != nil {
			log.E("file serve %s: %s", fp, err)
			http.Error(w, rp+": not found", http.StatusNotFound)
			return
		}
	}
	// send file content
	s.setHeaders(w, fp)
	if n, err := w.Write(body); err != nil {
		log.E("file serve write %s: %s", fp, err)
	} else {
		log.Printf("sent: %s %d bytes", fp, n)
	}
}

func (s *fileServer) setHeaders(w http.ResponseWriter, fp string) {
	log.D("file server setHeaders %s", fp)
	w.Header().Set("Content-Type", mime.TypeByExtension(path.Ext(fp)))
}
