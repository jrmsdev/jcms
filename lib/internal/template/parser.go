// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package template

import (
	htpl "html/template"
	"io"
	"io/ioutil"

	"github.com/jrmsdev/jcms/lib/log"
)

func Get(path string) string {
	return cfg.Get(path)
}

func Parse(dst io.Writer, src io.Reader, tpl io.ReadCloser) error {
	log.D("parse")
	if cfg == nil {
		log.Panic("nil template cfg")
	}
	if tpl != nil {
		return parseTpl(dst, src, tpl)
	}
	_, err := io.Copy(dst, src)
	return err
}

func parseTpl(dst io.Writer, src io.Reader, tpl io.ReadCloser) error {
	var (
		err   error
		main  *htpl.Template
		child *htpl.Template
		blob  []byte
	)
	defer tpl.Close()
	blob, err = ioutil.ReadAll(tpl)
	if err != nil {
		log.E("read main template %s", err)
		return err
	}
	main, err = htpl.New("main").Funcs(utils).Parse(string(blob))
	if err != nil {
		log.E("parse main template %s", err)
		return err
	}
	blob, err = ioutil.ReadAll(src)
	if err != nil {
		log.E("read template %s", err)
		return err
	}
	child, err = htpl.Must(main.Clone()).Parse(string(blob))
	if err != nil {
		log.E("parse template %s", err)
		return err
	}
	return child.Execute(dst, nil)
}
