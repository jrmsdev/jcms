// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package template

import (
	htpl "html/template"
	"io"
	"io/ioutil"
	"path/filepath"

	"github.com/jrmsdev/jcms/lib/internal/asset"
	"github.com/jrmsdev/jcms/lib/log"
)

func Parse(dst io.Writer, src io.Reader, path string) error {
	log.D("parse %s", path)
	if cfg == nil {
		log.Panic("nil template cfg")
	}
	tpl := cfg.Get(path)
	if tpl != "" {
		return parseTpl(dst, src, tpl)
	}
	_, err := io.Copy(dst, src)
	return err
}

func parseTpl(dst io.Writer, src io.Reader, n string) error {
	var (
		err   error
		main  *htpl.Template
		child *htpl.Template
		fh    io.ReadCloser
		mblob []byte
		cblob []byte
	)
	fn := filepath.Join("tpl", n+".html")
	log.D("parse template %s", fn)
	fh, err = asset.Open(fn)
	if err != nil {
		log.E("open template %s: %s", fn, err)
		return err
	}
	defer fh.Close()
	mblob, err = ioutil.ReadAll(fh)
	if err != nil {
		log.E("read template %s: %s", fn, err)
		return err
	}
	main, err = htpl.New(n).Funcs(utils).Parse(string(mblob))
	if err != nil {
		log.E("parse template %s: %s", fn, err)
		return err
	}
	cblob, err = ioutil.ReadAll(src)
	if err != nil {
		log.E("read template %s", err)
		return err
	}
	child, err = htpl.Must(main.Clone()).Parse(string(cblob))
	if err != nil {
		log.E("parse template %s", err)
		return err
	}
	return child.Execute(dst, nil)
}
