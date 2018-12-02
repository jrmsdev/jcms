#!/usr/bin/env python

import os
import sys
from subprocess import call

def _print(s):
	print(s)
	sys.stdout.flush()

def _exit(rc):
	if rc != 0:
		_print("check failed!")
	sys.exit(rc)

verbose = ""
race = ""
for a in sys.argv:
	if a == "-v":
		verbose = " -v"
	elif a == "-race":
		race = " -race"

tests = os.getenv("JCMS_TEST", "").split(",")
if "race" in tests and race == "":
	race = " -race"

prevcmd = {
	10: "go vet ./...",
	20: "go install -i ./cmd/jcms",
	30: "go get -v -t ./...",
}
gotest = "go test{}{} ./...".format(verbose, race)

for idx in sorted(prevcmd.keys()):
	cmd = prevcmd[idx]
	_print(cmd)
	rc = call(cmd.split())
	if rc != 0:
		_exit(rc)

_print(gotest)
rc = call(gotest.split())
if rc != 0:
	_exit(rc)

_exit(0)
