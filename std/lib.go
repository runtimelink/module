package std

import "unsafe"

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

	File LibraryFiles
	Jump LibraryJumps

	ASCII LibraryASCII

	Memory LibraryMemory
	System LibrarySystem

	Program LibraryProgram
	Signals LibrarySignals
	Strings LibraryStrings

	Division LibraryDivision
}

// LibraryASCII provides the functions from <ctype.h>.
type LibraryASCII struct {
	location

	IsAlphaNumeric func(c rune) rune `std:"int isalnum(int)"` // IsAlpha || IsDigit
	IsAlpha        func(c rune) rune `std:"int isalpha(int)"` // IsUpper || IsLower
	IsControl      func(c rune) rune `std:"int iscntrl(int)"`
	IsDigit        func(c rune) rune `std:"int isdigit(int)"`
	IsGraph        func(c rune) rune `std:"int isgraph(int)"`
	IsLower        func(c rune) rune `std:"int islower(int)"`
	IsPrintable    func(c rune) rune `std:"int isprint(int)"`
	IsPuncuation   func(c rune) rune `std:"int ispunct(int)"`
	IsSpace        func(c rune) rune `std:"int isspace(int)"` // space, formfeed, newline, carriage return, tab, vertical tab
	IsUpper        func(c rune) rune `std:"int isupper(int)"`
	IsHexDigit     func(c rune) rune `std:"int isxdigit(int)"`

	ToLower func(c rune) rune `std:"int tolower(int)"`
	ToUpper func(c rune) rune `std:"int toupper(int)"`
}

// LibraryMath provides numerical functions from <math.h> and <stdlib.h>.
// Names kept abbreviated as is in the Go math package.
type LibraryMath struct {
	location

	Abs func(x Int) Int `sym:"abs"`

	Sin   func(x float64) float64             `std:"double sin(double)"`
	Cos   func(x float64) float64             `std:"double cos(double)"`
	Tan   func(x float64) float64             `std:"double tan(double)"`
	Asin  func(x float64) float64             `std:"double asin(double)"`
	Atan2 func(x, y float64) float64          `std:"double atan2(double)"`
	Sinh  func(x float64) float64             `std:"double sinh(double,double)"`
	Cosh  func(x float64) float64             `std:"double cosh(double)"`
	Tanh  func(x float64) float64             `std:"double tanh(double)"`
	Exp   func(x float64) float64             `std:"double exp(double)"`
	Log   func(x float64) float64             `std:"double log(double)"`
	Log10 func(x float64) float64             `std:"double log10(double,double)"`
	Pow   func(x, y float64) float64          `std:"double pow(double,double)"`
	Sqrt  func(x float64) float64             `std:"double sqrt(double)"`
	Ceil  func(x float64) float64             `std:"double ceil(double)"`
	Floor func(x float64) float64             `std:"double floor(double)"`
	Fabs  func(x float64) float64             `std:"double fabs(double)"`
	Ldexp func(x float64, n int32) float64    `std:"double ldexp(double,int)"`
	Frexp func(x float64, exp *int32) float64 `std:"double frexp(double,&int)"`
	Modf  func(x float64, y *float64) float64 `std:"double modf(double,&double)"`

	Rand     func() int32            `std:"int rand"`
	SeedRand func(seed uint32) int32 `std:"int srand(unsigned_int)"`
}

// LibraryJumps provides the functions from <setjmp.h>.
type LibraryJumps struct {
	location

	Set  func(env *JumpBuffer) error            `std:"int setjmp(&void)"`
	Long func(env *JumpBuffer, err error) error `std:"int longjmp(&void,int)"`
}

// LibrarySignals provides the functions from <signal.h>.
type LibrarySignals struct {
	location

	Handle func(sig Signal, handler func(Signal)) `std:"-void signal(int,void)"`
	Raise  func(sig Signal) error                 `std:"int raise(int)"`
}

