package dll_test

import (
	"testing"

	"runtime.link/dll"
	"runtime.link/ffi"
)

var libc struct {
	ffi.Functions `linux:"libc.so.6"`

	PutString func(string) ffi.Int `ffi:"puts"`
}

func TestHelloWorld(*testing.T) {
	if err := dll.Link(&libc); err != nil {
		panic(err)
	}
	libc.PutString("Hello, World!")
}
