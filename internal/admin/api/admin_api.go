// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package api

import (
	"github.com/gorilla/mux"

	"github.com/jrmsdev/jcms/internal/admin/api/jcms"
	"github.com/jrmsdev/jcms/internal/log"
)

func Setup(r *mux.Router) {
	log.D("Setup")
	r.HandleFunc("/_/jcms.json", jcms.Handler)
}
