// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"fmt"
	"os"

	"github.com/jrmsdev/jcms"
	"github.com/jrmsdev/jcms/internal/cmd/flags"
	"github.com/jrmsdev/jcms/internal/log"
	"github.com/jrmsdev/jcms/webapp/config"
)

func main() {
	flags.Parse()
	if flags.ShowVersion {
		fmt.Printf("jcms-devel version %s\n", jcms.Version())
		os.Exit(0)
	}
	cfg := config.New(flags.Webapp)
	if flags.Quiet {
		cfg.Log = "quiet"
	}
	if flags.Debug {
		cfg.Log = "debug"
	}
	cfg.HandlerSetup["libhandler"] = libHandlerSetup
	cfg.HttpPort = fmt.Sprintf("%d", flags.HttpPort)
	log.Printf("%s %s", cfg.Name, jcms.Start(cfg))
	defer jcms.Stop()
	jcms.Serve()
}
