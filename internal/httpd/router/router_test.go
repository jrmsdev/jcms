// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package router_test

import (
	"fmt"
	"testing"

	"github.com/jrmsdev/jcms/internal/_t/check"
	"github.com/jrmsdev/jcms/internal/httpd/router"
)

func TestRouterInit(t *testing.T) {
	r := router.Init()
	typ := fmt.Sprintf("%T", r)
	if check.NotEqual(t, typ, "*mux.Router",
		"invalid httpd router type") {
		t.FailNow()
	}
}
