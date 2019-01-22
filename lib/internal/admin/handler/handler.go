// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package handler

import (
	"fmt"

	"github.com/gorilla/mux"

	"github.com/jrmsdev/jcms/lib/internal/admin/api"
	"github.com/jrmsdev/jcms/lib/log"
)

var sprintf = fmt.Sprintf

func Setup(r *mux.Router) {
	log.D("handler setup: zipmode(%t)", zipmode)
	api.Setup(r)
	if zipmode {
		setupZipServer(r)
	} else {
		setupFileServer(r)
	}
}
