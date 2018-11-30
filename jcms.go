// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jcms

import (
	"fmt"
	"os"

	"github.com/jrmsdev/jcms/internal/cmd/flags"
	"github.com/jrmsdev/jcms/internal/httpd"
	"github.com/jrmsdev/jcms/internal/log"
	"github.com/jrmsdev/jcms/internal/webapp"
	"github.com/jrmsdev/jcms/webapp/config"
)

func Main() {
	flags.Parse()
	if flags.ShowVersion {
		fmt.Printf("jcms version %s\n", Version())
		os.Exit(0)
	}
	cfg := config.New()
	if flags.Quiet {
		cfg.Log = "quiet"
	}
	if flags.Debug {
		cfg.Log = "debug"
	}
	log.Printf("jcms v%s %s", Version(), Start(cfg))
	defer Stop()
	Serve()
}

func Start(cfg *config.Config) string {
	config.SetDefaults(cfg)
	log.Init(cfg.Log)
	log.D("Start: %s", cfg.Name)
	webapp.Setup(cfg)
	httpd.Setup(cfg)
	return httpd.Listen()
}

func Serve() {
	log.D("Serve")
	httpd.Serve()
}

func Stop() {
	log.D("Stop")
	httpd.Stop()
}
