/*

  Simple DirectMedia Layer
  Copyright (C) 1997-2023 Sam Lantinga <slouken@libsdl.org>

  This software is provided 'as-is', without any express or implied
  warranty.  In no event will the authors be held liable for any damages
  arising from the use of this software.

  Permission is granted to anyone to use this software for any purpose,
  including commercial applications, and to alter it and redistribute it
  freely, subject to the following restrictions:

  1. The origin of this software must not be misrepresented; you must not
     claim that you wrote the original software. If you use this software
     in a product, an acknowledgment in the product documentation would be
     appreciated but is not required.
  2. Altered source versions must be plainly marked as such, and must not be
     misrepresented as being the original software.
  3. This notice may not be removed or altered from any source distribution.

*/

package sdl

import (
	"runtime.link/ffi"
)

type LogCategory ffi.Enum

const (
	LogApplication LogCategory = iota
	LogError
	LogAssert
	LogSystem
	LogAudio
	LogVideo
	LogRender
	LogInput
)

type LogPriority ffi.Enum

const (
	LogAsVerbose LogPriority = iota
	LogAsDebug
	LogAsInfo
	LogAsWarn
	LogAsError
	LogAsCritical
)

var Log struct {
	Lib

	Printf         func(string, ...ffi.UnsafePointer)                           `ffi:"SDL_LogPrintf"`
	Message        func(LogCategory, LogPriority, string, ...ffi.UnsafePointer) `ffi:"SDL_LogMessage"`
	SetAllPriority func(LogPriority)                                            `ffi:"SDL_LogSetAllPriority"`
	SetPriority    func(LogCategory, LogPriority)                               `ffi:"SDL_LogSetPriority"`

	Verbose  func(string, ...ffi.UnsafePointer) `ffi:"SDL_LogVerbose"`
	Debug    func(string, ...ffi.UnsafePointer) `ffi:"SDL_LogDebug"`
	Info     func(string, ...ffi.UnsafePointer) `ffi:"SDL_LogInfo"`
	Warn     func(string, ...ffi.UnsafePointer) `ffi:"SDL_LogWarn"`
	Error    func(string, ...ffi.UnsafePointer) `ffi:"SDL_LogError"`
	Critical func(string, ...ffi.UnsafePointer) `ffi:"SDL_LogCritical"`
}
