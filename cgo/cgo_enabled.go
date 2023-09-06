//go:build cgo

package cgo

import "C"
import "runtime.link/std"

func makeFunc(fn any, tag std.Tag) error {
	return ErrDisabled
}
