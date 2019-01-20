// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package schema

import (
	"github.com/jrmsdev/jcms/internal/db/schema/parser"
	"github.com/jrmsdev/jcms/internal/log"
)

var (
	dbs *Schema
)

type Schema struct {
	name string
	data parser.Data
}

func newSchema(n string) *Schema {
	return &Schema{name: n}
}

func Setup(wapp string) {
	log.D("Setup %s", wapp)
	var err error
	if dbs != nil {
		log.Panic("db schema setup already done: %s", dbs)
	}
	dbs = newSchema(wapp)
	if dbs.data, err = parser.Parse(); err != nil {
		log.Panic("parse db schema: %s", err.Error())
	}
	log.D("parse dbs %s done", dbs)
}

func (s *Schema) String() string {
	return s.name
}
