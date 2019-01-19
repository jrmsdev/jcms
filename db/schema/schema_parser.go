// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package schema

import (
	"errors"

	"github.com/jrmsdev/jcms/assets"
	"github.com/jrmsdev/jcms/internal/log"
)

func parse(s *Schema) error {
	log.D("parse %s", s)
	if _, err := assets.ReadFile("db.json"); err != nil {
		return errors.New(sprintf("parse %s: %s", s, err))
	}
	return nil
}
