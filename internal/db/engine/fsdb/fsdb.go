// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package fsdb

import (
	"path/filepath"

	"github.com/jrmsdev/jcms/internal/log"
)

type FSDB struct {
	wapp   string
	dbpath string
}

func New(wapp, dbdir, dbname string) *FSDB {
	return &FSDB{wapp, filepath.Join(dbdir, dbname)}
}

func (e *FSDB) String() string {
	return "FSDB " + e.wapp
}

func (e *FSDB) Webapp() string {
	return e.wapp
}

func (e *FSDB) Connect() error {
	log.D("Connect %s", e.dbpath)
	return nil
}

func (e *FSDB) Disconnect() error {
	log.D("Disconnect %s", e.dbpath)
	return nil
}
