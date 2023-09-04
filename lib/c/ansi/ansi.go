// Package ansi provides the ANSI C standard library functions.
// This package does not include any types, constants or macros.
// https://www.csse.uwa.edu.au/programming/ansic-library.html
package ansi

import (
	"runtime.link/lib/c/internal/libc"
	"runtime.link/std"
)

// Functions defined in the ANSI C standard library. Function names
// have been expanded to prefer full words over abbreviations. The
// functions have been organised into sensible categories.
type Functions struct {
	_ libc.Location

	IO IO

	Char Char
	Math Math
	Jump Jump
	File File
	Time Time

	Memory Memory
	System System

	Program Program
	Signals Signals
	Strings Strings

	Division Division
}

// Char provides the functions from <ctype.h>.
type Char struct {
	_ libc.Location

	IsAlphaNumeric func(c std.Int) std.Int `ffi:"isalnum"` // IsAlpha || IsDigit
	IsAlpha        func(c std.Int) std.Int `ffi:"isalpha"` // IsUpper || IsLower
	IsControl      func(c std.Int) std.Int `ffi:"iscntrl"`
	IsDigit        func(c std.Int) std.Int `ffi:"isdigit"`
	IsGraph        func(c std.Int) std.Int `ffi:"isgraph"`
	IsLower        func(c std.Int) std.Int `ffi:"islower"`
	IsPrintable    func(c std.Int) std.Int `ffi:"isprint"`
	IsPuncuation   func(c std.Int) std.Int `ffi:"ispunct"`
	IsSpace        func(c std.Int) std.Int `ffi:"isspace"` // space, formfeed, newline, carriage return, tab, vertical tab
	IsUpper        func(c std.Int) std.Int `ffi:"isupper"`
	IsHexDigit     func(c std.Int) std.Int `ffi:"isxdigit"`

	ToLower func(c std.Int) std.Int `ffi:"tolower"`
	ToUpper func(c std.Int) std.Int `ffi:"toupper"`
}

// Math provides numerical functions from <math.h> and <stdlib.h>.
// Names kept abbreviated as is in the Go math package.
type Math struct {
	_ libc.Location

	Abs func(x std.Int) std.Int `ffi:"abs"`

	Sin   func(x std.Double) std.Double                `ffi:"sin"`
	Cos   func(x std.Double) std.Double                `ffi:"cos"`
	Tan   func(x std.Double) std.Double                `ffi:"tan"`
	Asin  func(x std.Double) std.Double                `ffi:"asin"`
	Atan2 func(x, y std.Double) std.Double             `ffi:"atan2"`
	Sinh  func(x std.Double) std.Double                `ffi:"sinh"`
	Cosh  func(x std.Double) std.Double                `ffi:"cosh"`
	Tanh  func(x std.Double) std.Double                `ffi:"tanh"`
	Exp   func(x std.Double) std.Double                `ffi:"exp"`
	Log   func(x std.Double) std.Double                `ffi:"log"`
	Log10 func(x std.Double) std.Double                `ffi:"log10"`
	Pow   func(x, y std.Double) std.Double             `ffi:"pow"`
	Sqrt  func(x std.Double) std.Double                `ffi:"sqrt"`
	Ceil  func(x std.Double) std.Double                `ffi:"ceil"`
	Floor func(x std.Double) std.Double                `ffi:"floor"`
	Fabs  func(x std.Double) std.Double                `ffi:"fabs"`
	Ldexp func(x std.Double, n std.Int) std.Double     `ffi:"ldexp"`
	Frexp func(x std.Double, exp *std.Int) std.Double  `ffi:"frexp"`
	Modf  func(x std.Double, y *std.Double) std.Double `ffi:"modf"`

	Rand     func() std.Int                     `ffi:"rand"`
	SeedRand func(seed std.UnsignedInt) std.Int `ffi:"srand"`
}

// Jump provides the functions from <setjmp.h>.
type Jump struct {
	_ libc.Location

	Set  func(env *std.JumpBuffer) std.Int              `ffi:"setjmp"`
	Long func(env *std.JumpBuffer, val std.Int) std.Int `ffi:"longjmp"`
}

// Signals provides the functions from <signal.h>.
type Signals struct {
	_ libc.Location

	Handle func(sig std.Signal, handler func(std.Signal)) std.UnsafePointer `ffi:"signal"`
	Raise  func(sig std.Signal) std.Int                                     `ffi:"raise"`
}

