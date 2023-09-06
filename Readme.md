# runtime.link

The runtime.link project aims to encourage the wider adoption of leveraging Go to clearly 
represent and communicate software interfaces. The author(s) believe Go to be the best 
widely-available and supported language for this purpose.

At this time, the project includes a specification (and is working on an implementation) 
for a dynamic linker for Go binaries, which will be able to link to C shared libraries 
safely at runtime. When correct struct tags are in-place, runtime.link will provide 
excellent memory-safety guarantees. Such that, the ANSI C standard library will be 
completely safe to use from Go.

This repository is open to contributions, and the author(s) encourage anyone interested
in representing well known libraries and APIs to contribute to the project and add to
the lib registry.

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

### API Naming Conventions, Design Principles and Standards.

1. Prefer words over abbreviations ie. "PutString" over "puts".
   The only exceptions to this are package names and usage that
   appears in the Go standard library.
2. Acronyms reflect a compressed subject, they should only appear 
   at the end of a name, in isolation, or as a package name.