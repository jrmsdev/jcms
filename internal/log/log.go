// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package log

import (
	xlog "log"
	"os"
)

var l *xlog.Logger
var lflags = xlog.Ldate | xlog.Ltime | xlog.Lmicroseconds | xlog.Lshortfile

var Debug func(fmtstr string, args ...interface{})
var Error func(fmtstr string, args ...interface{})
var Printf func(fmtstr string, args ...interface{})

func init() {
	Debug = dummy
	Error = dummy
	Printf = dummy
}

func Init(level string) {
	if l == nil {
		l = xlog.New(os.Stderr, "", lflags)
		setLevel(level)
	}
}

func setLevel(level string) {
	Debug = dummy
	Error = printf
	Printf = printf
	if level == "debug" {
		Debug = printf
	} else if level == "quiet" {
		Printf = dummy
	}
}

func dummy(fmtstr string, args ...interface{}) {
}

func printf(fmtstr string, args ...interface{}) {
	l.Printf(fmtstr, args...)
}
