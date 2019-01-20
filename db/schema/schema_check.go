// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package schema

import (
	"container/heap"
	"errors"

	"github.com/jrmsdev/jcms/db"
	"github.com/jrmsdev/jcms/internal/db/schema/parser"
	"github.com/jrmsdev/jcms/internal/log"
)

func Check() error {
	log.D("Check %s", db.Webapp())
	if dbs == nil {
		return errors.New("nil schema")
	}
	for dbs.data.Len() > 0 {
		stmt := heap.Pop(&dbs.data).(*parser.Stmt)
		log.D("stmt: %s", stmt)
	}
	return nil
}
