// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jcms

import (
	"github.com/jrmsdev/jcms/internal/log"
	"github.com/jrmsdev/jcms/internal/webapp"
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

func (w *Webapp) Start() {
	log.D("Start")
	webapp.Start(w.cfg)
}

func (w *Webapp) Stop() {
	log.D("Stop")
	webapp.Stop(w.cfg)
}
