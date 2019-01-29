// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package handler

import (
	"fmt"

	"github.com/gorilla/mux"

	"github.com/jrmsdev/jcms/lib/internal/api"
	errh "github.com/jrmsdev/jcms/lib/internal/error/handler"
	"github.com/jrmsdev/jcms/lib/log"
)

var devel bool
var sprintf = fmt.Sprintf
var errhdlr = errh.Error

func init() {
	devel = true
}

func Setup(r *mux.Router) {
	log.D("handler setup: admin(%t) devel(%t)", admin, devel)
	r.Host("127.0.0.1")
	api.Setup(r)
	if devel {
		develFileServer(r)
	} else {
		setupZipServer(r)
	}
	setupAssetsServer(r)
}
