// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package handler

import (
	"net/http"
	"path"
	"strings"

	"github.com/jrmsdev/jcms/assets"
	"github.com/jrmsdev/jcms/internal/errors"
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
		err errors.Error
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
			log.D("view redirect static: %s", fp)
			errors.Redirect(rp, r, path.Join("/", "static", rp)).
				WriteResponse(w)
			return
		}
	} else {
		if strings.HasSuffix(fp, ".html") {
			log.D("denied .html access")
			errors.InvalidRequest(rp).WriteResponse(w)
			return
		}
	}
	if s.typ == "_lib" {
		// _lib files
		body, err = libReadFile(fp)
	} else {
		// asset (static / view) files
		body, err = assets.ReadFile(fp)
	}
	if err != nil {
		err.WriteResponse(w)
		return
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
