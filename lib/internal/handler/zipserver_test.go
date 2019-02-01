// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package handler

import (
	"testing"

	"github.com/jrmsdev/jcms/_t/check"
	"github.com/jrmsdev/jcms/_t/http"
)

func TestZHeaders(t *testing.T) {
	zs := &zipServer{}
	w := http.Writer()
	for _, x := range ht {
		zs.setHeaders(w, x.path)
		if check.NotEqual(t, w.Header().Get(x.key), x.expect, x.path + " " + x.key) {
			t.Fail()
		}
	}
}
