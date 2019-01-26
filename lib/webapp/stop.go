// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package webapp

import (
	"context"
	"time"

	"github.com/jrmsdev/jcms/lib/log"
)

func Stop(w *Webapp) error {
	log.D("stop")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := w.server.Shutdown(ctx)
	if err != nil {
		return err
	}
	return nil
}
