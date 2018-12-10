// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package schema

import (
	"github.com/jrmsdev/jcms/db"
	"github.com/jrmsdev/jcms/internal/log"
)

func Check() error {
	log.D("Check %s", db.Webapp())
	return nil
}
