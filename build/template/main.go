package main

import (
	"C"

	libHope "github.com/asimov1234/libnewv"
)

func main() {}

//export CGoCustomUUID
func CGoCustomUUID(base64Text *C.char) *C.char {
	text := C.GoString(base64Text)
	return C.CString(libHope.CustomUUID(text))
}

//export CGoRunHope
func CGoRunHope(base64Text *C.char) *C.char {
	text := C.GoString(base64Text)
	return C.CString(libHope.RunHope(text))
}

//export CGoStopHope
func CGoStopHope() *C.char {
	return C.CString(libHope.StopHope())
}
