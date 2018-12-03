// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package webapp_test

import (
	"fmt"
	"testing"

	"github.com/jrmsdev/jcms/internal/_t/check"
	"github.com/jrmsdev/jcms/internal/_t/test"
)

func TestMain(m *testing.M) {
	test.Main(m, "testing")
}

func TestWebapp(t *testing.T) {
	wapp := test.Webapp()
	typ := fmt.Sprintf("%T", wapp)
	if check.NotEqual(t, typ, "*webapp.Webapp", "webapp type") {
		t.FailNow()
	}
}

func TestName(t *testing.T) {
	wapp := test.Webapp()
	if check.NotEqual(t, wapp.Name(), "testing", "webapp name") {
		t.FailNow()
	}
}

func TestServerUri(t *testing.T) {
	wapp := test.Webapp()
	if check.NotMatch(t, "^http://127\\.0\\.0\\.1:\\d+$",
		wapp.ServerUri(), "server uri") {
		t.FailNow()
	}
}

func TestClient(t *testing.T) {
	wapp := test.Webapp()
	cli := wapp.Client()
	typ := fmt.Sprintf("%T", cli)
	if check.NotEqual(t, typ, "*client.Client", "webapp client type") {
		t.FailNow()
	}
}
