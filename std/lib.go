// Package ansi
package std

type location struct {
	linux   Location `libc.so.6 libm.so.6`
	darwin  Location `libSystem.dylib`
	windows Location `msvcrt.dll`
}

// Library provides provides the ANSI C standard library functions.
// https://www.csse.uwa.edu.au/programming/ansic-library.html
// Function names have been expanded to prefer full words over
// abbreviations. The functions have been organised into sensible
// categories.
type Library struct {
	location

	IO LibraryIO

	Math LibraryMath
	Time LibraryTime

	Char LibraryChars
	File LibraryFiles
	Jump LibraryJumps

	Memory LibraryMemory
	System LibrarySystem

	Program LibraryProgram
	Signals LibrarySignals
	Strings LibraryStrings

	Division LibraryDivision
}

// LibraryChars provides the functions from <ctype.h>.
type LibraryChars struct {
	location

	IsAlphaNumeric func(c Int) Int `sym:"isalnum"` // IsAlpha || IsDigit
	IsAlpha        func(c Int) Int `sym:"isalpha"` // IsUpper || IsLower
	IsControl      func(c Int) Int `sym:"iscntrl"`
	IsDigit        func(c Int) Int `sym:"isdigit"`
	IsGraph        func(c Int) Int `sym:"isgraph"`
	IsLower        func(c Int) Int `sym:"islower"`
	IsPrintable    func(c Int) Int `sym:"isprint"`
	IsPuncuation   func(c Int) Int `sym:"ispunct"`
	IsSpace        func(c Int) Int `sym:"isspace"` // space, formfeed, newline, carriage return, tab, vertical tab
	IsUpper        func(c Int) Int `sym:"isupper"`
	IsHexDigit     func(c Int) Int `sym:"isxdigit"`

	ToLower func(c Int) Int `sym:"tolower"`
	ToUpper func(c Int) Int `sym:"toupper"`
}

// LibraryMath provides numerical functions from <math.h> and <stdlib.h>.
// Names kept abbreviated as is in the Go math package.
type LibraryMath struct {
	location

	Abs func(x Int) Int `sym:"abs"`

	Sin   func(x Double) Double            `sym:"sin"`
	Cos   func(x Double) Double            `sym:"cos"`
	Tan   func(x Double) Double            `sym:"tan"`
	Asin  func(x Double) Double            `sym:"asin"`
	Atan2 func(x, y Double) Double         `sym:"atan2"`
	Sinh  func(x Double) Double            `sym:"sinh"`
	Cosh  func(x Double) Double            `sym:"cosh"`
	Tanh  func(x Double) Double            `sym:"tanh"`
	Exp   func(x Double) Double            `sym:"exp"`
	Log   func(x Double) Double            `sym:"log"`
	Log10 func(x Double) Double            `sym:"log10"`
	Pow   func(x, y Double) Double         `sym:"pow"`
	Sqrt  func(x Double) Double            `sym:"sqrt"`
	Ceil  func(x Double) Double            `sym:"ceil"`
	Floor func(x Double) Double            `sym:"floor"`
	Fabs  func(x Double) Double            `sym:"fabs"`
	Ldexp func(x Double, n Int) Double     `sym:"ldexp"`
	Frexp func(x Double, exp *Int) Double  `sym:"frexp"`
	Modf  func(x Double, y *Double) Double `sym:"modf"`

	Rand     func() Int                 `sym:"rand"`
	SeedRand func(seed UnsignedInt) Int `sym:"srand"`
}

// LibraryJumps provides the functions from <setjmp.h>.
type LibraryJumps struct {
	location

	Set  func(env *JumpBuffer) Int          `sym:"setjmp"`
	Long func(env *JumpBuffer, val Int) Int `sym:"longjmp"`
}

// LibrarySignals provides the functions from <signal.h>.
type LibrarySignals struct {
	location

	Handle func(sig Signal, handler func(Signal)) UnsafePointer `sym:"signal"`
	Raise  func(sig Signal) Int                                 `sym:"raise"`
}

