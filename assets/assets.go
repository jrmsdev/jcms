// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package assets

import (
	"io"
	"os"
	"path/filepath"

	"github.com/jrmsdev/jcms/internal/log"
)

type File interface {
	io.ReadSeeker
	io.Closer
}

type Manager interface {
	Open(filename string) (File, error)
	Stat(filename string) (os.FileInfo, error)
	ReadFile(name string) ([]byte, error)
}

var manager Manager

func SetManager(m Manager) {
	log.D("SetManager")
	if manager != nil {
		log.Panic("assets manager already set!")
	}
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
