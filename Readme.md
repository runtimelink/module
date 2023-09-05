# runtime.link

This module provides a dynamic linker for Go programs so they can load C standard ABI shared libraries at runtime.
Granted that struct tag semantics are provided, runtime.link provides good memory safety guarantees. Such that C
libraries (such as the ANSI C standard library) are completely safe to use from Go. 


Example:
```go
package main

import (
    "runtime.link/dll"
    "runtime.link/ffi"
)

var libc = dll.Import[struct {
    linux   std.Location `libc.so.6 libm.so.6`
	darwin  std.Location `libSystem.dylib`
	windows std.Location `msvcrt.dll`

    PutString func(string) error `std:"puts func(&char)int<0"`
}]()

func main() {
    libc.PutString("Hello, World!")
}

```