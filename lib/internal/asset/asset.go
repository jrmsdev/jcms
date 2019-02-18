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

var assetsdir string = flags.Assetsdir
var webapp string = flags.Webapp

func filename(n string) string {
	return filepath.Join(assetsdir, webapp, n)
}

func Exists(name string) bool {
	log.D("%s exists?", name)
	fp := filename(name)
	fi, err := os.Stat(fp)
	if err != nil {
		log.D("%s", err)
		return false
	}
	if fi.IsDir() {
		log.D("%s is a directory", fp)
		return false
	}
	return true
}

func Open(name string) (io.ReadCloser, error) {
	log.D("open %s", name)
	return os.Open(filename(name))
}

func InitTest() {
	assetsdir = "testdata"
	webapp = "wapp"
}
