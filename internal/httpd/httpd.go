// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package httpd

import (
	"github.com/jrmsdev/jcms/internal/httpd/handler"
	"github.com/jrmsdev/jcms/internal/httpd/router"
	"github.com/jrmsdev/jcms/internal/log"
	"github.com/jrmsdev/jcms/webapp/config"
)

func Setup(cfg *config.Config) {
	log.D("Setup")
	r := router.Init()
	handler.Setup(r, cfg)
}
