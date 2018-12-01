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

type TestResponse struct {
	t    *testing.T
	orig *http.Response
}

func newResponse(t *testing.T, r *http.Response) *TestResponse {
	return &TestResponse{t, r}
}

func (r *TestResponse) String() string {
	return fmt.Sprintf("%s", r.orig.Status)
}

func (r *TestResponse) Status(expect int) {
	r.t.Helper()
	if check.NotEqual(r.t, r.orig.StatusCode, expect, "response status") {
		r.t.FailNow()
	}
}

func (r *TestResponse) StatusInfo(expect string) {
	r.t.Helper()
	if check.NotEqual(r.t, r.orig.Status, expect, "response status info") {
		r.t.FailNow()
	}
}

func (r *TestResponse) getBody() string {
	r.t.Helper()
	body, err := ioutil.ReadAll(r.orig.Body)
	if err != nil {
		r.t.Fatal(err)
	}
	r.orig.Body.Close()
	return strings.TrimSpace(string(body))
}

func (r *TestResponse) Body(expect string) {
	r.t.Helper()
	if check.NotEqual(r.t, r.getBody(), expect, "response body") {
		r.t.FailNow()
	}
}

func (r *TestResponse) BodyMatch(pat string) {
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
