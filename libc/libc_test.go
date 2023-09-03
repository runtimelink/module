package libc_test

import (
	"fmt"
	"math"
	"testing"

	"runtime.link/ffi"
	"runtime.link/libc"
)

func init() {
	if err := libc.Link(); err != nil {
		panic(err)
	}
}

func TestMain(m *testing.M) {
	libc.Program.Exit(ffi.Int(m.Run()))
}

func TestLibc(t *testing.T) {
	fmt.Println(libc.Char.IsAlphaNumeric('a'))

	fmt.Println(libc.Double.Sqrt(2))

	fmt.Println(libc.Float.Frexp(2.2))

	fmt.Println(libc.Time.Now(nil))

	fmt.Println(libc.Locale.Get())

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
		libc.Double.Sqrt(2)
	}
}
