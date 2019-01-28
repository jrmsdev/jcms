// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package handler

import (
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/gorilla/mux"

	"github.com/jrmsdev/jcms/lib/internal/flags"
	"github.com/jrmsdev/jcms/lib/internal/mime"
	"github.com/jrmsdev/jcms/lib/log"
)

var admin bool = false
var htmldir string = "./webapp/devel/html"

func setupFileServer(r *mux.Router) {
	if admin {
		return
	}
	log.D("setup file server")
	dir := filepath.Join(flags.Assetsdir, flags.Webapp, "html")
	s := newFileServer(dir)
	s.defname = "index.html"
	r.PathPrefix("/").Handler(http.StripPrefix("/", s))
}

func develFileServer(r *mux.Router) {
	log.D("setup devel file server")
	r.PathPrefix("/_lib/").Handler(http.StripPrefix("/",
		newFileServer("./webapp")))
	r.PathPrefix("/_inc/").Handler(http.StripPrefix("/",
		newFileServer("./webapp/html")))
	s := newFileServer(htmldir)
	s.defname = "index.html"
	r.PathPrefix("/").Handler(http.StripPrefix("/", s))
}

type fileServer struct {
	dir     string
	defname string
}

func newFileServer(dir string) *fileServer {
	p, err := filepath.Abs(filepath.FromSlash(dir))
	if err != nil {
		log.Panic("%s", err)
	}
	log.D("new file server %s", p)
	return &fileServer{dir: p}
}

func (s *fileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rp := r.URL.Path
	if rp == "" && s.defname != "" {
		rp = path.Base(s.defname)
	}
	fp := filepath.Join(s.dir, filepath.FromSlash(rp))
	log.D("serve http '%s'", rp)
	if s.notFound(fp) {
		log.Printf("%s file not found", rp)
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	fh, err := s.open(fp)
	if err != nil {
		log.E("%s", err)
		http.Error(w, "open error", http.StatusInternalServerError)
		return
	}
	s.setHeaders(w, fp)
	if n, err := io.Copy(w, fh); err != nil {
		log.E("file serve write %s: %s", fp, err)
	} else {
		log.Printf("sent: %s %d bytes", rp, n)
	}
	if err := fh.Close(); err != nil {
		log.E("%s", err)
	}
}

func (s *fileServer) setHeaders(w http.ResponseWriter, fp string) {
	log.D("setHeaders %s", fp)
	w.Header().Set("Content-Type", mime.TypeByExtension(path.Ext(fp)))
}

func (s *fileServer) notFound(fp string) bool {
	fi, err := os.Stat(fp)
	if err != nil {
		log.E("%s", err)
		return true
	}
	if fi.IsDir() {
		log.E("%s is a directory", fp)
		return true
	}
	return false
}

func (s *fileServer) open(fp string) (io.ReadCloser, error) {
	return os.Open(fp)
}
