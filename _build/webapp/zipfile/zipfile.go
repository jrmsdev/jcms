// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package zipfile

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	fpath "path/filepath"
	"strings"
	"time"
)

type Glob struct {
	Dir  string
	Patt []string
}

var sprintf = fmt.Sprintf
var zbuf = new(bytes.Buffer)
var z = zip.NewWriter(zbuf)
var zfiles = make([]string, 0)
var b64 = base64.StdEncoding.EncodeToString

var (
	srcdir     string
	srcfn      string
	buildFlags string
)

func init() {
	var err error
	srcdir, err = fpath.Abs(fpath.FromSlash("../lib/internal/handler"))
	if err != nil {
		panic(err)
	}
	srcfn = fpath.Join(srcdir, "zipfile.go.in")
	buildFlags = "jcms"
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func Gen(id string, glob []Glob) {
	if z == nil {
		zfiles = nil
		z = zip.NewWriter(zbuf)
		zfiles = make([]string, 0)
		zbuf.Reset()
	}
	if id != "webapp" {
		buildFlags = "jcms" + id
	}
	_, err := os.Stat(srcfn)
	check(err)
	dstfn := fpath.Join(srcdir, "zipfile_" + id + ".go")
	println("generate " + dstfn)
	for _, g := range glob {
		dir := fpath.FromSlash(g.Dir)
		for _, patt := range g.Patt {
			files, err := fpath.Glob(dir + fpath.FromSlash(patt))
			check(err)
			for _, fn := range files {
				n, err := fpath.Rel(dir, fn)
				check(err)
				check(zfile(n, fn))
				println("     zip " + n)
				zfiles = append(zfiles, n)
			}
		}
	}
	check(z.Close())
	z = nil
	check(write(dstfn))
}

func zfile(name, fn string) error {
	fh, err := z.Create(name)
	if err != nil {
		return err
	}
	var src *os.File
	src, err = os.Open(fn)
	if err != nil {
		return err
	}
	_, err = io.Copy(fh, src)
	if err != nil {
		return err
	}
	err = z.Flush()
	if err != nil {
		return err
	}
	return nil
}

func write(fn string) error {
	var (
		err   error
		src   []byte
		sbuf  *bytes.Buffer
		dbuf  *bytes.Buffer
	)
	src, err = ioutil.ReadFile(srcfn)
	if err != nil {
		return err
	}
	dbuf = new(bytes.Buffer)
	sbuf = bytes.NewBuffer(src)
	for x := 1; x > 0; x++ {
		line, err := sbuf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		_, err = dbuf.WriteString(parse(line))
		if err != nil {
			return err
		}
	}
	_, err = dbuf.WriteString("\n")
	if err != nil {
		return err
	}
	for _, f := range zfiles {
		_, err = dbuf.WriteString(sprintf("// %s\n", f))
		if err != nil {
			return err
		}
	}
	if err := ioutil.WriteFile(fn, dbuf.Bytes(), 0640); err != nil {
		return err
	}
	return nil
}

func parse(line string) string {
	l := strings.ToLower(strings.TrimSpace(line))
	//~ println("parse line " + l)
	if strings.HasPrefix(l, "// generated on") {
		return sprintf("// generated on %s\n", time.Now().Format(time.RFC1123Z))
	} else if strings.HasPrefix(l, "// +build jcms") {
		return sprintf("// +build %s\n", buildFlags)
	} else if strings.HasPrefix(l, "zipfile = ") {
		return sprintf("\tzipfile = \"%s\"\n", b64(zbuf.Bytes()))
	}
	return line
}
