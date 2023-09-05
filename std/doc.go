/*
Package std provides standard types for cross-language interoperability.

# Standard Symbols

This package defines a standard string format for representing symbols
along with their type. This string always starts with comma seperated
symbol names, in order of preference. Next up is a space, followed by
the type of the symbol. The type either begins with 'func' for function
types, or the name of the standard C type. Similarly to Go, the return
type is placed after the parameter list.

	abs func(int)int // simple C function, no pointer semantics.

Assertions can be added to function parameters (and return value) to
document pointer ownership, error handling and memory safety assertions
to make when interacting with the function.

	fread func(&void[@3],size_t/@1,size_t,&FILE)size_t<@3; ferror(@4)
	           ^     ^         ^                      ^    ^
			   │     │         │                      │    └── error details.
			   │     |         |                      |
			   │     |         |                      └── error condition.
			   │     |         |
			   │     |         └── must equal the underlying value
			   │     |             sizeof the 1st argument (void).
			   │     |
			   │     └── memory capacity must be greater than the 3rd
			   │         argument.
			   |
			   └── fread borrows this buffer for the duration of the call.

# Ownership Assertions

These assertions are used to document the ownership semantics for
pointer types, their presence signals that the value is a pointer.

  - $type - memory ownership is sold to the reciever of the value,
    the receiver becomes responsible for freeing it and
    can do so immediately.
  - &type - the receiver borrows this pointer and will not keep a
    reference to it beyond the the lifetime of the
    function call. If specified on a return value, the
    receiver must copy the value.
  - *type - the value is a static pointer, neither the receiver
    nor the sender may free it.
  - +type - the receiver borrows this pointer so that they can
    initialize it. The existing value is overwritten.
  - -type - the receiver is granted ownership but can safely
    ignore this parameter as it copies an existing
    parameter that the receiver leased to the sender.

# Memory Safety Assertions

These assertions are used to document memory safety semantics. 'n'
is an unsigned integer that refers to the Nth type identifier in
the standard symbol.

  - type[@n]   - the underlying memory capacity of the pointer must be
    greater (and not equal) to '@', the '@' symbol can be
    omitted to refer to the literal integer value 'n'.
  - type/@n	 - the value must equal the underlying value sizeof '@',
    the '@' symbol can be omitted to refer to the literal
    integer value 'n'.
  - type^@n    - the value points within the memory buffer of '@',
    therefore the lifetime of this value must match the
    lifetime of '@'.
  - type...f@n - the value should be validated as a printf-style vararg
    list.
  - type:@n  	 - the value's type matches '@'.

# Error Handling Assertions

These assertions are used to document error handling semantics for
a function. They are only valid on return values. The 'sym' is a
symbol name that can be used to lookup the error message. It can
be omitted.

  - type>@n/n; sym  - when the value is greater than n/@n, see the given
    symbol for more information about the error.
  - type<@n/n; sym  - when the value is less than n/@n, see the given
    symbol for more information about the error.
  - type>=@n/n; sym - when the value is greater than or equal to n/@n,
    see the given symbol for more information about
    the error.
  - type<=@n/n; sym - when the value is less than or equal to n/@n, see
    the given symbol for more information about the error.
  - type!@n/n; sym - when the value is not n/@n, see the given symbol
    for more information about the error.
  - type=@n/n; sym - when the value is n/@n, see the given symbol
    for more information about the error.

# Macros

When a standard symbol is tagged on a Go function field, it conveys the C
representation of that field. Some standard macros are supported for mapping
the Go function signature to the C one. These macros are purely used for
convenience and do not change the semantics of the C function signature
(the C semantics of a standard symbol tag with macros remain constant
when the macros are removed).

  - type{n} - The constant integer value n is always used for this
    parameter.
  - type%v  - The Vth function argument is mapped against this
    parameter. Standard printf formatting rules apply
    as if each argument in the function was passed to
    the fmt.Sprintf function. Only %v and %[n]v verbs
    are supported.

# Structures

A struct is identified by an ordered sequence of standard symbols.

Structs passed across language boundaries must have their fields
tagged with the 'std' tag. This tag is used to specify the standard
type for each field. Pointer fields are typically tagged with an '&'.

	type MyStruct {
		Name string `std:"name &char"`
	}

# Deep Copies

By default, values are deep-copied between languages. In order to
avoid these copies, C ownership can be preserved with [String] and
[Pointer] types. Which need to be manually freed. Struct fields
can be accessed directly this way by specifying getter and setter
functions. Such types are safe to pass back and forth between
languages (although may panic when misused).

	// accessor methods to link.
	type MyStructs {
		Name(std.Pointer[MyStruct]) string     `std:"my_struct.Name"`
		SetName(std.Pointer[MyStruct], string) `std:"my_struct.Name"`
	}
*/
package std
