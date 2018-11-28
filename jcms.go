// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jcms

import (
	"github.com/jrmsdev/jcms/webapp/config"
)

func Main(cfg *config.Config) {
	w := NewWebapp(cfg)
	defer Stop(w)
	Start(w)
}

func Start(w *Webapp) string {
	return "127.0.0.1:6080"
}

func Stop(w *Webapp) {
}
