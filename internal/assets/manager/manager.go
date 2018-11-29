// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package manager

import (
	"io/ioutil"
	"os"

	"github.com/jrmsdev/jcms/assets"
)

type astman struct {
}

func New() *astman {
	return &astman{}
}

func (m *astman) Open(filename string) (assets.File, error) {
	return os.Open(filename)
}

func (m *astman) ReadFile(name string) ([]byte, error) {
	return ioutil.ReadFile(name)
}

func (m *astman) Stat(filename string) (os.FileInfo, error) {
	return os.Stat(filename)
}
