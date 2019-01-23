// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package webapp

import (
	"net"

	"github.com/jrmsdev/jcms/lib/internal/flags"
	"github.com/jrmsdev/jcms/lib/log"
)

func Listen(w *Webapp) (string, error) {
	log.D("listen %s port %s", flags.Webapp, flags.HttpPort)
	var err error
	addr := sprintf("127.0.0.1:%s", flags.HttpPort)
	w.listener, err = net.Listen("tcp4", addr)
	if err != nil {
		return "", err
	}
	uri := sprintf("http://%s/", addr)
	log.Printf("%s", uri)
	return uri, nil
}
