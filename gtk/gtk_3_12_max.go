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

// This file includes wrapers for symbols deprecated since GTK 3.14, and
// should not be included in a build intended to target any newer GTK versions.
//
// Warning, the build flag gtk_no_deprecated will soon be reversed to gtk_3_14_deprecated.
//
// This will allow to transition programs to strict 3.14 version, and
// ensure that those using deprecated objects know they do it at their
// own risks (of removal).

// +build gtk_3_6 gtk_3_8 gtk_3_10 gtk_3_12 !gtk_no_deprecated

package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
// #include "gtk_3_12_max.go.h"
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

func init() {
	cast_3_12_max_func = cast_3_12_max

	tm := []glib.TypeMarshaler{
		{glib.Type(C.gtk_alignment_get_type()), marshalAlignment},
		{glib.Type(C.gtk_arrow_get_type()), marshalArrow},
		{glib.Type(C.gtk_misc_get_type()), marshalMisc},
		{glib.Type(C.gtk_status_icon_get_type()), marshalStatusIcon},
	}
	glib.RegisterGValueMarshalers(tm)
}

func cast_3_12_max(class string, o *glib.Object) glib.IObject {
	var g glib.IObject
	switch class {
	case "GtkAlignment":
		g = wrapAlignment(o)
	case "GtkArrow":
		g = wrapArrow(o)
	case "GtkMisc":
		g = wrapMisc(o)
	case "GtkStatusIcon":
		g = wrapStatusIcon(o)
	}
	return g
}

// GetAlignment is a wrapper around gtk_button_get_alignment().
// Deprecated since 3.14.
func (v *Button) GetAlignment() (xalign, yalign float32) {
	var x, y C.gfloat
	C.gtk_button_get_alignment(v.native(), &x, &y)
	return float32(x), float32(y)
}

// SetAlignment is a wrapper around gtk_button_set_alignment().
// Deprecated since 3.14.
func (v *Button) SetAlignment(xalign, yalign float32) {
	C.gtk_button_set_alignment(v.native(), (C.gfloat)(xalign),
		(C.gfloat)(yalign))
}

// SetReallocateRedraws is a wrapper around
// gtk_container_set_reallocate_redraws().
// Deprecated since 3.14.
func (v *Container) SetReallocateRedraws(needsRedraws bool) {
	C.gtk_container_set_reallocate_redraws(v.native(), gbool(needsRedraws))
}

// GetDoubleBuffered is a wrapper around gtk_widget_get_double_buffered().
// Deprecated since 3.14.
func (v *Widget) GetDoubleBuffered() bool {
	c := C.gtk_widget_get_double_buffered(v.native())
	return gobool(c)
}

// SetDoubleBuffered is a wrapper around gtk_widget_set_double_buffered().
// Deprecated since 3.14.
func (v *Widget) SetDoubleBuffered(doubleBuffered bool) {
	C.gtk_widget_set_double_buffered(v.native(), gbool(doubleBuffered))
}

// Reparent() is a wrapper around gtk_widget_reparent().
// Deprecated since 3.14.
func (v *Widget) Reparent(newParent IWidget) {
	C.gtk_widget_reparent(v.native(), newParent.toWidget())
}

// ResizeGripIsVisible is a wrapper around gtk_window_resize_grip_is_visible().
// Deprecated since 3.14.
func (v *Window) ResizeGripIsVisible() bool {
	c := C.gtk_window_resize_grip_is_visible(v.native())
	return gobool(c)
}

// GetHasResizeGrip is a wrapper around gtk_window_get_has_resize_grip().
// Deprecated since 3.14.
func (v *Window) GetHasResizeGrip() bool {
	c := C.gtk_window_get_has_resize_grip(v.native())
	return gobool(c)
}

// SetHasResizeGrip is a wrapper around gtk_window_set_has_resize_grip().
// Deprecated since 3.14.
func (v *Window) SetHasResizeGrip(setting bool) {
	C.gtk_window_set_has_resize_grip(v.native(), gbool(setting))
}

/*
 * GtkAlignment
 * deprecated since version 3.14
 */

type Alignment struct {
	Bin
}

// native returns a pointer to the underlying GtkAlignment.
// Deprecated since 3.14.
func (v *Alignment) native() *C.GtkAlignment {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkAlignment(p)
}

func marshalAlignment(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapAlignment(obj), nil
}

