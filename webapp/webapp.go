// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package webapp

import (
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
