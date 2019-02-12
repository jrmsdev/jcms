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

func cfgLoad(cfg *Config, blob []byte) {
	log.D("load")
	err := json.Unmarshal(blob, cfg)
	if err != nil {
		log.Panic("%s", err)
	}
	if cfg.Default == "" {
		cfg.Default = "main"
	}
}

func (c *Config) Get(path string) string {
	n, ok := c.Templates[path]
	if !ok {
		n = c.Default
	}
	log.D("%s template: %s", path, n)
	return n
}
