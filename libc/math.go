package libc

import (
	"math"

	"runtime.link/ffi"
)

var Int struct {
	LibM

	Abs func(ffi.Int) ffi.Int               `ffi:"abs"`
	Div func(ffi.Int, ffi.Int) Div[ffi.Int] `ffi:"div"`

	Rand        func() ffi.Int `ffi:"rand"`
	SetRandSeed func(ffi.Int)  `ffi:"srand"`
}

var Long struct {
	LibM

	Abs func(ffi.Long) ffi.Long                `ffi:"labs"`
	Div func(ffi.Long, ffi.Long) Div[ffi.Long] `ffi:"ldiv"`

	RoundFloat func(ffi.Float) ffi.Long      `ffi:"lround"`
	Round      func(ffi.Double) ffi.Long     `ffi:"lround"`
	RoundLong  func(ffi.DoubleLong) ffi.Long `ffi:"lround"`
}

var LongLong struct {
	LibM

	Abs func(ffi.LongLong) ffi.LongLong                    `ffi:"llabs"`
	Div func(ffi.LongLong, ffi.LongLong) Div[ffi.LongLong] `ffi:"lldiv"`

	RoundFloat func(ffi.Float) ffi.Long      `ffi:"llround"`
	Round      func(ffi.Double) ffi.Long     `ffi:"llround"`
	RoundLong  func(ffi.DoubleLong) ffi.Long `ffi:"llround"`
}

var IntMax struct {
	LibM

	Abs func(ffi.IntMax) ffi.IntMax                  `ffi:"imaxabs"`
	Div func(ffi.IntMax, ffi.IntMax) Div[ffi.IntMax] `ffi:"imaxdiv"`
}

