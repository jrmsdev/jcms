// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package schema

import (
	"fmt"
	"encoding/json"

	"github.com/jrmsdev/jcms/db"
	"github.com/jrmsdev/jcms/internal/log"
)

var (
	dbs *Schema
)
var sprintf = fmt.Sprintf

type Schema struct {
	name string
	Data Data `json:"schema"`
}

func newSchema(n string) *Schema {
	return &Schema{name: n, Data: newData()}
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
		log.Panic("parse db schema: %s", err.Error())
	}
	log.D("parse dbs %s done", dbs)
}

func Check() error {
	log.D("Check %s", db.Webapp())
	if blob, err := json.MarshalIndent(dbs, "", "  "); err != nil {
		return err
	} else {
		log.D("JSON: %s", blob)
	}
	return nil
}
