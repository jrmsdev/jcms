// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

type Config struct {
	Name string
	Log string
}

func SetDefaults(cfg *Config) {
	if cfg.Log == "" {
		cfg.Log = "default"
	}
}
