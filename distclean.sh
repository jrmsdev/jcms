#!/bin/sh -eu
DIR=${GOPATH}
go clean -i ./..
go clean -cache
test -d ${DIR}/bin
test -d ${DIR}/pkg
rm -rfv ${DIR}/bin/jcms ${DIR}/pkg/*/github.com/jrmsdev/jcms*
