// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package schema

import (
	"errors"

	"github.com/jrmsdev/jcms/db"
	"github.com/jrmsdev/jcms/internal/log"
)

func Check() error {
	log.D("Check %s", db.Webapp())
	if dbs == nil {
		return errors.New("nil schema")
	}
	log.D("dbs %#v", dbs)
	return nil
}
