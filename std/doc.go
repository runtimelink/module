/*
Package std provides standard types for cross-language interoperability.

# Standard Types

This package defines a standard string format for representing symbols
along with their type. This string always starts with comma seperated
symbol names, in order of preference. Next up is a space, followed by
the type of the symbol. The type either begins with 'func' for function
types, or the name of the standard C type.

abs func(int)int

Assertions can be added to function parameters (and return value) to
document pointer ownership, error handling and memory safety assertions
to make when interacting with the function.

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

# Memory Safety Assertions

These assertions are used to document memory safety semantics. 'n'
is an unsigned integer that refers to the Nth argument of the function.

	- type[@n]  - the underlying memory capacity of the pointer must be
			      greater (and not equal) to '@', the '@' symbol can be
				  omitted to refer to the literal integer value 'n'.
	- type^@n   - the value points within the memory buffer of '@'.
	- type...@n - the value should be validated as a printf-style vararg
	              list.

# Error Handling Assertions

These assertions are used to document error handling semantics for
a function. They are only valid on return values. The 'sym' is a
symbol name that can be used to lookup the error message. It can
be omitted.


	- type>sym - when the value is greater than zero, see the given
				 symbol for more information about the error.
	- type<sym - when the value is less than zero, see the given
				 symbol for more information about the error.
	- type0sym - when the value is zero, see the given symbol
				 for more information about the error.




Import dynamically links to the the specified library.
If any symbols fail to load, the corresponding functions
will panic. Library locations provided to this function
will override the default ones to search for.

Library should be a struct of functions, each function
must clearly define a standard signature and symbol.
This can be achieved by sticking to std package types, or
by using a std tag that defines the signature.

For example:

	PutString func(std.String) std.Int `sym:"puts"`   // unsafe, pointer values passed directly.
	PutString func(string) int `std:"int puts(char)"` // safest, deep copy and convert all values.

The 'std' is similar to a C function signature, but
with *, [] symbols and the argument names omitted
(function arguments are specified using 'void').
Import will not free memory by default, as this
is the safest option. In order to prevent these
memory leaks, the function signature can have
appropriate parameter annotations.

Import may use these annotations to optimize calls
and decide how pointers are passed.

	'#type'     - tags this symbol as the destructor for
				  the given type. Which will be used
				  to track ownership disposal.
	'&type'     - the receiver borrows this pointer and
				  will not keep a reference to it.
	'-type'     - the receiver ignores this parameter.
	'type=0'    - set to zero
	'type=1'    - set to one
	'type<0'     - signals error when smaller than zero.
	'type!'     - signals error when zero.
	'$type'     - the receiver takes ownership of this
				  pointer and is responsible for freeing it.
	'type%v'    - the argument identified by the given
		          fmt parameter is mapped here. Must
		          come before other suffixed annotations.
	'type|%v'   - the argument must have greater length than
				  the argument identified by the given fmt
				  parameter.
	'type||%v'   - the argument must have greater capacity
				  to the length of the argumenent identified
				  by the given fmt parameter.
	'type?sym'  - (for return type only) when the value
				  is not empty, return the result from
				  the given symbol instead. Otherwise
				  return the zero value. Either directly
				  or in an additional return value (if specified).
	'type!sym'  - (for argument type only) when the value
				  is empty, return the result from the
				  given symbol instead, either directly
				  or in an additional return value (if specified).
				  if 'sym' is omitted, invert the output.
	'free@sym'  - frees the memory allocated because of
				  this parameter, right after the next time
				  the given symbol is called with a matching
			 	  pattern.
	'ptrdiff_t%v' - the argument identified by the given
				  fmt parameter is assumed to be a pointer
				  within that parameter.
	'null'	    - like void but a null char is appended to
				  the end of it. works only for []byte.
	'varg%v'   	- the arguments are validated to correspond
				  to the given fmt string.

'sym' name can have optional pattern {} where each
comma separated value is either a fmt parameter or
underscore (wildcard). The fmt parameters indicate
how arguments from the function are mapped to the
arguments of the sumbol.

Structs and struct pointers must either be entirely
composed of std typed fields, or have std tags on
each field that define the C type. Field order must
match the C struct definition. If there are layout
or alignment differences between the C and Go structs,
or non-std Go types are being used, then the struct
must embed a std.Struct field.

	// safest, deep copy all pointers to this struct.
	type MyStruct {
		std.Struct // if-in-doubt, embed this.

		Name string `std:"char"`
	}

	// fastest, struct pointers passed directly
	type MyStruct {
		Name std.String
	}

IMPORT IS FUNDAMENTALLY UNSAFE
Although it will validate what it can in order to
ensure safety. Callers unfamiliar with C should
stick to the 'std' tag and avoid libraries that
require C struct values to be accessed directly.

Alternatively, use a library with an existing
representation in Go, as can be found under
runtime.link/lib
*/

package std
