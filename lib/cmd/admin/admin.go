// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package admin

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jrmsdev/jcms"
	"github.com/jrmsdev/jcms/lib/internal/admin/handler"
	"github.com/jrmsdev/jcms/lib/internal/flags"
	"github.com/jrmsdev/jcms/lib/log"
)

func Main() {
	flags.Parse()
	if flags.ShowVersion {
		fmt.Fprintf(os.Stderr, "jcms-admin version %s\n", jcms.Version())
		os.Exit(0)
	}
	log.Init(flags.Log)
	log.Printf("%s version %s", filepath.Base(os.Args[0]), jcms.Version())
	log.Printf("http://127.0.0.1:%s/", flags.HttpPort)
	srv := handler.Setup()
	if err := srv.ListenAndServe(); err != nil {
		log.E("%s", err)
		os.Exit(2)
	}
}
