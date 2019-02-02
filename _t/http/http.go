// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package http

import (
	x "net/http"
	t "net/http/httptest"
)

func Writer() *t.ResponseRecorder {
	return t.NewRecorder()
}

func GET(path string) (*x.Request, *t.ResponseRecorder) {
	rp := "http://127.0.0.1:666/" + path
	return t.NewRequest("GET", rp, nil), Writer()
}
