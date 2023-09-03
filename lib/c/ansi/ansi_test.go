package ansi_test

import (
	"fmt"
	"math"
	"testing"

	"runtime.link/dll"
	"runtime.link/ffi"
	"runtime.link/lib/c/ansi"
)

var libc = dll.Import[ansi.Functions]()

func TestMain(m *testing.M) {
	libc.Program.Exit(ffi.Int(m.Run()))
}

func TestLibc(t *testing.T) {
	fmt.Println(libc.Char.IsAlphaNumeric('a'))

	fmt.Println(libc.Math.Sqrt(2))

	var i ffi.Int
	var d = libc.Math.Frexp(2.2, &i)
	fmt.Println(d, i)

	fmt.Println(libc.System.Time(nil))

	libc.Program.OnExit(func() {
		fmt.Println("exiting...")
	})
}

func BenchmarkGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.Sqrt(2)
	}
}

func BenchmarkC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		libc.Math.Sqrt(2)
	}
}
