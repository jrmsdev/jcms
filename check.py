#!/usr/bin/env python3

import os
import sys
from subprocess import check_output, getstatusoutput

verbose = ""
if '-v' in sys.argv:
	verbose = " -v"

prevcmd = {
	10: "go install -i ./cmd/jcms",
	20: "go get -v -t ./...",
	30: "go vet ./...",
}
gotest = f"go test{verbose} ./..."

for idx in sorted(prevcmd.keys()):
	cmd = prevcmd[idx]
	print(cmd)
	outs = check_output(cmd.split()).decode().strip()
	if outs != "":
		print(outs)

print(gotest)
rc, outs = getstatusoutput(gotest)
print(outs)
sys.exit(rc)
