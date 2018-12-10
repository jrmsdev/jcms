// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package db

import (
	"github.com/jrmsdev/jcms/internal/log"
)

var eng Engine

type Engine interface {
	String() string
	Webapp() string
	Connect() error
	Disconnect() error
}

func SetEngine(e Engine) {
	log.D("SetEngine %s", e)
	if eng != nil {
		log.Panic("db engine already set: %s", eng)
	}
	eng = e
}

func CheckEngine() {
	log.D("CheckEngine")
	if eng == nil {
		log.Panic("db engine not set!")
	}
}

func Webapp() string {
	return eng.Webapp()
}

func Connect() error {
	return eng.Connect()
}

func Disconnect() error {
	return eng.Disconnect()
}
