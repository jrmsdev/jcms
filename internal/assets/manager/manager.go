// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package manager

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/jrmsdev/jcms/assets"
	"github.com/jrmsdev/jcms/internal/log"
)

type astman struct {
	basedir string
}

func New(wapp, basedir string) *astman {
	log.D("New: %s %s", wapp, basedir)
	return &astman{filepath.Join(basedir, wapp)}
}

func (m *astman) Open(relname string) (assets.File, error) {
	return os.Open(filepath.Join(m.basedir, relname))
}

func (m *astman) ReadFile(relname string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join(m.basedir, relname))
}

func (m *astman) Stat(relname string) (os.FileInfo, error) {
	return os.Stat(filepath.Join(m.basedir, relname))
}