// LibraryFiles provides file-related functions from <stdio.h>.
type LibraryFiles struct {
	location

	Open     func(filename string, mode string) *File               `std:"$void fopen(&char,&char)"`
	Reopen   func(filename string, mode string, stream *File) *File `std:"$void freopen(&char,&char,#void)"`
	Flush    func(stream *File) Int                                 `std:"int fflush(&void)"`
	Close    func(stream *File) error                               `std:"int fclose(#void)"`
	Remove   func(filename string) error                            `std:"int remove(&char)"`
	Rename   func(oldname, newname string) error                    `std:"int rename(&char,&char)"`
	Temp     func() *File                                           `std:"$void tmpfile"`
	TempName func(*[TempNameLength]byte) string                     `std:"char tmpnam(&char)"`

	SetBufferMode func(stream *File, buf *[BufferSize]byte, mode BufferMode) error `std:"int setvbuf(&void,free@fclose{%v},int,size_t%[2]v)"`
	SetBuffer     func(stream *File, buf *[BufferSize]byte) error                  `std:"int setbuf(&void,free@fclose{%v})"`

	Printf    func(stream *File, format string, args ...any) (int, error) `std:"int<0 fprintf(&void,&char,&vfmt%[1]v...)"`
	Scanf     func(stream *File, format string, args ...any) (int, error) `std:"int<0 fscanf(&void,&char,&vfmt%[1]v...)"`
	GetChar   func(stream *File) rune                                     `std:"int fgetc(&void)"`
	GetString func(s []byte, stream *File) string                         `std:"char fgets(&char,int,&void)"`
	PutChar   func(c rune, stream *File) rune                             `std:"int fputc(int,&void)"`
	Unget     func(c rune, stream *File) rune                             `std:"int ungetc(int,&void)"`

	Read  func(ptr []byte, stream *File) int `std:"int fread(&void,size_t=1,size_t%[1]v,&void)"`
	Write func(ptr []byte, stream *File) int `std:"int fwrite(&void,size_t=1,size_t%[1]v,&void)"`

	Seek func(stream *File, offset int, origin SeekMode) error `std:"int fseek(&void,long,int)"`
	Tell func(stream *File) int                                `std:"long ftell(&void)"`

	Rewind func(stream *File) error                    `std:"int rewind(&void)"`
	GetPos func(stream *File, ptr *FilePosition) error `std:"int fgetpos(&void,&void)"`
	SetPos func(stream *File, ptr *FilePosition) error `std:"int fsetpos(&void,&void)"`

	ClearError func(stream *File)      `std:"clearerr(&void)"`
	IsEOF      func(stream *File) bool `std:"int feof(&void)"`
	Error      func(stream *File) bool `std:"int ferror(&void)"`
}

// LibraryIO provides stdin/stdout functions from <stdio.h>.
type LibraryIO struct {
	location

	Printf    func(format string, args ...any) (int, error) `std:"int<0 printf(&char,&vfmt%v...)"`
	Scanf     func(format string, args ...any) (int, error) `std:"int<0 scanf(&char,&vfmt%v...)"`
	GetChar   func() rune                                   `std:"int getchar"`
	GetString func(s unsafe.Pointer) string                 `std:"char gets"`
	PutChar   func(c rune) error                            `std:"int<0 putchar(int)"`
	PutString func(s string) error                          `std:"int<0 puts(&char)"`

	Error func(s string) `sym:"perror(&char)"`
}

// LibraryStrings provides string-related functions from <string.h>, <stdio.h> and <stdlib.h>.
type LibraryStrings struct {
	location

	Printf func(s unsafe.Pointer, fmt string, args ...any) (int, error) `std:"int<0 sprintf(&char,&char,&vfmt%v...)"`
	Scanf  func(s, fmt string, args ...any) (int, error)                `std:"int<0 sscanf(&char,&char,&vfmt%v...)"`

	ToFloat64    func(s string) float64                 `std:"double atof(&char)"`
	ToInt32      func(s string) int32                   `std:"int atoi(&char)"`
	ToInt64      func(s string) int64                   `std:"long atol(&char)"`
	ParseFloat64 func(s string) (float64, int)          `std:"double atof(&char,+ptrdiff%v)"`
	ParseInt64   func(s string, base int) (int64, int)  `std:"long strtol(&char,+ptrdiff(%v),int)"`
	ParseUint64  func(s string, base int) (uint64, int) `std:"unsigned_long strtoul(&char,+ptrdiff(%v),int)"`

	Copy           func([]byte, string) string    `std:"char strcpy(&char|%[1]v,&char)"`
	CopyLimited    func([]byte, string) string    `std:"char strncpy(&char|%[1]v,&char)"`
	Cat            func([]byte, string) string    `std:"char strcat(&char||%[1]v,&char)"`
	CatLimited     func([]byte, string) string    `std:"char strncat(&char||%[1]v,&char,size_t%[2]v)"`
	Compare        func(cs, ct string) int        `std:"int strcmp(&char,&char)"`
	CompareLimited func(cs, ct string, n int) int `std:"int strncmp(&char,&char,int)"`

	Index             func(cs string, c rune) int `std:"ptrdiff%v strchr(&char,int)"`
	IndexLast         func(cs string, c rune) int `std:"ptrdiff%v strrchr(&char,int)"`
	Span              func(cs, ct string) int     `std:"size_t strspn(&char,&char)"`
	ComplimentarySpan func(cs, ct string) int     `std:"size_t strcspn(&char,&char)"`
	PointerBreak      func(cs, ct string) int     `std:"ptrdiff%v strpbrk(&char,&char)"`

	Search func(cs, ct string) int `std:"ptrdiff%v strstr(&char,&char)"`
	Length func(cs string) int     `std:"size_t strlen(&char)"`

	Error  func(n error) string                `std:"$char strerror(int)"`
	Tokens func(s []byte, delim string) string `std:"char strtok(&char,&char)"`
}

