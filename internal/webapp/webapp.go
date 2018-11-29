// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package webapp

import (
	"github.com/jrmsdev/jcms/internal/assets"
	"github.com/jrmsdev/jcms/internal/log"
	"github.com/jrmsdev/jcms/webapp/config"
)

func Setup(cfg *config.Config) {
	log.D("Setup")
	assets.SetManager(cfg.AssetsManager)
}

func Start(cfg *config.Config) {
	log.D("Start")
}

func Stop(cfg *config.Config) {
	log.D("Stop")
}
