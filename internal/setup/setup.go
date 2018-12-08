// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package setup

import (
	"github.com/jrmsdev/jcms/assets"
	"github.com/jrmsdev/jcms/internal/log"
	"github.com/jrmsdev/jcms/storage"
	"github.com/jrmsdev/jcms/webapp/config"
)

func Webapp(cfg *config.Config) {
	log.D("Webapp: %s", cfg.Name)
	assets.SetManager(cfg.GetAssetsManager())
	storage.SetDriver(cfg.GetStorageDriver())
}
