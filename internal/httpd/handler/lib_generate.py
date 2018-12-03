#!/usr/bin/env python

import os
import sys
import md5
from base64 import b64encode
from time import asctime, gmtime
from subprocess import check_output

W3JS = "https://www.w3schools.com/lib/w3.js"
W3CSS = "https://www.w3schools.com/w3css/4/w3.css"

def _now():
	return "%s UTC" % asctime(gmtime())

def _print(s):
	print(s)
	sys.stdout.flush()

def _exit(rc):
	if rc != 0:
		_print("lib generate failed!")
	sys.exit(rc)

def _call(cmd):
	_print(cmd)
	rc = os.system(cmd)
	if rc != 0:
		_exit(rc)

_cwd = os.getcwd()

def _path(fn):
	return os.path.join(_cwd, fn)

def _load(fn):
	_print("    load %s" % fn)
	with open(fn, "r") as fh:
		return b64encode(fh.read())

def _md5(s):
	m = md5.new()
	m.update(s)
	return m.hexdigest()

def _genDone():
	check = (
		"lib_files.go.in",
		"lib/w3.js",
		"lib/w3.css",
	)
	s = ""
	for fn in check:
		with open(fn, "r") as fh:
			s = "%s%s %s " % (s, _md5(fh.read()), fn)
			fh.close()
	x = _md5(s)
	_call("echo done >.gen.%s" % x)
	_call("echo %s >.gen.done" % x)

def _checkDone():
	if not os.path.isfile(".gen.done"):
		return False
	with open(".gen.done", "r") as fh:
		return os.path.isfile(".gen.%s" % fh.readline().strip())

def _gen():
	if _checkDone():
		return
	dst = _path("lib_files.go")
	_print("generate %s" % dst)
	with open(_path("lib_files.go.in"), "r") as src:
		with open(dst, "w") as fh:
			s = src.read().\
				replace("[[LIB_W3JS]]", _load("lib/w3.js"), 1).\
				replace("[[LIB_W3CSS]]", _load("lib/w3.css"), 1).\
				replace("[[GEN_DATE]]", _now(), 1)
			fh.write(s)
			fh.close()
		src.close()
	_genDone()

if "--update" in sys.argv:
	_call("wget -nv -c -O lib/w3.js %s" % W3JS)
	_call("wget -nv -c -O lib/w3.css %s" % W3CSS)

_gen()
_exit(0)
