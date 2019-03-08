// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package handler

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/gorilla/mux"

	"github.com/jrmsdev/jcms/lib/internal/asset"
	"github.com/jrmsdev/jcms/lib/internal/mime"
	"github.com/jrmsdev/jcms/lib/internal/request"
	"github.com/jrmsdev/jcms/lib/internal/template"
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
	rp := req.Filename()
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
	defer fh.Close()
	resp := new(bytes.Buffer)
	err = s.parseTpl(resp, fh, req.Path())
	if err != nil {
		errhdlr(w, "read/write error", http.StatusInternalServerError)
		return
	}
	s.setHeaders(w, fp)
	if n, err := io.Copy(w, resp); err != nil {
		log.E("write %s: %s", fp, err)
	} else {
		log.Response(req, n)
		resp.Reset()
	}
}

func (s *fileServer) setHeaders(w http.ResponseWriter, fp string) {
	log.D("setHeaders %s", fp)
	w.Header().Set("Content-Type", mime.TypeByExtension(path.Ext(fp)))
}

func (s *fileServer) notFound(fp string) bool {
	if s.assets {
		if !asset.Exists(fp) {
			log.E("%s asset not found", fp)
			return true
		}
		return false
	}
	fi, err := os.Stat(fp)
	if err != nil {
		log.D("%s", err)
		log.E("%s file not found", fp)
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

func (s *fileServer) parseTpl(resp *bytes.Buffer, src io.Reader, rp string) error {
	var (
		tpl io.ReadCloser
		err error
	)
	tname := template.Get(rp)
	if tname != "" {
		fn := filepath.Join("tpl", tname+".html")
		log.D("parse template %s", fn)
		tpl, err = s.open(fn)
		if err != nil {
			log.E("%s", err)
			return err
		}
		return template.Parse(resp, src, tpl)
	}
	log.D("copy")
	_, err = io.Copy(resp, src)
	return err
}
