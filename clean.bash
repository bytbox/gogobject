#!/bin/bash

set -e

clean_in_dir() {
	pushd . > /dev/null
	cd $1
	go clean
	popd > /dev/null
}

go clean
clean_in_dir glib-2.0
clean_in_dir gobject-2.0
clean_in_dir cairo-1.0
clean_in_dir atk-1.0
clean_in_dir gio-2.0
clean_in_dir gdkpixbuf-2.0
clean_in_dir gdk-3.0
clean_in_dir pango-1.0
clean_in_dir pangocairo-1.0
clean_in_dir gtk-3.0
clean_in_dir gtksource-3.0
clean_in_dir gtk-2.0
clean_in_dir gdk-2.0

rm -f glib-2.0/glib.go
rm -f gobject-2.0/gobject.go
rm -f atk-1.0/atk.go
rm -f gio-2.0/gio.go
rm -f gdkpixbuf-2.0/gdkpixbuf.go
rm -f gdk-3.0/gdk.go
rm -f pango-1.0/pango.go
rm -f pangocairo-1.0/pangocairo.go
rm -f gtk-3.0/gtk.go
rm -f gtksource-3.0/gtksource.go
rm -f gtk-2.0/gtk.go
rm -f gdk-2.0/gdk.go
