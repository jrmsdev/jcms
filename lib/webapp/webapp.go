// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package webapp

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"

	"github.com/jrmsdev/jcms"
	"github.com/jrmsdev/jcms/lib/internal/flags"
	"github.com/jrmsdev/jcms/lib/log"
)

var sprintf = fmt.Sprintf
var fprintf = fmt.Fprintf

type Webapp struct {
	admin    bool
	listener net.Listener
	router   *mux.Router
	server   *http.Server
}

func wapp() *Webapp {
	cmd := filepath.Base(os.Args[0])
	flags.Parse()
	log.Init(flags.Log)
	if ! flags.ShowVersion {
		log.Printf("%s version %s", cmd, jcms.Version())
	}
	return &Webapp{}
}

func New() *Webapp {
	return setup(wapp())
}

func Admin() *Webapp {
	w := wapp()
	w.admin = true
	return setup(w)
}
