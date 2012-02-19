#!/bin/sh

set -e

cd gogtk-demo
for f in *.go; do
	go build $f
done
