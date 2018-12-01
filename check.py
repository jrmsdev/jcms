#!/usr/bin/env python

import os
import sys
from subprocess import call

verbose = ""
race = ""
for a in sys.argv:
	if a == "-v":
		verbose = " -v"
	elif a == "-race":
		race = " -race"

prevcmd = {
	10: "go vet ./...",
	20: "go install -i ./cmd/jcms",
	30: "go get -v -t ./...",
}
gotest = "go test{}{} ./...".format(verbose, race)

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

sys.exit(0)
