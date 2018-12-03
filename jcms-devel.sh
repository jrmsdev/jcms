#!/bin/sh -eu
srcdir=`pwd`
mainsrc=./internal/cmd/_devel
develcmd=./build/jcms-devel.bin
go install -i .
mkdir -p build
rm -f $develcmd
go build -o $develcmd $mainsrc
exec $develcmd -d -n devel -p 6080