// LibraryFiles provides file-related functions from <stdio.h>.
type LibraryFiles struct {
	location

	Open     func(filename string, mode string) *File               `sym:"fopen"`
	Reopen   func(filename string, mode string, stream *File) *File `sym:"freopen"`
	Flush    func(stream *File) Int                                 `sym:"fflush"`
	Close    func(stream *File) Int                                 `sym:"fclose"`
	Remove   func(filename string) Int                              `sym:"remove"`
	Rename   func(oldname, newname string) Int                      `sym:"rename"`
	Temp     func() *File                                           `sym:"tmpfile"`
	TempName func(*[TempNameLength]byte) string                     `sym:"tmpnam"`

	SetBufferMode func(stream *File, buf *[BufferSize]byte, mode BufferMode, size Size) Int `sym:"setvbuf"`
	SetBuffer     func(stream *File, buf *[BufferSize]byte) Int                             `sym:"setbuf"`

	Printf    func(stream *File, format string, args ...UnsafePointer) Int `sym:"fprintf"`
	Scanf     func(stream *File, format string, args ...UnsafePointer) Int `sym:"fscanf"`
	GetChar   func(stream *File) Int                                       `sym:"fgetc"`
	GetString func(s Buffer, stream *File) UnsafePointer                   `sym:"fgets"`
	PutChar   func(c Int, stream *File) Int                                `sym:"fputc"`
	Unget     func(c Int, stream *File) Int                                `sym:"ungetc"`

	Read  func(ptr UnsafePointer, size Size, nobj Size, stream *File) Size `sym:"fread"`
	Write func(ptr UnsafePointer, size Size, nobj Size, stream *File) Size `sym:"fwrite"`

	Seek func(stream *File, offset Long, origin SeekMode) Int `sym:"fseek"`
	Tell func(stream *File) Long                              `sym:"ftell"`

	Rewind func(stream *File) Int                    `sym:"rewind"`
	GetPos func(stream *File, ptr *FilePosition) Int `sym:"fgetpos"`
	SetPos func(stream *File, ptr *FilePosition) Int `sym:"fsetpos"`

	ClearError func(stream *File) Int `sym:"clearerr"`
	IsEOF      func(stream *File) Int `sym:"feof"`
	Error      func(stream *File) Int `sym:"ferror"`
}

// LibraryIO provides stdin/stdout functions from <stdio.h>.
type LibraryIO struct {
	location

	Printf    func(format string, args ...UnsafePointer) Int `sym:"printf"`
	Scanf     func(format string, args ...UnsafePointer) Int `sym:"scanf"`
	GetChar   func() Int                                     `sym:"getchar"`
	GetString func(s UnsafePointer) string                   `sym:"gets"`
	PutChar   func(c Int) Int                                `sym:"putchar"`
	PutString func(s string) Int                             `sym:"puts"`

	Error func(s string) Int `sym:"perror"`
}

