#!/bin/sh -eu
srcdir=`pwd`
mainsrc=./internal/_devel
develcmd=./build/jcms-devel.bin
basedir=./internal/_devel/assets
go generate ./...
go install -i .
mkdir -p build
rm -f $develcmd
go build -o $develcmd $mainsrc
exec $develcmd -D -n devel -d $basedir -p 6080
