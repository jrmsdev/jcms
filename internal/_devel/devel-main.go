// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"fmt"
	"os"

	"github.com/jrmsdev/jcms"
	"github.com/jrmsdev/jcms/internal/cmd/flags"
	"github.com/jrmsdev/jcms/internal/log"
)

func main() {
	cfg := flags.Parse()
	if flags.ShowVersion {
		fmt.Printf("jcms-devel version %s\n", jcms.Version())
		os.Exit(0)
	}
	cfg.HandlerSetup["libhandler"] = libHandlerSetup
	cfg.HandlerSetup["pprof"] = pprofSetup
	log.Printf("%s %s", cfg.Name, jcms.Start(cfg))
	defer jcms.Stop()
	jcms.Serve()
}
