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

var wapp *webapp.Webapp

func init() {
	wapp = webapp.New("testing")
}

func Main(m *testing.M) {
	println("test.Main")
	os.Exit(m.Run())
}
