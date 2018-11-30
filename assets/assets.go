// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package assets

import (
	"io"
	"os"

	"github.com/jrmsdev/jcms/internal/log"
)

type File interface {
	io.ReadSeeker
	io.Closer
}

type Manager interface {
	Open(relname string) (File, error)
	Stat(relname string) (os.FileInfo, error)
	ReadFile(relname string) ([]byte, error)
}

var manager Manager

func SetManager(m Manager) {
	log.D("SetManager")
	if manager != nil {
		log.Panic("assets manager already set!")
	}
	manager = m
}

func ReadFile(relname string) ([]byte, error) {
	log.D("ReadFile: %s", relname)
	return manager.ReadFile(relname)
}

func Open(relname string) (File, error) {
	log.D("Open: %s", relname)
	log.Printf("asset open: %s", relname)
	return manager.Open(relname)
}

func Stat(relname string) (os.FileInfo, error) {
	log.D("Stat: %s", relname)
	return manager.Stat(relname)
}
