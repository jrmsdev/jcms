// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package handler

import (
	"encoding/base64"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
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

// struct to serve static files common to static and view handlers

type staticFile struct {
	assets.File
	name string
}

func newFile(fn string, fh assets.File) http.File {
	log.D("newFile %s", fn)
	return &staticFile{fh, fn}
}

func (f *staticFile) Readdir(count int) ([]os.FileInfo, error) {
	log.D("Readdir %s", f.name)
	return []os.FileInfo{}, nil
}

func (f *staticFile) Stat() (os.FileInfo, error) {
	log.D("Stat %s", f.name)
	return assets.Stat(f.name)
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
	rp := r.URL.Path
	if rp == "" {
		rp = "/"
	}
	fp := path.Join("/", s.typ, rp)
	log.D("ServeHTTP %s", fp)
	if s.typ == "view" {
		if rp != "index.html" {
			fp = filepath.Join(fp, "index.html")
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
	if n, err := io.WriteString(w, string(body)); err != nil {
		log.E("file serve write %s: %s", fp, err)
	} else {
		log.Printf("sent: %s %d bytes", fp, n)
	}
}

func (s *fileServer) setHeaders(w http.ResponseWriter, fp string) {
	log.D("file server setHeaders %s", fp)
	h := w.Header()
	if strings.HasSuffix(fp, ".js") {
		h.Set("Content-Type", "application/x-javascript; charset=utf-8")
	} else if strings.HasSuffix(fp, ".css") {
		h.Set("Content-Type", "text/css; charset=utf-8")
	}
}
