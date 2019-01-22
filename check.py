#!/usr/bin/env python

import os
import sys
from subprocess import check_output

def _print(s):
	print(s)
	sys.stdout.flush()

def _exit(rc):
	if rc != 0:
		_print("jcms check failed!")
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

verbose = ""
race = ""
coverage = ""
test_only = False
test_coverage = False

for a in sys.argv:
	if a == "-v":
		verbose = " -v"
	elif a == "-race":
		race = " -race"
	elif a == "-test":
		test_only = True
	elif a == "-coverage":
		test_only = True
		test_coverage = True
		coverage = " -coverprofile coverage.out"

tests = os.getenv("JCMS_TEST", "").split(",")
if "race" in tests and race == "":
	race = " -race"

prevcmd = {
	0: "go generate ./webapp/...",
}

if not test_only:
	prevcmd[10] = "go vet ./bin/... ./lib/..."
	prevcmd[20] = "go install{} ./bin/...".format(install_args)
	prevcmd[30] = "go get -v -t ./lib/..."

for idx in sorted(prevcmd.keys()):
	_call(prevcmd[idx])

_call("go test{}{}{} ./lib/...".format(verbose, race, coverage))

if test_coverage:
	_call("go tool cover -html coverage.out -o coverage.html")

_exit(0)
