package libc

import "runtime.link/dll"

// Location defines the DLL locations for the C standard library
// on supported operating systems.
type Location struct {
	linux   dll.Tag `dll:"libc.so.6 libm.so.6"`
	darwin  dll.Tag `dll:"libSystem.dylib"`
	windows dll.Tag `dll:"msvcrt.dll"`
}