var Double struct {
	LibM

	Abs                func(ffi.Double) ffi.Double                         `ffi:"fabs"`
	Mod                func(ffi.Double, ffi.Double) ffi.Double             `ffi:"fmod"`
	Remainder          func(ffi.Double, ffi.Double) ffi.Double             `ffi:"remainder"`
	RemainderQuotient  func(ffi.Double, ffi.Double) (ffi.Double, ffi.Int)  `ffi:"remquo"`
	FusedMuliplyAdd    func(ffi.Double, ffi.Double, ffi.Double) ffi.Double `ffi:"fma"`
	Max                func(ffi.Double, ffi.Double) ffi.Double             `ffi:"fmax"`
	Min                func(ffi.Double, ffi.Double) ffi.Double             `ffi:"fmin"`
	PositiveDifference func(ffi.Double, ffi.Double) ffi.Double             `ffi:"fdim"`
	Nan                func(ffi.String) ffi.Double                         `ffi:"nan"`

	Exp   func(ffi.Double) ffi.Double `ffi:"exp"`
	Exp2  func(ffi.Double) ffi.Double `ffi:"exp2"`
	Expm1 func(ffi.Double) ffi.Double `ffi:"expm1"`
	Log   func(ffi.Double) ffi.Double `ffi:"log"`
	Log10 func(ffi.Double) ffi.Double `ffi:"log10"`
	Log2  func(ffi.Double) ffi.Double `ffi:"log2"`
	Log1p func(ffi.Double) ffi.Double `ffi:"log1p"`

	Pow   func(ffi.Double, ffi.Double) ffi.Double `ffi:"pow"`
	Sqrt  func(ffi.Double) ffi.Double             `ffi:"sqrt"`
	Cbrt  func(ffi.Double) ffi.Double             `ffi:"cbrt"`
	Hypot func(ffi.Double, ffi.Double) ffi.Double `ffi:"hypot"`

	Sin   func(ffi.Double) ffi.Double             `ffi:"sin"`
	Cos   func(ffi.Double) ffi.Double             `ffi:"cos"`
	Tan   func(ffi.Double) ffi.Double             `ffi:"tan"`
	Asin  func(ffi.Double) ffi.Double             `ffi:"asin"`
	Acos  func(ffi.Double) ffi.Double             `ffi:"acos"`
	Atan  func(ffi.Double) ffi.Double             `ffi:"atan"`
	Atan2 func(ffi.Double, ffi.Double) ffi.Double `ffi:"atan2"`

	Sinh  func(ffi.Double) ffi.Double `ffi:"sinh"`
	Cosh  func(ffi.Double) ffi.Double `ffi:"cosh"`
	Tanh  func(ffi.Double) ffi.Double `ffi:"tanh"`
	Asinh func(ffi.Double) ffi.Double `ffi:"asinh"`
	Acosh func(ffi.Double) ffi.Double `ffi:"acosh"`
	Atanh func(ffi.Double) ffi.Double `ffi:"atanh"`

	Erf    func(ffi.Double) ffi.Double `ffi:"erf"`
	Erfc   func(ffi.Double) ffi.Double `ffi:"erfc"`
	GammaT func(ffi.Double) ffi.Double `ffi:"tgamma"`
	GammaL func(ffi.Double) ffi.Double `ffi:"lgamma"`

	Ceil      func(ffi.Double) ffi.Double   `ffi:"ceil"`
	Floor     func(ffi.Double) ffi.Double   `ffi:"floor"`
	Trunc     func(ffi.Double) ffi.Double   `ffi:"trunc"`
	Round     func(ffi.Double) ffi.Double   `ffi:"round"`
	NearbyInt func(ffi.Double) ffi.Double   `ffi:"nearbyint"`
	Int       func(ffi.Double) ffi.Double   `ffi:"rint"`
	Long      func(ffi.Double) ffi.Long     `ffi:"lrint"`
	LongLong  func(ffi.Double) ffi.LongLong `ffi:"llrint"`

	Frexp      func(ffi.Double) (ffi.Double, ffi.Int)      `ffi:"frexp"`
	Ldexp      func(ffi.Double, ffi.Int) ffi.Double        `ffi:"ldexp"`
	Modf       func(ffi.Double) (ffi.Double, ffi.Double)   `ffi:"modf"`
	Scale      func(ffi.Double, ffi.Double) ffi.Double     `ffi:"scalbn"`
	ScaleLong  func(ffi.Double, ffi.Long) ffi.Double       `ffi:"scalbln"`
	LogInt     func(ffi.Double) ffi.Int                    `ffi:"logb"`
	Logb       func(ffi.Double) ffi.Double                 `ffi:"logb"`
	NextAfter  func(ffi.Double, ffi.Double) ffi.Double     `ffi:"nextafter"`
	NextToward func(ffi.Double, ffi.DoubleLong) ffi.Double `ffi:"nexttoward"`
	CopySign   func(ffi.Double, ffi.Double) ffi.Double     `ffi:"copysign"`
}

