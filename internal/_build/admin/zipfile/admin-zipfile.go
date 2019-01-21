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
	"path"
	fpath "path/filepath"
	"strings"
	"time"
)

type fdef struct {
	dir     string
	prefix  string
	pattern string
}

var sprintf = fmt.Sprintf
var zbuf = new(bytes.Buffer)
var z = zip.NewWriter(zbuf)
var b64 = base64.StdEncoding.EncodeToString

var glob = []fdef{
	{"../httpd/handler/lib/", "_lib", "*.css"},
	{"../httpd/handler/lib/", "_lib", "*.js"},
	{"./html/", "", "*.html"},
}

var (
	dstfn string
	srcfn string
)

func init() {
	var err error
	dstfn, err = fpath.Abs(fpath.FromSlash("./handler/zipfile.go"))
	if err != nil {
		panic(err)
	}
	srcfn, err = fpath.Abs(fpath.FromSlash("./handler/zipfile.go.in"))
	if err != nil {
		panic(err)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func Gen() {
	_, err := os.Stat(srcfn)
	check(err)
	println("generate " + dstfn)
	for _, g := range glob {
		dir := fpath.FromSlash(g.dir)
		files, err := fpath.Glob(dir + g.pattern)
		check(err)
		for _, fn := range files {
			n, err := fpath.Rel(dir, fn)
			check(err)
			if g.prefix != "" {
				n = path.Join(g.prefix, fpath.ToSlash(n))
			}
			check(zfile(n, fn))
			println("     zip " + n)
		}
	}
	check(z.Close())
	check(write())
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

func write() error {
	var (
		err  error
		src  []byte
		sbuf *bytes.Buffer
		dbuf *bytes.Buffer
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
	if err := ioutil.WriteFile(dstfn, dbuf.Bytes(), 0640); err != nil {
		return err
	}
	return nil
}

func parse(line string) string {
	l := strings.ToLower(strings.TrimSpace(line))
	//~ println("parse line " + l)
	if strings.HasPrefix(l, "// generated on") {
		return sprintf("// generated on %s\n", time.Now().Format(time.RFC1123Z))
	} else if strings.HasPrefix(l, "zipfile = ") {
		return sprintf("\tzipfile = \"%s\"\n", b64(zbuf.Bytes()))
	}
	return line
}
