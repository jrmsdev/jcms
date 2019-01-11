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
)

type FileServer struct {
	typ string
}

func NewFileServer(typ string) *FileServer {
	return &FileServer{typ}
}

func (s *FileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		body []byte
		err  errors.Error
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
	} else if s.typ == "_lib_devel" {
		// _lib devel files
		fp = path.Join("./internal/httpd/handler/lib", r.URL.Path)
		rp = path.Join("/", "_lib", r.URL.Path)
		body, err = libDevelReadFile(fp)
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

func (s *FileServer) setHeaders(w http.ResponseWriter, fp string) {
	log.D("file server setHeaders %s", fp)
	w.Header().Set("Content-Type", mime.TypeByExtension(path.Ext(fp)))
}