// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package webapp

import (
	"os"
	"path/filepath"

	"github.com/jrmsdev/jcms"
	"github.com/jrmsdev/jcms/lib/internal/flags"
	"github.com/jrmsdev/jcms/lib/log"
)

func Main(w *Webapp) {
	cmd := filepath.Base(os.Args[0])
	if flags.ShowVersion {
		fprintf(os.Stderr, "%s version %s\n", cmd, jcms.Version())
		os.Exit(0)
	}
	if _, err := Listen(w); err != nil {
		log.E("%s", err)
		os.Exit(1)
	}
	defer stop(w)
	if err := Start(w); err != nil {
		log.E("%s", err)
		os.Exit(2)
	}
}

func stop(w *Webapp) {
	if err := Stop(w); err != nil {
		log.E("%s", err)
	}
}
