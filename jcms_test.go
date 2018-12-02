// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jcms_test

import (
	"testing"
	"time"

	"github.com/jrmsdev/jcms"
	"github.com/jrmsdev/jcms/internal/_t/test"
)

func TestJCMS(t *testing.T) {
	cfg := test.Config("testing")
	jcms.Start(cfg)
	go func() {
		jcms.Serve()
	}()
	time.Sleep(300 * time.Millisecond)
	jcms.Stop()
}
