#!/bin/bash

set -e

error_exit() {
	echo -e "\033[1;31m!!!!!!!!! ERROR !!!!!!!!!!\033[0m"
	exit $?
}

build() {
	./go-gobject-gen -config config.json $1
	make -C $1 install || error_exit
}

./clean.bash

go build -o go-gobject-gen
build glib-2.0
build gobject-2.0

make -C cairo-1.0 install || error_exit
build atk-1.0
build gio-2.0
build gdkpixbuf-2.0
build pango-1.0
build pangocairo-1.0
build gdk-2.0
build gtk-2.0
