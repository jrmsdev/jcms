// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package handler

import (
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"

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
	typ     string
	basedir string
}

func newFileServer(typ, basedir string) *fileServer {
	return &fileServer{typ, basedir}
}

func (s *fileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rp := r.URL.String()
	fp := filepath.FromSlash(path.Join(s.basedir, rp))
	if s.typ == "html" {
		if rp != "index.html" {
			fp = filepath.Join(fp, "index.html")
		}
	}
	log.D("ServeHTTP '%s'", rp)
	log.D("filepath '%s'", fp)
	blob, err := assets.ReadFile(fp)
	if err != nil {
		log.E("serve file %s: %s", fp, err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if n, err := io.WriteString(w, string(blob)); err != nil {
		log.E("serve file write %s: %s", fp, err)
	} else {
		log.Printf("sent: %s %d bytes", fp, n)
	}
}