// LibraryMemory provides memory-related functions from <stdlib.h>.
type LibraryMemory struct {
	location

	AllocateZeros func(int) []byte         `std:"$void|%v calloc(size_t=1,size_t)"`
	Allocate      func(int) []byte         `std:"$void|%v malloc(size_t)"`
	Reallocate    func([]byte, int) []byte `std:"$void|%[2]v realloc(#void,size_t)"`
	Free          func([]byte)             `std:"free(#void)"`

	Sort   func(base unsafe.Pointer, n, size Size, cmp func(a, b unsafe.Pointer) int)                              `std:"qsort(&void,size_t,size_t,&func(&void,&void)int)"`
	Search func(key, base unsafe.Pointer, n, size Size, cmp func(keyval, datum unsafe.Pointer) int) unsafe.Pointer `std:"void bsearch(&void,&void,size_t,size_t,&func(&void,&void)int)"`

	Copy    func(dst, src []byte) []byte `std:"void|%v memcpy(&void|%[2]v,&void,size_t%[2]v)"`
	Move    func(dst, src []byte) []byte `std:"void|%v memmove(&void|%[2]v,&void,size_t%[2]v)"`
	Compare func(cs, ct []byte) int      `std:"int memcmp(&void|%[2]v,&void,size_t%[2]v)"`

	Index func([]byte, byte) int `std:"ptrdiff%v memchr(&void,int,size_t%[2]v)"`

	Set func([]byte, byte) []byte `std:"void|%v memset(&void,int,size_t%[2]v)"`
}

// LibraryProgram provides program-related functions from <stdlib.h>.
type LibraryProgram struct {
	location

	Abort  func()                   `std:"abort"`
	Exit   func(status ExitStatus)  `std:"exit(int)"`
	OnExit func(func())             `std:"atexit($func)" sym:"atexit,__cxa_atexit"`
	Getenv func(name string) string `std:"char getenv(&char)"`
}

// LibrarySystem provides system-related functions from <stdlib.h>.
type LibrarySystem struct {
	location

	Command func(command string) Int `sym:"system" std:"int func(&char)"`

	Clock func() Clock       `sym:"clock" std:"clock_t func"`
	Time  func(t *Time) Time `sym:"time" std:"time_t func(&time_t)"`
}

// LibraryDivision provides division-related functions from <stdlib.h>.
type LibraryDivision struct {
	location

	Int32 func(num, denom int32) DivisionInt  `sym:"div" std:"div_t func(int,int)"`
	Int64 func(num, denom int64) DivisionLong `sym:"ldiv" std:"ldiv_t func(long,long)"`
}

// LibraryTime provides time-related functions from <time.h>.
type LibraryTime struct {
	location

	Sub    func(t1, t2 Time) Time `sym:"difftime" std:"time_t func(time_t,time_t)"`
	String func(t Time) string    `sym:"ctime" std:"$char func(time_t)"`

	UTC   func(t Time) *Date `sym:"timegm" std:"$void func(time_t)"`
	Local func(t Time) *Date `sym:"localtime" std:"$void func(time_t)"`
}

// LibraryDates provides date-related functions from <time.h>.
type LibraryDates struct {
	location

	Time   func(t *Date) Time                           `sym:"mktime" std:"time_t func(&void)"`
	String func(t *Date) string                         `sym:"asctime" std:"$char func(&void)"`
	Format func(s []byte, format string, tp *Date) Size `sym:"strftime" std:"size_t func(&char,size_t%v,&char,&void)"`
}
