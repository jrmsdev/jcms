// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package handler

import (
	"io/ioutil"
	gohttp "net/http"
	"testing"

	"github.com/jrmsdev/jcms/_t/check"
	"github.com/jrmsdev/jcms/_t/http"
)

func TestHandler(t *testing.T) {
	w := http.Writer()
	Error(w, "testing", gohttp.StatusInternalServerError)
	res := w.Result()
	if check.NotEqual(t, res.StatusCode, gohttp.StatusInternalServerError, "response status") {
		t.FailNow()
	}
	if check.NotEqual(t, res.Header.Get("content-type"), "text/plain; charset=utf-8", "response content-type") {
		t.FailNow()
	}
	if blob, err := ioutil.ReadAll(res.Body); err != nil {
		t.Log(err)
		t.FailNow()
	} else {
		if check.NotEqual(t, string(blob), "testing\n", "response body") {
			t.Fail()
		}
	}
}
