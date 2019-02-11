// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package log

import (
	"bytes"
	"fmt"
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

var callerSkip int
var codeInfo bool
var shortIdx int
var quiet bool

func init() {
	callerSkip = 2
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
	quiet = false
	D = dummy
	E = printError
	Printf = printf
	if level == "debug" {
		D = printDebug
	} else if level == "quiet" {
		Printf = dummy
		quiet = true
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
		_, fn, ln, ok := runtime.Caller(callerSkip)
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

// log response

type reqInfo interface {
	Path() string
}

func Response(r reqInfo, size int64) {
	if !quiet {
		l.Printf("%s %d bytes", r.Path(), size)
	}
}

// testing mode

var lbuf []byte
var tbuf *bytes.Buffer

func testlog(tag, fmtstr string, args ...interface{}) {
	m := getCodeInfo()
	if tag != "" {
		m += fmt.Sprintf("[%s] ", tag)
	}
	m += fmt.Sprintf(fmtstr, args...)
	m += "\n"
	c := []byte(m)
	writeLog(c)
}

func writeLog(c []byte) {
	fh, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0640)
	if err != nil {
		panic(err)
	}
	if _, err := fh.Write(c); err != nil {
		panic(err)
	}
	if err := fh.Close(); err != nil {
		panic(err)
	}
}

func InitTest() {
	if l == nil {
		tbuf = bytes.NewBuffer(lbuf)
		l = xlog.New(tbuf, "", lflags)
		setLevel("quiet")
	} else {
		if tbuf == nil {
			panic("log testing mode was not initialized")
		}
	}
	callerSkip = 3
	codeInfo = true
	D = testD
	E = testE
	Panic = testPanic
	Printf = testPrintf
	tbuf.Reset()
	if _, err := os.Stat("test.log"); err == nil {
		if err := os.Remove("test.log"); err != nil {
			panic(err)
		}
	}
	testlog("test", "init")
}

func testD(fmtstr string, args ...interface{}) {
	testlog("D", fmtstr, args...)
}

func testE(fmtstr string, args ...interface{}) {
	testlog("E", fmtstr, args...)
}

func testPanic(fmtstr string, args ...interface{}) {
	testlog("Panic", fmtstr, args...)
}

func testPrintf(fmtstr string, args ...interface{}) {
	testlog("", fmtstr, args...)
}
