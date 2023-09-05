// Package dll provides an interface for loading C shared libraries at runtime.
package dll

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"unsafe"

	"runtime.link/dll/internal/dyncall"
	"runtime.link/std"
)

// #include <internal/dyncall/dyncall.h>
import "C"

// Import dynamically links to the the specified library.
// If any symbols fail to load, the corresponding functions
// will panic. Library locations provided to this function
// will override the default ones to search for.
//
// Library should be a struct of functions, each function
// must clearly define a standard signature and symbol.
// This can be achieved by sticking to std package types, or
// by using a std tag that defines the signature.
//
// For example:
//
//	PutString func(std.String) std.Int `sym:"puts"`   // unsafe, pointer values passed directly.
//	PutString func(string) int `std:"int puts(char)"` // safest, deep copy and convert all values.
//
// The 'std' is similar to a C function signature, but
// with *, [] symbols and the argument names omitted
// (function arguments are specified using 'void').
// Import will not free memory by default, as this
// is the safest option. In order to prevent these
// memory leaks, the function signature can have
// appropriate parameter annotations.
//
// Import may use these annotations to optimize calls
// and decide how pointers are passed.
//
//	'#type'     - tags this symbol as the destructor for
//				  the given type. Which will be used
//				  to track ownership disposal.
//	'&type'     - the receiver borrows this pointer and
//				  will not keep a reference to it.
//	'-type'     - the receiver ignores this parameter.
//	'type=0'    - set to zero
//	'type=1'    - set to one
//	'type<0'     - signals error when smaller than zero.
//	'type!'     - signals error when zero.
//	'$type'     - the receiver takes ownership of this
//				  pointer and is responsible for freeing it.
//	'type%v'    - the argument identified by the given
//		          fmt parameter is mapped here. Must
//		          come before other suffixed annotations.
//	'type|%v'   - the argument must have greater length than
//				  the argument identified by the given fmt
//				  parameter.
//	'type||%v'   - the argument must have greater capacity
//				  to the length of the argumenent identified
//				  by the given fmt parameter.
//	'type?sym'  - (for return type only) when the value
//				  is not empty, return the result from
//				  the given symbol instead. Otherwise
//				  return the zero value. Either directly
//				  or in an additional return value (if specified).
//	'type!sym'  - (for argument type only) when the value
//				  is empty, return the result from the
//				  given symbol instead, either directly
//				  or in an additional return value (if specified).
//				  if 'sym' is omitted, invert the output.
//	'free@sym'  - frees the memory allocated because of
//				  this parameter, right after the next time
//				  the given symbol is called with a matching
//			 	  pattern.
//	'ptrdiff%v' - the argument identified by the given
//				  fmt parameter is assumed to be a pointer
//				  within that parameter.
//	'null'	    - like void but a null char is appended to
//				  the end of it. works only for []byte.
//	'vfmt%v'   	- the arguments are validated to correspond
//				  to the given fmt string.
//
// 'sym' name can have optional pattern {} where each
// comma separated value is either a fmt parameter or
// underscore (wildcard). The fmt parameters indicate
// how arguments from the function are mapped to the
// arguments of the sumbol.
//
// Structs and struct pointers must either be entirely
// composed of std typed fields, or have std tags on
// each field that define the C type. Field order must
// match the C struct definition. If there are layout
// or alignment differences between the C and Go structs,
// or non-std Go types are being used, then the struct
// must embed a std.Struct field.
//
//	// safest, deep copy all pointers to this struct.
//	type MyStruct {
//		std.Struct // if-in-doubt, embed this.
//
//		Name string `std:"char"`
//	}
//
//	// fastest, struct pointers passed directly
//	type MyStruct {
//		Name std.String
//	}
//
// IMPORT IS FUNDAMENTALLY UNSAFE
// Although it will validate what it can in order to
// ensure safety. Callers unfamiliar with C should
// stick to the 'std' tag and avoid libraries that
// require C struct values to be accessed directly.
//
// Alternatively, use a library with an existing
// representation in Go, as can be found under
// runtime.link/lib
func Import[Library any](names ...string) Library {
	var lib Library
	for _, name := range names {
		if err := set(&lib, name); err == nil {
			return lib
		}
	}
	location := reflect.TypeOf(&lib).Elem()
	found, ok := location.FieldByName(runtime.GOOS)
	if !ok {
		found, ok = location.Field(0).Type.FieldByName(runtime.GOOS)
	}
	if !ok && len(names) == 0 {
		panic(fmt.Sprintf("library for %T not available on %s", lib, runtime.GOOS))
	}
	if ok {
		if err := set(&lib, string(found.Tag)); err != nil {
			log.Println(err)
		}
	}
	return lib
}

var vm4096 sync.Pool
var vm8 sync.Pool

