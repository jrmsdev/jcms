// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package template

import (
	"bytes"
	"testing"

	"github.com/jrmsdev/jcms/_t/check"
)

var tt = map[string]string{
	"testing": "testing",
}

func TestTemplate(t *testing.T) {
	for s, d := range tt {
		t.Log(s, d)
		src := bytes.NewBufferString(s)
		dst := new(bytes.Buffer)
		err := Parse(dst, src)
		if err != nil {
			t.Log(err)
			t.FailNow()
		}
		if check.NotEqual(t, dst.String(), d, s) {
			t.Fail()
		}
	}
}
