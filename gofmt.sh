#!/bin/sh -eu
extra_files='
./internal/httpd/handler/lib_files.go.in
./internal/admin/handler/zipfile.go.in
'
gofmt -w -l -s .
for f in $extra_files; do
	gofmt -w -l -s $f
done