var Float struct {
	LibM

	Abs                func(ffi.Float) ffi.Float                       `ffi:"fabsf"`
	Mod                func(ffi.Float, ffi.Float) ffi.Float            `ffi:"fmodf"`
	Remainder          func(ffi.Float, ffi.Float) ffi.Float            `ffi:"remainderf"`
	RemainderQuotient  func(ffi.Float, ffi.Float) (ffi.Float, ffi.Int) `ffi:"remquof"`
	FusedMuliplyAdd    func(ffi.Float, ffi.Float, ffi.Float) ffi.Float `ffi:"fmaf"`
	Max                func(ffi.Float, ffi.Float) ffi.Float            `ffi:"fmaxf"`
	Min                func(ffi.Float, ffi.Float) ffi.Float            `ffi:"fminf"`
	PositiveDifference func(ffi.Float, ffi.Float) ffi.Float            `ffi:"fdimf"`
	Nan                func(ffi.String) ffi.Float                      `ffi:"nanf"`

	Exp   func(ffi.Float) ffi.Float `ffi:"expf"`
	Exp2  func(ffi.Float) ffi.Float `ffi:"exp2f"`
	Expm1 func(ffi.Float) ffi.Float `ffi:"expm1f"`
	Log   func(ffi.Float) ffi.Float `ffi:"logf"`
	Log10 func(ffi.Float) ffi.Float `ffi:"log10f"`
	Log2  func(ffi.Float) ffi.Float `ffi:"log2f"`
	Log1p func(ffi.Float) ffi.Float `ffi:"log1pf"`

	Pow   func(ffi.Float, ffi.Float) ffi.Float `ffi:"powf"`
	Sqrt  func(ffi.Float) ffi.Float            `ffi:"sqrtf"`
	Cbrt  func(ffi.Float) ffi.Float            `ffi:"cbrtf"`
	Hypot func(ffi.Float, ffi.Float) ffi.Float `ffi:"hypotf"`

	Sin   func(ffi.Float) ffi.Float            `ffi:"sinf"`
	Cos   func(ffi.Float) ffi.Float            `ffi:"cosf"`
	Tan   func(ffi.Float) ffi.Float            `ffi:"tanf"`
	Asin  func(ffi.Float) ffi.Float            `ffi:"asinf"`
	Acos  func(ffi.Float) ffi.Float            `ffi:"acosf"`
	Atan  func(ffi.Float) ffi.Float            `ffi:"atanf"`
	Atan2 func(ffi.Float, ffi.Float) ffi.Float `ffi:"atan2f"`

	Sinh  func(ffi.Float) ffi.Float `ffi:"sinhf"`
	Cosh  func(ffi.Float) ffi.Float `ffi:"coshf"`
	Tanh  func(ffi.Float) ffi.Float `ffi:"tanhf"`
	Asinh func(ffi.Float) ffi.Float `ffi:"asinhf"`
	Acosh func(ffi.Float) ffi.Float `ffi:"acoshf"`
	Atanh func(ffi.Float) ffi.Float `ffi:"atanhf"`

	Erf    func(ffi.Float) ffi.Float `ffi:"erff"`
	Erfc   func(ffi.Float) ffi.Float `ffi:"erfcf"`
	GammaT func(ffi.Float) ffi.Float `ffi:"tgammaf"`
	GammaL func(ffi.Float) ffi.Float `ffi:"lgammaf"`

	Ceil      func(ffi.Float) ffi.Float    `ffi:"ceilf"`
	Floor     func(ffi.Float) ffi.Float    `ffi:"floorf"`
	Trunc     func(ffi.Float) ffi.Float    `ffi:"truncf"`
	Round     func(ffi.Float) ffi.Float    `ffi:"roundf"`
	NearbyInt func(ffi.Float) ffi.Float    `ffi:"nearbyintf"`
	Int       func(ffi.Float) ffi.Float    `ffi:"rintf"`
	Long      func(ffi.Float) ffi.Long     `ffi:"lrintf"`
	LongLong  func(ffi.Float) ffi.LongLong `ffi:"llrintf"`

	Frexp      func(ffi.Float) (ffi.Float, ffi.Int)      `ffi:"frexpf"`
	Ldexp      func(ffi.Float, ffi.Int) ffi.Float        `ffi:"ldexpf"`
	Modf       func(ffi.Float) (ffi.Float, ffi.Float)    `ffi:"modff"`
	Scale      func(ffi.Float, ffi.Float) ffi.Float      `ffi:"scalbnf"`
	ScaleLong  func(ffi.Float, ffi.Long) ffi.Float       `ffi:"scalblnf"`
	LogInt     func(ffi.Float) ffi.Int                   `ffi:"logbf"`
	Logb       func(ffi.Float) ffi.Float                 `ffi:"logbf"`
	NextAfter  func(ffi.Float, ffi.Float) ffi.Float      `ffi:"nextafterf"`
	NextToward func(ffi.Float, ffi.DoubleLong) ffi.Float `ffi:"nexttowardf"`
	CopySign   func(ffi.Float, ffi.Float) ffi.Float      `ffi:"copysignf"`
}

