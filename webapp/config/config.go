// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"os"

	"github.com/jrmsdev/jcms/assets"
	"github.com/jrmsdev/jcms/internal/assets/manager"
)

type Config struct {
	Name          string
	Log           string
	AssetsManager assets.Manager
	StaticEnable  bool
	StaticURL     string
}

var defDone bool
var defName string
var defLog string

func init() {
	defDone = false
	defName = os.Getenv("JCMS_WEBAPP")
	if defName == "" {
		defName = "default"
	}
	defLog = os.Getenv("JCMS_LOG")
	if defLog == "" {
		defLog = "default"
	}
}

func New() *Config {
	return &Config{}
}

func SetDefaults(cfg *Config) {
	if defDone {
		panic("config.SetDefaults was already called!")
	}
	defDone = true
	if cfg.Name == "" {
		cfg.Name = defName
	}
	if cfg.Log == "" {
		cfg.Log = defLog
	}
	if cfg.AssetsManager == nil {
		cfg.AssetsManager = manager.New()
	}
	cfg.StaticEnable = true
	cfg.StaticURL = "/static/"
}
