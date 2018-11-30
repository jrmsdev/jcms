// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"os"
	"path/filepath"

	"github.com/jrmsdev/jcms/assets"
	"github.com/jrmsdev/jcms/internal/assets/manager"
)

type Config struct {
	Name          string
	Log           string
	Basedir       string
	AssetsManager assets.Manager
	StaticEnable  bool
	HttpPort      string
}

var (
	defDone    bool
	defName    string
	defLog     string
	defBasedir string
)

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
	defBasedir = os.Getenv("JCMS_BASEDIR")
	if defBasedir == "" {
		defBasedir = filepath.FromSlash("/srv/jcms")
	}
}

func New(name string) *Config {
	return &Config{
		Name: name,
		StaticEnable: true,
	}
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
	if cfg.Basedir == "" {
		cfg.Basedir = defBasedir
	}
	if cfg.AssetsManager == nil {
		cfg.AssetsManager = manager.New(cfg.Name, cfg.Basedir)
	}
	if cfg.HttpPort == "" {
		cfg.HttpPort = "0"
	}
}
