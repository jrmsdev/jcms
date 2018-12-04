// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package manager

import (
	"io/ioutil"
	"path"
	"path/filepath"

	"github.com/jrmsdev/jcms/internal/log"
)

type astman struct {
	basedir string
}

func New(wapp, basedir string) *astman {
	log.D("New: %s %s", wapp, basedir)
	return &astman{filepath.Join(basedir, wapp)}
}

func (m *astman) getFilename(relname string) string {
	return filepath.FromSlash(path.Join(m.basedir, relname))
}

func (m *astman) ReadFile(relname string) ([]byte, error) {
	log.D("ReadFile %s", relname)
	return ioutil.ReadFile(m.getFilename(relname))
}
