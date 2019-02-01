// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package handler

type headerTest struct {
	path   string
	key    string
	expect string
}

var ht = []headerTest{
	{"/",           "content-type", "application/octet-stream"},
	{"/index.html", "content-type", "text/html; charset=utf-8"},
	{"/jcms.json",  "content-type", "application/json"},
	{"/jcms.js",  "content-type", "application/javascript; charset=utf-8"},
	{"/jcms.css",  "content-type", "text/css; charset=utf-8"},
	{"/jcms.ico",  "content-type", "image/vnd.microsoft.icon"},
}
