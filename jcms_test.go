// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jcms_test

import (
	"testing"

	"github.com/jrmsdev/jcms/internal/_t/check"
	"github.com/jrmsdev/jcms/internal/_t/test"
)

func TestMain(m *testing.M) {
	test.Main(m, "testing")
}

func TestWebappName(t *testing.T) {
	wapp := test.Webapp()
	if check.NotEqual(t, wapp.Name(), "testing", "webapp name") {
		t.FailNow()
	}
}

func TestServerUri(t *testing.T) {
	wapp := test.Webapp()
	if check.NotMatch(t, "^http://127\\.0\\.0\\.1:\\d+$",
		wapp.ServerUri(), "webapp server uri") {
		t.FailNow()
	}
}
