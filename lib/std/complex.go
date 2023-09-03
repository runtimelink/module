package std

import (
	"runtime.link/ffi"
)

var Complex struct {
	LibM

	Real func(ffi.ComplexDouble) ffi.Double        `ffi:"creal"`
	Imag func(ffi.ComplexDouble) ffi.Double        `ffi:"cimag"`
	Abs  func(ffi.ComplexDouble) ffi.Double        `ffi:"cabs"`
	Arg  func(ffi.ComplexDouble) ffi.Double        `ffi:"carg"`
	Conj func(ffi.ComplexDouble) ffi.ComplexDouble `ffi:"conj"`
	Proj func(ffi.ComplexDouble) ffi.ComplexDouble `ffi:"cproj"`

	Exp func(ffi.ComplexDouble) ffi.ComplexDouble                    `ffi:"cexp"`
	Log func(ffi.ComplexDouble) ffi.ComplexDouble                    `ffi:"clog"`
	Pow func(ffi.ComplexDouble, ffi.ComplexDouble) ffi.ComplexDouble `ffi:"cpow"`

	Sqrt func(ffi.ComplexDouble) ffi.ComplexDouble `ffi:"csqrt"`
	Sin  func(ffi.ComplexDouble) ffi.ComplexDouble `ffi:"csin"`
	Cos  func(ffi.ComplexDouble) ffi.ComplexDouble `ffi:"ccos"`
	Tan  func(ffi.ComplexDouble) ffi.ComplexDouble `ffi:"ctan"`
	Asin func(ffi.ComplexDouble) ffi.ComplexDouble `ffi:"casin"`
	Acos func(ffi.ComplexDouble) ffi.ComplexDouble `ffi:"cacos"`
	Atan func(ffi.ComplexDouble) ffi.ComplexDouble `ffi:"catan"`

	Sinh  func(ffi.ComplexDouble) ffi.ComplexDouble `ffi:"csinh"`
	Cosh  func(ffi.ComplexDouble) ffi.ComplexDouble `ffi:"ccosh"`
	Tanh  func(ffi.ComplexDouble) ffi.ComplexDouble `ffi:"ctanh"`
	Asinh func(ffi.ComplexDouble) ffi.ComplexDouble `ffi:"casinh"`
	Acosh func(ffi.ComplexDouble) ffi.ComplexDouble `ffi:"cacosh"`
	Atanh func(ffi.ComplexDouble) ffi.ComplexDouble `ffi:"catanh"`
}

var ComplexFloat struct {
	LibM

	Real func(ffi.ComplexFloat) ffi.Float        `ffi:"crealf"`
	Imag func(ffi.ComplexFloat) ffi.Float        `ffi:"cimagf"`
	Abs  func(ffi.ComplexFloat) ffi.Float        `ffi:"cabsf"`
	Arg  func(ffi.ComplexFloat) ffi.Float        `ffi:"cargf"`
	Conj func(ffi.ComplexFloat) ffi.ComplexFloat `ffi:"conjf"`
	Proj func(ffi.ComplexFloat) ffi.ComplexFloat `ffi:"cprojf"`

	Exp func(ffi.ComplexFloat) ffi.ComplexFloat                   `ffi:"cexpf"`
	Log func(ffi.ComplexFloat) ffi.ComplexFloat                   `ffi:"clogf"`
	Pow func(ffi.ComplexFloat, ffi.ComplexFloat) ffi.ComplexFloat `ffi:"cpowf"`

	Sqrt func(ffi.ComplexFloat) ffi.ComplexFloat `ffi:"csqrtf"`
	Sin  func(ffi.ComplexFloat) ffi.ComplexFloat `ffi:"csinf"`
	Cos  func(ffi.ComplexFloat) ffi.ComplexFloat `ffi:"ccosf"`
	Tan  func(ffi.ComplexFloat) ffi.ComplexFloat `ffi:"ctanf"`
	Asin func(ffi.ComplexFloat) ffi.ComplexFloat `ffi:"casinf"`
	Acos func(ffi.ComplexFloat) ffi.ComplexFloat `ffi:"cacosf"`
	Atan func(ffi.ComplexFloat) ffi.ComplexFloat `ffi:"catanf"`

	Sinh  func(ffi.ComplexFloat) ffi.ComplexFloat `ffi:"csinhf"`
	Cosh  func(ffi.ComplexFloat) ffi.ComplexFloat `ffi:"ccoshf"`
	Tanh  func(ffi.ComplexFloat) ffi.ComplexFloat `ffi:"ctanhf"`
	Asinh func(ffi.ComplexFloat) ffi.ComplexFloat `ffi:"casinhf"`
	Acosh func(ffi.ComplexFloat) ffi.ComplexFloat `ffi:"cacoshf"`
	Atanh func(ffi.ComplexFloat) ffi.ComplexFloat `ffi:"catanhf"`
}

var ComplexDoubleLong struct {
	LibM

	Real func(ffi.ComplexDoubleLong) ffi.DoubleLong        `ffi:"creall"`
	Imag func(ffi.ComplexDoubleLong) ffi.DoubleLong        `ffi:"cimagl"`
	Abs  func(ffi.ComplexDoubleLong) ffi.DoubleLong        `ffi:"cabsl"`
	Arg  func(ffi.ComplexDoubleLong) ffi.DoubleLong        `ffi:"cargl"`
	Conj func(ffi.ComplexDoubleLong) ffi.ComplexDoubleLong `ffi:"conjl"`
	Proj func(ffi.ComplexDoubleLong) ffi.ComplexDoubleLong `ffi:"cprojl"`

	Exp func(ffi.ComplexDoubleLong) ffi.ComplexDoubleLong                        `ffi:"cexpl"`
	Log func(ffi.ComplexDoubleLong) ffi.ComplexDoubleLong                        `ffi:"clogl"`
	Pow func(ffi.ComplexDoubleLong, ffi.ComplexDoubleLong) ffi.ComplexDoubleLong `ffi:"cpowl"`

	Sqrt func(ffi.ComplexDoubleLong) ffi.ComplexDoubleLong `ffi:"csqrtl"`
	Sin  func(ffi.ComplexDoubleLong) ffi.ComplexDoubleLong `ffi:"csinl"`
	Cos  func(ffi.ComplexDoubleLong) ffi.ComplexDoubleLong `ffi:"ccosl"`
	Tan  func(ffi.ComplexDoubleLong) ffi.ComplexDoubleLong `ffi:"ctanl"`
	Asin func(ffi.ComplexDoubleLong) ffi.ComplexDoubleLong `ffi:"casinl"`
	Acos func(ffi.ComplexDoubleLong) ffi.ComplexDoubleLong `ffi:"cacosl"`
	Atan func(ffi.ComplexDoubleLong) ffi.ComplexDoubleLong `ffi:"catanl"`

	Sinh  func(ffi.ComplexDoubleLong) ffi.ComplexDoubleLong `ffi:"csinhl"`
	Cosh  func(ffi.ComplexDoubleLong) ffi.ComplexDoubleLong `ffi:"ccoshl"`
	Tanh  func(ffi.ComplexDoubleLong) ffi.ComplexDoubleLong `ffi:"ctanhl"`
	Asinh func(ffi.ComplexDoubleLong) ffi.ComplexDoubleLong `ffi:"casinhl"`
	Acosh func(ffi.ComplexDoubleLong) ffi.ComplexDoubleLong `ffi:"cacoshl"`
	Atanh func(ffi.ComplexDoubleLong) ffi.ComplexDoubleLong `ffi:"catanhl"`
}
