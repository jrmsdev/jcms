// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jcms

import (
	"fmt"
	"runtime"
)

const (
	VMAJOR = 0
	VMINOR = 0
	VPATCH = 0
)

func Version() string {
	v := fmt.Sprintf("%d.%d", VMAJOR, VMINOR)
	if VPATCH > 0 {
		v = fmt.Sprintf("%s.%d", v, VPATCH)
	}
	return fmt.Sprintf("%s %s/%s", v, runtime.GOOS, runtime.GOARCH)
}
