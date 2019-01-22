// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/jrmsdev/jcms/lib/internal/admin/api"
	"github.com/jrmsdev/jcms/lib/internal/flags"
	"github.com/jrmsdev/jcms/lib/log"
)

var sprintf = fmt.Sprintf

func Setup() *http.Server {
	log.D("handler setup: zipmode(%t)", zipmode)
	r := mux.NewRouter()
	r.Host("127.0.0.1")
	api.Setup(r)
	if zipmode {
		setupZipServer(r)
	} else {
		setupFileServer(r)
	}
	return initServer(r, flags.HttpPort)
}

func initServer(rtr *mux.Router, port string) *http.Server {
	log.D("init server")
	return &http.Server{
		Handler:        rtr,
		Addr:           "127.0.0.1:" + port,
		WriteTimeout:   10 * time.Second,
		ReadTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
