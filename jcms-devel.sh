#!/bin/sh -eu
. ./jcms-devel.env
go generate ./...
mkdir -p build
rm -f $develcmd
go build -o $develcmd $mainsrc
$develcmd -D -n devel -p 6080
