// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package jcms

type Webapp struct {
	Name string
}

func NewWebapp(name string) *Webapp {
	return &Webapp{
		Name: name,
	}
}
