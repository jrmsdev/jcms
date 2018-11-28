// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package log

import (
	xlog "log"
	"os"
)

var l *xlog.Logger
var lflags = xlog.Ldate | xlog.Ltime | xlog.Lshortfile

var D func(fmtstr string, args ...interface{})
var E func(fmtstr string, args ...interface{})
var Panic func(fmtstr string, args ...interface{})
var Printf func(fmtstr string, args ...interface{})

func init() {
	D = dummy
	E = dummy
	Panic = dummy
	Printf = dummy
}

func Init(level string) {
	if l == nil {
		l = xlog.New(os.Stderr, "", lflags)
		setLevel(level)
	}
}

func setLevel(level string) {
	D = dummy
	E = l.Printf
	Panic = l.Panicf
	Printf = l.Printf
	if level == "debug" {
		D = l.Printf
	} else if level == "quiet" {
		Printf = dummy
	}
}

func dummy(fmtstr string, args ...interface{}) {
}
