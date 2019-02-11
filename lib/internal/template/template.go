// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package template

import (
	"io"
)

func Parse(dst io.Writer, src io.Reader) error {
	_, err := io.Copy(dst, src)
	return err
}
