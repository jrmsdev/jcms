// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package test

import (
	"os"
	"testing"

	"github.com/jrmsdev/jcms/webapp"
	"github.com/jrmsdev/jcms/webapp/config"
)

func Config() *config.Config {
	return &config.Config{Name: "testing"}
}

var (
	wapp      *webapp.Webapp
	serverURI string
)

func Main(m *testing.M, name string) {
	if wapp != nil {
		panic("wapp is not nil")
	}
	wapp = webapp.New(newConfig(name))
	serverURI = wapp.Start()
	defer wapp.Stop()
	go func() {
		wapp.Serve()
	}()
	rc := m.Run()
	wapp = nil
	os.Exit(rc)
}

func Webapp() *webapp.Webapp {
	return wapp
}

func newConfig(name string) *config.Config {
	return &config.Config{Name: name}
}
