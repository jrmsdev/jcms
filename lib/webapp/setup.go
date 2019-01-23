// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package webapp

import (
	"github.com/gorilla/mux"

	"github.com/jrmsdev/jcms/lib/internal/handler/admin"
)

func setup(w *Webapp) *Webapp {
	w.router = mux.NewRouter()
	if w.admin {
		admin.Setup(w.router)
	}
	return w
}
