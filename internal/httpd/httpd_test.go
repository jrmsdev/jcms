// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package httpd_test

import (
	"testing"
	"time"

	"github.com/jrmsdev/jcms/internal/httpd"
	"github.com/jrmsdev/jcms/internal/_t/test"
)

func TestHttpd(t *testing.T) {
	cfg := test.Config("testing")
	httpd.Setup(cfg)
	httpd.Listen()
	go func() {
		httpd.Serve()
	}()
	time.Sleep(300 * time.Millisecond)
	httpd.Stop()
}
