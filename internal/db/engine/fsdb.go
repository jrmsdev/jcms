// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package engine

type FSDB struct {
	wapp    string
	datadir string
}

func newFSDB(name, datadir string) *FSDB {
	return &FSDB{name, datadir}
}

func (e *FSDB) String() string {
	return "FSDB " + e.wapp
}
