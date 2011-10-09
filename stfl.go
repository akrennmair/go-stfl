package stfl

/*
#cgo pkg-config: stfl
#cgo LDFLAGS: -lncursesw -lpthread
#include <stdlib.h>
#include <stfl.h>
#include "stfl_wrapper.h"
#include <locale.h>
*/
import "C"
import "unsafe"

type Form C.stfl_form

func Init() {
	emptystr := C.CString("")
	defer C.free(unsafe.Pointer(emptystr))
	C.setlocale(C.LC_ALL, emptystr)
}

func Create(text string) *Form {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	return (*Form)(C.stfl_create_wrapper(ctext))
}

func(form *Form) Free() {
	C.stfl_free((*C.stfl_form)(form))
}

func(form *Form) Run(timeout int) string {
	return C.GoString(C.stfl_run_wrapper((*C.stfl_form)(form), C.int(timeout)))
}

func Reset() {
	C.stfl_reset()
}

func(form *Form) Get(name string) string {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cvalue := C.stfl_get_wrapper((*C.stfl_form)(form), cname)
	return C.GoString(cvalue)
}

func(form *Form) Set(name string, value string) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cvalue := C.CString(value)
	defer C.free(unsafe.Pointer(cvalue))
	C.stfl_set_wrapper((*C.stfl_form)(form), cname, cvalue)
}

func(form *Form) GetFocus() string {
	return C.GoString(C.stfl_get_focus_wrapper((*C.stfl_form)(form)))
}

func(form *Form) SetFocus(name string) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.stfl_set_focus_wrapper((*C.stfl_form)(form), cname)
}

func Quote(text string) string {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	return C.GoString(C.stfl_quote_wrapper(ctext))
}

func(form *Form) Modify(name string, mode string, text string) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cmode := C.CString(mode)
	defer C.free(unsafe.Pointer(cmode))
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	C.stfl_modify_wrapper((*C.stfl_form)(form), cname, cmode, ctext)
}

func(form *Form) Lookup(path string, newname string) string {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	cnewname := C.CString(newname)
	defer C.free(unsafe.Pointer(cnewname))
	return C.GoString(C.stfl_lookup_wrapper((*C.stfl_form)(form), cpath, cnewname))

}

func(form *Form) Dump(name string, prefix string, focus bool) string {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cprefix := C.CString(prefix)
	defer C.free(unsafe.Pointer(cprefix))
	cfocus := 0
	if focus {
		cfocus = 1
	}
	return C.GoString(C.stfl_dump_wrapper((*C.stfl_form)(form), cname, cprefix, C.int(cfocus)))
}
