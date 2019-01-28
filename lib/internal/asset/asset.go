// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package asset

import (
	"io"
	"os"

	"github.com/jrmsdev/jcms/lib/log"
)

func Open(filename string) (io.ReadCloser, error) {
	log.D("open %s", filename)
	return os.Open(filename)
}
