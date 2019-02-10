// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package asset

import (
	fp "path/filepath"
	"testing"

	"github.com/jrmsdev/jcms/_t/check"
)

func init() {
	assetsdir = "testdata"
	webapp = "wapp"
}

func TestFilename(t *testing.T) {
	n := "test.txt"
	if check.NotEqual(t, filename(n), fp.Join(assetsdir, webapp, n), "asset filename") {
		t.Fail()
	}
}

func TestExists(t *testing.T) {
	n := fp.FromSlash("test.txt")
	if check.NotTrue(t, Exists(n), n+" exists") {
		t.Fail()
	}
}

func TestNotExists(t *testing.T) {
	n := fp.FromSlash("nofile.txt")
	if check.NotFalse(t, Exists(n), n+" not exists") {
		t.Fail()
	}
}

func TestIsDir(t *testing.T) {
	n := fp.FromSlash("testdir")
	if check.NotFalse(t, Exists(n), n+" exists") {
		t.Fail()
	}
}

func TestOpen(t *testing.T) {
	n := fp.FromSlash("test.txt")
	_, err := Open(n)
	if check.NotNil(t, err, n+" open error") {
		t.Fail()
	}
}

func TestOpenError(t *testing.T) {
	n := fp.FromSlash("nofile.txt")
	_, err := Open(n)
	if check.IsNil(t, err, n+" open error") {
		t.Fail()
	}
}
