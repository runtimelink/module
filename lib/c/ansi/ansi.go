// Package ansi provides the ANSI C standard library functions.
// This package does not include any types, constants or macros.
// https://www.csse.uwa.edu.au/programming/ansic-library.html
package ansi

import (
	"runtime.link/ffi"
	"runtime.link/lib/c/internal/libc"
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

	IsAlphaNumeric func(c ffi.Int) ffi.Int `ffi:"isalnum"` // IsAlpha || IsDigit
	IsAlpha        func(c ffi.Int) ffi.Int `ffi:"isalpha"` // IsUpper || IsLower
	IsControl      func(c ffi.Int) ffi.Int `ffi:"iscntrl"`
	IsDigit        func(c ffi.Int) ffi.Int `ffi:"isdigit"`
	IsGraph        func(c ffi.Int) ffi.Int `ffi:"isgraph"`
	IsLower        func(c ffi.Int) ffi.Int `ffi:"islower"`
	IsPrintable    func(c ffi.Int) ffi.Int `ffi:"isprint"`
	IsPuncuation   func(c ffi.Int) ffi.Int `ffi:"ispunct"`
	IsSpace        func(c ffi.Int) ffi.Int `ffi:"isspace"` // space, formfeed, newline, carriage return, tab, vertical tab
	IsUpper        func(c ffi.Int) ffi.Int `ffi:"isupper"`
	IsHexDigit     func(c ffi.Int) ffi.Int `ffi:"isxdigit"`

	ToLower func(c ffi.Int) ffi.Int `ffi:"tolower"`
	ToUpper func(c ffi.Int) ffi.Int `ffi:"toupper"`
}

// Math provides numerical functions from <math.h> and <stdlib.h>.
// Names kept abbreviated as is in the Go math package.
type Math struct {
	_ libc.Location

	Abs func(x ffi.Int) ffi.Int `ffi:"abs"`

	Sin   func(x ffi.Double) ffi.Double                `ffi:"sin"`
	Cos   func(x ffi.Double) ffi.Double                `ffi:"cos"`
	Tan   func(x ffi.Double) ffi.Double                `ffi:"tan"`
	Asin  func(x ffi.Double) ffi.Double                `ffi:"asin"`
	Atan2 func(x, y ffi.Double) ffi.Double             `ffi:"atan2"`
	Sinh  func(x ffi.Double) ffi.Double                `ffi:"sinh"`
	Cosh  func(x ffi.Double) ffi.Double                `ffi:"cosh"`
	Tanh  func(x ffi.Double) ffi.Double                `ffi:"tanh"`
	Exp   func(x ffi.Double) ffi.Double                `ffi:"exp"`
	Log   func(x ffi.Double) ffi.Double                `ffi:"log"`
	Log10 func(x ffi.Double) ffi.Double                `ffi:"log10"`
	Pow   func(x, y ffi.Double) ffi.Double             `ffi:"pow"`
	Sqrt  func(x ffi.Double) ffi.Double                `ffi:"sqrt"`
	Ceil  func(x ffi.Double) ffi.Double                `ffi:"ceil"`
	Floor func(x ffi.Double) ffi.Double                `ffi:"floor"`
	Fabs  func(x ffi.Double) ffi.Double                `ffi:"fabs"`
	Ldexp func(x ffi.Double, n ffi.Int) ffi.Double     `ffi:"ldexp"`
	Frexp func(x ffi.Double, exp *ffi.Int) ffi.Double  `ffi:"frexp"`
	Modf  func(x ffi.Double, y *ffi.Double) ffi.Double `ffi:"modf"`

	Rand     func() ffi.Int                     `ffi:"rand"`
	SeedRand func(seed ffi.IntUnsigned) ffi.Int `ffi:"srand"`
}

// Jump provides the functions from <setjmp.h>.
type Jump struct {
	_ libc.Location

	Set  func(env *ffi.JumpBuffer) ffi.Int              `ffi:"setjmp"`
	Long func(env *ffi.JumpBuffer, val ffi.Int) ffi.Int `ffi:"longjmp"`
}

// Signals provides the functions from <signal.h>.
type Signals struct {
	_ libc.Location

	Handle func(sig ffi.Signal, handler func(ffi.Signal)) ffi.UnsafePointer `ffi:"signal"`
	Raise  func(sig ffi.Signal) ffi.Int                                     `ffi:"raise"`
}

