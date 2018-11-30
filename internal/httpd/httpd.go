// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package httpd

import (
	"context"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/jrmsdev/jcms/internal/httpd/handler"
	"github.com/jrmsdev/jcms/internal/httpd/router"
	"github.com/jrmsdev/jcms/internal/log"
	"github.com/jrmsdev/jcms/webapp/config"
)

var (
	server     *http.Server
	serverAddr string
	listener   net.Listener
)

func Setup(cfg *config.Config) {
	log.D("Setup")
	if server != nil {
		log.Panic("httpd setup was done already!")
	}
	serverAddr = "127.0.0.1:" + cfg.HttpPort
	r := router.Init()
	handler.Setup(r, cfg)
	server = &http.Server{
		Handler:        r,
		Addr:           serverAddr,
		WriteTimeout:   10 * time.Second,
		ReadTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func Listen() string {
	var err error
	listener, err = net.Listen("tcp4", serverAddr)
	if err != nil {
		log.E("httpd listen: %s", serverAddr)
		log.Panic(err.Error())
	}
	url := &url.URL{}
	url.Scheme = "http"
	url.Host = listener.Addr().String()
	url.Path = ""
	return url.String()
}

func Serve() {
	log.D("Serve")
	if listener == nil {
		log.Panic("nil listener... call httpd.Listen() first")
	}
	var err error
	err = server.Serve(listener)
	if err != nil {
		log.E("httpd serve: %s", err.Error())
	}
}

func Stop() {
	log.D("Stop")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		log.E("httpd stop: %s", err.Error())
	}
}
