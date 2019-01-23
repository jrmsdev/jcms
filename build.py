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
VERSION = "0.0"
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

argc = len(sys.argv) - 1

goos = check_output("go env GOOS".split()).strip()
goarch = check_output("go env GOARCH".split()).strip()

if not "--all" in sys.argv:
	if argc == 0:
		BUILDS = {goos: [goarch]}
	elif argc >= 1:
		n = sys.argv[1]
		l = sys.argv[2:]
		if len(l) == 0:
			try:
				l = BUILDS[n]
			except KeyError:
				print("unknown os")
				sys.exit(1)
		else:
			for x in l:
				if x not in BUILDS[n]:
					print("unknown arch: %s" % x)
					sys.exit(1)
		BUILDS = {n: l}

if os.system("rm -rf build") != 0:
	_exit(1)
if os.system("mkdir build") != 0:
	_exit(1)

_call("go generate ./webapp/...")
goversion = check_output(["go", "version"]).strip().split()[2].strip()

_call("go vet ./bin/... ./lib/...")

for goos in sorted(BUILDS.keys()):
	os.environ["GOOS"] = goos
	for goarch in BUILDS[goos]:
		os.environ["GOARCH"] = goarch
		version = "{}-{}-{}-{}".format(VERSION, goversion, goos, goarch)
		for n in CMDBIN:
			cmd = "go build -o build/{}-{}.bin ./bin/{}".format(n, version, n)
			_call(cmd)

_exit(0)
