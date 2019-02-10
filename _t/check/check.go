// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package check

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/jrmsdev/jcms/lib/log"
)

func init() {
	log.InitTest()
}

func NotNil(t *testing.T, got interface{}, errmsg string) bool {
	t.Helper()
	if got != nil {
		t.Logf("%s not nil: %T", errmsg, got)
		return true
	}
	return false
}

func IsNil(t *testing.T, got interface{}, errmsg string) bool {
	t.Helper()
	if got == nil {
		t.Logf("%s is nil", errmsg)
		return true
	}
	return false
}

func NotTrue(t *testing.T, got bool, errmsg string) bool {
	t.Helper()
	if !got {
		t.Logf("%s: is false (should be true)", errmsg)
		return true
	}
	return false
}

func NotFalse(t *testing.T, got bool, errmsg string) bool {
	t.Helper()
	if got {
		t.Logf("%s: is true (should be false)", errmsg)
		return true
	}
	return false
}

func NotEqual(t *testing.T, got, expect interface{}, errmsg string) bool {
	t.Helper()
	if got != expect {
		t.Logf("%s: got: '%v' - expect: '%v'", errmsg, got, expect)
		return true
	}
	return false
}

func NotMatch(t *testing.T, pat, s, errmsg string) bool {
	t.Helper()
	m, err := regexp.MatchString(pat, s)
	if err != nil {
		t.Fatalf("%s: %s", errmsg, err)
		return true
	}
	if m {
		return false
	}
	t.Logf("%s: '%s' not match '%s'", errmsg, s, pat)
	return true
}

func NotFileChecksum(t *testing.T, got []byte, fn string) bool {
	t.Helper()
	fh, err := os.Open(filepath.FromSlash(fn))
	if err != nil {
		t.Fatalf("%s: %s", fn, err.Error())
		return true
	}
	defer fh.Close()
	h := md5.New()
	if _, err := io.Copy(h, fh); err != nil {
		t.Fatalf("%s: %s", fn, err.Error())
		return true
	}
	return NotEqual(t, fmt.Sprintf("%x", md5.Sum(got)),
		fmt.Sprintf("%x", h.Sum(nil)), "checksum "+fn)
}
