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
	redirect string
	r *http.Request
}

func (e *err) Error() string {
	return sprintf("%s %s", e.typ, e.msg)
}

func (e *err) WriteResponse(w http.ResponseWriter) {
	if e.redirect != "" && e.r != nil {
		log.E("redirect %s -> %s", e.msg, e.redirect)
		http.Redirect(w, e.r, e.redirect, e.status)
		return
	}
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

func Redirect(path string, r *http.Request, location string) Error {
	return &err{
		typ: "Redirect",
		status: http.StatusMovedPermanently,
		redirect: location,
		r: r,
		msg: path,
	}
}
