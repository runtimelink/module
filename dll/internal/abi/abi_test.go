package abi_test

import (
	"math"
	"testing"

	"runtime.link/dll"
	"runtime.link/dll/internal/abi"
)

func TestRegisters(t *testing.T) {

	//libc := dll.Open("libm.so.6")
	//sym := dll.Sym(libc, "sqrt")

	/*
		add(1, 2.2)

		var evil func(float64, int64) = *(*func(float64, int64))(unsafe.Pointer(&add))
		evil(1.1, 2)*/
}

func BenchmarkSqrt(b *testing.B) {
	libc := dll.Open("libm.so.6")
	sym := dll.Sym(libc, "sqrt")

	var sqrt func(float64) float64 = math.Sqrt

	sqrt = func(f float64) float64 {
		abi.PushFunc(uintptr(sym))
		return abi.Fcallf(f)
	}

	for i := 0; i < b.N; i++ {
		sqrt(2.0)
	}

}
