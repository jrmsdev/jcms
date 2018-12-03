#!/bin/sh -eu
eval `go env`
DIR=${GOPATH}
OSARCH=${GOOS}_${GOARCH}
go clean -i ./..
go clean -cache
test -d ${DIR}/bin
test -d ${DIR}/pkg
rm -rfv ${DIR}/bin/jcms ${DIR}/pkg/${OSARCH}/github.com/jrmsdev/jcms* | sort
git clean -xfd
