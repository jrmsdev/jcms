#!/bin/sh -eu
. ./jcms-devel.env
go generate ./webapp
if test "${1:-'zipmode'}" = '--devel'; then
	rm -vf lib/internal/handler/zipfile_webapp.go
fi
mkdir -p build
mainsrc=./bin/jcms
admincmd=./build/jcms.bin
rm -f $admincmd
go build -tags jcms -o $admincmd $mainsrc
$admincmd -debug -port 6080
