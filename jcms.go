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
	Start(cfg)
	defer Stop(cfg)
	Serve(cfg)
}

func Start(cfg *config.Config) string {
	config.SetDefaults(cfg)
	log.Init(cfg.Log)
	log.D("Start: %s", cfg.Name)
	webapp.Setup(cfg)
	httpd.Setup(cfg)
	return "127.0.0.1:6080"
}

func Serve(cfg *config.Config) {
	log.D("Serve")
}

func Stop(cfg *config.Config) {
	log.D("Stop")
}
