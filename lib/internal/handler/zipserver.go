// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package handler

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"errors"
	"io"
	"net/http"
	"path"

	"github.com/gorilla/mux"

	"github.com/jrmsdev/jcms/lib/internal/mime"
	"github.com/jrmsdev/jcms/lib/internal/request"
	"github.com/jrmsdev/jcms/lib/log"
)

var zipfile string = ""

var b64 = base64.StdEncoding.DecodeString

func echeck(err error) {
	if err != nil {
		panic(err)
	}
}

func setupZipServer(r *mux.Router) {
	log.D("setup zip server")
	zs := newZipServer()
	if admin {
		r.PathPrefix("/").Handler(http.StripPrefix("/", zs))
	} else {
		r.PathPrefix("/_lib/").Handler(http.StripPrefix("/", zs))
		r.PathPrefix("/_inc/").Handler(http.StripPrefix("/", zs))
	}
}

type zipServer struct {
	rdr   *zip.Reader
	files map[string]*zip.File
}

func newZipServer() *zipServer {
	zf := make(map[string]*zip.File)
	zdata := zipLoad()
	zr, err := zip.NewReader(zdata, int64(zdata.Len()))
	echeck(err)
	for _, f := range zr.File {
		log.D("zip file: %s", f.Name)
		zf[f.Name] = f
	}
	return &zipServer{
		rdr:   zr,
		files: zf,
	}
}

func (s *zipServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := request.New(r)
	rp := req.Path()
	log.D("serve %s", rp)
	if s.notFound(rp) {
		log.Printf("'%s' zip file not found", rp)
		errhdlr(w, "not found", http.StatusNotFound)
		return
	}
	fh, err := s.open(rp)
	if err != nil {
		errhdlr(w, "open error", http.StatusInternalServerError)
		return
	}
	s.setHeaders(w, rp)
	if n, err := io.Copy(w, fh); err != nil {
		log.E("zip file '%s' write: %s", rp, err)
	} else {
		log.Response(r, n)
	}
	if err := fh.Close(); err != nil {
		log.E("%s", err)
	}
}

func (s *zipServer) setHeaders(w http.ResponseWriter, rp string) {
	log.D("set headers %s", rp)
	w.Header().Set("Content-Type", mime.TypeByExtension(path.Ext(rp)))
}

func (s *zipServer) notFound(rp string) bool {
	log.D("zip find: %s", rp)
	_, ok := s.files[rp]
	return !ok
}

func (s *zipServer) open(rp string) (io.ReadCloser, error) {
	log.D("open '%s'", rp)
	f, ok := s.files[rp]
	if !ok {
		err := errors.New(sprintf("'%s' zip file open error", rp))
		log.E("%s", err)
		return nil, err
	}
	return f.Open()
}

func zipLoad() *bytes.Reader {
	blob, err := b64(zipfile)
	echeck(err)
	return bytes.NewReader(blob)
}
