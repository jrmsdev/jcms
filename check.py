#!/usr/bin/env python3

import os
import sys
from subprocess import getstatusoutput

verbose = ""
if '-v' in sys.argv:
	verbose = " -v"

prevcmd = {
	10: "go install -i ./cmd/jcms",
	20: "go get -v -t ./...",
	30: "go vet ./...",
}
gotest = "go test{} ./...".format(verbose)

for idx in sorted(prevcmd.keys()):
	cmd = prevcmd[idx]
	print(cmd)
	rc, outs = getstatusoutput(cmd)
	if outs != "":
		print(outs)
	if rc != 0:
		sys.exit(rc)

print(gotest)
rc, outs = getstatusoutput(gotest)

print(outs)
sys.exit(rc)
