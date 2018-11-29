// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package flags

import "flag"

var ShowVersion bool

func init() {
	flag.BoolVar(&ShowVersion, "version", false, "show version")
}

func Parse() {
	flag.Parse()
}
