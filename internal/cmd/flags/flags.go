// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package flags

import "flag"

var ShowVersion bool
var Debug bool
var Quiet bool

func init() {
	flag.BoolVar(&ShowVersion, "V", false, "show version")
	flag.BoolVar(&Debug, "d", false, "enable debug")
	flag.BoolVar(&Quiet, "q", false, "quiet mode")
}

func Parse() {
	flag.Parse()
}
