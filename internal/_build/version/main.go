// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"fmt"

	"github.com/jrmsdev/jcms"
)

func main() {
	v := fmt.Sprintf("%d.%d", jcms.VMAJOR, jcms.VMINOR)
	if jcms.VPATCH > 0 {
		v = fmt.Sprintf("%s.%d", v, jcms.VPATCH)
	}
	fmt.Println(v)
}
