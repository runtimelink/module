package std

import (
	"runtime/debug"

	"runtime.link/dll"
	"runtime.link/ffi"
)

func Link() error {
	return dll.Link(
		&Char,
		&FloatingPoint,
		&Locale,
		&Program,
		&Files,
		&IO,
		&String,
		&Memory,
		&Time,
		&Complex,
		&ComplexFloat,
		&ComplexDoubleLong,
		&Int,
		&Long,
		&LongLong,
		&IntMax,
		&Double,
		&DoubleLong,
		&Float,
	)
}

type LibC struct {
	dll.Library `linux:"libc.so.6" darwin:"libSystem.dylib"`
}

type LibM struct {
	dll.Library `linux:"libm.so.6" darwin:"libSystem.dylib"`
}

// Assert aborts the program if val is zero.
func Assert[T comparable](val T) {
	var zero T
	if val == zero {
		debug.PrintStack()
		Program.Abort()
	}
}

// FloatRoundingMode returns the current rounding mode.
// -1	the default rounding direction is not known
// 0	toward zero; same meaning as FE_TOWARDZERO
// 1	to nearest; same meaning as FE_TONEAREST
// 2	towards positive infinity; same meaning as FE_UPWARD
// 3	towards negative infinity; same meaning as FE_DOWNWARD
// other values	implementation-defined behavior
func FloatRoundingMode() ffi.Int {
	switch FloatingPoint.GetRoundingMode() {
	case ffi.FloatRoundTowardZero:
		return 0
	case ffi.FloatRoundToNearest:
		return 1
	case ffi.FloatRoundUpward:
		return 2
	case ffi.FloatRoundDownward:
		return 3
	default:
		return -1
	}
}

var Char struct {
	LibC

	IsAlphaNumeric func(ffi.Int) ffi.Int `ffi:"isalnum"`
	IsAlpha        func(ffi.Int) ffi.Int `ffi:"isalpha"`
	IsUpper        func(ffi.Int) ffi.Int `ffi:"isupper"`
	IsLower        func(ffi.Int) ffi.Int `ffi:"islower"`
	IsDigit        func(ffi.Int) ffi.Int `ffi:"isdigit"`
	IsHexDigit     func(ffi.Int) ffi.Int `ffi:"isxdigit"`
	IsControl      func(ffi.Int) ffi.Int `ffi:"iscntrl"`
	IsGraph        func(ffi.Int) ffi.Int `ffi:"isgraph"`
	IsSpace        func(ffi.Int) ffi.Int `ffi:"isspace"`
	IsBlank        func(ffi.Int) ffi.Int `ffi:"isblank"`
	IsPrint        func(ffi.Int) ffi.Int `ffi:"isprint"`
	IsPuncuation   func(ffi.Int) ffi.Int `ffi:"ispunct"`

	ToLower func(ffi.Int) ffi.Int `ffi:"tolower"`
	ToUpper func(ffi.Int) ffi.Int `ffi:"toupper"`
}

var FloatingPoint struct {
	LibM

	ClearExceptions   func(ffi.FloatException) ffi.Error                                `ffi:"feclearexcept"`
	Exceptions        func(ffi.FloatException) ffi.FloatException                       `ffi:"fetestexcept"`
	RaiseExceptions   func(ffi.FloatException) ffi.Error                                `ffi:"feraiseexcept"`
	GetExceptionFlag  func(*ffi.FloatingPointEnvironment, ffi.FloatException) ffi.Error `ffi:"fegetexceptflag"`
	SetExceptionFlag  func(*ffi.FloatingPointEnvironment, ffi.FloatException) ffi.Error `ffi:"fesetexceptflag"`
	SetRoundingMode   func(ffi.FloatRoundingMode) ffi.Error                             `ffi:"fesetround"`
	GetRoundingMode   func() ffi.FloatRoundingMode                                      `ffi:"fegetround"`
	GetEnvironment    func(*ffi.FloatingPointEnvironment) ffi.Error                     `ffi:"fegetenv"`
	SetEnvironment    func(*ffi.FloatingPointEnvironment) ffi.Error                     `ffi:"fesetenv"`
	UpdateEnvironment func(*ffi.FloatingPointEnvironment) ffi.Error                     `ffi:"feupdateenv"`
	HoldExceptions    func(*ffi.FloatingPointEnvironment) ffi.Error                     `ffi:"feholdexcept"`
}

