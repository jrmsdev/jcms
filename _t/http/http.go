// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package http

import (
	x "net/http"
	t "net/http/httptest"
)

func Writer() x.ResponseWriter {
	return t.NewRecorder()
}

func GET(path string) (*x.Request, x.ResponseWriter) {
	rp := "http://127.0.0.1:666/" + path
	return t.NewRequest("GET", rp, nil), t.NewRecorder()
}
