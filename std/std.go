// Package std provides standard C types for cross-language interoperability.
// The types in this package reflect the default ABI for the current platform.
package std

import (
	"errors"
	"strconv"
	"unsafe"
)

// Boolean constants.
const (
	True  = c_true
	False = c_false
)

// Limits.
const (
	BitsInChar          = c_CHAR_BIT
	MaxMultiByte        = c_MB_LEN_MAX
	MinChar             = c_CHAR_MIN
	MaxChar             = c_CHAR_MAX
	MinSignedChar       = c_SCHAR_MIN
	MinShort            = c_SHRT_MIN
	MinInt              = c_INT_MIN
	MinLong             = c_LONG_MIN
	MinLongLong         = c_LLONG_MIN
	MaxSignedChar       = c_SCHAR_MAX
	MaxShort            = c_SHRT_MAX
	MaxInt              = c_INT_MAX
	MaxLong             = c_LONG_MAX
	MaxLongLong         = c_LLONG_MAX
	MaxUnsignedChar     = c_UCHAR_MAX
	MaxUnsignedShort    = c_USHRT_MAX
	MaxUnsignedInt      = c_UINT_MAX
	MaxUnsignedLong     = c_ULONG_MAX
	MaxUnsignedLongLong = c_ULLONG_MAX
	MaxPtrdiff          = c_PTRDIFF_MAX
	MaxSize             = c_SIZE_MAX
	MinInt8             = c_INT8_MIN
	MinInt16            = c_INT16_MIN
	MinInt32            = c_INT32_MIN
	MinInt64            = c_INT64_MIN
	MaxInt8             = c_INT8_MAX
	MaxInt16            = c_INT16_MAX
	MaxInt32            = c_INT32_MAX
	MaxInt64            = c_INT64_MAX
	MaxUint8            = c_UINT8_MAX
	MaxUint16           = c_UINT16_MAX
	MaxUint32           = c_UINT32_MAX
	MaxUint64           = c_UINT64_MAX
	MinIntFast8         = c_INT_FAST8_MIN
	MinIntFast16        = c_INT_FAST16_MIN
	MinIntFast32        = c_INT_FAST32_MIN
	MinIntFast64        = c_INT_FAST64_MIN
	MaxIntFast8         = c_INT_FAST8_MAX
	MaxIntFast16        = c_INT_FAST16_MAX
	MaxIntFast32        = c_INT_FAST32_MAX
	MaxIntFast64        = c_INT_FAST64_MAX
	MaxUintFast8        = c_UINT_FAST8_MAX
	MaxUintFast16       = c_UINT_FAST16_MAX
	MaxUintFast32       = c_UINT_FAST32_MAX
	MaxUintFast64       = c_UINT_FAST64_MAX
	MinIntLeast8        = c_INT_LEAST8_MIN
	MinIntLeast16       = c_INT_LEAST16_MIN
	MinIntLeast32       = c_INT_LEAST32_MIN
	MinIntLeast64       = c_INT_LEAST64_MIN
	MaxIntLeast8        = c_INT_LEAST8_MAX
	MaxIntLeast16       = c_INT_LEAST16_MAX
	MaxIntLeast32       = c_INT_LEAST32_MAX
	MaxIntLeast64       = c_INT_LEAST64_MAX
	MaxUintLeast8       = c_UINT_LEAST8_MAX
	MaxUintLeast16      = c_UINT_LEAST16_MAX
	MaxUintLeast32      = c_UINT_LEAST32_MAX
	MaxUintLeast64      = c_UINT_LEAST64_MAX
	MinWideChar         = c_WCHAR_MIN
	MaxWideChar         = c_WCHAR_MAX
	MinWideInt          = c_WINT_MIN
	MaxWideInt          = c_WINT_MAX
	MinIntptr           = c_INTPTR_MIN
	MaxIntptr           = c_INTPTR_MAX
	MaxUintptr          = c_UINTPTR_MAX
	MinIntmax           = c_INTMAX_MIN
	MaxIntmax           = c_INTMAX_MAX
	MaxUintmax          = c_UINTMAX_MAX
	MinSignedAtomic     = c_SIG_ATOMIC_MIN
	MaxSignedAtomic     = c_SIG_ATOMIC_MAX
)

// Complex constants.
const (
	I = 1i
)

