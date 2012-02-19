#!/bin/bash

set -e

cd gogtk-demo
for f in *.go; do
	go build $f
done
