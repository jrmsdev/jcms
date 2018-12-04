// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package assets

import (
	"github.com/jrmsdev/jcms/internal/log"
)

type Manager interface {
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
