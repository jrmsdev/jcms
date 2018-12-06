// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package errors

import (
	"fmt"
	"net/http"
)

var sprintf = fmt.Sprintf

type Error interface {
	Error() string
	WriteResponse(http.ResponseWriter)
}

type err struct {
	typ    string
	status int
	msg    string
}

func (e *err) Error() string {
	return sprintf("%s %s", e.typ, e.msg)
}

func (e *err) WriteResponse(w http.ResponseWriter) {
	http.Error(w, sprintf("%s %s", e.typ, e.msg), e.status)
}

func IOError(msg string) Error {
	return &err{
		typ: "IOError",
		status: http.StatusInternalServerError,
		msg: msg,
	}
}

func FileNotFound(name string) Error {
	return &err{
		typ: "FileNotFound",
		status: http.StatusNotFound,
		msg: name,
	}
}
