// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package assets

import (
	"os"
	"path/filepath"

	. "github.com/jrmsdev/jcms/assets"
	"github.com/jrmsdev/jcms/internal/log"
)

var manager Manager

func SetManager(m Manager) {
	log.D("SetManager")
	manager = m
}

func ReadFile(parts ...string) ([]byte, error) {
	fn := filepath.Join(parts...)
	return manager.ReadFile(fn)
}

func Open(parts ...string) (File, error) {
	fn := filepath.Join(parts...)
	return manager.Open(fn)
}

func Stat(parts ...string) (os.FileInfo, error) {
	fn := filepath.Join(parts...)
	return manager.Stat(fn)
}
