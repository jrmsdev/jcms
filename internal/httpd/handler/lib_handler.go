// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

//go:generate python lib_generate.py

package handler

import (
	"fmt"
	"encoding/base64"
	"net/http"
	"path"

	"github.com/jrmsdev/jcms/internal/errors"
	"github.com/jrmsdev/jcms/internal/log"

	"github.com/gorilla/mux"
)

var sprintf = fmt.Sprintf

func setupLib(r *mux.Router) {
	log.D("setupLib")
	if r.Get("_lib") == nil {
		r.PathPrefix("/_lib/").Handler(http.StripPrefix("/_lib",
			newFileServer("_lib"))).Name("_lib")
	}
}

func libReadFile(fp string) ([]byte, errors.Error) {
	log.D("libReadFile %s", fp)
	errp := path.Join("/", fp)
	var (
		body []byte
		err error
	)
	encBody, found := libFiles[fp]
	if found {
		body, err = base64.StdEncoding.DecodeString(encBody)
		if err != nil {
			log.E("lib read file %s: %s", fp, err)
			return nil, errors.IOError(sprintf("%s: %s", errp, err))
		}
	} else {
		log.E("lib file %s: not found", fp)
		return nil, errors.FileNotFound(errp)
	}
	return body, nil
}
