// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package template

import (
	"io"
	"path/filepath"

	"github.com/jrmsdev/jcms/lib/log"
)

func Parse(dst io.Writer, src io.Reader, path string) error {
	fn := filepath.Join("tpl", cfg.Get(path) + ".html")
	log.D("parse %s", fn)
	_, err := io.Copy(dst, src)
	return err
}
