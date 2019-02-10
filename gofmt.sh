#!/bin/sh -eu
extra_files='
./lib/internal/handler/zipfile.go.in
'
gofmt -w -l -s ./bin ./lib ./_t ./_build ./webapp
for f in $extra_files; do
	gofmt -w -l -s $f
done
