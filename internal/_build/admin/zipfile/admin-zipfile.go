// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package zipfile

import (
	"os"
	"path"
	fpath "path/filepath"
)

type fdef struct {
	dir     string
	prefix  string
	pattern string
}

var srcfn = fpath.FromSlash("./handler/zipfile.go.in")
var glob = []fdef{
	{"../httpd/handler/lib/", "_lib", "*.css"},
	{"../httpd/handler/lib/", "_lib", "*.js"},
	{"./html/", "", "*.html"},
}

func Gen() {
	println("gen admin zipfile")
	if _, err := os.Stat(srcfn); err != nil {
		panic(err)
	}
	for _, g := range glob {
		files, err := fpath.Glob(fpath.FromSlash(g.dir + g.pattern))
		if err != nil {
			panic(err)
		}
		for _, fn := range files {
			if n, err := fpath.Rel(fpath.FromSlash(g.dir), fn); err != nil {
				panic(err)
			} else {
				if g.prefix != "" {
					n = path.Join(g.prefix, fpath.ToSlash(n))
				}
				println(n)
			}
		}
	}
}
