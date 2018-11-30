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

func init() {
	codeInfo = false
	l = xlog.New(&buf, "", 0)
	setLevel("testing")
}

func TestPrintf(t *testing.T) {
	Printf("testing printf")
	defer buf.Reset()
	if check.NotEqual(t, buf.String(), "testing printf\n", "") {
		t.FailNow()
	}
}

func TestQuiet(t *testing.T) {
	setLevel("quiet")
	Printf("testing quiet")
	defer buf.Reset()
	if check.NotEqual(t, buf.String(), "", "") {
		t.FailNow()
	}
	setLevel("testing")
}

func TestError(t *testing.T) {
	E("testing")
	defer buf.Reset()
	if check.NotEqual(t, buf.String(), "ERROR testing\n", "") {
		t.FailNow()
	}
}

func TestDebugDisabled(t *testing.T) {
	D("testing debug disabled")
	defer buf.Reset()
	if check.NotEqual(t, buf.String(), "", "") {
		t.FailNow()
	}
}

func TestDebug(t *testing.T) {
	setLevel("debug")
	D("testing debug")
	defer buf.Reset()
	if check.NotEqual(t, buf.String(), "[D] testing debug\n", "") {
		t.FailNow()
	}
	setLevel("testing")
}
