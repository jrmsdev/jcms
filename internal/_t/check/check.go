// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package check

import (
	"regexp"
	"testing"
)

func NotEqual(t *testing.T, got, expect interface{}, errmsg string) bool {
	t.Helper()
	if got != expect {
		t.Logf("%s: got: '%v' - expect: '%v'", errmsg, got, expect)
		return true
	}
	return false
}

func NotMatch(t *testing.T, pat, s, errmsg string) bool {
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
