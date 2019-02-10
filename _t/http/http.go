// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package http

import (
	"errors"

	x "net/http"
	t "net/http/httptest"
)

type errw struct {
	t.ResponseRecorder
}

func ErrorWriter() *errw {
	return &errw{}
}

func (w *errw) Write(p []byte) (int, error) {
	return -1, errors.New("testing error")
}

func Writer() *t.ResponseRecorder {
	return t.NewRecorder()
}

func Request(path string) *x.Request {
	rp := "http://127.0.0.1:666" + path
	return t.NewRequest("GET", rp, nil)
}

func GET(path string) (*x.Request, *t.ResponseRecorder) {
	return Request(path), Writer()
}
