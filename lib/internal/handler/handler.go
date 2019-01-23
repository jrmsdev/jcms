// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package handler

import (
	"fmt"

	"github.com/gorilla/mux"

	"github.com/jrmsdev/jcms/lib/internal/api"
	"github.com/jrmsdev/jcms/lib/log"
)

var sprintf = fmt.Sprintf

func Setup(r *mux.Router) {
	log.D("handler setup: zipmode(%t)", zipmode)
	r.Host("127.0.0.1")
	api.Setup(r)
	if zipmode {
		setupZipServer(r)
	} else {
		setupFileServer(r)
	}
}

func Admin(r *mux.Router) {
	log.D("admin setup")
}
