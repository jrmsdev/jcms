// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package template

import (
	"encoding/json"

	"github.com/jrmsdev/jcms/lib/log"
)

type Config struct {
	Templates map[string]string `json:"templates"`
}

var initcfg = []byte(`{"templates": {}}`)

var admincfg = []byte(`{
	"templates": {
		"/": "main"
	}
}`)

func cfgLoad(n string, blob []byte) {
	log.D("load %s", n)
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
		return ""
	}
	return n
}