var DoubleLong struct {
	LibM

	Abs                func(ffi.DoubleLong) ffi.DoubleLong                                 `ffi:"fabsl"`
	Mod                func(ffi.DoubleLong, ffi.DoubleLong) ffi.DoubleLong                 `ffi:"fmodl"`
	Remainder          func(ffi.DoubleLong, ffi.DoubleLong) ffi.DoubleLong                 `ffi:"remainderl"`
	RemainderQuotient  func(ffi.DoubleLong, ffi.DoubleLong) (ffi.DoubleLong, ffi.Int)      `ffi:"remquol"`
	FusedMuliplyAdd    func(ffi.DoubleLong, ffi.DoubleLong, ffi.DoubleLong) ffi.DoubleLong `ffi:"fmal"`
	Max                func(ffi.DoubleLong, ffi.DoubleLong) ffi.DoubleLong                 `ffi:"fmaxl"`
	Min                func(ffi.DoubleLong, ffi.DoubleLong) ffi.DoubleLong                 `ffi:"fminl"`
	PositiveDifference func(ffi.DoubleLong, ffi.DoubleLong) ffi.DoubleLong                 `ffi:"fdiml"`
	Nan                func(ffi.String) ffi.DoubleLong                                     `ffi:"nanl"`

	Exp   func(ffi.DoubleLong) ffi.DoubleLong `ffi:"expl"`
	Exp2  func(ffi.DoubleLong) ffi.DoubleLong `ffi:"exp2l"`
	Expm1 func(ffi.DoubleLong) ffi.DoubleLong `ffi:"expm1l"`
	Log   func(ffi.DoubleLong) ffi.DoubleLong `ffi:"logl"`
	Log10 func(ffi.DoubleLong) ffi.DoubleLong `ffi:"log10l"`
	Log2  func(ffi.DoubleLong) ffi.DoubleLong `ffi:"log2l"`
	Log1p func(ffi.DoubleLong) ffi.DoubleLong `ffi:"log1pl"`

	Pow   func(ffi.DoubleLong, ffi.DoubleLong) ffi.DoubleLong `ffi:"powl"`
	Sqrt  func(ffi.DoubleLong) ffi.DoubleLong                 `ffi:"sqrtl"`
	Cbrt  func(ffi.DoubleLong) ffi.DoubleLong                 `ffi:"cbrtl"`
	Hypot func(ffi.DoubleLong, ffi.DoubleLong) ffi.DoubleLong `ffi:"hypotl"`

	Sin   func(ffi.DoubleLong) ffi.DoubleLong                 `ffi:"sinl"`
	Cos   func(ffi.DoubleLong) ffi.DoubleLong                 `ffi:"cosl"`
	Tan   func(ffi.DoubleLong) ffi.DoubleLong                 `ffi:"tanl"`
	Asin  func(ffi.DoubleLong) ffi.DoubleLong                 `ffi:"asinl"`
	Acos  func(ffi.DoubleLong) ffi.DoubleLong                 `ffi:"acosl"`
	Atan  func(ffi.DoubleLong) ffi.DoubleLong                 `ffi:"atanl"`
	Atan2 func(ffi.DoubleLong, ffi.DoubleLong) ffi.DoubleLong `ffi:"atan2l"`

	Sinh  func(ffi.DoubleLong) ffi.DoubleLong `ffi:"sinhl"`
	Cosh  func(ffi.DoubleLong) ffi.DoubleLong `ffi:"coshl"`
	Tanh  func(ffi.DoubleLong) ffi.DoubleLong `ffi:"tanhl"`
	Asinh func(ffi.DoubleLong) ffi.DoubleLong `ffi:"asinhl"`
	Acosh func(ffi.DoubleLong) ffi.DoubleLong `ffi:"acoshl"`
	Atanh func(ffi.DoubleLong) ffi.DoubleLong `ffi:"atanhl"`

	Erf    func(ffi.DoubleLong) ffi.DoubleLong `ffi:"erfl"`
	Erfc   func(ffi.DoubleLong) ffi.DoubleLong `ffi:"erfcl"`
	GammaT func(ffi.DoubleLong) ffi.DoubleLong `ffi:"tgammal"`
	GammaL func(ffi.DoubleLong) ffi.DoubleLong `ffi:"lgammal"`

	Ceil      func(ffi.DoubleLong) ffi.DoubleLong `ffi:"ceill"`
	Floor     func(ffi.DoubleLong) ffi.DoubleLong `ffi:"floorl"`
	Trunc     func(ffi.DoubleLong) ffi.DoubleLong `ffi:"truncl"`
	Round     func(ffi.DoubleLong) ffi.DoubleLong `ffi:"roundl"`
	NearbyInt func(ffi.DoubleLong) ffi.DoubleLong `ffi:"nearbyintl"`
	Int       func(ffi.DoubleLong) ffi.DoubleLong `ffi:"rintl"`
	Long      func(ffi.DoubleLong) ffi.Long       `ffi:"lrintl"`
	LongLong  func(ffi.DoubleLong) ffi.LongLong   `ffi:"llrintl"`

	Frexp      func(ffi.DoubleLong) (ffi.DoubleLong, ffi.Int)        `ffi:"frexpl"`
	Ldexp      func(ffi.DoubleLong, ffi.Int) ffi.DoubleLong          `ffi:"ldexpl"`
	Modf       func(ffi.DoubleLong) (ffi.DoubleLong, ffi.DoubleLong) `ffi:"modfl"`
	Scale      func(ffi.DoubleLong, ffi.DoubleLong) ffi.DoubleLong   `ffi:"scalbnl"`
	ScaleLong  func(ffi.DoubleLong, ffi.Long) ffi.DoubleLong         `ffi:"scalblnl"`
	LogInt     func(ffi.DoubleLong) ffi.Int                          `ffi:"logbl"`
	Logb       func(ffi.DoubleLong) ffi.DoubleLong                   `ffi:"logbl"`
	NextAfter  func(ffi.DoubleLong, ffi.DoubleLong) ffi.DoubleLong   `ffi:"nextafterl"`
	NextToward func(ffi.DoubleLong, ffi.DoubleLong) ffi.DoubleLong   `ffi:"nexttowardl"`
	CopySign   func(ffi.DoubleLong, ffi.DoubleLong) ffi.DoubleLong   `ffi:"copysignl"`
}