var Locale struct {
	LibC

	Set func(ffi.LocaleCategory, *ffi.Locale) ffi.String `ffi:"setlocale"`
	Get func() *ffi.Locale                               `ffi:"localeconv"`
}

var Program struct {
	LibC

	Abort              func()                             `ffi:"abort"`
	Exit               func(ffi.Int)                      `ffi:"exit"`
	ExitFast           func(ffi.Int)                      `ffi:"quick_exit"`
	ExitWithoutCleanup func(ffi.Int)                      `ffi:"_Exit"`
	OnExit             func(func())                       `ffi:"atexit,__cxa_atexit"`
	OnExitFast         func(func())                       `ffi:"at_quick_exit,__cxa_at_quick_exit"`
	LongJump           func(ffi.JumpBuffer, ffi.Int)      `ffi:"longjmp"`
	OnSignal           func(ffi.Signal, func(ffi.Signal)) `ffi:"signal"`
	Raise              func(ffi.Signal)                   `ffi:"raise"`
	Getenv             func(ffi.String) ffi.String        `ffi:"getenv"`
	Exec               func(ffi.String) ffi.Error         `ffi:"system"`
}

var Files struct {
	LibC

	Open          func(ffi.String, ffi.String) *ffi.File                               `ffi:"fopen"`
	Reopen        func(ffi.String, ffi.String, *ffi.File) *ffi.File                    `ffi:"freopen"`
	Flush         func(*ffi.File) ffi.Int                                              `ffi:"fflush"`
	SetBuffer     func(*ffi.File, ffi.UnsafePointer) ffi.Int                           `ffi:"setbuf"`
	SetBufferMode func(*ffi.File, ffi.UnsafePointer, ffi.BufferMode, ffi.Size) ffi.Int `ffi:"setvbuf"`
	SetCharWide   func(*ffi.File, ffi.Int) ffi.Int                                     `ffi:"fwide"`

	Read  func(ffi.Pointer[ffi.Char], ffi.Size, ffi.Size, *ffi.File) ffi.Int `ffi:"fread"`
	Write func(ffi.Pointer[ffi.Char], ffi.Size, ffi.Size, *ffi.File) ffi.Int `ffi:"fwrite"`

	GetChar   func(*ffi.File) ffi.Int                                               `ffi:"fgetc"`
	GetString func(ffi.Pointer[ffi.Char], ffi.Int, *ffi.File) ffi.Pointer[ffi.Char] `ffi:"fgets"`
	PutChar   func(ffi.Int, *ffi.File) ffi.Int                                      `ffi:"fputc"`
	PutString func(ffi.String, *ffi.File) ffi.Int                                   `ffi:"fputs"`
	UngetChar func(ffi.Int, *ffi.File) ffi.Int                                      `ffi:"ungetc"`

	GetCharWide   func(*ffi.File) ffi.CharWide                            `ffi:"fgetwc"`
	GetStringWide func(ffi.StringWide, ffi.Int, *ffi.File) ffi.StringWide `ffi:"fgetws"`
	PutCharWide   func(ffi.CharWide, *ffi.File) ffi.CharWide              `ffi:"fputwc"`
	PutStringWide func(ffi.StringWide, *ffi.File) ffi.Int                 `ffi:"fputws"`
	UngetCharWide func(ffi.CharWide, *ffi.File) ffi.CharWide              `ffi:"ungetwc"`

	Scanf      func(*ffi.File, ffi.String, ...ffi.UnsafePointer) ffi.Int     `ffi:"fscanf"`
	Printf     func(*ffi.File, ffi.String, ...ffi.UnsafePointer) ffi.Int     `ffi:"fprintf"`
	ScanWidef  func(*ffi.File, ffi.StringWide, ...ffi.UnsafePointer) ffi.Int `ffi:"fwscanf"`
	PrintWidef func(*ffi.File, ffi.StringWide, ...ffi.UnsafePointer) ffi.Int `ffi:"fwprintf"`

	Tell   func(*ffi.File) ffi.Long                        `ffi:"ftell"`
	GetPos func(*ffi.File, *ffi.FilePosition) ffi.Int      `ffi:"fgetpos"`
	Seek   func(*ffi.File, ffi.Long, ffi.SeekMode) ffi.Int `ffi:"fseek"`
	SetPos func(*ffi.File, *ffi.FilePosition) ffi.Int      `ffi:"fsetpos"`
	Rewind func(*ffi.File)                                 `ffi:"rewind"`

	ClearErr func(*ffi.File)         `ffi:"clearerr"`
	IsEOF    func(*ffi.File) ffi.Int `ffi:"feof"`
	IsErr    func(*ffi.File) ffi.Int `ffi:"ferror"`
	Error    func(*ffi.String)       `ffi:"perror"`

	Remove   func(ffi.String) ffi.Int             `ffi:"remove"`
	Rename   func(ffi.String, ffi.String) ffi.Int `ffi:"rename"`
	Temp     func() *ffi.File                     `ffi:"tmpfile"`
	TempName func(ffi.String) ffi.String          `ffi:"tmpnam"`
}

