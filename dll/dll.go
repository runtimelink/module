// Package dll provides an interface for loading C shared libraries at runtime.
package dll

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"runtime"
	"strings"
	"unsafe"

	"runtime.link/cgo"
	"runtime.link/std"
)

// Import dynamically links to the the specified library.
// If any symbols fail to load, the corresponding functions
// will panic. Library locations provided to this function
// will override the default ones to search for.
//
// Library should be a struct of functions, each function must
// clearly tag a standard symbol (see the [std] package).
//
// For example:
//
//	PutString func(string) error `std:"puts func(&char)int<0"`
//
// As long as the [std] tags correctly describe the standard
// function's signature and memory semantics, the function
// is safe to use from Go.
//
// Packages under runtime.link/lib are specifically designed
// to be memory safe.
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
		if err := set(&lib, string(found.Tag.Get("std"))); err != nil {
			log.Println(err)
		}
	}
	return lib
}

/*func sigRune(t reflect.Type) rune {
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
}*/

func set(library any, tag string) error {
	var (
		libs []unsafe.Pointer
	)
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
	var (
		rtype  = reflect.TypeOf(library).Elem()
		rvalue = reflect.ValueOf(library).Elem()
	)
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
		var ptr any
		if field.IsExported() {
			ptr = value.Addr().Interface()
		} else {
			ptr = reflect.NewAt(field.Type, unsafe.Add(rvalue.Addr().UnsafePointer(), field.Offset)).Interface()
		}

		if err := cgo.Linker(func(name string) unsafe.Pointer {
			for _, lib := range libs {
				if ptr := dlsym(lib, name); ptr != nil {
					return ptr
				}
			}
			return nil
		}).MakeFunc(ptr, std.Tag(field.Tag.Get("std"))); err != nil {
			log.Println(err)
		}
	}
	return nil
}
