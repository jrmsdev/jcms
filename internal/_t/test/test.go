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
	wapp = webapp.New(Config(name))
	serverURI = wapp.Start()
	go func() {
		wapp.Serve()
		wapp.Stop()
	}()
	cli = wapp.Client()
	os.Exit(m.Run())
}

func Config(name string) *config.Config {
	cfg := config.New(name)
	cfg.Basedir = filepath.FromSlash("./testdata/assets")
	cfg.Datadir = filepath.FromSlash("./testdata/data")
	return cfg
}

func Webapp() *webapp.Webapp {
	return wapp
}
