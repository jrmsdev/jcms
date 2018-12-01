// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/jrmsdev/jcms/internal/_t/check"
	"github.com/jrmsdev/jcms/webapp/config"
)

func TestDefaults(t *testing.T) {
	cfg := config.New("")
	if check.NotEqual(t, cfg.Name, "default", "config name") {
		t.Fail()
	}
	expect := os.Getenv("JCMS_LOG")
	if expect == "" {
		expect = "default"
	}
	if check.NotEqual(t, cfg.Log, expect, "log level") {
		t.Fail()
	}
	if check.NotEqual(t, cfg.Basedir,
		filepath.FromSlash("/srv/jcms"), "basedir") {
		t.Fail()
	}
	if check.NotTrue(t, cfg.StaticEnable, "static enable") {
		t.Fail()
	}
	if check.NotEqual(t, cfg.HttpPort, "0", "http port") {
		t.Fail()
	}
}

func TestAssetsManager(t *testing.T) {
	cfg := config.New("")
	m := cfg.GetAssetsManager()
	typ := fmt.Sprintf("%T", m)
	if check.NotEqual(t, typ, "*manager.astman",
		"invalid assets manager type") {
		t.FailNow()
	}
}
