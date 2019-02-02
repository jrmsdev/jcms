// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package response

import (
	"io/ioutil"
	"testing"

	"github.com/jrmsdev/jcms/_t/check"
	"github.com/jrmsdev/jcms/_t/http"
)

type headerTest struct {
	path   string
	key    string
	expect string
}

var ht = []headerTest{
	{"/",           "content-type", "application/octet-stream"},
	{"/index.html", "content-type", "text/html; charset=utf-8"},
	{"/jcms.json",  "content-type", "application/json"},
}

func TestHeaders(t *testing.T) {
	w := http.Writer()
	for _, x := range ht {
		setHeaders(w, x.path)
		if check.NotEqual(t, w.Header().Get(x.key), x.expect, x.path + " " + x.key) {
			t.Fail()
		}
	}
}

type jsonTest struct {
	Testing string      `json:"testing"`
	status  int
	Data    interface{} `json:"testdata"`
}

var jt = []jsonTest{
	{"test nil", 200, nil},
	{"test", 200, "testing"},
}

func TestSend(t *testing.T) {
	for _, x := range jt {
		r, w := http.GET("/t.json")
		Send(w, r, x)
		res := w.Result()
		//~ t.Log(res)
		if check.NotEqual(t, res.StatusCode, x.status, "response status") {
			t.FailNow()
		}
		if blob, err := ioutil.ReadAll(res.Body); err != nil {
			t.Log(err)
			t.FailNow()
		} else {
			//~ t.Log(string(blob))
			if check.NotJSON(t, blob, "testdata", x.Data, x.Testing) {
				t.Fail()
			}
		}
	}
}
