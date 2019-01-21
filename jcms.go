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
	"github.com/jrmsdev/jcms/internal/errors"
	"github.com/jrmsdev/jcms/internal/httpd"
	"github.com/jrmsdev/jcms/internal/log"
	"github.com/jrmsdev/jcms/internal/webapp"
	"github.com/jrmsdev/jcms/webapp/config"
)

var done bool

func Main() {
	cfg := flags.Parse()
	if flags.ShowVersion {
		fmt.Fprintf(os.Stderr, "jcms version %s\n", Version())
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
	log.Printf("assets %s", cfg.Assetsdir)
	log.Printf("data %s", cfg.Datadir)
	webapp.Setup(cfg)
	db.CheckEngine()
	httpd.Setup(cfg)
	trapSignals()
	return httpd.Listen()
}

func Serve() {
	log.D("Serve")
	// connect
	if err := db.Connect(); err != nil {
		httpd.Serve(errors.DBError(err.Error()))
	}
	// check schema
	if err := schema.Check(); err != nil {
		httpd.Serve(errors.DBError(err.Error()))
	}
	httpd.Serve(nil)
}

func Stop() {
	if done {
		log.D("Stop done!")
	} else {
		log.D("Stop")
		if err := db.Disconnect(); err != nil {
			log.E("DB disconnect: %s", err)
		}
		httpd.Stop()
		done = true
	}
}
