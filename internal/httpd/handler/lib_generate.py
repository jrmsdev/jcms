#!/usr/bin/env python

import os
import sys
from time import asctime, gmtime

W3JS = "https://www.w3schools.com/lib/w3.js"

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

def _update():
	_call("wget -nv -c -O lib/w3.js %s" % W3JS)

_cwd = os.getcwd()

def _path(fn):
	return os.path.join(_cwd, fn)

def _gen():
	dst = _path("lib_files.go")
	with open(_path("lib_files.go.in"), "r") as src:
		with open(dst, "w") as fh:
			s = src.read().\
				replace("[[LIB_W3JS]]", "testing", 1).\
				replace("[[GEN_DATE]]", _now(), 1)
			fh.write(s)
			fh.close()
		src.close()
	_print("created %s" % dst)

if "--update" in sys.argv:
	_update()

_gen()
_exit(0)
