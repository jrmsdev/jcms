// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"testing"

	"github.com/jrmsdev/jcms/internal/_t/check"
)

type Response struct {
	t    *testing.T
	orig *http.Response
	body []byte
}

func newResponse(t *testing.T, r *http.Response) *Response {
	return &Response{t: t, orig: r}
}

func (r *Response) String() string {
	return fmt.Sprintf("%s", r.orig.Status)
}

func (r *Response) Status(expect int) {
	r.t.Helper()
	if check.NotEqual(r.t, r.orig.StatusCode, expect, "response status") {
		r.t.FailNow()
	}
}

func (r *Response) StatusInfo(expect string) {
	r.t.Helper()
	if check.NotEqual(r.t, r.orig.Status, expect, "response status info") {
		r.t.FailNow()
	}
}

func (r *Response) Header(k, expect string) {
	r.t.Helper()
	if check.NotEqual(r.t, r.orig.Header.Get(k),
		expect, "response header "+k) {
		r.t.FailNow()
	}
}

func (r *Response) ReadBody() []byte {
	var err error
	r.t.Helper()
	if r.body == nil {
		r.body, err = ioutil.ReadAll(r.orig.Body)
		if err != nil {
			r.t.Fatal(err)
		}
		r.orig.Body.Close()
	}
	return r.body
}

func (r *Response) getBody() string {
	return strings.TrimSpace(string(r.ReadBody()))
}

func (r *Response) Body(expect string) {
	r.t.Helper()
	if check.NotEqual(r.t, r.getBody(), expect, "response body") {
		r.t.FailNow()
	}
}

func (r *Response) BodyMatch(pat string) {
	r.t.Helper()
	body := r.getBody()
	m, err := regexp.MatchString(pat, body)
	if err != nil {
		r.t.Fatal(err)
	}
	if !m {
		r.t.Fatalf("'%s' not match body '%s'", pat, body)
	}
}

func (r *Response) BodyChecksumMatch(fn string) {
	r.t.Helper()
	if check.NotFileChecksum(r.t, r.ReadBody(), fn) {
		r.t.FailNow()
	}
}

func (r *Response) ContentType(typ string) {
	r.t.Helper()
	r.Header("content-type", typ)
}

func (r *Response) Check(status int, ctyp string) {
	r.Status(status)
	t := ctyp
	if strings.HasPrefix(t, "text/") {
		t = fmt.Sprintf("%s; charset=utf-8", t)
	}
	r.ContentType(t)
}
