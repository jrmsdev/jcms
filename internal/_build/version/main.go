// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/jrmsdev/jcms"
)

func main() {
	v := fmt.Sprintf("%d.%d", jcms.VMAJOR, jcms.VMINOR)
	if jcms.VPATCH > 0 {
		v = fmt.Sprintf("%s.%d", v, jcms.VPATCH)
	}
	rv := strings.TrimSpace(runtime.Version())
	if strings.HasPrefix(rv, "devel") {
		rv = "godev"
	}
	fmt.Printf("%s-%s", v, rv)
}
