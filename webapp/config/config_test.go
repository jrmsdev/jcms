// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package config

import (
	"os"
	"testing"

	"github.com/jrmsdev/jcms/internal/_t/check"
)

func TestDefaults(t *testing.T) {
	cfg := New("")
	if check.NotEqual(t, cfg.Name, "default", "config name") {
		t.Fail()
	}
	expect := os.Getenv("JCMS_LOG")
	if expect == "" {
		expect = "default"
	}
	if check.NotEqual(t, cfg.Log, expect, "config log") {
		t.Fail()
	}
}
