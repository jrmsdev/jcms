// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jcms

import (
	"github.com/jrmsdev/jcms/webapp/config"
)

type Webapp struct {
	Name string
}

func NewWebapp(cfg *config.Config) *Webapp {
	if cfg.Name == "" {
		panic("empty webpp name")
	}
	return &Webapp{
		Name: cfg.Name,
	}
}
