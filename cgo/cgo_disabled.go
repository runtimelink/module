//go:build !cgo

package cgo

import (
	"runtime.link/std"
)

func makeFunc(fn any, tag std.Tag) error {
	return ErrDisabled
}
