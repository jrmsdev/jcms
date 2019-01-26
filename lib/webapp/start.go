// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package webapp

import (
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	"github.com/jrmsdev/jcms/lib/internal/flags"
	"github.com/jrmsdev/jcms/lib/log"
)

func Start(w *Webapp) error {
	log.D("start")
	if w.listener == nil {
		return errors.New("nil webapp listener")
	}
	// TODO: if w.haserr then serve error
	w.server = initServer(w.router)
	trapSignals(w)
	return w.server.Serve(w.listener)
}

func initServer(rtr *mux.Router) *http.Server {
	log.D("init server")
	return &http.Server{
		Handler:        rtr,
		Addr:           "127.0.0.1:" + flags.HttpPort,
		WriteTimeout:   10 * time.Second,
		ReadTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func trapSignals(w *Webapp) {
	log.D("trapSignals")
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		s := <-sig
		log.Printf("got os signal: %s", s)
		Stop(w)
	}()
}
