// Package cgo can be used to efficiently call C ABI functions from Go.
package cgo

import (
	"reflect"
	"unsafe"

	"runtime.link/std"
)

const (
	ErrDisabled errorString = "cgo is disabled" // returned when CGO_ENABLED=0 and the function requires CGO to call.
)

// MissingSymbolError is returned when the linker
// couldn't find the symbol for the given name.
type MissingSymbolError string

func (e MissingSymbolError) Error() string { return string("symbol " + string(e) + " not found") }

type TagCompatiblityError struct {
	tag   std.Tag
	err   error
	ftype reflect.Type
}

func (e TagCompatiblityError) Error() string {
	return "incompatible tag '" + string(e.tag) + "' for function " + e.ftype.String() + ": " + e.err.Error()
}

type errorString string

func (e errorString) Error() string { return string(e) }

// Linker can return the function pointer for
// the given symbol, or nil if no symbol is found.
type Linker func(sym string) unsafe.Pointer

// MakeFunc takes a pointer to Go function 'fn' and a runtime.link
// standard tag and implements fn, such that it calls the platform-native
// ABI function. An error is returned if the linker couldn't find the symbol,
// or if the Go function signature is incompatible with the tag. MakeFunc
// cannot assert the correctness of the tag, so it is very important that
// the tag correctly describes the function signature and memory behaviour.
// Incorrect tags can lead to undefined behaviour, memory corruption and
// unpredictable crashes. Treat the tag as you would unsafe code.
func (ln Linker) MakeFunc(fn any, tag std.Tag) error {
	return ln.makeFunc(fn, tag)
}
