// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package schema

import (
	"fmt"

	"github.com/jrmsdev/jcms/db"
	"github.com/jrmsdev/jcms/internal/log"
)

var (
	dbs *Schema
)
var sprintf = fmt.Sprintf

type Schema struct {
	name string
}

func newSchema(n string) *Schema {
	return &Schema{n}
}

func (s *Schema) String() string {
	return s.name
}

func Setup(wapp string) {
	log.D("Setup %s %s", wapp)
	if dbs != nil {
		log.Panic("db schema setup already done: %s", dbs)
	}
	dbs = newSchema(wapp)
	if err := parse(dbs); err != nil {
		log.Panic("schema setup: %s", err.Error())
	}
	log.D("parse dbs %s done", dbs)
}

func Check() error {
	log.D("Check %s", db.Webapp())
	return nil
}
