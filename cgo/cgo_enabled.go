//go:build cgo

package cgo

// #include <internal/dyncall/dyncall.h>
import "C"
import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"

	"runtime.link/cgo/internal/dyncall"
	"runtime.link/std"
)

func (ln Linker) makeFunc(fn any, tag std.Tag) error {
	var (
		rtype = reflect.TypeOf(fn).Elem()
		value = reflect.ValueOf(fn).Elem()
	)
	symbols, ctype, err := tag.Parse()
	if err != nil {
		return err
	}
	if ctype.Func == nil {
		return TagCompatiblityError{tag, errorString("symbol is not a function"), rtype}
	}
	var symbol unsafe.Pointer
	for _, sym := range symbols {
		symbol = ln(sym)
		if symbol != nil {
			break
		}
	}
	if symbol == nil {
		return MissingSymbolError(strings.Join(symbols, ","))
	}
	value.Set(reflect.MakeFunc(rtype, func(args []reflect.Value) []reflect.Value {
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
			/*case reflect.Func:
			signature := newSignature(value.Type())
			ptr := dyncall.NewCallback(signature, newCallback(signature, value))
			vm.PushPointer(unsafe.Pointer(ptr))*/
			default:
				panic("unsupported type " + value.Type().String())
			}
		}
		for _, carg := range ctype.Args {
			fmt.Println(carg.Maps)
			push(args[carg.Maps-1])
		}
		var results = make([]reflect.Value, rtype.NumOut())
		for i := 0; i < rtype.NumOut(); i++ {
			results[i] = reflect.New(rtype.Out(i)).Elem()
		}
		var returnsError bool
		_ = returnsError
		switch rtype.NumOut() {
		default:
			if rtype.NumOut() > 1 {
				length := rtype.NumOut()
				if rtype.Out(length-1) == reflect.TypeOf([0]error{}).Elem() {
					returnsError = true
					length--
				}
				for i := 1; i < length; i++ {
					push(results[i].Addr())
				}
			}
			fallthrough
		case 1:
			result := rtype.Out(0)
			switch result.Kind() {
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
				results[0] = reflect.NewAt(rtype.Out(0).Elem(), unsafe.Pointer(vm.CallPointer(symbol)))
			case reflect.Struct:
				if rtype.Implements(reflect.TypeOf([0]std.IsPointer{}).Elem()) {
					*(*unsafe.Pointer)(results[0].Addr().UnsafePointer()) = vm.CallPointer(symbol)
				} else {
					panic("unsupported struct " + rtype.Out(0).String())
				}
			default:
				panic("unsupported type " + rtype.Out(0).String())
			}
		case 0:
			vm.Call(symbol)
		}
		/*if returnsError {
			if results[0].IsZero() {
				if !getErr.IsValid() {
					panic("an error occured")
				}
				switch fn := getErr.Interface().(type) {
				case func() string:
					results[len(results)-1] = reflect.ValueOf(errors.New(fn()))
				}
			}
		}*/
		return results
	}))
	return nil
}
