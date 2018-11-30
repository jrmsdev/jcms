// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package webapp

type Webapp struct {
	name string
}

func New(name string) *Webapp {
	return &Webapp{name}
}
