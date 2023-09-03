// Package std provides standard C types for cross-language interoperability.
// The types in this package reflect the default ABI for the current platform.
package std

import (
	"errors"
	"strconv"
	"unsafe"
)

// Fixed width numeric types.
type (
	Int8    c_int8_t
	Int16   c_int16_t
	Int32   c_int32_t
	Int64   c_int64_t
	Uint8   c_uint8_t
	Uint16  c_uint16_t
	Uint32  c_uint32_t
	Uint64  c_uint64_t
	Uintptr c_uintptr_t
)

// Variable width numeric types.
type (
	Char     c_char
	Short    c_short
	Int      c_int
	Long     c_long
	LongLong c_longlong

	SignedChar c_signed_char

	UnsignedChar     c_unsigned_char
	UnsignedShort    c_unsigned_short
	UnsignedInt      c_unsigned_int
	UnsignedLong     c_unsigned_long
	UnsignedLongLong c_unsigned_longlong

	Float  c_float
	Double c_double

	Size c_size_t

	Time  c_time_t
	Clock c_clock_t

	Bool c_bool
)

// String is a mutable null-terminated array of characters.
type String struct {
	ptr *Char
}

// StringOf returns a String from a Go string.
func StringOf(s string) String {
	if len(s) > 0 && s[len(s)-1] != 0 {
		s += "\x00"
	}
	return String{
		ptr: (*Char)(unsafe.Pointer(unsafe.StringData(s))),
	}
}

// String returns the Go string representation of s.
func (s String) String() string {
	if s.ptr == nil {
		return ""
	}
	var from = 0
	var upto = 0
	for *(*Char)(unsafe.Add(unsafe.Pointer(s.ptr), upto)) != 0 {
		from++
	}
	return unsafe.String((*byte)(unsafe.Pointer(s.ptr)), upto)
}

// Buffer is a mutable array of bytes along with a length.
// When passed to a C function, will be passed together as
// two arguments: a pointer to the first byte and the length.
type Buffer struct {
	ptr *Char
	len Size
}

// Make returns a Buffer of the given size.
func Make(size Size) Buffer {
	buf := make([]byte, size)
	return Buffer{
		ptr: (*Char)(unsafe.Pointer(unsafe.SliceData(buf))),
		len: size,
	}
}

// Bytes returns the Go byte slice representation of buf.
func (buf Buffer) Bytes() []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(buf.ptr)), buf.len)
}

// Error represents a C int error, where 0 is success
// and any other value is an error.
type Error c_int

// Err returns nil if err is 0, otherwise it returns an error
func (err Error) Err() error {
	if err == 0 {
		return nil
	}
	return errors.New(strconv.Itoa(int(err)))
}

// BooleanInt is a C int that is either 0 or 1.
type BooleanInt c_int

// Bool returns true if b is 1, otherwise false.
func (b BooleanInt) Bool() bool {
	return b == 1
}

// Func points to a function of the specified type.
type Func[T any] c_uintptr_t

// Handle is an opaque pointer to a C object.
// It cannot be dereferenced and can only be
// manipulated when passed to C functions.
//
// A handle should not be used directly, but
// instead should be used as the underlying
// type for a named Go type. Use the defined
// type as the parameter.
//
// For example:
//
//	type Texture std.Handle[Texture]
type Handle[T any] struct {
	_ [0]*T
	c_uintptr_t
}