// File provides file-related functions from <stdio.h>.
type File struct {
	_ libc.Location

	Open     func(filename string, mode string) *std.File                   `ffi:"fopen"`
	Reopen   func(filename string, mode string, stream *std.File) *std.File `ffi:"freopen"`
	Flush    func(stream *std.File) std.Int                                 `ffi:"fflush"`
	Close    func(stream *std.File) std.Int                                 `ffi:"fclose"`
	Remove   func(filename string) std.Int                                  `ffi:"remove"`
	Rename   func(oldname, newname string) std.Int                          `ffi:"rename"`
	Temp     func() *std.File                                               `ffi:"tmpfile"`
	TempName func(*[std.TempNameLength]byte) string                         `ffi:"tmpnam"`

	SetBufferMode func(stream *std.File, buf *[std.BufferSize]byte, mode std.BufferMode, size std.Size) std.Int `ffi:"setvbuf"`
	SetBuffer     func(stream *std.File, buf *[std.BufferSize]byte) std.Int                                     `ffi:"setbuf"`

	Printf    func(stream *std.File, format string, args ...std.UnsafePointer) std.Int `ffi:"fprintf"`
	Scanf     func(stream *std.File, format string, args ...std.UnsafePointer) std.Int `ffi:"fscanf"`
	GetChar   func(stream *std.File) std.Int                                           `ffi:"fgetc"`
	GetString func(s std.Buffer, stream *std.File) std.UnsafePointer                   `ffi:"fgets"`
	PutChar   func(c std.Int, stream *std.File) std.Int                                `ffi:"fputc"`
	Unget     func(c std.Int, stream *std.File) std.Int                                `ffi:"ungetc"`

	Read  func(ptr std.UnsafePointer, size std.Size, nobj std.Size, stream *std.File) std.Size `ffi:"fread"`
	Write func(ptr std.UnsafePointer, size std.Size, nobj std.Size, stream *std.File) std.Size `ffi:"fwrite"`

	Seek func(stream *std.File, offset std.Long, origin std.SeekMode) std.Int `ffi:"fseek"`
	Tell func(stream *std.File) std.Long                                      `ffi:"ftell"`

	Rewind func(stream *std.File) std.Int                        `ffi:"rewind"`
	GetPos func(stream *std.File, ptr *std.FilePosition) std.Int `ffi:"fgetpos"`
	SetPos func(stream *std.File, ptr *std.FilePosition) std.Int `ffi:"fsetpos"`

	ClearError func(stream *std.File) std.Int `ffi:"clearerr"`
	IsEOF      func(stream *std.File) std.Int `ffi:"feof"`
	Error      func(stream *std.File) std.Int `ffi:"ferror"`
}

// IO provides stdin/stdout functions from <stdio.h>.
type IO struct {
	_ libc.Location

	Printf    func(format string, args ...std.UnsafePointer) std.Int `ffi:"printf"`
	Scanf     func(format string, args ...std.UnsafePointer) std.Int `ffi:"scanf"`
	GetChar   func() std.Int                                         `ffi:"getchar"`
	GetString func(s std.UnsafePointer) string                       `ffi:"gets"`
	PutChar   func(c std.Int) std.Int                                `ffi:"putchar"`
	PutString func(s string) std.Int                                 `ffi:"puts"`

	Error func(s string) std.Int `ffi:"perror"`
}

