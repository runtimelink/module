package dll_test

import (
	"fmt"
	"testing"

	"runtime.link/dll"
	"runtime.link/lib"
)

var libc = dll.Import[struct {
	linux   lib.Location `std:"libc.so.6 libm.so.6"`
	darwin  lib.Location `std:"libSystem.dylib"`
	windows lib.Location `std:"msvcrt.dll"`

	puts func(string) error    `std:"puts func(&char)int<0"`
	sqrt func(float64) float64 `std:"sqrt func(double)double"`
}]()

func TestHelloWorld(*testing.T) {
	//libc.PutString("Hello, World!")

	fmt.Println(libc.sqrt == nil)
	fmt.Println(libc.sqrt(2))
}