func init() {
	vm8.New = func() any {
		return dyncall.NewVM(8)
	}
	vm4096.New = func() any {
		return dyncall.NewVM(4096)
	}
}

func sigRune(t reflect.Type) rune {
	switch t.Kind() {
	case reflect.TypeOf(std.Bool(0)).Kind():
		return dyncall.Bool
	case reflect.TypeOf(std.Char(0)).Kind():
		return dyncall.Char
	case reflect.TypeOf(std.UnsignedChar(0)).Kind():
		return dyncall.UnsignedChar
	case reflect.TypeOf(std.Short(0)).Kind():
		return dyncall.Short
	case reflect.TypeOf(std.UnsignedShort(0)).Kind():
		return dyncall.UnsignedShort
	case reflect.TypeOf(std.Int(0)).Kind():
		return dyncall.Int
	case reflect.TypeOf(std.UnsignedInt(0)).Kind():
		return dyncall.Uint
	case reflect.TypeOf(std.Long(0)).Kind():
		return dyncall.Long
	case reflect.TypeOf(std.UnsignedLong(0)).Kind():
		return dyncall.UnsignedLong
	case reflect.TypeOf(std.LongLong(0)).Kind():
		return dyncall.LongLong
	case reflect.TypeOf(std.UnsignedLongLong(0)).Kind():
		return dyncall.UnsignedLongLong
	case reflect.TypeOf(std.Float(0)).Kind():
		return dyncall.Float
	case reflect.TypeOf(std.Double(0)).Kind():
		return dyncall.Double
	case reflect.String:
		return dyncall.String
	case reflect.Pointer:
		return dyncall.Pointer
	case reflect.Struct:
		if t.Implements(reflect.TypeOf([0]std.IsPointer{}).Elem()) {
			return dyncall.Pointer
		} else {
			panic("unsupported struct " + t.String())
		}
	default:
		panic("unsupported type " + t.String())
	}
}

func newSignature(ftype reflect.Type) dyncall.Signature {
	var sig dyncall.Signature
	for i := 0; i < ftype.NumIn(); i++ {
		sig.Args = append(sig.Args, sigRune(ftype.In(i)))
	}
	if ftype.NumOut() > 1 {
		sig.Returns = sigRune(ftype.Out(0))
	} else {
		sig.Returns = dyncall.Void
	}
	return sig
}

func newCallback(signature dyncall.Signature, function reflect.Value) dyncall.CallbackHandler {
	return func(cb *dyncall.Callback, args *dyncall.Args, result unsafe.Pointer) rune {
		var values = make([]reflect.Value, len(signature.Args))
		for i := range values {
			values[i] = reflect.New(function.Type().In(i)).Elem()
		}
		for i := 0; i < len(signature.Args); i++ {
			switch signature.Args[i] {
			case dyncall.Bool:
				switch args.Bool() {
				case 0:
					values[i].SetBool(false)
				case 1:
					values[i].SetBool(true)
				}
			case dyncall.Char:
				values[i].SetInt(int64(args.Char()))
			case dyncall.UnsignedChar:
				values[i].SetUint(uint64(args.UnsignedChar()))
			case dyncall.Short:
				values[i].SetInt(int64(args.Short()))
			case dyncall.UnsignedShort:
				values[i].SetUint(uint64(args.UnsignedShort()))
			case dyncall.Int:
				values[i].SetInt(int64(args.Int()))
			case dyncall.Uint:
				values[i].SetUint(uint64(args.UnsignedInt()))
			case dyncall.Long:
				values[i].SetInt(int64(args.Long()))
			case dyncall.UnsignedLong:
				values[i].SetUint(uint64(args.UnsignedLong()))
			case dyncall.LongLong:
				values[i].SetInt(int64(args.LongLong()))
			case dyncall.UnsignedLongLong:
				values[i].SetUint(uint64(args.UnsignedLongLong()))
			case dyncall.Float:
				values[i].SetFloat(float64(args.Float()))
			case dyncall.Double:
				values[i].SetFloat(float64(args.Double()))
			case dyncall.String:
				ptr := args.Pointer()
				switch values[i].Kind() {
				case reflect.String:
					values[i].SetString(C.GoString((*C.char)(ptr)))
				case reflect.Struct:
					values[i].Set(reflect.ValueOf(*(*std.String)(unsafe.Pointer(&ptr))))
				default:
					panic("unsupported type " + values[i].Type().String())
				}
			case dyncall.Pointer:
				switch values[i].Kind() {
				case reflect.UnsafePointer:
					values[i].SetPointer(unsafe.Pointer(args.Pointer()))
				default:
					settable, ok := values[i].Addr().Interface().(interface {
						SetPointer(unsafe.Pointer)
					})
					if !ok {
						panic("unsupported type " + values[i].Type().String())
					}
					settable.SetPointer(unsafe.Pointer(args.Pointer()))
				}
			default:
				panic("unsupported type " + string(signature.Args[i]))
			}
		}
		results := function.Call(values)
		switch signature.Returns {
		case dyncall.Void:
		case dyncall.Bool:
			var b std.Bool
			if results[0].Bool() {
				b = std.True
			}
			*(*std.Bool)(result) = b
		case dyncall.Char:
			*(*std.Char)(result) = std.Char(results[0].Int())
		case dyncall.UnsignedChar:
			*(*std.UnsignedChar)(result) = std.UnsignedChar(results[0].Uint())
		case dyncall.Short:
			*(*std.Short)(result) = std.Short(results[0].Int())
		case dyncall.UnsignedShort:
			*(*std.UnsignedShort)(result) = std.UnsignedShort(results[0].Uint())
		case dyncall.Int:
			*(*std.Int)(result) = std.Int(results[0].Int())
		case dyncall.Uint:
			*(*std.UnsignedInt)(result) = std.UnsignedInt(results[0].Uint())
		case dyncall.Long:
			*(*std.Long)(result) = std.Long(results[0].Int())
		case dyncall.UnsignedLong:
			*(*std.UnsignedLong)(result) = std.UnsignedLong(results[0].Uint())
		case dyncall.LongLong:
			*(*std.LongLong)(result) = std.LongLong(results[0].Int())
		case dyncall.UnsignedLongLong:
			*(*std.UnsignedLongLong)(result) = std.UnsignedLongLong(results[0].Uint())
		case dyncall.Float:
			*(*std.Float)(result) = std.Float(results[0].Float())
		case dyncall.Double:
			*(*std.Double)(result) = std.Double(results[0].Float())
		case dyncall.String:
			*(*std.String)(result) = std.StringOf(results[0].String()) // FIXME allocate in C memory?
		case dyncall.Pointer:
			*(*unsafe.Pointer)(result) = results[0].UnsafePointer()
		default:
			panic("unsupported type " + results[0].Type().String())
		}
		return signature.Returns
	}
}

