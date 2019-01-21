#!/bin/sh -eu
. ./jcms-devel.env
go generate ./...
if test "${1:-'zipmode'}" = '--devel'; then
	rm -vf internal/admin/handler/zipfile.go
fi
mkdir -p build
mainsrc=./cmd/jcms-admin
admincmd=./build/jcms-admin.bin
rm -f $admincmd
go build -o $admincmd $mainsrc
$admincmd -D -p 6080
