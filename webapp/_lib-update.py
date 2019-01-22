#!/usr/bin/env python

import os
import os.path
import sys

W3JS = "https://www.w3schools.com/lib/w3.js"
W3CSS = "https://www.w3schools.com/w3css/4/w3.css"

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

def _write(n):
	orig = os.path.join("_lib", ".orig.%s" % n)
	dst = os.path.join("_lib", n)
	with open(orig, "r") as src:
		with open(dst, "w") as fh:
			for l in src.readlines():
				# unify line endinds (LF)
				l = l.rstrip()
				fh.write("%s\n" % l)

_call("wget -nv -c -O _lib/.orig.w3.js %s" % W3JS)
_write("w3.js")

_call("wget -nv -c -O _lib/.orig.w3.css %s" % W3CSS)
_write("w3.css")

_exit(0)
