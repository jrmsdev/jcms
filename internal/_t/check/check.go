// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package check

import (
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
