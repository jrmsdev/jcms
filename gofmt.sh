#!/bin/sh -eu
extra_files='
./lib/internal/admin/handler/zipfile.go.in
'
gofmt -w -l -s ./bin ./lib
for f in $extra_files; do
	gofmt -w -l -s $f
done
