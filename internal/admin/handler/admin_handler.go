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

	"github.com/jrmsdev/jcms/internal/log"
	"github.com/jrmsdev/jcms/internal/mime"
)

func Setup(r *mux.Router) {
	log.D("handler setup")
	r.PathPrefix("/_lib/").Handler(http.StripPrefix("/_lib/",
		newFileServer("./internal/httpd/handler/lib")))
	r.PathPrefix("/").Handler(http.StripPrefix("/",
		newFileServer("./internal/admin/html")))
}

type fileServer struct {
	dir string
}

func newFileServer(dir string) *fileServer {
	p, err := filepath.Abs(filepath.FromSlash(dir))
	if err != nil {
		log.Panic("%s", err)
	}
	log.D("new file server %s", p)
	return &fileServer{p}
}

func (s *fileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rp := r.URL.Path
	fp := filepath.Join(s.dir, filepath.FromSlash(rp))
	log.D("ServeHTTP %s", rp)
	fh, err := os.Open(fp)
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
}

func (s *fileServer) setHeaders(w http.ResponseWriter, fp string) {
	log.D("setHeaders %s", fp)
	w.Header().Set("Content-Type", mime.TypeByExtension(path.Ext(fp)))
}
