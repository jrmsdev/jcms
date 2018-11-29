// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package log

import (
	xlog "log"
	"os"
	"runtime"
	"fmt"
	"path/filepath"
	"strings"
)

var l *xlog.Logger
var lflags = xlog.Ldate | xlog.Ltime

var D func(fmtstr string, args ...interface{})
var E func(fmtstr string, args ...interface{})
var Panic func(fmtstr string, args ...interface{})
var Printf func(fmtstr string, args ...interface{})

var codeInfo bool
var shortIdx int

func init() {
	codeInfo = true
	shortIdx = 0
	D = dummy
	E = dummy
	Panic = dummy
	Printf = dummy
}

func Init(level string) {
	if l == nil {
		l = xlog.New(os.Stderr, "", lflags)
		setLevel(level)
	} else {
		panic("log pkg was initialized already")
	}
}

func setLevel(level string) {
	D = dummy
	E = printf
	Panic = panicf
	Printf = printf
	if level == "debug" {
		D = printf
	} else if level == "quiet" {
		Printf = dummy
	}
}

func shortFile(name string) string {
	if shortIdx == 0 {
		shortIdx = strings.Index(name, "jcms")
		shortIdx += 4 + len(string(filepath.Separator))
	}
	return name[shortIdx:]
}

func dummy(fmtstr string, args ...interface{}) {
}

func printf(fmtstr string, args ...interface{}) {
	prefix := ""
	if codeInfo {
		_, fn, ln, ok := runtime.Caller(1)
		if ok {
			prefix = fmt.Sprintf("%s:%d: ", shortFile(fn), ln)
		}
	}
	l.Printf(prefix+fmtstr, args...)
}

func panicf(fmtstr string, args ...interface{}) {
	prefix := ""
	if codeInfo {
		_, fn, ln, ok := runtime.Caller(1)
		if ok {
			prefix = fmt.Sprintf("%s:%d: ", fn, ln)
		}
	}
	l.Panicf(prefix+fmtstr, args...)
}