func set(library any, tag string) error {
	var libs []unsafe.Pointer

	for _, location := range strings.Split(tag, ",") {
		for _, name := range strings.Split(location, " ") {
			lib := dlopen(name)
			if lib == nil {
				continue
			}
			libs = append(libs, lib)
		}
	}

	if len(libs) == 0 {
		return errors.New(tag + " not found")
	}

	rtype := reflect.TypeOf(library).Elem()
	rvalue := reflect.ValueOf(library).Elem()

	for i := 0; i < rtype.NumField(); i++ {
		field := rtype.Field(i)
		value := rvalue.Field(i)

		if field.IsExported() && field.Type.Kind() == reflect.Struct {
			if err := set(value.Addr().Interface(), tag); err != nil {
				return err
			}
		}

		if field.Type.Kind() != reflect.Func {
			continue
		}

		name := field.Tag.Get("ffi")
		if name == "" {
			name = field.Tag.Get("sym")
			if name == "" {
				name = field.Name
			}
		}
		symbols := strings.Split(name, ",")
		var symbol unsafe.Pointer
		for _, name := range symbols {
			for _, lib := range libs {
				symbol = dlsym(lib, name)
				if symbol == nil {
					continue
				}
			}
			if symbol == nil {
				continue
			}
		}
		if symbol == nil {
			log.Println(errors.New(dlerror()))
			continue
		}
		getErr := rvalue.FieldByName("Error")

		switch fn := value.Addr().Interface().(type) {
		case *func(std.Double) std.Double:
			*fn = func(a std.Double) std.Double {
				vm := vm8.Get().(*dyncall.VM)
				vm.Reset()
				defer vm8.Put(vm)
				vm.PushFloat64(float64(a))
				return std.Double(vm.CallFloat64(symbol))
			}
		case *func():
			*fn = func() {
				vm := dyncall.NewVM(0)
				defer vm.Free()
				vm.Call(symbol)
			}
		default:
			_ = fn
			value.Set(reflect.MakeFunc(field.Type, func(args []reflect.Value) []reflect.Value {
				var vm = dyncall.NewVM(4096)
				defer vm.Free()
				push := func(value reflect.Value) {
					switch value.Kind() {
					case reflect.Bool:
						vm.PushBool(value.Bool())
					case reflect.Int8:
						vm.PushInt8(int8(value.Int()))
					case reflect.Int16:
						vm.PushInt16(int16(value.Int()))
					case reflect.Int32:
						vm.PushInt32(int32(value.Int()))
					case reflect.Int64:
						vm.PushInt64(value.Int())
					case reflect.Uint8:
						u8 := uint8(value.Uint())
						vm.PushInt8(*(*int8)(unsafe.Pointer(&u8)))
					case reflect.Uint16:
						u16 := uint16(value.Uint())
						vm.PushInt16(*(*int16)(unsafe.Pointer(&u16)))
					case reflect.Uint32:
						u32 := uint32(value.Uint())
						vm.PushInt32(*(*int32)(unsafe.Pointer(&u32)))
					case reflect.Uint64:
						u64 := uint64(value.Uint())
						vm.PushInt64(*(*int64)(unsafe.Pointer(&u64)))
					case reflect.Float32:
						vm.PushFloat32(float32(value.Float()))
					case reflect.Float64:
						vm.PushFloat64(value.Float())
					case reflect.Pointer, reflect.UnsafePointer:
						vm.PushPointer(value.UnsafePointer())
					case reflect.String:
						s := std.StringOf(value.String())
						vm.PushPointer(unsafe.Pointer(s.UnsafePointer()))
					case reflect.Struct:
						if value.Type().Implements(reflect.TypeOf([0]std.IsPointer{}).Elem()) {
							vm.PushPointer(unsafe.Pointer(value.Interface().(std.IsPointer).Pointer()))
						} else {
							panic("unsupported struct " + value.Type().String())
						}
					case reflect.Func:
						signature := newSignature(value.Type())
						ptr := dyncall.NewCallback(signature, newCallback(signature, value))
						vm.PushPointer(unsafe.Pointer(ptr))
					default:
						panic("unsupported type " + value.Type().String())
					}
				}
				for _, arg := range args {
					push(arg)
				}
				var results = make([]reflect.Value, field.Type.NumOut())
				for i := 0; i < field.Type.NumOut(); i++ {
					results[i] = reflect.New(field.Type.Out(i)).Elem()
				}
				var returnsError bool
				switch field.Type.NumOut() {
				default:
					if field.Type.NumOut() > 1 {
						length := field.Type.NumOut()
						if field.Type.Out(length-1) == reflect.TypeOf([0]error{}).Elem() {
							returnsError = true
							length--
						}
						for i := 1; i < length; i++ {
							push(results[i].Addr())
						}
					}
					fallthrough
				case 1:
					rtype := field.Type.Out(0)
					switch field.Type.Out(0).Kind() {
					case reflect.Bool:
						results[0].SetBool(vm.CallBool(symbol))
					case reflect.Int8:
						results[0].SetInt(int64(vm.CallInt8(symbol)))
					case reflect.Int16:
						results[0].SetInt(int64(vm.CallInt16(symbol)))
					case reflect.Int32:
						results[0].SetInt(int64(vm.CallInt32(symbol)))
					case reflect.Int64:
						results[0].SetInt(int64(vm.CallInt64(symbol)))
					case reflect.Uint8:
						u8 := vm.CallInt8(symbol)
						results[0].SetUint(uint64(*(*uint8)(unsafe.Pointer(&u8))))
					case reflect.Uint16:
						u16 := vm.CallInt16(symbol)
						results[0].SetUint(uint64(*(*uint16)(unsafe.Pointer(&u16))))
					case reflect.Uint32:
						u32 := vm.CallInt32(symbol)
						results[0].SetUint(uint64(*(*uint32)(unsafe.Pointer(&u32))))
					case reflect.Uint64:
						u64 := vm.CallInt64(symbol)
						results[0].SetUint(uint64(*(*uint64)(unsafe.Pointer(&u64))))
					case reflect.Float32:
						results[0].SetFloat(float64(vm.CallFloat32(symbol)))
					case reflect.Float64:
						results[0].SetFloat(float64(vm.CallFloat64(symbol)))
					case reflect.String:
						results[0].SetString(C.GoString((*C.char)(vm.CallPointer(symbol))))
					case reflect.UnsafePointer:
						results[0].SetPointer(vm.CallPointer(symbol))
					case reflect.Pointer:
						results[0] = reflect.NewAt(field.Type.Out(0).Elem(), unsafe.Pointer(vm.CallPointer(symbol)))
					case reflect.Struct:
						if rtype.Implements(reflect.TypeOf([0]std.IsPointer{}).Elem()) {
							*(*unsafe.Pointer)(results[0].Addr().UnsafePointer()) = vm.CallPointer(symbol)
						} else {
							panic("unsupported struct " + field.Type.Out(0).String())
						}
					default:
						panic("unsupported type " + field.Type.Out(0).String())
					}
				case 0:
					vm.Call(symbol)
				}
				if returnsError {
					if results[0].IsZero() {
						if !getErr.IsValid() {
							panic("an error occured")
						}
						switch fn := getErr.Interface().(type) {
						case func() string:
							results[len(results)-1] = reflect.ValueOf(errors.New(fn()))
						}
					}
				}
				return results
			}))
		}
	}

	return nil
}
