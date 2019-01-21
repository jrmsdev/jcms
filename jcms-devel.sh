#!/bin/sh -eux
. ./jcms-devel.env
go generate ./...
mkdir -p build
mainsrc=./internal/_devel
develcmd=./build/jcms-devel.bin
rm -f $develcmd
go build -o $develcmd $mainsrc
$develcmd -D -p 6080
