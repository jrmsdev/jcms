// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package httpd

import (
	"net/http"
	"time"

	"github.com/jrmsdev/jcms/internal/httpd/handler"
	"github.com/jrmsdev/jcms/internal/httpd/router"
	"github.com/jrmsdev/jcms/internal/log"
	"github.com/jrmsdev/jcms/webapp/config"
)

var server *http.Server

func Setup(cfg *config.Config) {
	log.D("Setup")
	if server != nil {
		log.Panic("httpd setup was done already!")
	}
	r := router.Init()
	handler.Setup(r, cfg)
	server = &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:" + cfg.HttpPort,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}
