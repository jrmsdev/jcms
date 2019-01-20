// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package parser

type Data []*Stmt

func newData() Data {
	return make(Data, 0)
}

func (d Data) Len() int {
	return len(d)
}

func (d Data) Less(i, j int) bool {
	return d[i].priority < d[j].priority
}

func (d Data) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
	d[i].index = i
	d[j].index = j
}

func (d *Data) Push(x interface{}) {
	n := len(*d)
	s := x.(*Stmt)
	s.index = n
	*d = append(*d, s)
}

func (d *Data) Pop() interface{} {
	old := *d
	n := len(old)
	s := old[n-1]
	s.index = -1
	*d = old[0:n-1]
	return s
}

type Stmt struct {
	cmd      *Cmd
	priority float32
	index    int
}

func (s *Stmt) String() string {
	return sprintf("%f %s", s.priority, s.cmd)
}

type Cmd struct {
	name string
	tbl  *Table
}

func newCmd(n string, t *Table) *Cmd {
	return &Cmd{n, t}
}

func (c *Cmd) String() string {
	return c.name
}

type Table struct {
	name string
	info *Info
}

func newTable() *Table {
	return &Table{}
}

type Info struct {
	field string
	args  string
}