var IO struct {
	LibC

	GetChar   func() ffi.Int                                    `ffi:"getchar"`
	GetString func(ffi.Pointer[ffi.Char]) ffi.Pointer[ffi.Char] `ffi:"gets"`
	PutChar   func(ffi.Int) ffi.Int                             `ffi:"putchar"`
	PutString func(ffi.String) ffi.Int                          `ffi:"puts"`

	GetCharWide func() ffi.CharWide             `ffi:"getwchar"`
	PutCharWide func(ffi.CharWide) ffi.CharWide `ffi:"putwchar"`

	Scanf      func(ffi.String, ...ffi.UnsafePointer) ffi.Int     `ffi:"scanf"`
	Printf     func(ffi.String, ...ffi.UnsafePointer) ffi.Int     `ffi:"printf"`
	ScanWidef  func(ffi.StringWide, ...ffi.UnsafePointer) ffi.Int `ffi:"wscanf"`
	PrintWidef func(ffi.StringWide, ...ffi.UnsafePointer) ffi.Int `ffi:"wprintf"`
}

var String struct {
	LibC

	Error func(ffi.Error) ffi.String `ffi:"strerror"`

	Scanf      func(ffi.String, ffi.String, ...ffi.UnsafePointer) ffi.Int         `ffi:"sscanf"`
	Printf     func(ffi.String, ffi.String, ...ffi.UnsafePointer) ffi.Int         `ffi:"sprintf"`
	ScanWidef  func(ffi.StringWide, ffi.StringWide, ...ffi.UnsafePointer) ffi.Int `ffi:"swscanf"`
	PrintWidef func(ffi.StringWide, ffi.StringWide, ...ffi.UnsafePointer) ffi.Int `ffi:"swprintf"`

	ToFloat               func(ffi.String) ffi.Float                                `ffi:"atof"`
	ToInt                 func(ffi.String) ffi.Int                                  `ffi:"atoi"`
	ToLong                func(ffi.String) ffi.Long                                 `ffi:"atol"`
	ToLongLong            func(ffi.String) ffi.LongLong                             `ffi:"atoll"`
	ParseLong             func(ffi.String, *ffi.Char, ffi.Int) ffi.Long             `ffi:"strtol"`
	ParseLongLong         func(ffi.String, *ffi.Char, ffi.Int) ffi.LongLong         `ffi:"strtoll"`
	ParseUnsignedLong     func(ffi.String, *ffi.Char, ffi.Int) ffi.LongUnsigned     `ffi:"strtoul"`
	ParseUnsignedLongLong func(ffi.String, *ffi.Char, ffi.Int) ffi.LongLongUnsigned `ffi:"strtoull"`
	ParseFloat            func(ffi.String, *ffi.Char) ffi.Float                     `ffi:"strtof"`
	ParseDouble           func(ffi.String, *ffi.Char) ffi.Double                    `ffi:"strtod"`
	ParseDoubleLong       func(ffi.String, *ffi.Char) ffi.DoubleLong                `ffi:"strtold"`
	ParseIntmax           func(ffi.String, *ffi.Char, ffi.Int) ffi.IntMax           `ffi:"strtoimax"`
	ParseUintmax          func(ffi.String, *ffi.Char, ffi.Int) ffi.UIntMax          `ffi:"strtoumax"`

	Copy           func(ffi.String, ffi.String) ffi.Error          `ffi:"strcpy"`
	CopyRange      func(ffi.String, ffi.String) ffi.Error          `ffi:"strncpy"`
	Append         func(ffi.String, ffi.String) ffi.Error          `ffi:"strcat"`
	AppendRange    func(ffi.String, ffi.String) ffi.Error          `ffi:"strncat"`
	Localize       func(ffi.String, ffi.String, ffi.Size) ffi.Size `ffi:"strxfrm"`
	Duplicate      func(ffi.String) ffi.String                     `ffi:"strdup"`
	DuplicateRange func(ffi.String, ffi.Size) ffi.String           `ffi:"strndup"`

	Length          func(ffi.String) ffi.Size              `ffi:"strlen"`
	Compare         func(ffi.String, ffi.String) ffi.Int   `ffi:"strcmp"`
	CompareInLocale func(ffi.String, ffi.String) ffi.Int   `ffi:"strcoll"`
	FindFirst       func(ffi.String, ffi.Int) *ffi.Char    `ffi:"strchr"`
	FindLast        func(ffi.String, ffi.Int) *ffi.Char    `ffi:"strrchr"`
	MatchLength     func(ffi.String, ffi.String) ffi.Size  `ffi:"strspn"`
	Match           func(ffi.String, ffi.String) ffi.Size  `ffi:"strcspn"`
	MatchFirst      func(ffi.String, ffi.String) *ffi.Char `ffi:"strpbrk"`
	Contains        func(ffi.String, ffi.String) *ffi.Char `ffi:"strstr"`
	ScanToken       func(ffi.String, ffi.String) *ffi.Char `ffi:"strtok"`
}

