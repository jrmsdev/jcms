#!/usr/bin/env python

import os
import sys
from time import asctime, gmtime

def _print(s):
	print(s)
	sys.stdout.flush()

def _printerr(s):
	sys.stderr.write("%s\n" % s)
	sys.stderr.flush()

try:
	fn = sys.argv[1]
except IndexError:
	fn = "/etc/mime.types"

types = dict()
tn = 0
exts = dict()

def _load(typ, e):
	global tn
	if types.get(typ, 0) > 0:
		_printerr("duplicate mime type: %s" % typ)
		return
	tn += 1
	types[tn] = typ
	for x in e:
		e_tn = exts.get(x, 0)
		if e_tn > 0:
			regas = types[e_tn]
			_printerr("duplicate extension: %s for %s" % (x, typ))
			_printerr("                     %s registered as %s" % (x, regas))
			continue
		exts[x] = tn

with open(fn, "r") as fh:
	for l in fh.readlines():
		l = l.strip()
		if l == "":
			continue
		elif l.startswith("#"):
			continue
		typ = l.split()[0]
		elist = l.split()[1:]
		elen = len(elist)
		if elen > 0:
			_load(typ, elist)
	fh.close()

_print("// generated with %s %s" % (sys.argv[0], fn))
_print("")
_print("package mime")

_print("")
_print("var dbTypes = map[int]string{")
for tn in sorted(types.keys()):
	_print('\t%d: "%s",' % (tn, types[tn]))
_print("}")

_print("")
_print("var dbExts = map[string]int{")
for x in sorted(exts.keys()):
	_print('\t".%s": %d,' % (x, exts[x]))
_print("}")
