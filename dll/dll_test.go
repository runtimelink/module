package dll_test

import (
	"testing"

	"runtime.link/dll"
	"runtime.link/std"
)

var libc = dll.Import[struct {
	linux   std.Location `libc.so.6 libm.so.6`
	darwin  std.Location `libSystem.dylib`
	windows std.Location `msvcrt.dll`

	PutString func(string) std.Int `sym:"puts"`
}]()

func TestHelloWorld(*testing.T) {
	libc.PutString("Hello, World!")
}
