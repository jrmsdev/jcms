// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package fsdb

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/jrmsdev/jcms/internal/log"
)

type FSDB struct {
	wapp   string
	dbpath string
	lockfn string
	lockfh *os.File
}

func New(wapp, dbdir, dbname string) *FSDB {
	return &FSDB{
		wapp:   wapp,
		dbpath: filepath.Join(dbdir, dbname),
	}
}

func (e *FSDB) String() string {
	return "FSDB " + e.wapp
}

func (e *FSDB) Webapp() string {
	return e.wapp
}

func (e *FSDB) Connect() error {
	log.D("Connect %s", e.dbpath)
	var (
		err error
	)
	if e.lockfn != "" {
		log.E("already locked: %s", e.lockfn)
		return errors.New("fsdb already locked")
	}
	e.lockfn = filepath.Join(e.dbpath, "db.lock")
	e.lockfh, err = os.OpenFile(e.lockfn, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0750)
	if err != nil {
		log.E("%s", err)
		return errors.New("fsdb lock failed")
	}
	log.D("locked: %s", e.lockfn)
	return nil
}

func (e *FSDB) Disconnect() error {
	log.D("Disconnect %s", e.dbpath)
	if err := e.lockfh.Close(); err != nil {
		log.E("fsdb unlock: %s", err)
		return err
	}
	if err := os.Remove(e.lockfn); err != nil {
		log.E("fsdb remove lock file: %s", err)
		return err
	}
	log.D("unlocked: %s", e.lockfn)
	return nil
}
