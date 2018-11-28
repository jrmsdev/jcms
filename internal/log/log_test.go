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
	if l == nil {
		l = xlog.New(&buf, "", 0)
		setLevel("testing")
	}
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
	Debug("testing quiet")
	defer buf.Reset()
	if check.NotEqual(t, buf.String(), "", "") {
		t.FailNow()
	}
	setLevel("testing")
}

func TestError(t *testing.T) {
	Error("testing error")
	defer buf.Reset()
	if check.NotEqual(t, buf.String(), "testing error\n", "") {
		t.FailNow()
	}
}

func TestDebug(t *testing.T) {
	Debug("testing debug")
	defer buf.Reset()
	if check.NotEqual(t, buf.String(), "", "") {
		t.FailNow()
	}
}

func TestEnableDebug(t *testing.T) {
	setLevel("debug")
	Debug("testing debug")
	defer buf.Reset()
	if check.NotEqual(t, buf.String(), "testing debug\n", "") {
		t.FailNow()
	}
	setLevel("testing")
}
