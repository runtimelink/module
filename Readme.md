# runtime.link

This module provides a dynamic linker for Go programs so they can load C standard ABI shared libraries at runtime.
Granted that correct struct tag semantics are provided, runtime.link provides good memory safety guarantees. Such 
that C libraries (such as the ANSI C standard library) are completely safe to use from Go. 


Example:
```go
package main

import (
    "runtime.link/dll"
    "runtime.link/lib"
)

var libc = dll.Import[struct {
    linux   lib.Location `std:"libc.so.6 libm.so.6"`
	darwin  lib.Location `std:"libSystem.dylib"`
	windows lib.Location `std:"msvcrt.dll"`

    PutString func(string) error `std:"puts func(&char)int<0"
                                                 ^        ^
                                                 |        |
                                                 |        └── error condition
                                                 |
                                                 └── puts borrows the string but
                                                     it doesn't keep a reference.`
}]()

func main() {
    if err := libc.PutString("Hello, World!"); err != nil {
        panic(err)
    }
}
```