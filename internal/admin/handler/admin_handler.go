// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package handler

import (
	"fmt"

	"github.com/gorilla/mux"

	"github.com/jrmsdev/jcms/internal/log"
)

var sprintf = fmt.Sprintf

func Setup(r *mux.Router) {
	log.D("handler setup: zipmode(%t)", zipmode)
	if zipmode {
		setupZipServer(r)
	} else {
		setupFileServer(r)
	}
}
