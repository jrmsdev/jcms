// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package schema

import (
	"bytes"
	"text/scanner"
	"fmt"
	"strconv"
	"errors"

	"github.com/jrmsdev/jcms/assets"
	"github.com/jrmsdev/jcms/internal/log"
)

type Data map[float32]Cmd
type Cmd map[string]Table
type Table map[string]Info
type Info map[string]string

func newData() Data {
	return make(Data)
}

type cmdFunc func(Table, *scanner.Scanner) error
var cmdToken = map[string]cmdFunc{
	"create": createTable,
	"index":  dummy,
	"int":    dummy,
	"string": dummy,
	"join":   dummy,
	"remove": dummy,
}

func parse(s *Schema) error {
	log.D("parse %s", s)
	if blob, err := assets.ReadFile("db.schema"); err != nil {
		return err
	} else {
		var (
			x   scanner.Scanner
			cur float32
		)
		x.Init(bytes.NewReader(blob))
		x.Filename = "db.schema"
		x.Whitespace ^= 1<<'\n'
		for tok := x.Scan(); tok != scanner.EOF; tok = x.Scan() {
			typ := scanner.TokenString(tok)
			text := x.TokenText()
			fmt.Printf("%s(%s): %s\n", x.Position, typ, text)
			if typ == "Float" {
				if val, err := strconv.ParseFloat(text, 32); err != nil {
					return err
				} else {
					cur = float32(val)
					log.Printf("parse init %f", cur)
					if cmd, err := readStatement(&x); err != nil {
						return err
					} else {
						s.Data[cur] = cmd
					}
				}
			} else if tok == '\n' {
				continue
			} else {
				return errors.New(sprintf("unkown token %s: %s(%s)",
					x.Position, typ, text))
			}
		}
	}
	return nil
}

func readStatement(x *scanner.Scanner) (Cmd, error) {
	var cmd string
	for tok := x.Scan(); tok != ':'; tok = x.Scan() {
		typ := scanner.TokenString(tok)
		text := x.TokenText()
		fmt.Printf("READ: %s(%s)\n", text, typ)
		if typ == "Ident" {
			cmd = text
			_, ok := cmdToken[cmd]
			if !ok {
				return nil, errors.New(sprintf("invalid action token: %s", cmd))
			}
		} else {
			return nil, errors.New(sprintf("invalid type token: %s(%s)", typ, cmd))
		}
	}
	t := make(Table)
	f := cmdToken[cmd]
	if err := f(t, x); err != nil {
		return nil, err
	}
	r := make(Cmd)
	r[cmd] = t
	return r, nil
}

func dummy(t Table, x *scanner.Scanner) error {
	for tok := x.Scan(); tok != '\n'; tok = x.Scan() {
		continue
	}
	return nil
}

func createTable(t Table, x *scanner.Scanner) error {
	for tok := x.Scan(); tok != '\n'; tok = x.Scan() {
		typ := scanner.TokenString(tok)
		tbl := x.TokenText()
		fmt.Printf("CREATE: %s %s\n", typ, tbl)
		if typ == "Ident" {
			t[tbl] = nil
		}
	}
	return nil
}
