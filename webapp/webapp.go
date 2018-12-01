// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package webapp

import (
	"github.com/jrmsdev/jcms"
	"github.com/jrmsdev/jcms/webapp/client"
	"github.com/jrmsdev/jcms/webapp/config"
)

type Webapp struct {
	cfg *config.Config
	uri string
}

func New(cfg *config.Config) *Webapp {
	return &Webapp{cfg, ""}
}

func (w *Webapp) Name() string {
	return w.cfg.Name
}

func (w *Webapp) Start() string {
	w.uri = jcms.Start(w.cfg)
	return w.uri
}

func (w *Webapp) Serve() {
	jcms.Serve()
}

func (w *Webapp) Stop() {
	jcms.Stop()
}

func (w *Webapp) Client() *client.Client {
	if w.uri == "" {
		panic("webapp not started yet!")
	}
	return client.New(w.uri)
}

func (w *Webapp) ServerUri() string {
	return w.uri
}
