// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package schema

import (
	"path/filepath"
	"testing"

	"github.com/jrmsdev/jcms/internal/_t/check"
)

func TestParser(t *testing.T) {
	fn := filepath.FromSlash("./testdata/db.json")
	s := newSchema("testing")
	parse(s, fn)
	if check.NotEqual(t, s.String(), "testing", "schema name") {
		t.FailNow()
	}
}
