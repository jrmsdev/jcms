// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package assets

import (
	"io"
	"io/ioutil"

	"github.com/jrmsdev/jcms/internal/log"
)

type File interface {
	io.ReadSeeker
	io.Closer
}

type Manager interface {
	Open(relname string) (File, error)
}

var manager Manager

func SetManager(m Manager) {
	log.D("SetManager")
	if manager != nil {
		log.Panic("assets manager already set!")
	}
	manager = m
}

func ReadFile(relname string) ([]byte, error) {
	log.D("ReadFile: %s", relname)
	fh, err := manager.Open(relname)
	if err != nil {
		return nil, err
	}
	defer fh.Close()
	return ioutil.ReadAll(fh)
}
