// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package handler

type headerTest struct {
	path   string
	key    string
	expect string
}

var ht = []headerTest{
	{"/",       "content-type", "application/octet-stream"},
	{"/t.html", "content-type", "text/html; charset=utf-8"},
	{"/t.json", "content-type", "application/json"},
	{"/t.js",   "content-type", "application/javascript; charset=utf-8"},
	{"/t.css",  "content-type", "text/css; charset=utf-8"},
	{"/t.ico",  "content-type", "image/vnd.microsoft.icon"},
}
