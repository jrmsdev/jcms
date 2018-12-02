#!/usr/bin/env python

import os
import sys
from subprocess import check_output

BUILDS = (
	("linux",   ("386", "amd64", "arm")),
	("freebsd", ("386", "amd64", "arm")),
	("openbsd", ("386", "amd64", "arm")),
	("darwin",  ("386", "amd64")),
	("windows", ("386", "amd64")),
)

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

version = check_output("go run ./internal/_build/version/main.go".split()).strip()

_call("rm -rfv build")
_call("mkdir -v build")

_call("go vet ./...")
_call("go test ./...")

for b in BUILDS:
	goos = b[0]
	os.environ["GOOS"] = goos
	for goarch in b[1]:
		os.environ["GOARCH"] = goarch
		cmd = "go build -o build/jcms-{}-{}-{}.bin ./cmd/jcms".format(version, goos, goarch)
		_call(cmd)

_exit(0)
