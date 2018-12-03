// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package handler

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/jrmsdev/jcms/assets"
	"github.com/jrmsdev/jcms/internal/log"

	"github.com/gorilla/mux"
)

func setupStatic(r *mux.Router) {
	log.D("setupStatic")
	r.PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", http.FileServer(staticFS{}))).
		Name("static")
}

type staticFS struct{}

func (fs staticFS) Open(name string) (http.File, error) {
	fn := filepath.Join("static", name)
	log.D("Open: %s", fn)
	if strings.HasSuffix(fn, ".html") {
		log.D("denied .html access")
		return nil, &os.PathError{
			Op:   "open",
			Path: fn,
			Err:  os.ErrNotExist,
		}
	}
	fh, err := assets.Open(fn)
	if err != nil {
		log.E(err.Error())
		return nil, err
	}
	return newFile(fn, fh), nil
}

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
