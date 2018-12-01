// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package log

import (
	"bytes"
	xlog "log"
	"testing"

	"github.com/jrmsdev/jcms/internal/_t/check"
)

var buf bytes.Buffer

func testInit(lvl string) {
	if l == nil {
		codeInfo = false
		l = xlog.New(&buf, "", 0)
	}
	buf.Reset()
	setLevel(lvl)
}

func TestPrintf(t *testing.T) {
	testInit("default")
	Printf("testing printf")
	if check.NotEqual(t, buf.String(), "testing printf\n", "") {
		t.FailNow()
	}
}

func TestQuiet(t *testing.T) {
	testInit("quiet")
	Printf("testing quiet")
	if check.NotEqual(t, buf.String(), "", "") {
		t.FailNow()
	}
}

func TestError(t *testing.T) {
	testInit("default")
	E("testing")
	if check.NotEqual(t, buf.String(), "ERROR testing\n", "") {
		t.FailNow()
	}
}

func TestDebugDisabled(t *testing.T) {
	testInit("default")
	D("testing debug disabled")
	if check.NotEqual(t, buf.String(), "", "") {
		t.FailNow()
	}
}

func TestDebug(t *testing.T) {
	testInit("debug")
	D("testing debug")
	if check.NotEqual(t, buf.String(), "[D] testing debug\n", "") {
		t.FailNow()
	}
}
