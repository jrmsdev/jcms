// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package asset

import (
	"io"
	"os"
	"path/filepath"

	"github.com/jrmsdev/jcms/lib/internal/flags"
	"github.com/jrmsdev/jcms/lib/log"
)

func filename(n string) string {
	return filepath.Join(flags.Assetsdir, flags.Webapp, n)
}

func Exists(name string) bool {
	log.D("exists %s", name)
	fp := filename(name)
	fi, err := os.Stat(fp)
	if err != nil {
		log.E("%s", err)
		return false
	}
	if fi.IsDir() {
		log.E("%s is a directory", fp)
		return false
	}
	return true
}

func Open(name string) (io.ReadCloser, error) {
	log.D("open %s", name)
	return os.Open(filename(name))
}