// Strings provides string-related functions from <string.h>, <stdio.h> and <stdlib.h>.
type Strings struct {
	_ libc.Location

	Printf func(s std.UnsafePointer, args ...std.UnsafePointer) std.Int `ffi:"sprintf"`
	Scanf  func(s std.UnsafePointer, args ...std.UnsafePointer) std.Int `ffi:"sscanf"`

	ToDouble          func(s string) std.Double                                      `ffi:"atof"`
	ToInt             func(s string) std.Int                                         `ffi:"atoi"`
	ToLong            func(s string) std.Long                                        `ffi:"atol"`
	ParseDouble       func(s string, endp **std.Char) std.Double                     `ffi:"atof"`
	ParseLong         func(s string, endp **std.Char, base std.Int) std.Long         `ffi:"strtol"`
	ParseUnsignedLong func(s string, endp **std.Char, base std.Int) std.UnsignedLong `ffi:"strtoul"`

	Copy           func(dest, src std.UnsafePointer) string             `ffi:"strcpy"`
	CopyLimited    func(dest, src std.UnsafePointer, n std.Size) string `ffi:"strncpy"`
	Cat            func(dest, src std.UnsafePointer) string             `ffi:"strcat"`
	CatLimited     func(dest, src std.UnsafePointer, n std.Size) string `ffi:"strncat"`
	Compare        func(cs, ct string) std.Int                          `ffi:"strcmp"`
	CompareLimited func(cs, ct string, n std.Size) std.Int              `ffi:"strncmp"`

	SearchChar        func(cs string, c std.Int) *std.Char `ffi:"strchr"`
	SearchCharReverse func(cs string, c std.Int) *std.Char `ffi:"strrchr"`
	Span              func(cs, ct string) std.Size         `ffi:"strspn"`
	ComplimentarySpan func(cs, ct string) std.Size         `ffi:"strcspn"`
	PointerBreak      func(cs, ct string) *std.Char        `ffi:"strpbrk"`

	Search func(cs, ct string) *std.Char `ffi:"strstr"`
	Length func(cs string) std.Size      `ffi:"strlen"`

	Error  func(n std.Int) string                  `ffi:"strerror"`
	Tokens func(s std.String, delim string) string `ffi:"strtok"`
}

// Memory provides memory-related functions from <stdlib.h>.
type Memory struct {
	_ libc.Location

	AllocateZeros func(nobj std.Size, size std.Size) std.UnsafePointer       `ffi:"calloc"`
	Allocate      func(size std.Size) std.UnsafePointer                      `ffi:"malloc"`
	Reallocate    func(p std.UnsafePointer, size std.Size) std.UnsafePointer `ffi:"realloc"`
	Free          func(p std.UnsafePointer)                                  `ffi:"free"`

	Sort   func(base std.UnsafePointer, n, size std.Size, cmp func(a, b std.UnsafePointer) std.Int)                                 `ffi:"qsort"`
	Search func(key, base std.UnsafePointer, n, size std.Size, cmp func(keyval, datum std.UnsafePointer) std.Int) std.UnsafePointer `ffi:"bsearch"`

	Copy    func(dst, src std.UnsafePointer, n std.Size) std.UnsafePointer `ffi:"memcpy"`
	Move    func(dst, src std.UnsafePointer, n std.Size) std.UnsafePointer `ffi:"memmove"`
	Compare func(cs, ct std.UnsafePointer, n std.Size) std.Int             `ffi:"memcmp"`

	SearchChar func(cs std.UnsafePointer, c std.Int, n std.Size) *std.Char `ffi:"memchr"`

	Set func(s std.UnsafePointer, c std.Int, n std.Size) std.UnsafePointer `ffi:"memset"`
}

// Program provides program-related functions from <stdlib.h>.
type Program struct {
	_ libc.Location

	Abort  func()                   `ffi:"abort"`
	Exit   func(status std.Int)     `ffi:"exit"`
	OnExit func(func())             `ffi:"atexit,__cxa_atexit"`
	Getenv func(name string) string `ffi:"getenv"`
}

// System provides system-related functions from <stdlib.h>.
type System struct {
	_ libc.Location

	Command func(command string) std.Int `ffi:"system"`

	Clock func() std.Clock           `ffi:"clock"`
	Time  func(t *std.Time) std.Time `ffi:"time"`
}

// Division provides division-related functions from <stdlib.h>.
type Division struct {
	_ libc.Location

	Int  func(num, denom std.Int) std.DivisionInt   `ffi:"div"`
	Long func(num, denom std.Long) std.DivisionLong `ffi:"ldiv"`
}

// Time provides time-related functions from <time.h>.
type Time struct {
	_ libc.Location

	Sub    func(t1, t2 std.Time) std.Time `ffi:"difftime"`
	String func(t std.Time) string        `ffi:"ctime"`

	UTC   func(t std.Time) *std.Date `ffi:"timegm"`
	Local func(t std.Time) *std.Date `ffi:"localtime"`
}

// Date provides date-related functions from <time.h>.
type Date struct {
	_ libc.Location

	Time   func(t *std.Date) std.Time                               `ffi:"mktime"`
	String func(t *std.Date) string                                 `ffi:"asctime"`
	Format func(s std.Buffer, format string, tp *std.Date) std.Size `ffi:"strftime"`
}
