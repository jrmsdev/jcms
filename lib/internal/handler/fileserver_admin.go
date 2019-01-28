// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

// +build jcmsadmin

package handler

func init() {
	admin = true
	htmldir = "./webapp/admin"
}
