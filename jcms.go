// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jcms

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/jrmsdev/jcms/db"
	"github.com/jrmsdev/jcms/db/schema"
	"github.com/jrmsdev/jcms/internal/cmd/flags"
	"github.com/jrmsdev/jcms/internal/httpd"
	"github.com/jrmsdev/jcms/internal/log"
	"github.com/jrmsdev/jcms/internal/webapp"
	"github.com/jrmsdev/jcms/webapp/config"
)

var done bool

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

func trapSignals() {
	log.D("trapSignals")
	done = false
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		s := <-sigs
		log.Printf("got signal: %s", s)
		Stop()
	}()
}

func Start(cfg *config.Config) string {
	log.Init(cfg.Log)
	log.D("Start: %s", cfg.Name)
	log.Printf("jcms v%s", Version())
	log.Printf("basedir %s", cfg.Basedir)
	webapp.Setup(cfg)
	db.CheckEngine()
	httpd.Setup(cfg)
	trapSignals()
	return httpd.Listen()
}

func Serve() {
	log.D("Serve")
	db.Connect()
	schema.Check()
	httpd.Serve()
}

func Stop() {
	if done {
		log.D("Stop done!")
	} else {
		log.D("Stop")
		db.Disconnect()
		httpd.Stop()
		done = true
	}
}
