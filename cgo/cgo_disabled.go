//go:build !cgo

package cgo

import (
	"runtime.link/std"
)

func (ln Linker) makeFunc(fn any, tag std.Tag) error {
	return ErrDisabled
}