// File provides file-related functions from <stdio.h>.
type File struct {
	_ libc.Location

	Open     func(filename string, mode string) *ffi.File                   `ffi:"fopen"`
	Reopen   func(filename string, mode string, stream *ffi.File) *ffi.File `ffi:"freopen"`
	Flush    func(stream *ffi.File) ffi.Int                                 `ffi:"fflush"`
	Close    func(stream *ffi.File) ffi.Int                                 `ffi:"fclose"`
	Remove   func(filename string) ffi.Int                                  `ffi:"remove"`
	Rename   func(oldname, newname string) ffi.Int                          `ffi:"rename"`
	Temp     func() *ffi.File                                               `ffi:"tmpfile"`
	TempName func(*[ffi.TempNameSize]byte) string                           `ffi:"tmpnam"`

	SetBufferMode func(stream *ffi.File, buf *[ffi.BufferSize]byte, mode ffi.BufferMode, size ffi.Size) ffi.Int `ffi:"setvbuf"`
	SetBuffer     func(stream *ffi.File, buf *[ffi.BufferSize]byte) ffi.Int                                     `ffi:"setbuf"`

	Printf    func(stream *ffi.File, format string, args ...ffi.UnsafePointer) ffi.Int `ffi:"fprintf"`
	Scanf     func(stream *ffi.File, format string, args ...ffi.UnsafePointer) ffi.Int `ffi:"fscanf"`
	GetChar   func(stream *ffi.File) ffi.Int                                           `ffi:"fgetc"`
	GetString func(s ffi.Buffer, stream *ffi.File) ffi.UnsafePointer                   `ffi:"fgets"`
	PutChar   func(c ffi.Int, stream *ffi.File) ffi.Int                                `ffi:"fputc"`
	Unget     func(c ffi.Int, stream *ffi.File) ffi.Int                                `ffi:"ungetc"`

	Read  func(ptr ffi.UnsafePointer, size ffi.Size, nobj ffi.Size, stream *ffi.File) ffi.Size `ffi:"fread"`
	Write func(ptr ffi.UnsafePointer, size ffi.Size, nobj ffi.Size, stream *ffi.File) ffi.Size `ffi:"fwrite"`

	Seek func(stream *ffi.File, offset ffi.Long, origin ffi.SeekMode) ffi.Int `ffi:"fseek"`
	Tell func(stream *ffi.File) ffi.Long                                      `ffi:"ftell"`

	Rewind func(stream *ffi.File) ffi.Int                        `ffi:"rewind"`
	GetPos func(stream *ffi.File, ptr *ffi.FilePosition) ffi.Int `ffi:"fgetpos"`
	SetPos func(stream *ffi.File, ptr *ffi.FilePosition) ffi.Int `ffi:"fsetpos"`

	ClearError func(stream *ffi.File) ffi.Int `ffi:"clearerr"`
	IsEOF      func(stream *ffi.File) ffi.Int `ffi:"feof"`
	Error      func(stream *ffi.File) ffi.Int `ffi:"ferror"`
}

// IO provides stdin/stdout functions from <stdio.h>.
type IO struct {
	_ libc.Location

	Printf    func(format string, args ...ffi.UnsafePointer) ffi.Int `ffi:"printf"`
	Scanf     func(format string, args ...ffi.UnsafePointer) ffi.Int `ffi:"scanf"`
	GetChar   func() ffi.Int                                         `ffi:"getchar"`
	GetString func(s ffi.UnsafePointer) string                       `ffi:"gets"`
	PutChar   func(c ffi.Int) ffi.Int                                `ffi:"putchar"`
	PutString func(s string) ffi.Int                                 `ffi:"puts"`

	Error func(s string) ffi.Int `ffi:"perror"`
}