func wrapAlignment(obj *glib.Object) *Alignment {
	return &Alignment{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

// AlignmentNew is a wrapper around gtk_alignment_new().
// Deprecated since 3.14.
func AlignmentNew(xalign, yalign, xscale, yscale float32) (*Alignment, error) {
	c := C.gtk_alignment_new(C.gfloat(xalign), C.gfloat(yalign), C.gfloat(xscale),
		C.gfloat(yscale))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapAlignment(obj), nil
}

// GetPadding is a wrapper around gtk_alignment_get_padding().
// Deprecated since 3.14.
func (v *Alignment) GetPadding() (top, bottom, left, right uint) {
	var ctop, cbottom, cleft, cright C.guint
	C.gtk_alignment_get_padding(v.native(), &ctop, &cbottom, &cleft,
		&cright)
	return uint(ctop), uint(cbottom), uint(cleft), uint(cright)
}

// SetPadding is a wrapper around gtk_alignment_set_padding().
// Deprecated since 3.14.
func (v *Alignment) SetPadding(top, bottom, left, right uint) {
	C.gtk_alignment_set_padding(v.native(), C.guint(top), C.guint(bottom),
		C.guint(left), C.guint(right))
}

// Set is a wrapper around gtk_alignment_set().
// Deprecated since 3.14.
func (v *Alignment) Set(xalign, yalign, xscale, yscale float32) {
	C.gtk_alignment_set(v.native(), C.gfloat(xalign), C.gfloat(yalign),
		C.gfloat(xscale), C.gfloat(yscale))
}

/*
 * GtkArrow
 */

// Arrow is a representation of GTK's GtkArrow.
// Deprecated since 3.14.
type Arrow struct {
	Misc
}

// native returns a pointer to the underlying GtkButton.
func (v *Arrow) native() *C.GtkArrow {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkArrow(p)
}

func marshalArrow(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapArrow(obj), nil
}

func wrapArrow(obj *glib.Object) *Arrow {
	return &Arrow{Misc{Widget{glib.InitiallyUnowned{obj}}}}
}

// ArrowNew is a wrapper around gtk_arrow_new().
// Deprecated since 3.14.
func ArrowNew(arrowType ArrowType, shadowType ShadowType) (*Arrow, error) {
	c := C.gtk_arrow_new(C.GtkArrowType(arrowType),
		C.GtkShadowType(shadowType))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapArrow(obj), nil
}

// Set is a wrapper around gtk_arrow_set().
// Deprecated since 3.14.
func (v *Arrow) Set(arrowType ArrowType, shadowType ShadowType) {
	C.gtk_arrow_set(v.native(), C.GtkArrowType(arrowType), C.GtkShadowType(shadowType))
}

/*
 * GtkStatusIcon
 * deprecated since version 3.14
 */

// StatusIcon is a representation of GTK's GtkStatusIcon
// Deprecated since 3.14.
type StatusIcon struct {
	*glib.Object
}

func marshalStatusIcon(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapStatusIcon(obj), nil
}

func wrapStatusIcon(obj *glib.Object) *StatusIcon {
	return &StatusIcon{obj}
}

func (v *StatusIcon) native() *C.GtkStatusIcon {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkStatusIcon(p)
}

// StatusIconNew is a wrapper around gtk_status_icon_new()
// Deprecated since 3.14.
func StatusIconNew() (*StatusIcon, error) {
	c := C.gtk_status_icon_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	obj.RefSink()
	e := wrapStatusIcon(obj)
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return e, nil
}

// StatusIconNewFromFile is a wrapper around gtk_status_icon_new_from_file()
// Deprecated since 3.14.
func StatusIconNewFromFile(filename string) (*StatusIcon, error) {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_status_icon_new_from_file((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	obj.RefSink()
	e := wrapStatusIcon(obj)
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return e, nil
}

// StatusIconNewFromIconName is a wrapper around gtk_status_icon_new_from_name()
// Deprecated since 3.14.
func StatusIconNewFromIconName(iconName string) (*StatusIcon, error) {
	cstr := C.CString(iconName)
	defer C.free(unsafe.Pointer(cstr))
	s := C.gtk_status_icon_new_from_icon_name((*C.gchar)(cstr))
	if s == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(s))}
	obj.RefSink()
	e := wrapStatusIcon(obj)
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return e, nil
}

// SetFromFile is a wrapper around gtk_status_icon_set_from_file()
// Deprecated since 3.14.
func (v *StatusIcon) SetFromFile(filename string) {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_status_icon_set_from_file(v.native(), (*C.gchar)(cstr))
}

// SetFromIconName is a wrapper around gtk_status_icon_set_from_icon_name()
// Deprecated since 3.14.
func (v *StatusIcon) SetFromIconName(iconName string) {
	cstr := C.CString(iconName)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_status_icon_set_from_icon_name(v.native(), (*C.gchar)(cstr))
}

// GetStorageType is a wrapper around gtk_status_icon_get_storage_type()
// Deprecated since 3.14.
func (v *StatusIcon) GetStorageType() ImageType {
	return (ImageType)(C.gtk_status_icon_get_storage_type(v.native()))
}

// SetTooltipText is a wrapper around gtk_status_icon_set_tooltip_text()
// Deprecated since 3.14.
func (v *StatusIcon) SetTooltipText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_status_icon_set_tooltip_text(v.native(), (*C.gchar)(cstr))
}

// GetTooltipText is a wrapper around gtk_status_icon_get_tooltip_text()
// Deprecated since 3.14.
func (v *StatusIcon) GetTooltipText() string {
	cstr := (*C.char)(C.gtk_status_icon_get_tooltip_text(v.native()))
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// SetTooltipMarkup is a wrapper around gtk_status_icon_set_tooltip_markup()
// Deprecated since 3.14.
func (v *StatusIcon) SetTooltipMarkup(markup string) {
	cstr := (*C.gchar)(C.CString(markup))
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_status_icon_set_tooltip_markup(v.native(), cstr)
}

// GetTooltipMarkup is a wrapper around gtk_status_icon_get_tooltip_markup()
func (v *StatusIcon) GetTooltipMarkup() string {
	cstr := (*C.char)(C.gtk_status_icon_get_tooltip_markup(v.native()))
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// SetHasTooltip is a wrapper around gtk_status_icon_set_has_tooltip()
// Deprecated since 3.14.
func (v *StatusIcon) SetHasTooltip(hasTooltip bool) {
	C.gtk_status_icon_set_has_tooltip(v.native(), gbool(hasTooltip))
}

// GetTitle is a wrapper around gtk_status_icon_get_title()
// Deprecated since 3.14.
func (v *StatusIcon) GetTitle() string {
	cstr := (*C.char)(C.gtk_status_icon_get_title(v.native()))
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// SetName is a wrapper around gtk_status_icon_set_name()
// Deprecated since 3.14.
func (v *StatusIcon) SetName(name string) {
	cstr := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_status_icon_set_name(v.native(), cstr)
}

// SetVisible is a wrapper around gtk_status_icon_set_visible()
// Deprecated since 3.14.
func (v *StatusIcon) SetVisible(visible bool) {
	C.gtk_status_icon_set_visible(v.native(), gbool(visible))
}

// GetVisible is a wrapper around gtk_status_icon_get_visible()
// Deprecated since 3.14.
func (v *StatusIcon) GetVisible() bool {
	return gobool(C.gtk_status_icon_get_visible(v.native()))
}

// IsEmbedded is a wrapper around gtk_status_icon_is_embedded()
// Deprecated since 3.14.
func (v *StatusIcon) IsEmbedded() bool {
	return gobool(C.gtk_status_icon_is_embedded(v.native()))
}

// GetX11WindowID is a wrapper around gtk_status_icon_get_x11_window_id()
// Deprecated since 3.14.
func (v *StatusIcon) GetX11WindowID() int {
	return int(C.gtk_status_icon_get_x11_window_id(v.native()))
}

// GetHasTooltip is a wrapper around gtk_status_icon_get_has_tooltip()
// Deprecated since 3.14.
func (v *StatusIcon) GetHasTooltip() bool {
	return gobool(C.gtk_status_icon_get_has_tooltip(v.native()))
}

// SetTitle is a wrapper around gtk_status_icon_set_title()
// Deprecated since 3.14.
func (v *StatusIcon) SetTitle(title string) {
	cstr := (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_status_icon_set_title(v.native(), cstr)
}

// GetIconName is a wrapper around gtk_status_icon_get_icon_name()
// Deprecated since 3.14.
func (v *StatusIcon) GetIconName() string {
	cstr := (*C.char)(C.gtk_status_icon_get_icon_name(v.native()))
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// GetSize is a wrapper around gtk_status_icon_get_size()
// Deprecated since 3.14.
func (v *StatusIcon) GetSize() int {
	return int(C.gtk_status_icon_get_size(v.native()))
}

/*
 * GtkMisc
 */

// Misc is a representation of GTK's GtkMisc.
// Deprecated since 3.14.
type Misc struct {
	Widget
}

// native returns a pointer to the underlying GtkMisc.
func (v *Misc) native() *C.GtkMisc {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkMisc(p)
}

func marshalMisc(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapMisc(obj), nil
}

func wrapMisc(obj *glib.Object) *Misc {
	return &Misc{Widget{glib.InitiallyUnowned{obj}}}
}

// GetAlignment is a wrapper around gtk_misc_get_alignment().
// Deprecated since 3.14.
func (v *Misc) GetAlignment() (xAlign, yAlign float32) {
	var x, y C.gfloat
	C.gtk_misc_get_alignment(v.native(), &x, &y)
	return float32(x), float32(y)
}

// SetAlignment is a wrapper around gtk_misc_set_alignment().
// Deprecated since 3.14.
func (v *Misc) SetAlignment(xAlign, yAlign float32) {
	C.gtk_misc_set_alignment(v.native(), C.gfloat(xAlign), C.gfloat(yAlign))
}

// GetPadding is a wrapper around gtk_misc_get_padding().
// Deprecated since 3.14.
func (v *Misc) GetPadding() (xpad, ypad int) {
	var x, y C.gint
	C.gtk_misc_get_padding(v.native(), &x, &y)
	return int(x), int(y)
}

// SetPadding is a wrapper around gtk_misc_set_padding().
// Deprecated since 3.14.
func (v *Misc) SetPadding(xPad, yPad int) {
	C.gtk_misc_set_padding(v.native(), C.gint(xPad), C.gint(yPad))
}
