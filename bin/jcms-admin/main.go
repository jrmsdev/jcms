// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"github.com/jrmsdev/jcms/lib/webapp"
)

func main() {
	wapp := webapp.Admin()
	webapp.Main(wapp)
}
