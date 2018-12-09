// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package fsdb

type FSDB struct {
	wapp    string
	datadir string
}

func New(name, datadir string) *FSDB {
	return &FSDB{name, datadir}
}

func (e *FSDB) String() string {
	return "FSDB " + e.wapp
}
