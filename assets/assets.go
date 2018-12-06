// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package assets

import (
	"fmt"
	"io"
	"io/ioutil"
	"path"

	"github.com/jrmsdev/jcms/internal/errors"
	"github.com/jrmsdev/jcms/internal/log"
)

var sprintf = fmt.Sprintf

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

func ReadFile(relname string) ([]byte, errors.Error) {
	log.D("ReadFile: %s", relname)
	var (
		fh   File
		body []byte
		err  error
	)
	errp := path.Join("/", relname)
	fh, err = manager.Open(relname)
	if err != nil {
		log.E("assets file %s: not found", relname)
		return nil, errors.FileNotFound(errp)
	}
	defer fh.Close()
	body, err = ioutil.ReadAll(fh)
	if err != nil {
		log.D("assets %s: %s", relname, err)
		return nil, errors.PathError(errp, err)
	}
	return body, nil
}
