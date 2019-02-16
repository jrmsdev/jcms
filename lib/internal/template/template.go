// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package template

import (
	"io/ioutil"

	"github.com/jrmsdev/jcms/lib/internal/asset"
	"github.com/jrmsdev/jcms/lib/log"
)

var cfg *Config

func Setup() {
	log.D("setup")
	if cfg != nil {
		log.Panic("templates setup already done!")
	}
	cfg = new(Config)
	if asset.Exists("templates.json") {
		var blob []byte
		fh, err := asset.Open("templates.json")
		if err != nil {
			log.Panic("%s", err)
		}
		defer fh.Close()
		blob, err = ioutil.ReadAll(fh)
		if err != nil {
			log.Panic("%s", err)
		}
		cfgLoad(blob)
	} else {
		cfgLoad(nil)
	}
}
