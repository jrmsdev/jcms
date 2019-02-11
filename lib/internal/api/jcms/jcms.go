// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jcms

import (
	"net/http"

	"github.com/jrmsdev/jcms"
	"github.com/jrmsdev/jcms/lib/internal/api/response"
	"github.com/jrmsdev/jcms/lib/internal/request"
	"github.com/jrmsdev/jcms/lib/log"
)

type resp struct {
	Version string `json:"jcms.version"`
}

func newResp() *resp {
	return &resp{
		Version: jcms.Version(),
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	req := request.New(r)
	log.D("handler %s", req.Path())
	response.Send(w, req, newResp())
}
