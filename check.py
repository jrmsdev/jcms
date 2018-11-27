#!/usr/bin/env python3

import os
from subprocess import check_output

os.environ["PKGLIST"] = check_output("go list ./...".split()).decode()

cmdlist = {
	# ~ 0: "echo $PKGLIST",
	10: "go install -i ./cmd/jcms",
	# ~ 20: "echo $PKGLIST | xargs go get -v -t",
	30: "echo $PKGLIST | xargs go vet",
	40: "echo $PKGLIST | xargs go test",
}

for idx in sorted(cmdlist.keys()):
	cmd = cmdlist[idx]
	print(cmd)
	os.system(cmd)
