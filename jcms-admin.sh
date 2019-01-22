#!/bin/sh -eu
. ./jcms-devel.env
go generate ./...
if test "${1:-'zipmode'}" = '--devel'; then
	rm -vf lib/internal/admin/handler/zipfile.go
fi
mkdir -p build
mainsrc=./bin/jcms-admin
admincmd=./build/jcms-admin.bin
rm -f $admincmd
go build -o $admincmd $mainsrc
$admincmd -debug
