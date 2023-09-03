# runtime.link

This module provides a dynamic linker for Go programs so they can load shared libraries at runtime.
C libraries to be loaded this way are represented using a set of Go structs. 

Example:
```go
package main

import (
    "runtime.link/dll"
    "runtime.link/ffi"
)

var libc struct {
    ffi.Functions

    PutString func(string) ffi.Int `ffi:"puts"`
}

func main() {
    if err := dll.Link(&libc); err != nil {
        panic(err)
    }
    libc.PutString("Hello, World!")
}

```