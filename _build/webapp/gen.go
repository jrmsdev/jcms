// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"github.com/jrmsdev/jcms/_build/webapp/zipfile"
)

// webapp files
var glob = []zipfile.Glob{
	// _lib
	{"./",
		[]string{"_lib/*.css", "_lib/*.js"}},
	// html files
	{"./html/",
		[]string{"*.html", "inc/*.html"}},
}

// admin files
var adminGlob = []zipfile.Glob{
	// _lib
	{"./",
		[]string{"_lib/*.css", "_lib/*.js"}},
	// webapp html files
	{"./html/",
		[]string{"inc/*.html"}},
	// admin html files
	{"./admin/",
		[]string{"*.html", "_admin/*.js"}},
}

func main() {
	zipfile.Gen("webapp", glob)
	zipfile.Gen("admin", adminGlob)
}
