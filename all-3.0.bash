#!/bin/bash

set -e

error_exit() {
	echo -e "\033[1;31m!!!!!!!!! ERROR !!!!!!!!!!\033[0m"
	exit $?
}

build() {
	echo $1
	./go-gobject-gen -config config.json $1
	pushd . > /dev/null
	cd $1
	go install
	popd > /dev/null
}

./clean.bash

go build -o go-gobject-gen
build glib-2.0
build gobject-2.0

pushd . > /dev/null
cd cairo-1.0
go install
popd > /dev/null
build atk-1.0
build gio-2.0
build gdkpixbuf-2.0
build pango-1.0
build pangocairo-1.0
build gdk-3.0
build gtk-3.0
build gtksource-3.0
