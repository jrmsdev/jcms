#!/usr/bin/env python

import os
import sys
from time import asctime, gmtime
from subprocess import check_output

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

def _read(fn):
	fn = _path(fn)
	_print("load %s" % fn)
	cmd = "base64 -w 0 %s" % fn
	return check_output(cmd.split())

def _gen():
	dst = _path("lib_files.go")
	_print("generate %s" % dst)
	with open(_path("lib_files.go.in"), "r") as src:
		with open(dst, "w") as fh:
			s = src.read().\
				replace("[[LIB_W3JS]]", _read("lib/w3.js"), 1).\
				replace("[[GEN_DATE]]", _now(), 1)
			fh.write(s)
			fh.close()
		src.close()
	_print("%s done!" % dst)

if "--update" in sys.argv:
	_update()

_gen()
_exit(0)
