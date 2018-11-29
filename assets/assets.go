// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package assets

import (
	"io"
	"os"
)

type File interface {
	io.ReadSeeker
	io.Closer
}

type Manager interface {
	Open(filename string) (File, error)
	Stat(filename string) (os.FileInfo, error)
	ReadFile(name string) ([]byte, error)
}
