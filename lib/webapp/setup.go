// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package webapp

import (
	"github.com/gorilla/mux"

	"github.com/jrmsdev/jcms/lib/internal/handler"
)

func setup(w *Webapp) *Webapp {
	w.router = mux.NewRouter()
	handler.Setup(w.router)
	return w
}
