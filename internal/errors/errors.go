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
	typ      string
	status   int
	msg      string
	redirect string
	r        *http.Request
}

func (e *err) Error() string {
	return sprintf("%s %s", e.typ, e.msg)
}

func (e *err) WriteResponse(w http.ResponseWriter) {
	if e.redirect != "" && e.r != nil {
		http.Redirect(w, e.r, e.redirect, e.status)
		return
	}
	http.Error(w, sprintf("%s %s", e.typ, e.msg), e.status)
}

func PathError(path string, x error) Error {
	if e, ok := x.(*os.PathError); ok {
		log.E("PathError %#T(%d) %s", e.Err, e.Err, e.Err)
		es := sprintf("%d", e.Err)
		if e.Op == "read" {
			// 21 -> unix "is a directory"
			// 6  -> windows "invalid handle"
			if es == "21" || es == "6" {
				return InvalidRequest(path)
			}
		}
	}
	return IOError(path, x.Error())
}

func IOError(path, msg string) Error {
	m := sprintf("%s: %s", path, msg)
	log.E("I/O %s", m)
	return &err{
		typ:    "IOError",
		status: http.StatusInternalServerError,
		msg:    m,
	}
}

func FileNotFound(name string) Error {
	return &err{
		typ:    "FileNotFound",
		status: http.StatusNotFound,
		msg:    name,
	}
}

func InvalidRequest(path string) Error {
	log.E("invalid request %s", path)
	return &err{
		typ:    "InvalidRequest",
		status: http.StatusBadRequest,
		msg:    path,
	}
}

func Redirect(path string, r *http.Request, location string) Error {
	log.E("redirect %s -> %s", path, location)
	return &err{
		typ:      "Redirect",
		status:   http.StatusMovedPermanently,
		redirect: location,
		r:        r,
		msg:      path,
	}
}
