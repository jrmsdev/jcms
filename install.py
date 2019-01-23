#!/usr/bin/env python

import os
import sys
from subprocess import check_output

CMDBIN = ['jcms', 'jcms-admin']

def _print(s):
	print(s)
	sys.stdout.flush()

def _exit(rc):
	if rc != 0:
		_print("jcms build failed!")
		sys.exit(rc)

def _call(cmd):
	_print(cmd)
	rc = os.system(cmd)
	if rc != 0:
		_exit(rc)
	sys.stdout.flush()

install_args = " -i"
goversion = check_output("go version".split()).strip()
if "1.9" in goversion:
	install_args = ""

_call("go generate ./webapp/...")
_call("go vet ./bin/... ./lib/...")

for n in CMDBIN:
	t = n.replace("-", "")
	_call("go install{} -tags {} ./bin/{}".format(install_args, t, n))

_exit(0)
