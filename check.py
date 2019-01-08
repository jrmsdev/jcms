#!/usr/bin/env python

import os
import sys
from subprocess import call, check_output

def _print(s):
	print(s)
	sys.stdout.flush()

def _exit(rc):
	if rc != 0:
		_print("check failed!")
	sys.exit(rc)

install_args = " -i"
goversion = check_output("go version".split()).strip()
if "1.9" in goversion:
	install_args = ""

verbose = ""
race = ""
test_only = False
for a in sys.argv:
	if a == "-v":
		verbose = " -v"
	elif a == "-race":
		race = " -race"
	elif a == "-test":
		test_only = True

tests = os.getenv("JCMS_TEST", "").split(",")
if "race" in tests and race == "":
	race = " -race"

prevcmd = {
	0: "go generate ./...",
}

if not test_only:
	prevcmd[10] = "go vet ./..."
	prevcmd[20] = "go install{} ./cmd/jcms".format(install_args)
	prevcmd[30] = "go get -v -t ./..."

for idx in sorted(prevcmd.keys()):
	cmd = prevcmd[idx]
	_print(cmd)
	rc = call(cmd.split())
	if rc != 0:
		_exit(rc)

gotest = "go test{}{} ./...".format(verbose, race)
_print(gotest)
rc = call(gotest.split())
if rc != 0:
	_exit(rc)

_exit(0)
