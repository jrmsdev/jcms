// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"os"
	"path/filepath"

	"github.com/jrmsdev/jcms/assets"
	"github.com/jrmsdev/jcms/internal/assets/manager"

	"github.com/gorilla/mux"
)

type HandlerSetupFunc func(*mux.Router)

type Config struct {
	Name          string
	Log           string
	Basedir       string
	AssetsManager assets.Manager
	StaticEnable  bool
	HttpPort      string
	HandlerSetup  map[string]HandlerSetupFunc
}

var (
	defName    string
	defLog     string
	defBasedir string
)

func init() {
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
	if name == "" {
		name = defName
	}
	return &Config{
		Name:         name,
		Log:          defLog,
		Basedir:      defBasedir,
		StaticEnable: true,
		HttpPort:     "0",
		HandlerSetup: make(map[string]HandlerSetupFunc),
	}
}

func (cfg *Config) GetAssetsManager() assets.Manager {
	if cfg.AssetsManager == nil {
		cfg.AssetsManager = manager.New(cfg.Name, cfg.Basedir)
	}
	return cfg.AssetsManager
}
