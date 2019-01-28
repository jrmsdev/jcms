// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package response

import (
	"encoding/json"
	"net/http"
	"path"

	"github.com/jrmsdev/jcms/lib/internal/mime"
	"github.com/jrmsdev/jcms/lib/log"
)

func Send(w http.ResponseWriter, r *http.Request, data interface{}) {
	log.D("send %s", r.URL.Path)
	blob, err := json.MarshalIndent(&data, "", "  ")
	if err != nil {
		log.E("%s", err)
		http.Error(w, "json error", http.StatusInternalServerError)
		return
	}
	setHeaders(w, r.URL.Path)
	if n, err := w.Write(blob); err != nil {
		log.E("%s", err)
	} else {
		log.Printf("sent: %s %d bytes", r.URL.Path, n)
	}
}

func setHeaders(w http.ResponseWriter, rp string) {
	log.D("setHeaders %s", rp)
	w.Header().Set("Content-Type", mime.TypeByExtension(path.Ext(rp)))
}
