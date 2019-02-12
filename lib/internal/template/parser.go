// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package template

import (
	"io"

	"github.com/jrmsdev/jcms/lib/log"
)

func Parse(dst io.Writer, src io.Reader, path string) error {
	log.D("parse")
	_, err := io.Copy(dst, src)
	return err
}
