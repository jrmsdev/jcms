// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package driver

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/jrmsdev/jcms/internal/log"
	"github.com/jrmsdev/jcms/storage"
)

var sprintf = fmt.Sprintf

type osdrv struct {
	datadir string
}

func New(name, datadir string) storage.Driver {
	return &osdrv{filepath.Join(datadir, name)}
}

func (d *osdrv) String() string {
	return "OS storage driver"
}

func (d *osdrv) getFilename(relname string) string {
	return filepath.FromSlash(path.Join(d.datadir, relname))
}

func (d *osdrv) Open(relname, mode string) (storage.File, error) {
	log.D("Open (%s) %s", mode, relname)
	f := os.O_RDONLY
	if mode == "rw" {
		f = os.O_RDWR
	} else if mode != "r" {
		return nil, errors.New(sprintf("invalid open access mode: %s", mode))
	}
	return os.OpenFile(d.getFilename(relname), f, 0600)
}

func (d *osdrv) Create(relname string) (storage.File, error) {
	log.D("Create %s", relname)
	f := os.O_WRONLY | os.O_CREATE | os.O_EXCL
	return os.OpenFile(d.getFilename(relname), f, 0600)
}

func (d *osdrv) Remove(relname string) error {
	log.D("Remove %s", relname)
	return os.Remove(d.getFilename(relname))
}
