#!/usr/bin/env python

import os
import sys
from subprocess import call

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
	rc = call(cmd.split())
	if rc != 0:
		sys.exit(rc)

print(gotest)
rc = call(gotest.split())
sys.exit(rc)
