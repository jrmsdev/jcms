// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package db

import (
	"github.com/jrmsdev/jcms/internal/log"
)

var eng Engine

type Engine interface {
	String() string
}

func SetEngine(e Engine) {
	log.D("SetEngine %s", e)
	if eng != nil {
		log.Panic("db engine already set: %s", eng)
	}
	eng = e
}
