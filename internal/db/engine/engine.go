// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package engine

import (
	"net/url"

	"github.com/jrmsdev/jcms/db"
	"github.com/jrmsdev/jcms/internal/log"
)

func New(uri, wapp, datadir string) db.Engine {
	log.D("New %s %s", wapp, uri)
	x, err := url.Parse(uri)
	if err != nil {
		log.Panic("parse DatabaseURI: %s", err)
	}
	if x.Scheme == "fs" {
		return newFSDB(wapp, datadir)
	} else {
		log.Panic("%s invalid database engine: %s", wapp, x.Scheme)
	}
	return nil
}