// Floating point constants.
const (
	FloatingPointRadix       = c_FLT_RADIX
	MaxFloatingPointDigits   = c_DECIMAL_DIG
	MaxFloatDigits           = c_FLT_DECIMAL_DIG
	MaxDoubleDigits          = c_DBL_DECIMAL_DIG
	MaxLongDoubleDigits      = c_LDBL_DECIMAL_DIG
	MinFloat                 = c_FLT_MIN
	MinDouble                = c_DBL_MIN
	MinLongDouble            = c_LDBL_MIN
	MaxFloat                 = c_FLT_MAX
	MaxDouble                = c_DBL_MAX
	MaxLongDouble            = c_LDBL_MAX
	FloatEpsilon             = c_FLT_EPSILON
	DoubleEpsilon            = c_DBL_EPSILON
	LongDoubleEpsilon        = c_LDBL_EPSILON
	FloatDigits              = c_FLT_DIG
	DoubleDigits             = c_DBL_DIG
	LongDoubleDigits         = c_LDBL_DIG
	FloatMantissaDigits      = c_FLT_MANT_DIG
	DoubleMantissaDigits     = c_DBL_MANT_DIG
	LongDoubleMantissaDigits = c_LDBL_MANT_DIG
	MinFloatExp              = c_FLT_MIN_EXP
	MinDoubleExp             = c_DBL_MIN_EXP
	MinLongDoubleExp         = c_LDBL_MIN_EXP
	MinFloatExp10            = c_FLT_MIN_10_EXP
	MinDoubleExp10           = c_DBL_MIN_10_EXP
	MinLongDoubleExp10       = c_LDBL_MIN_10_EXP

	MaxFloatExp        = c_FLT_MAX_EXP
	MaxDoubleExp       = c_DBL_MAX_EXP
	MaxLongDoubleExp   = c_LDBL_MAX_EXP
	MaxFloatExp10      = c_FLT_MAX_10_EXP
	MaxDoubleExp10     = c_DBL_MAX_10_EXP
	MaxLongDoubleExp10 = c_LDBL_MAX_10_EXP

	FloatingPointRoundingMode     = c_FLT_ROUNDS
	FloatingPointEvaluationMethod = c_FLT_EVAL_METHOD

	FloatHasSubnormal      = c_FLT_HAS_SUBNORM
	DoubleHasSubnormal     = c_DBL_HAS_SUBNORM
	LongDoubleHasSubnormal = c_LDBL_HAS_SUBNORM

	FloatingPointErrorHandling      = c_math_errhandling // equals FloatingPointWillError and/or FloatingPointWillRaiseException
	FloatingPointWillError          = c_MATH_ERRNO
	FloatingPointWillRaiseException = c_MATH_ERREXCEPT
)

// FloatingPointException bitmask.
type FloatingPointException c_int

// Floating point exceptions.
const (
	FloatingPointExceptionsDefault FloatingPointException = c_FE_DFL_ENV
	FloatingPointDivisionByZero    FloatingPointException = c_FE_DIVBYZERO
	FloatingPointInexact           FloatingPointException = c_FE_INEXACT
	FloatingPointInvalid           FloatingPointException = c_FE_INVALID
	FloatingPointOverflow          FloatingPointException = c_FE_OVERFLOW
	FloatingPointUnderflow         FloatingPointException = c_FE_UNDERFLOW
	FloatingPointExceptionsAll     FloatingPointException = c_FE_ALL_EXCEPT
)

// RoundingMode used for floating point operations.
type RoundingMode c_int

const (
	FloatRoundDefault    RoundingMode = c_fegetround
	FloatRoundDownward   RoundingMode = c_FE_DOWNWARD
	FloatRoundToNearest  RoundingMode = c_FE_TONEAREST
	FloatRoundTowardZero RoundingMode = c_FE_TOWARDZERO
	FloatRoundUpward     RoundingMode = c_FE_UPWARD
)

// FloatingPointClassification of a floating point number.
type FloatingPointClassification c_int

const (
	FloatingPointIsNormal    FloatingPointClassification = c_FP_NORMAL
	FloatingPointIsSubnormal FloatingPointClassification = c_FP_SUBNORMAL
	FloatingPointIsZero      FloatingPointClassification = c_FP_ZERO
	FloatingPointIsInfinite  FloatingPointClassification = c_FP_INFINITE
	FloatingPointIsNaN       FloatingPointClassification = c_FP_NAN
)

