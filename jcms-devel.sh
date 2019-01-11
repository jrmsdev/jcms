#!/bin/sh -eu
srcdir=`pwd`
mainsrc=./internal/_devel
develcmd=./build/jcms-devel.bin
basedir=./internal/_devel/assets
datadir=./internal/_devel/data
go generate ./...
mkdir -p build
rm -f $develcmd
go build -o $develcmd $mainsrc
export JCMS_BASEDIR=$basedir
export JCMS_DATADIR=$datadir
$develcmd -D -n devel -p 6080
