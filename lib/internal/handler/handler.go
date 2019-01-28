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
var devel bool

func init() {
	devel = true
}

func Setup(r *mux.Router) {
	log.D("handler setup: admin(%t) devel(%t)", adminSetup != nil, devel)
	r.Host("127.0.0.1")
	api.Setup(r)
	if devel {
		develFileServer(r)
	} else {
		setupZipServer(r)
	}
	setupFileServer(r)
}
