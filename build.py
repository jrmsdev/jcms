#!/usr/bin/env python

import os
import sys
from subprocess import check_output

BUILDS = {
	"linux":   ("386", "amd64", "arm"),
	"freebsd": ("386", "amd64", "arm"),
	"openbsd": ("386", "amd64", "arm"),
	"darwin":  ("386", "amd64"),
	"windows": ("386", "amd64"),
}

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

if len(sys.argv) >= 2:
	n = sys.argv[1]
	l = sys.argv[2:]
	if len(l) == 0:
		try:
			l = BUILDS[n]
		except KeyError:
			print("unknown os")
			sys.exit(1)
	BUILDS = {n: l}

version = check_output("go run ./internal/_build/version/main.go".split()).strip()

_call("rm -rf build")
_call("mkdir build")

_call("go vet ./...")
_call("go test ./...")

for goos in BUILDS.keys():
	os.environ["GOOS"] = goos
	for goarch in BUILDS[goos]:
		os.environ["GOARCH"] = goarch
		cmd = "go build -o build/jcms-{}-{}-{}.bin ./cmd/jcms".format(version, goos, goarch)
		_call(cmd)

_exit(0)
