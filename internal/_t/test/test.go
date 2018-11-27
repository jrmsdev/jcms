// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package test

import (
	"github.com/jrmsdev/jcms/webapp/config"
)

func Config() *config.Config {
	return &config.Config{Name: "testing"}
}
