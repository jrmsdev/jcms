// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package webapp

import (
	"github.com/jrmsdev/jcms/assets"
	"github.com/jrmsdev/jcms/internal/log"
	"github.com/jrmsdev/jcms/webapp/config"
)

func Setup(cfg *config.Config) {
	log.D("Setup: %s", cfg.Name)
	assets.SetManager(cfg.AssetsManager)
}
