// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"os"
	"path/filepath"

	"github.com/jrmsdev/jcms/assets"
	"github.com/jrmsdev/jcms/internal/assets/manager"
	"github.com/jrmsdev/jcms/internal/storage/driver"
	"github.com/jrmsdev/jcms/storage"

	"github.com/gorilla/mux"
)

type HandlerSetupFunc func(*mux.Router)

type Config struct {
	Name          string
	Log           string
	Assetsdir     string
	Datadir       string
	AssetsManager assets.Manager
	StorageDriver storage.Driver
	DatabaseURI   string
	StaticEnable  bool
	HttpPort      string
	HandlerSetup  map[string]HandlerSetupFunc
}

var (
	defName      string
	defLog       string
	defAssetsdir string
	defDatadir   string
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
	defAssetsdir = os.Getenv("JCMS_ASSETSDIR")
	if defAssetsdir == "" {
		defAssetsdir = filepath.FromSlash("/srv/jcms")
	}
	defDatadir = os.Getenv("JCMS_DATADIR")
	if defDatadir == "" {
		defDatadir = filepath.FromSlash("/srv/jcms")
	}
}

func New(name string) *Config {
	if name == "" {
		name = defName
	}
	return &Config{
		Name:         name,
		Log:          defLog,
		Assetsdir:    defAssetsdir,
		Datadir:      defDatadir,
		StaticEnable: true,
		HttpPort:     "0",
		HandlerSetup: make(map[string]HandlerSetupFunc),
		DatabaseURI:  "fs://jcms",
	}
}

func (cfg *Config) GetAssetsManager() assets.Manager {
	if cfg.AssetsManager == nil {
		cfg.AssetsManager = manager.New(cfg.Name, cfg.Assetsdir)
	}
	return cfg.AssetsManager
}

func (cfg *Config) GetStorageDriver() storage.Driver {
	if cfg.StorageDriver == nil {
		cfg.StorageDriver = driver.New(cfg.Name, cfg.Datadir)
	}
	return cfg.StorageDriver
}
