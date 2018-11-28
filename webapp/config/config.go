// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import "os"

type Config struct {
	Name string
	Log  string
}

var defName string
var defLog string

func init() {
	defName = os.Getenv("JCMS_WEBAPP")
	if defName == "" {
		defName = "default"
	}
	defLog = os.Getenv("JCMS_LOG")
	if defLog == "" {
		defLog = "default"
	}
}

func SetDefaults(cfg *Config) {
	if cfg.Name == "" {
		cfg.Name = defName
	}
	if cfg.Log == "" {
		cfg.Log = defLog
	}
}
