#!/usr/bin/env python

import os
import sys
from subprocess import call

verbose = ""
if '-v' in sys.argv:
	verbose = " -v"

gocmd = "go"
prevcmd = dict()
if sys.platform.startswith('win'):
	gocmd = "go.exe"
	prevcmd.update({
		0: "choco install golang",
	})

gotest = "{} test{} ./...".format(gocmd, verbose)
prevcmd.update({
	10: "{} install -i ./cmd/jcms".format(gocmd),
	20: "{} get -v -t ./...".format(gocmd),
	30: "{} vet ./...".format(gocmd),
})


rc = call([gocmd, "version"])
if rc != 0:
	sys.exit(rc)

for idx in sorted(prevcmd.keys()):
	cmd = prevcmd[idx]
	print(cmd)
	rc = call(cmd.split())
	if rc != 0:
		sys.exit(rc)

print(gotest)
rc = call(gotest.split())
if rc != 0:
	sys.exit(rc)

rc = call("jcms -version".split())
if rc != 0:
	sys.exit(rc)

sys.exit(0)
