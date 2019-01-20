// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package parser

import (
	"bytes"
	"container/heap"
	"errors"
	"fmt"
	"strconv"
	"text/scanner"

	"github.com/jrmsdev/jcms/assets"
	"github.com/jrmsdev/jcms/internal/log"
)

var sprintf = fmt.Sprintf

type cmdFunc func(*Table, *scanner.Scanner) error

var cmdToken = map[string]cmdFunc{
	"create": createTable,
	"index":  dummy,
	"int":    dummy,
	"string": dummy,
	"join":   dummy,
	"remove": dummy,
}

func Parse() (Data, error) {
	log.D("Parse")
	d := newData()
	heap.Init(&d)
	if blob, err := assets.ReadFile("db.schema"); err != nil {
		return nil, err
	} else {
		var (
			x   scanner.Scanner
			cur float32
		)
		x.Init(bytes.NewReader(blob))
		x.Filename = "db.schema"
		x.Whitespace ^= 1 << '\n'
		idx := 0
		for tok := x.Scan(); tok != scanner.EOF; tok = x.Scan() {
			typ := scanner.TokenString(tok)
			text := x.TokenText()
			if typ == "Float" {
				if val, err := strconv.ParseFloat(text, 32); err != nil {
					return nil, err
				} else {
					cur = float32(val)
					log.D("parse init %f", cur)
					if cmd, err := readStatement(&x); err != nil {
						return nil, err
					} else {
						stmt := &Stmt{
							cmd:      cmd,
							priority: cur,
							index:    idx,
						}
						heap.Push(&d, stmt)
						idx++
					}
				}
			} else if tok == '\n' {
				continue
			} else {
				return nil, errors.New(sprintf("unkown token %s: %s(%s)",
					x.Position, typ, text))
			}
		}
	}
	return d, nil
}

func readStatement(x *scanner.Scanner) (*Cmd, error) {
	var cmd string
	for tok := x.Scan(); tok != ':'; tok = x.Scan() {
		typ := scanner.TokenString(tok)
		text := x.TokenText()
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
	t := newTable()
	f := cmdToken[cmd]
	if err := f(t, x); err != nil {
		return nil, err
	}
	return newCmd(cmd, t), nil
}

func dummy(t *Table, x *scanner.Scanner) error {
	for tok := x.Scan(); tok != '\n'; tok = x.Scan() {
		continue
	}
	return nil
}

func createTable(t *Table, x *scanner.Scanner) error {
	for tok := x.Scan(); tok != '\n'; tok = x.Scan() {
		typ := scanner.TokenString(tok)
		tbl := x.TokenText()
		if typ == "Ident" {
			t.name = tbl
			t.info = nil
		}
	}
	return nil
}
