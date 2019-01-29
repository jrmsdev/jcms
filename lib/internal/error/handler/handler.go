// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

 package handler

 import (
	"net/http"
)

func Error(w http.ResponseWriter, msg string, status int) {
	http.Error(w, msg, status)
}
