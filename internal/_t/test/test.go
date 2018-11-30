// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package test

import (
	"os"
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
	defer wapp.Stop()
	go func() {
		wapp.Serve()
	}()
	cli = wapp.Client()
	rc := m.Run()
	wapp = nil
	os.Exit(rc)
}

func newConfig(name string) *config.Config {
	return &config.Config{Name: name}
}

func Webapp() *webapp.Webapp {
	return wapp
}

func Client() *client.Client {
	return cli
}