// LibraryStrings provides string-related functions from <string.h>, <stdio.h> and <stdlib.h>.
type LibraryStrings struct {
	location

	Printf func(s UnsafePointer, args ...UnsafePointer) Int `sym:"sprintf"`
	Scanf  func(s UnsafePointer, args ...UnsafePointer) Int `sym:"sscanf"`

	ToDouble          func(s string) Double                              `sym:"atof"`
	ToInt             func(s string) Int                                 `sym:"atoi"`
	ToLong            func(s string) Long                                `sym:"atol"`
	ParseDouble       func(s string, endp **Char) Double                 `sym:"atof"`
	ParseLong         func(s string, endp **Char, base Int) Long         `sym:"strtol"`
	ParseUnsignedLong func(s string, endp **Char, base Int) UnsignedLong `sym:"strtoul"`

	Copy           func(dest, src UnsafePointer) string         `sym:"strcpy"`
	CopyLimited    func(dest, src UnsafePointer, n Size) string `sym:"strncpy"`
	Cat            func(dest, src UnsafePointer) string         `sym:"strcat"`
	CatLimited     func(dest, src UnsafePointer, n Size) string `sym:"strncat"`
	Compare        func(cs, ct string) Int                      `sym:"strcmp"`
	CompareLimited func(cs, ct string, n Size) Int              `sym:"strncmp"`

	SearchChar        func(cs string, c Int) *Char `sym:"strchr"`
	SearchCharReverse func(cs string, c Int) *Char `sym:"strrchr"`
	Span              func(cs, ct string) Size     `sym:"strspn"`
	ComplimentarySpan func(cs, ct string) Size     `sym:"strcspn"`
	PointerBreak      func(cs, ct string) *Char    `sym:"strpbrk"`

	Search func(cs, ct string) *Char `sym:"strstr"`
	Length func(cs string) Size      `sym:"strlen"`

	Error  func(n Int) string                  `sym:"strerror"`
	Tokens func(s String, delim string) string `sym:"strtok"`
}

// LibraryMemory provides memory-related functions from <stdlib.h>.
type LibraryMemory struct {
	location

	AllocateZeros func(nobj Size, size Size) UnsafePointer       `sym:"calloc"`
	Allocate      func(size Size) UnsafePointer                  `sym:"malloc"`
	Reallocate    func(p UnsafePointer, size Size) UnsafePointer `sym:"realloc"`
	Free          func(p UnsafePointer)                          `sym:"free"`

	Sort   func(base UnsafePointer, n, size Size, cmp func(a, b UnsafePointer) Int)                             `sym:"qsort"`
	Search func(key, base UnsafePointer, n, size Size, cmp func(keyval, datum UnsafePointer) Int) UnsafePointer `sym:"bsearch"`

	Copy    func(dst, src UnsafePointer, n Size) UnsafePointer `sym:"memcpy"`
	Move    func(dst, src UnsafePointer, n Size) UnsafePointer `sym:"memmove"`
	Compare func(cs, ct UnsafePointer, n Size) Int             `sym:"memcmp"`

	SearchChar func(cs UnsafePointer, c Int, n Size) *Char `sym:"memchr"`

	Set func(s UnsafePointer, c Int, n Size) UnsafePointer `sym:"memset"`
}

// LibraryProgram provides program-related functions from <stdlib.h>.
type LibraryProgram struct {
	location

	Abort  func()                   `sym:"abort"`
	Exit   func(status Int)         `sym:"exit"`
	OnExit func(func())             `sym:"atexit,__cxa_atexit"`
	Getenv func(name string) string `sym:"getenv"`
}

// LibrarySystem provides system-related functions from <stdlib.h>.
type LibrarySystem struct {
	location

	Command func(command string) Int `sym:"system"`

	Clock func() Clock       `sym:"clock"`
	Time  func(t *Time) Time `sym:"time"`
}

// LibraryDivision provides division-related functions from <stdlib.h>.
type LibraryDivision struct {
	location

	Int  func(num, denom Int) DivisionInt   `sym:"div"`
	Long func(num, denom Long) DivisionLong `sym:"ldiv"`
}

// LibraryTime provides time-related functions from <time.h>.
type LibraryTime struct {
	location

	Sub    func(t1, t2 Time) Time `sym:"difftime"`
	String func(t Time) string    `sym:"ctime"`

	UTC   func(t Time) *Date `sym:"timegm"`
	Local func(t Time) *Date `sym:"localtime"`
}

// LibraryDates provides date-related functions from <time.h>.
type LibraryDates struct {
	location

	Time   func(t *Date) Time                           `sym:"mktime"`
	String func(t *Date) string                         `sym:"asctime"`
	Format func(s Buffer, format string, tp *Date) Size `sym:"strftime"`
}
