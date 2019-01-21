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

	"github.com/jrmsdev/jcms/internal/log"
	"github.com/jrmsdev/jcms/internal/mime"
)

var zipmode bool = false
var zipfile string = ""

var b64 = base64.StdEncoding.DecodeString

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func setupZipServer(r *mux.Router) {
	log.D("setup zip server")
	r.PathPrefix("/").Handler(http.StripPrefix("/", newZipServer()))
}

type zipServer struct {
	rdr *zip.Reader
	files map[string]*zip.File
}

func newZipServer() *zipServer {
	zf := make(map[string]*zip.File)
	zdata := zipLoad()
	zr, err := zip.NewReader(zdata, int64(zdata.Len()))
	check(err)
	for _, f := range zr.File {
		log.D("zip server: %s", f.Name)
		zf[f.Name] = f
	}
	return &zipServer{
		rdr:     zr,
		files: zf,
	}
}

func (s *zipServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rp := r.URL.Path
	if rp == "" {
		rp = "index.html"
	}
	log.D("ServeHTTP %s", rp)
	if s.notFound(rp) {
		log.Printf("%s file not found", rp)
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	fh, err := s.open(rp)
	if err != nil {
		log.E("%s", err)
		http.Error(w, "open error", http.StatusInternalServerError)
		return
	}
	s.setHeaders(w, rp)
	if n, err := io.Copy(w, fh); err != nil {
		log.E("file serve write %s: %s", rp, err)
	} else {
		log.Printf("sent: %s %d bytes", rp, n)
	}
}

func (s *zipServer) setHeaders(w http.ResponseWriter, rp string) {
	log.D("setHeaders %s", rp)
	w.Header().Set("Content-Type", mime.TypeByExtension(path.Ext(rp)))
}

func (s *zipServer) notFound(rp string) bool {
	_, ok := s.files[rp]
	return !ok
}

func (s *zipServer) open(rp string) (io.ReadCloser, error) {
	f, ok := s.files[rp]
	if !ok {
		return nil, errors.New(sprintf("invalid zip file: %s", rp))
	}
	return f.Open()
}

func zipLoad() *bytes.Reader {
	blob, err := b64(zipfile)
	check(err)
	return bytes.NewReader(blob)
}
