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
	config.SetDefaults(cfg)
	return &Webapp{cfg}
}

func (w *Webapp) Name() string {
	return w.cfg.Name
}

func (w *Webapp) Log() string {
	return w.cfg.Log
}
