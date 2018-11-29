// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package handler

import (
	"net/http"
	"path/filepath"
	"errors"
	"os"

	"github.com/jrmsdev/jcms/internal/log"
	xassets "github.com/jrmsdev/jcms/assets"
	"github.com/jrmsdev/jcms/internal/assets"
	"github.com/jrmsdev/jcms/webapp/config"

	"github.com/gorilla/mux"
)

func setupStatic(r *mux.Router, cfg *config.Config) {
	log.D("setupStatic")
	r.PathPrefix(cfg.StaticURL).Handler(http.StripPrefix(cfg.StaticURL,
		     http.FileServer(staticFS{"static"})))
}

type staticFS struct {
	name string
}

func (fs staticFS) Open(name string) (http.File, error) {
	fn := filepath.Join(fs.name, name)
	fh, err := assets.Open(fn)
	if err != nil {
		log.E("Open %s", err)
		return nil, err
	}
	return newFile(fn, fh), nil
}

type staticFile struct {
	xassets.File
	name string
}

func newFile(fn string, fh xassets.File) http.File {
	log.D("newFile %s", fn)
	return &staticFile{fh, fn}
}

func (f *staticFile) Readdir(count int) ([]os.FileInfo, error) {
	log.D("Readdir %s", f.name)
	return nil, errors.New("dir not found")
}

func (f *staticFile) Stat() (os.FileInfo, error) {
	log.D("Stat %s", f.name)
	return assets.Stat(f.name)
}