// Signal types.
type Signal c_int

const (
	SignalTermination         Signal = c_SIGTERM
	SignalInvalidMemoryAccess Signal = c_SIGSEGV
	SignalInterrupt           Signal = c_SIGINT
	SignalInvalidInstruction  Signal = c_SIGILL
	SignalAbnormalTermination Signal = c_SIGABRT
	SignalFloatingPointError  Signal = c_SIGFPE
)

// File constants.
const (
	EOF            = c_EOF
	MaxFiles       = c_FOPEN_MAX
	MaxFileNameLen = c_FILENAME_MAX
	BufferSize     = c_BUFSIZ
	MaxTempFiles   = c_TMP_MAX
	TempNameLength = c_L_tmpnam
)

const ClocksPerSecond Clock = c_CLOCKS_PER_SEC

// BufferMode for buffering a file.
type BufferMode c_int

const (
	BufferFully    BufferMode = c__IOFBF
	BufferLine     BufferMode = c__IOLBF
	BufferDisabled BufferMode = c__IONBF
)

// SeekMode for seeking in a file.
type SeekMode c_int

const (
	SeekStart   SeekMode = c_SEEK_SET
	SeekCurrent SeekMode = c_SEEK_CUR
	SeekEnd     SeekMode = c_SEEK_END
)

// Fixed width numeric types.
type (
	Int8   c_int8_t
	Int16  c_int16_t
	Int32  c_int32_t
	Int64  c_int64_t
	Uint8  c_uint8_t
	Uint16 c_uint16_t
	Uint32 c_uint32_t
	Uint64 c_uint64_t
)

// Variable width numeric types.
type (
	Char     c_char
	Short    c_short
	Int      c_int
	Long     c_long
	LongLong c_longlong

	Char16     c_char16_t
	Char32     c_char32_t
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

	Uintptr c_uintptr_t
)

// Less common types.
type (
	Intptr   c_intptr_t
	Ptrdiff  c_ptrdiff_t
	MaxAlign c_max_align_t

	Longest         c_intmax_t
	UnsignedLongest c_uintmax_t

	IntAtLeast8  c_int_least8_t
	IntAtLeast16 c_int_least16_t
	IntAtLeast32 c_int_least32_t
	IntAtLeast64 c_int_least64_t

	UintAtLeast8  c_uint_least8_t
	UintAtLeast16 c_uint_least16_t
	UintAtLeast32 c_uint_least32_t
	UintAtLeast64 c_uint_least64_t

	FastInt8  c_int_fast8_t
	FastInt16 c_int_fast16_t
	FastInt32 c_int_fast32_t
	FastInt64 c_int_fast64_t

	FastUint8  c_uint_fast8_t
	FastUint16 c_uint_fast16_t
	FastUint32 c_uint_fast32_t
	FastUint64 c_uint_fast64_t

	FastFloat  c_float_t
	FastDouble c_double_t

	WideChar c_wchar_t
	WideInt  c_wint_t

	SignedAtomic c_sig_atomic_t
)

// Error represents a C int error, where 0 is success
// and any other value is an error.
type Error c_int

// Errors types.
const (
	ErrDomain              Error = c_EDOM
	ErrIllegalByteSequence Error = c_EILSEQ
	ErrResultTooLarge      Error = c_ERANGE
)

// Err returns nil if err is 0, otherwise it returns an error
func (err Error) Err() error {
	if err == 0 {
		return nil
	}
	return errors.New(strconv.Itoa(int(err)))
}

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
	handle
}

// Enum represents a type derived from a C enum.
type Enum c_int

// UnsafePointer to C memory, cannot contain Go
// pointers and cannot be dereferenced.
type UnsafePointer struct {
	ptr c_uintptr_t
}

// Structures.
type (
	Locale   = c_lconv
	Date     = c_tm
	NanoTime = c_timespec
)

// LocaleCategory for setting the locale.
type LocaleCategory c_int

const (
	LocaleAll      LocaleCategory = c_LC_ALL
	LocaleCollate  LocaleCategory = c_LC_COLLATE
	LocaleCharType LocaleCategory = c_LC_CTYPE
	LocaleMonetary LocaleCategory = c_LC_MONETARY
	LocaleNumeric  LocaleCategory = c_LC_NUMERIC
	LocaleTime     LocaleCategory = c_LC_TIME
)

