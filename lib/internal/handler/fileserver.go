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

	"github.com/jrmsdev/jcms/lib/internal/asset"
	"github.com/jrmsdev/jcms/lib/internal/mime"
	"github.com/jrmsdev/jcms/lib/internal/request"
	"github.com/jrmsdev/jcms/lib/log"
)

var admin bool = false
var htmldir string = "./webapp/devel/html"

func setupAssetsServer(r *mux.Router) {
	if admin {
		return
	}
	log.D("setup assets server")
	s := newFileServer("html")
	s.assets = true
	r.PathPrefix("/").Handler(http.StripPrefix("/", s))
}

func develFileServer(r *mux.Router) {
	log.D("setup devel file server")
	r.PathPrefix("/_lib/").Handler(http.StripPrefix("/",
		newFileServer("./webapp")))
	r.PathPrefix("/_inc/").Handler(http.StripPrefix("/",
		newFileServer("./webapp/html")))
	s := newFileServer(htmldir)
	r.PathPrefix("/").Handler(http.StripPrefix("/", s))
}

type fileServer struct {
	assets  bool
	dir     string
	defname string
}

func newFileServer(dir string) *fileServer {
	p := filepath.FromSlash(dir)
	log.D("new file server %s", p)
	return &fileServer{dir: p}
}

func (s *fileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := request.New(r)
	rp := req.Path()
	log.D("serve '%s' (assets:%t)", rp, s.assets)
	fp := filepath.Join(s.dir, filepath.FromSlash(rp))
	if s.notFound(fp) {
		errhdlr(w, "not found", http.StatusNotFound)
		return
	}
	fh, err := s.open(fp)
	if err != nil {
		log.E("%s", err)
		errhdlr(w, "open error", http.StatusInternalServerError)
		return
	}
	s.setHeaders(w, fp)
	if n, err := io.Copy(w, fh); err != nil {
		log.E("file serve write %s: %s", fp, err)
	} else {
		log.Response(r, n)
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
	if s.assets {
		return !asset.Exists(fp)
	}
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
	log.D("open (asset %t) %s", s.assets, fp)
	if s.assets {
		return asset.Open(fp)
	}
	return os.Open(fp)
}