func ClassifyFloat(f ffi.Float) ffi.FloatClass {
	switch {
	case f == 0:
		return ffi.FloatIsZero
	case math.IsNaN(float64(f)):
		return ffi.FloatIsNaN
	case math.IsInf(float64(f), 1), math.IsInf(float64(f), -1):
		return ffi.FloatIsInfinite
	case f == -0 || f == +0:
		return ffi.FloatIsSubnormal
	default:
		return ffi.FloatIsNormal
	}
}

func IsFinite(f ffi.Float) bool {
	return !math.IsNaN(float64(f)) && !math.IsInf(float64(f), 1) && !math.IsInf(float64(f), -1)
}

func IsInf(f ffi.Float) bool {
	return math.IsInf(float64(f), 1) || math.IsInf(float64(f), -1)
}

func IsNaN(f ffi.Float) bool {
	return math.IsNaN(float64(f))
}

func IsNormal(f ffi.Float) bool {
	return !math.IsNaN(float64(f)) && !math.IsInf(float64(f), 1) && !math.IsInf(float64(f), -1) && f != -0 && f != +0
}

func SignBit(f ffi.Float) bool {
	return math.Signbit(float64(f))
}

func IsGreater(a, b ffi.Float) bool {
	return a > b
}

func IsGreaterEqual(a, b ffi.Float) bool {
	return a >= b
}

func IsLess(a, b ffi.Float) bool {
	return a < b
}

func IsLessEqual(a, b ffi.Float) bool {
	return a <= b
}

func IsLessGreater(a, b ffi.Float) bool {
	return a != b
}

func IsUnordered(a, b ffi.Float) bool {
	return math.IsNaN(float64(a)) || math.IsNaN(float64(b))
}

func HugeFloat() ffi.Float {
	return ffi.Float(math.Inf(1))
}

func HugeDouble() ffi.Double {
	return ffi.Double(math.Inf(1))
}

func NaN() ffi.Float {
	return ffi.Float(math.NaN())
}

func Infinity() ffi.Float {
	return ffi.Float(math.Inf(1))
}
