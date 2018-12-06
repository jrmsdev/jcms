// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package errors

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jrmsdev/jcms/internal/log"
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

func PathError(path string, x error) Error {
	if e, ok := x.(*os.PathError); ok {
		if e.Op == "read" && e.Err.Error() == "is a directory" {
			log.E("invalid request %s: %s", path, e.Err)
			return InvalidRequest(path)
		}
	}
	return InvalidRequest(path)
}

func IOError(msg string) Error {
	st := http.StatusInternalServerError
	return &err{
		typ: "IOError",
		status: st,
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

func InvalidRequest(path string) Error {
	return &err{
		typ: "InvalidRequest",
		status: http.StatusBadRequest,
		msg: path,
	}
}
