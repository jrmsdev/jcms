// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jcms

import (
	"github.com/jrmsdev/jcms/webapp/config"
)

type Webapp struct {
	cfg *config.Config
}

func NewWebapp(cfg *config.Config) *Webapp {
	if cfg.Name == "" {
		panic("empty webpp name")
	}
	return &Webapp{cfg}
}

func (w *Webapp) Name() string {
	return w.cfg.Name
}
