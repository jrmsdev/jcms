// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package manager

import (
	"errors"
	"os"
	"path"
	"path/filepath"

	"github.com/jrmsdev/jcms/assets"
	"github.com/jrmsdev/jcms/internal/log"
)

type astman struct {
	assetsdir string
}

func New(wapp, assetsdir string) *astman {
	p := path.Join(assetsdir, wapp)
	log.D("%s", p)
	return &astman{p}
}

func (m *astman) getFilename(relname string) string {
	return filepath.FromSlash(path.Join(m.assetsdir, relname))
}

func (m *astman) Open(relname string) (assets.File, error) {
	log.D("Open %s", relname)
	fn := m.getFilename(relname)
	fi, err := os.Stat(fn)
	if err != nil {
		return nil, err
	}
	if fi.IsDir() {
		log.E("%s is a directory", fn)
		return nil, errors.New("is dir")
	}
	return os.Open(fn)
}
