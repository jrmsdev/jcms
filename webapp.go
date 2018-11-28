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
	setDefaults(cfg)
	return &Webapp{cfg}
}

func setDefaults(cfg *config.Config) {
	if cfg.Log == "" {
		cfg.Log = "default"
	}
}

func (w *Webapp) Name() string {
	return w.cfg.Name
}

func (w *Webapp) Log() string {
	return w.cfg.Log
}
