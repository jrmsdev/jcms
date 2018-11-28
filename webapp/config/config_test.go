// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"testing"

	"github.com/jrmsdev/jcms/internal/_t/check"
)

func TestDefaults(t *testing.T) {
	cfg := &Config{}
	SetDefaults(cfg)
	if check.NotEqual(t, cfg.Name, "default", "config name") {
		t.Fail()
	}
	if check.NotEqual(t, cfg.Log, "default", "config log") {
		t.Fail()
	}
}