var Memory struct {
	LibC

	Calloc  func(ffi.Size, ffi.Size) ffi.UnsafePointer          `ffi:"calloc"`
	Free    func(ffi.UnsafePointer)                             `ffi:"free"`
	Malloc  func(ffi.Size) ffi.UnsafePointer                    `ffi:"malloc"`
	Realloc func(ffi.UnsafePointer, ffi.Size) ffi.UnsafePointer `ffi:"realloc"`

	BinarySearch func(ffi.UnsafePointer, ffi.UnsafePointer, ffi.Size, ffi.Size, func(ffi.UnsafePointer, ffi.UnsafePointer) ffi.Int) ffi.UnsafePointer `ffi:"bsearch"`

	Sort func(ffi.UnsafePointer, ffi.Size, ffi.Size, func(ffi.UnsafePointer, ffi.UnsafePointer) ffi.Int) ffi.UnsafePointer `ffi:"qsort"`

	Compare func(ffi.UnsafePointer, ffi.UnsafePointer, ffi.Size) ffi.Int                     `ffi:"memcmp"`
	Copy    func(ffi.UnsafePointer, ffi.Size, ffi.UnsafePointer, ffi.Size) ffi.UnsafePointer `ffi:"memcpy"`
	Move    func(ffi.UnsafePointer, ffi.UnsafePointer, ffi.Size) ffi.UnsafePointer           `ffi:"memmove"`
	Set     func(ffi.UnsafePointer, ffi.Int, ffi.Size) ffi.UnsafePointer                     `ffi:"memset"`
	Find    func(ffi.UnsafePointer, ffi.Int, ffi.Size) ffi.UnsafePointer                     `ffi:"memchr"`
}

var Time struct {
	LibC

	Diff          func(ffi.Time, ffi.Time) ffi.Double       `ffi:"difftime"`
	Now           func(*ffi.Time) ffi.Time                  `ffi:"time"`
	Clock         func() ffi.Clock                          `ffi:"clock"`
	Nanos         func(*ffi.NanoTime, ffi.TimeType)         `ffi:"timespec_get"`
	GetResolution func(*ffi.NanoTime, ffi.TimeType) ffi.Int `ffi:"clock_getres"`

	DateString     func(ffi.String, ffi.Size, ffi.String, *ffi.Date) ffi.Size         `ffi:"strftime"`
	DateStringWide func(ffi.StringWide, ffi.Size, ffi.StringWide, *ffi.Date) ffi.Size `ffi:"wcsftime"`

	UTC   func(ffi.Time) *ffi.Date `ffi:"gmtime"`
	Local func(ffi.Time) *ffi.Date `ffi:"localtime"`
	Value func(*ffi.Date) ffi.Time `ffi:"mktime"`
}

type Div[T ffi.Int | ffi.Long | ffi.LongLong | ffi.IntMax] struct {
	Quo T
	Rem T
}
