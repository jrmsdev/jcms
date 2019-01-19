// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package schema_test

import (
	"os"
	"testing"

	"github.com/jrmsdev/jcms/db/schema"
	"github.com/jrmsdev/jcms/internal/_t/check"
	"github.com/jrmsdev/jcms/internal/_t/test"
	"github.com/jrmsdev/jcms/internal/log"
	"github.com/jrmsdev/jcms/internal/webapp"
)

func TestMain(m *testing.M) {
	cfg := test.Config("testing")
	log.Init(cfg.Log)
	webapp.Setup(cfg)
	os.Exit(m.Run())
}

func TestParser(t *testing.T) {
	if check.NotNil(t, schema.Check(), "check error") {
		t.FailNow()
	}
}
