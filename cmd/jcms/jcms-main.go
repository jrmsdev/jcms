// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"github.com/jrmsdev/jcms"
	"github.com/jrmsdev/jcms/webapp/config"
)

func main() {
	cfg := &config.Config{
		Name: "default",
	}
	jcms.Main(cfg)
}
