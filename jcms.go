// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jcms

import (
	"github.com/jrmsdev/jcms/internal/log"
	"github.com/jrmsdev/jcms/webapp/config"
)

func Main() {
	w := NewWebapp(&config.Config{})
	Start(w)
	defer Stop(w)
	Serve(w)
}

func Start(w *Webapp) string {
	log.Init(w.Log())
	log.D("Start")
	w.Start()
	return "127.0.0.1:6080"
}

func Serve(w *Webapp) {
	log.D("Serve")
}

func Stop(w *Webapp) {
	log.D("Stop")
	w.Stop()
}
