package main

import "C"

func main() {}


//export CGoCustomUUID
func CGoCustomUUID(base64Text *C.char) *C.char {
	text := C.GoString(base64Text)
	return C.CString(CustomUUID(text))
}


//export CGoRunHope
func CGoRunHope(base64Text *C.char) *C.char {
	text := C.GoString(base64Text)
	return C.CString(RunHope(text))
}

//export CGoStopHope
func CGoStopHope() *C.char {
	return C.CString(StopHope())
}
