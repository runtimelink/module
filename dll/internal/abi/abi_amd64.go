package abi

// AMD64 registers
// general purpose: 15
// floating point: 8

func PushFunc(u uintptr) {
	pushFunc(u)
}

func pushFunc(fn uintptr)

func Fcallf(f float64) float64 {
	return fcallf(f)
}

func fcallf(f1 float64) float64
