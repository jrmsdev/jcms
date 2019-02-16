// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package template

import (
	"encoding/json"

	"github.com/jrmsdev/jcms/lib/log"
)

type Config struct {
	Default   string            `json:"default"`
	Templates map[string]string `json:"templates"`
}

var initcfg = []byte(`{
	"default": "",
	"templates": {}
}`)

func cfgLoad(blob []byte) {
	log.D("load")
	if blob == nil {
		blob = initcfg
	}
	err := json.Unmarshal(blob, cfg)
	if err != nil {
		log.Panic("%s", err)
	}
}

func (c *Config) Get(path string) string {
	n, ok := c.Templates[path]
	if !ok {
		n = c.Default
	}
	return n
}
