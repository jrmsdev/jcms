#!/usr/bin/env python

import os
import sys
import md5
from base64 import b64encode
from time import asctime, gmtime
from glob import glob
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
_libfiles = glob("lib/*.*")

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
	check = ["lib_files.go.in", "lib_files.go"]
	check.extend(_libfiles)
	s = ""
	for fn in check:
		with open(fn, "r") as fh:
			s = "%s%s %s " % (s, _md5(fh.read()), fn)
			fh.close()
	x = _md5(s)
	if os.system("echo done >.gen.%s" % x) != 0:
		_exit(1)
	if os.system("echo %s >.gen.done" % x) != 0:
		_exit(1)

def _checkDone():
	if not os.path.isfile(".gen.done"):
		return False
	with open(".gen.done", "r") as fh:
		return os.path.isfile(".gen.%s" % fh.readline().strip())

def _genFiles(fh):
	fh.write("var libFiles = map[string]string{\n")
	for fn in sorted(_libfiles):
		fn = "/".join(os.path.split(fn))
		fh.write('\t"_%s": "%s",\n' % (fn, _load(fn)))
	fh.write("}\n")

def _gen():
	if _checkDone():
		return
	dst = _path("lib_files.go")
	_print("generate %s" % dst)
	with open(_path("lib_files.go.in"), "r") as src:
		with open(dst, "w") as fh:
			for l in src.readlines():
				l = l.rstrip()
				if l.startswith("// generated on"):
					s = "%s\n" % l.replace("[[GEN_DATE]]", _now(), 1)
					fh.write(s)
				elif l.startswith("var libFiles ="):
					_genFiles(fh)
				else:
					fh.write("%s\n" % l)
			fh.flush()
			fh.close()
		src.close()
	_gofmt()
	_genDone()

def _gofmt():
	if os.system("which gofmt >/dev/null") == 0:
		_call("gofmt -w -s lib_files.go")

def _write(n):
	orig = os.path.join("lib", ".orig.%s" % n)
	dst = os.path.join("lib", n)
	with open(orig, "r") as src:
		with open(dst, "w") as fh:
			for l in src.readlines():
				# unify line endinds (LF)
				l = l.rstrip()
				fh.write("%s\n" % l)

if "--update" in sys.argv:
	if os.system("rm -f .gen.*") != 0:
		_exit(1)
	_call("wget -nv -c -O lib/.orig.w3.js %s" % W3JS)
	_write("w3.js")
	_call("wget -nv -c -O lib/.orig.w3.css %s" % W3CSS)
	_write("w3.css")

_gen()
_exit(0)