// Integer division types that hold both the quotient and remainder.
type (
	DivisionInt      c_div_t
	DivisionLong     c_ldiv_t
	DivisionLongLong c_lldiv_t
	DivisionIntMax   c_imaxdiv_t
)

// ExitStatus values
type ExitStatus c_int

const (
	ExitSuccess ExitStatus = c_EXIT_SUCCESS
	ExitFailure ExitStatus = c_EXIT_FAILURE
)

// Handles.
type (
	File              c_FILE      // File stream.
	FilePosition      c_fpos_t    // FilePosition is a file position.
	JumpBuffer        c_jmp_buf   // JumpBuffer used for non-local jumps.
	ArgumentList      c_va_list   // ArgumentList is a list of arguments.
	MultiByteIterator c_mbstate_t // state for multi-byte characters.
)

// Atomic constants.
const (
	AtomicBoolLockFree     = c_ATOMIC_BOOL_LOCK_FREE
	AtomicCharLockFree     = c_ATOMIC_CHAR_LOCK_FREE
	AtomicChar16LockFree   = c_ATOMIC_CHAR16_T_LOCK_FREE
	AtomicChar32LockFree   = c_ATOMIC_CHAR32_T_LOCK_FREE
	AtomicWideCharLockFree = c_ATOMIC_WCHAR_T_LOCK_FREE
	AtomicShortLockFree    = c_ATOMIC_SHORT_LOCK_FREE
	AtomicIntLockFree      = c_ATOMIC_INT_LOCK_FREE
	AtomicLongLockFree     = c_ATOMIC_LONG_LOCK_FREE
	AtomicLongLongLockFree = c_ATOMIC_LLONG_LOCK_FREE
	AtomicPointerLockFree  = c_ATOMIC_POINTER_LOCK_FREE
)

type atomic interface {
	Bool | Char | SignedChar | UnsignedChar | Short | UnsignedShort | Int | UnsignedInt | Long | UnsignedLong | LongLong | UnsignedLongLong |
		Char16 | Char32 | WideChar | IntAtLeast8 | IntAtLeast16 | IntAtLeast32 | IntAtLeast64 | UintAtLeast8 | UintAtLeast16 | UintAtLeast32 | UintAtLeast64 |
		FastInt8 | FastInt16 | FastInt32 | FastInt64 | FastUint8 | FastUint16 | FastUint32 | FastUint64 | Intptr | Uintptr | Size | Ptrdiff | Longest | UnsignedLongest
}

// Atomic signals the type as an atomic type.
type Atomic[T atomic] struct {
	val T
}

type IsPointer interface {
	Pointer() uintptr
}

type handle uintptr

func (p handle) Pointer() uintptr {
	return uintptr(p)
}

func (p *handle) SetPointer(val unsafe.Pointer) {
	*p = handle(val)
}

func (s String) UnsafePointer() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

// Location can be added to a Library struct to specify
// the standard location or name of that library on a
// specific GOOS.
//
// For example:
//
//	type Library struct {
//		linux   Location `libc.so.6 libm.so.6`
//		darwin  Location `libSystem.dylib`
//		windows Location `msvcrt.dll`
//	}
type Location struct{}

// Memory is like [Handle] except it can be freed.
type Memory[T any] struct {
	_ [0]*T
	*memory
}

func (m memory) Pointer() uintptr {
	return uintptr(m.ptr)
}

type memory struct {
	ptr  Uintptr
	free func()
}

// Pointer is a typed pointer to C memory that can
// be freed, dereferenced, and passed to C functions.
type Pointer[T any] struct {
	_ [0]*T
	*pointer[T]
}

type pointer[T any] struct {
	ptr  unsafe.Pointer
	free func()
}

func (p pointer[T]) Get() T {
	return *(*T)(p.ptr)
}

func (p pointer[T]) Set(val T) {
	*(*T)(p.ptr) = val
}

func (p pointer[T]) Free() {
	p.free()
}

func (p pointer[T]) UnsafePointer() unsafe.Pointer {
	return p.ptr
}

type Struct[T any] struct {
	_      [0]*T
	ptr    Uintptr   // pointer to C
	free   func()    // free function
	layout []uintptr // Go -> C layout.
}
