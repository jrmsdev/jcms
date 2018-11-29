// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package router

import (
	"github.com/jrmsdev/jcms/internal/log"

	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	log.D("Init")
	return mux.NewRouter()
}
