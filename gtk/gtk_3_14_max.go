// Copyright (c) 2013-2014 Conformal Systems <info@conformal.com>
//
// This file originated from: http://opensource.conformal.com/
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

// This file includes wrapers for symbols deprecated since GTK 3.16, and
// should not be included in a build intended to target any newer GTK versions.
//
// Warning, the build flag gtk_no_deprecated will soon be reversed to gtk_3_16_deprecated.
//
// This will allow to transition programs to strict 3.16 version, and
// ensure that those using deprecated objects know they do it at their
// own risks (of removal).
//
// +build gtk_3_6 gtk_3_8 gtk_3_10 gtk_3_12 gtk_3_14 !gtk_no_deprecated

package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/gdk"
)

// OverrideColor is a wrapper around gtk_widget_override_color().
// Deprecated since 3.16.
func (v *Widget) OverrideColor(state StateFlags, color *gdk.RGBA) {
	var cColor *C.GdkRGBA
	if color != nil {
		cColor = (*C.GdkRGBA)(unsafe.Pointer((&color.RGBA)))
	}
	C.gtk_widget_override_color(v.native(), C.GtkStateFlags(state), cColor)
}

// OverrideFont is a wrapper around gtk_widget_override_font().
// Deprecated since 3.16.
func (v *Widget) OverrideFont(description string) {
	cstr := C.CString(description)
	defer C.free(unsafe.Pointer(cstr))
	c := C.pango_font_description_from_string(cstr)
	C.gtk_widget_override_font(v.native(), c)
}
