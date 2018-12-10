// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jcms

import (
	"fmt"
	"os"

	"github.com/jrmsdev/jcms/db"
	"github.com/jrmsdev/jcms/db/schema"
	"github.com/jrmsdev/jcms/internal/cmd/flags"
	"github.com/jrmsdev/jcms/internal/httpd"
	"github.com/jrmsdev/jcms/internal/log"
	"github.com/jrmsdev/jcms/internal/webapp"
	"github.com/jrmsdev/jcms/webapp/config"
)

func Main() {
	cfg := flags.Parse()
	if flags.ShowVersion {
		fmt.Printf("jcms version %s\n", Version())
		os.Exit(0)
	}
	log.Printf("%s %s", cfg.Name, Start(cfg))
	defer Stop()
	Serve()
}

func Start(cfg *config.Config) string {
	log.Init(cfg.Log)
	log.D("Start: %s", cfg.Name)
	log.Printf("jcms v%s", Version())
	log.Printf("basedir %s", cfg.Basedir)
	webapp.Setup(cfg)
	db.CheckEngine()
	httpd.Setup(cfg)
	return httpd.Listen()
}

func Serve() {
	log.D("Serve")
	db.Connect()
	schema.Check()
	httpd.Serve()
	//~ db.Disconnect()
}

func Stop() {
	log.D("Stop")
	httpd.Stop()
}
