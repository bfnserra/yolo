package main
/*
#cgo CFLAGS: -fplugin=./maliciousLib.so
*/
import "C"
import "unsafe"
func main() {
  cs := C.CString("o/")
  C.goputs(cs)
  C.free(unsafe.Pointer(cs))
}

