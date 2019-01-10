// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package storage

import (
	"io"
	"io/ioutil"
	"path"

	"github.com/jrmsdev/jcms/internal/errors"
	"github.com/jrmsdev/jcms/internal/log"
)

type File interface {
	io.ReadSeeker
	io.Writer
	io.Closer
}

type Driver interface {
	Open(relname, mode string) (File, error)
	Create(relname string) (File, error)
	Remove(relname string) error
	String() string
}

var drv Driver

func SetDriver(d Driver) {
	log.D("SetDriver %s", d)
	if drv != nil {
		log.Panic("storage driver already set!")
	}
	drv = d
}

func Open(relname, mode string) (File, errors.Error) {
	log.D("Open (%s) %s", mode, relname)
	var (
		fh  File
		err error
	)
	errp := path.Join("/", relname)
	fh, err = drv.Open(relname, mode)
	if err != nil {
		log.E("storage open %s: %s", relname, err)
		return nil, errors.FileNotFound(errp)
	}
	return fh, nil
}

func Create(relname string) (File, errors.Error) {
	log.D("Create %s", relname)
	var (
		fh  File
		err error
	)
	errp := path.Join("/", relname)
	fh, err = drv.Create(relname)
	if err != nil {
		return nil, errors.IOError(errp, "storage create")
	}
	return fh, nil
}

func Remove(relname string) errors.Error {
	log.D("Remove %s", relname)
	errp := path.Join("/", relname)
	err := drv.Remove(relname)
	if err != nil {
		return errors.IOError(errp, "storage remove")
	}
	return nil
}

func ReadFile(relname string) ([]byte, errors.Error) {
	log.D("ReadFile: %s", relname)
	var (
		fh   File
		body []byte
		err  error
	)
	errp := path.Join("/", relname)
	fh, err = drv.Open(relname, "r")
	if err != nil {
		log.E("storage file %s: not found", relname)
		return nil, errors.FileNotFound(errp)
	}
	defer fh.Close()
	body, err = ioutil.ReadAll(fh)
	if err != nil {
		log.D("storage %s: %s", relname, err)
		return nil, errors.IOError(errp, err.Error())
	}
	return body, nil
}