// Strings provides string-related functions from <string.h>, <stdio.h> and <stdlib.h>.
type Strings struct {
	_ libc.Location

	Printf func(s ffi.UnsafePointer, args ...ffi.UnsafePointer) ffi.Int `ffi:"sprintf"`
	Scanf  func(s ffi.UnsafePointer, args ...ffi.UnsafePointer) ffi.Int `ffi:"sscanf"`

	ToDouble          func(s string) ffi.Double                                      `ffi:"atof"`
	ToInt             func(s string) ffi.Int                                         `ffi:"atoi"`
	ToLong            func(s string) ffi.Long                                        `ffi:"atol"`
	ParseDouble       func(s string, endp **ffi.Char) ffi.Double                     `ffi:"atof"`
	ParseLong         func(s string, endp **ffi.Char, base ffi.Int) ffi.Long         `ffi:"strtol"`
	ParseUnsignedLong func(s string, endp **ffi.Char, base ffi.Int) ffi.LongUnsigned `ffi:"strtoul"`

	Copy           func(dest, src ffi.UnsafePointer) string             `ffi:"strcpy"`
	CopyLimited    func(dest, src ffi.UnsafePointer, n ffi.Size) string `ffi:"strncpy"`
	Cat            func(dest, src ffi.UnsafePointer) string             `ffi:"strcat"`
	CatLimited     func(dest, src ffi.UnsafePointer, n ffi.Size) string `ffi:"strncat"`
	Compare        func(cs, ct string) ffi.Int                          `ffi:"strcmp"`
	CompareLimited func(cs, ct string, n ffi.Size) ffi.Int              `ffi:"strncmp"`

	SearchChar        func(cs string, c ffi.Int) *ffi.Char `ffi:"strchr"`
	SearchCharReverse func(cs string, c ffi.Int) *ffi.Char `ffi:"strrchr"`
	Span              func(cs, ct string) ffi.Size         `ffi:"strspn"`
	ComplimentarySpan func(cs, ct string) ffi.Size         `ffi:"strcspn"`
	PointerBreak      func(cs, ct string) *ffi.Char        `ffi:"strpbrk"`

	Search func(cs, ct string) *ffi.Char `ffi:"strstr"`
	Length func(cs string) ffi.Size      `ffi:"strlen"`

	Error  func(n ffi.Int) string                  `ffi:"strerror"`
	Tokens func(s ffi.String, delim string) string `ffi:"strtok"`
}

// Memory provides memory-related functions from <stdlib.h>.
type Memory struct {
	_ libc.Location

	AllocateZeros func(nobj ffi.Size, size ffi.Size) ffi.UnsafePointer       `ffi:"calloc"`
	Allocate      func(size ffi.Size) ffi.UnsafePointer                      `ffi:"malloc"`
	Reallocate    func(p ffi.UnsafePointer, size ffi.Size) ffi.UnsafePointer `ffi:"realloc"`
	Free          func(p ffi.UnsafePointer)                                  `ffi:"free"`

	Sort   func(base ffi.UnsafePointer, n, size ffi.Size, cmp func(a, b ffi.UnsafePointer) ffi.Int)                                 `ffi:"qsort"`
	Search func(key, base ffi.UnsafePointer, n, size ffi.Size, cmp func(keyval, datum ffi.UnsafePointer) ffi.Int) ffi.UnsafePointer `ffi:"bsearch"`

	Copy    func(dst, src ffi.UnsafePointer, n ffi.Size) ffi.UnsafePointer `ffi:"memcpy"`
	Move    func(dst, src ffi.UnsafePointer, n ffi.Size) ffi.UnsafePointer `ffi:"memmove"`
	Compare func(cs, ct ffi.UnsafePointer, n ffi.Size) ffi.Int             `ffi:"memcmp"`

	SearchChar func(cs ffi.UnsafePointer, c ffi.Int, n ffi.Size) *ffi.Char `ffi:"memchr"`

	Set func(s ffi.UnsafePointer, c ffi.Int, n ffi.Size) ffi.UnsafePointer `ffi:"memset"`
}

// Program provides program-related functions from <stdlib.h>.
type Program struct {
	_ libc.Location

	Abort  func()                   `ffi:"abort"`
	Exit   func(status ffi.Int)     `ffi:"exit"`
	OnExit func(func())             `ffi:"atexit,__cxa_atexit"`
	Getenv func(name string) string `ffi:"getenv"`
}

// System provides system-related functions from <stdlib.h>.
type System struct {
	_ libc.Location

	Command func(command string) ffi.Int `ffi:"system"`

	Clock func() ffi.Clock           `ffi:"clock"`
	Time  func(t *ffi.Time) ffi.Time `ffi:"time"`
}

// Division provides division-related functions from <stdlib.h>.
type Division struct {
	_ libc.Location

	Int  func(num, denom ffi.Int) ffi.DivisionInt   `ffi:"div"`
	Long func(num, denom ffi.Long) ffi.DivisionLong `ffi:"ldiv"`
}

// Time provides time-related functions from <time.h>.
type Time struct {
	_ libc.Location

	Sub    func(t1, t2 ffi.Time) ffi.Time `ffi:"difftime"`
	String func(t ffi.Time) string        `ffi:"ctime"`

	UTC   func(t ffi.Time) *ffi.Date `ffi:"timegm"`
	Local func(t ffi.Time) *ffi.Date `ffi:"localtime"`
}

// Date provides date-related functions from <time.h>.
type Date struct {
	_ libc.Location

	Time   func(t *ffi.Date) ffi.Time                               `ffi:"mktime"`
	String func(t *ffi.Date) string                                 `ffi:"asctime"`
	Format func(s ffi.Buffer, format string, tp *ffi.Date) ffi.Size `ffi:"strftime"`
}
