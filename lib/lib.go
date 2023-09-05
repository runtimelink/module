package lib

// Location can be added to a library structure to specify
// the standard location or name of that library on a
// specific GOOS.
//
// For example:
//
//	type Library struct {
//		linux   lib.Location `std:"libc.so.6 libm.so.6"`
//		darwin  lib.Location `std:"libSystem.dylib"`
//		windows lib.Location `std:"msvcrt.dll"`
//	}
type Location struct{}
