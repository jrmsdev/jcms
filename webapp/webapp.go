// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package webapp

import (
	"github.com/jrmsdev/jcms"
	"github.com/jrmsdev/jcms/webapp/config"
)

type Webapp struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Webapp {
	return &Webapp{cfg}
}

func (w *Webapp) Name() string {
	return w.cfg.Name
}

func (w *Webapp) Start() string {
	return jcms.Start(w.cfg)
}

func (w *Webapp) Serve() {
	jcms.Serve()
}

func (w *Webapp) Stop() {
	jcms.Stop()
}
