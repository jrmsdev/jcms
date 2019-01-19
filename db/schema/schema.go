// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package schema

import (
	"path/filepath"

	"github.com/jrmsdev/jcms/db"
	"github.com/jrmsdev/jcms/internal/log"
)

var (
	fn  string
	dbs *Schema
)

type Schema struct {
	name string
}

func newSchema(n string) *Schema {
	return &Schema{n}
}

func (s *Schema) String() string {
	return s.name
}

func Setup(wapp, basedir string) {
	log.D("Setup %s %s", wapp, basedir)
	if fn != "" {
		log.Panic("db schema setup already done: %s")
	}
	fn = filepath.Join(basedir, wapp, "db.json")
	log.D("schema definition %s", fn)
	dbs = newSchema(wapp)
	if err := parse(dbs, fn); err != nil {
		log.Panic("schema setup: %s", err)
	}
	dbs.name = wapp
	log.D("dbs %s", dbs)
}

func Check() error {
	log.D("Check %s", db.Webapp())
	return nil
}
