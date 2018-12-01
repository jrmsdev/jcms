// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jrmsdev/jcms/webapp"
	"github.com/jrmsdev/jcms/webapp/client"
	"github.com/jrmsdev/jcms/webapp/config"
)

var (
	wapp      *webapp.Webapp
	cli       *client.Client
	serverURI string
)

func Main(m *testing.M, name string) {
	if wapp != nil {
		panic("wapp is not nil")
	}
	wapp = webapp.New(newConfig(name))
	serverURI = wapp.Start()
	go func() {
		wapp.Serve()
		wapp.Stop()
	}()
	cli = wapp.Client()
	os.Exit(m.Run())
}

func newConfig(name string) *config.Config {
	srcdir := filepath.Join(os.Getenv("GOPATH"), "src")
	cfg := config.New(name)
	cfg.Basedir = filepath.Join(srcdir,
		"github.com", "jrmsdev", "jcms", "testdata", "basedir")
	return cfg
}

func Webapp() *webapp.Webapp {
	return wapp
}
