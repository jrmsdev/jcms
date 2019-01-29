// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package log

import (
	"fmt"
	"net/http"
	xlog "log"
	"os"
	"path/filepath"
	"runtime"
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
	Panic = panicf
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
	E = printError
	Printf = printf
	if level == "debug" {
		D = printDebug
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

func getCodeInfo() string {
	if codeInfo {
		_, fn, ln, ok := runtime.Caller(2)
		if ok {
			return fmt.Sprintf("%s:%d: ", shortFile(fn), ln)
		}
	}
	return ""
}

func dummy(fmtstr string, args ...interface{}) {
}

func printf(fmtstr string, args ...interface{}) {
	l.Printf(fmtstr, args...)
}

func panicf(fmtstr string, args ...interface{}) {
	l.Panicf(getCodeInfo()+fmtstr, args...)
}

func printDebug(fmtstr string, args ...interface{}) {
	l.Printf("[D] "+getCodeInfo()+fmtstr, args...)
}

func printError(fmtstr string, args ...interface{}) {
	l.Printf("ERROR "+fmtstr, args...)
}

func Response(r *http.Request, size int64) {
	printf("sent: %s %d bytes", r.URL.Path, size)
}
