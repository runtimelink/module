# runtime.link

The runtime.link project aims to encourage the wider adoption of using Go to clearly 
represent software interfaces. The author(s) believe Go to be the best 
widely-available language for this purpose.

At this time, the project includes a specification (and is working on an implementation) 
for a dynamic linker for Go binaries which will be able to link to C shared libraries 
safely at runtime. When correct struct tags are in-place, runtime.link will provide 
excellent memory-safety guarantees. Such that, the ANSI C standard library will be 
completely safe to use from Go. 

**Loading Shared Libraries at Runtime**
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

    puts func(string) error `std:"puts func(&char)int<0"
                                            ^        ^
                                            |        |
                                            |        └── error condition
                                            |
                                            └── puts borrows the string but
                                                it doesn't keep a reference.`
}]()

func main() {
    if err := libc.puts("Hello, World!"); err != nil {
        panic(err)
    }
}
```