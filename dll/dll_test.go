package dll_test

import (
	"testing"

	"runtime.link/dll"
	"runtime.link/lib"
)

var libc = dll.Import[struct {
	linux   lib.Location `std:"libc.so.6 libm.so.6"`
	darwin  lib.Location `std:"libSystem.dylib"`
	windows lib.Location `std:"msvcrt.dll"`

	PutString func(string) error `std:"puts func(&char)int<0"`
}]()

func TestHelloWorld(*testing.T) {
	libc.PutString("Hello, World!")
}
