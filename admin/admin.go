// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package admin

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"

	"github.com/jrmsdev/jcms"
	"github.com/jrmsdev/jcms/internal/admin/handler"
	"github.com/jrmsdev/jcms/internal/cmd/flags"
	"github.com/jrmsdev/jcms/internal/log"
	"github.com/jrmsdev/jcms/webapp/config"
)

func Main() {
	cfg := flags.Parse()
	if flags.ShowVersion {
		fmt.Fprintf(os.Stderr, "jcms-admin version %s\n", jcms.Version())
		os.Exit(0)
	}
	log.Init(cfg.Log)
	log.Printf("jcms-admin version %s", jcms.Version())
	rtr := newRouter()
	handler.Setup(rtr)
	srv := initServer(cfg, rtr)
	if err := srv.ListenAndServe(); err != nil {
		log.E("%s", err)
		os.Exit(2)
	}
}

func newRouter() *mux.Router {
	log.D("init router")
	r := mux.NewRouter()
	return r.Host("127.0.0.1").
		Methods("GET").
		Schemes("http").
		Subrouter().
		StrictSlash(true)
}

func initServer(cfg *config.Config, rtr *mux.Router) *http.Server {
	log.D("init server")
	return &http.Server{
		Handler:        rtr,
		Addr:           "127.0.0.1:" + cfg.HttpPort,
		WriteTimeout:   10 * time.Second,
		ReadTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
