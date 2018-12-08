// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

//go:generate python lib_generate.py

package handler

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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
			NewFileServer("_lib"))).Name("_lib")
	}
}

func libReadFile(fp string) ([]byte, errors.Error) {
	log.D("libReadFile %s", fp)
	errp := path.Join("/", fp)
	var (
		body []byte
		err  error
	)
	encBody, found := libFiles[fp]
	if found {
		body, err = base64.StdEncoding.DecodeString(encBody)
		if err != nil {
			log.E("lib read file %s: %s", fp, err)
			return nil, errors.IOError(errp, "base64 decode")
		}
	} else {
		log.E("lib file %s: not found", fp)
		return nil, errors.FileNotFound(errp)
	}
	return body, nil
}

func libDevelReadFile(fp string) ([]byte, errors.Error) {
	log.D("libDevelReadFile %s", fp)
	errp := path.Join("/", fp)
	var (
		fh   *os.File
		body []byte
		err  error
	)
	fh, err = os.Open(fp)
	if err != nil {
		log.E("devel lib file %s: not found", fp)
		return nil, errors.FileNotFound(errp)
	}
	defer fh.Close()
	body, err = ioutil.ReadAll(fh)
	if err != nil {
		log.E("devel lib read file %s: %s", fp, err)
		return nil, errors.IOError(errp, err.Error())
	}
	return body, nil
}
