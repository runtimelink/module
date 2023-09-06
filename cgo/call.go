// Package cgo can be used to efficiently call C ABI functions from Go.
package cgo

import (
	"unsafe"

	"runtime.link/std"
)

// Linker can return the function pointer for
// the given symbol, or nil if no symbol is found.
type Linker func(sym string) unsafe.Pointer

// MakeFunc takes a pointer to Go function 'fn' and a runtime.link
// standard symbol string and implements fn, such that it calls
// the C ABI function. An error is returned if the linker couldn't
// find the symbol, or if the Go function signature is incompatible
// with the tag. MakeFunc cannot assert the correctness of the tag,
// so it is very important that the tag correctly matches the C
// function signature and memory behaviour. Incorrect tags can lead
// to undefined behaviour, memory corruption and unpredictable crashes.
// Treat the tag as you would unsafe code.
func (ln Linker) MakeFunc(fn any, tag std.Tag) error {
	return nil
}
