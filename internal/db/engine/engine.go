// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package engine

import (
	"net/url"
	"path/filepath"

	"github.com/jrmsdev/jcms/db"
	"github.com/jrmsdev/jcms/internal/db/engine/fsdb"
	"github.com/jrmsdev/jcms/internal/log"
)

func New(uri, wapp, datadir string) db.Engine {
	log.D("New %s %s", wapp, uri)
	x, err := url.Parse(uri)
	if err != nil {
		log.Panic("parse DatabaseURI: %s", err)
	}
	if x.Scheme == "fs" {
		dbdir := filepath.Join(datadir, wapp, "db")
		log.D("db dir %s", dbdir)
		return fsdb.New(wapp, dbdir, x.Path)
	} else {
		log.Panic("%s invalid database engine: %s", wapp, x.Scheme)
	}
	return nil
}
