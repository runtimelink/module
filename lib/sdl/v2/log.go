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
	"runtime.link/std"
)

type LogCategory std.Enum

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

type LogPriority std.Enum

const (
	LogAsVerbose LogPriority = iota
	LogAsDebug
	LogAsInfo
	LogAsWarn
	LogAsError
	LogAsCritical
)

type Log struct {
	location

	Printf         func(string, ...std.UnsafePointer)                           `ffi:"SDL_Log"`
	Message        func(LogCategory, LogPriority, string, ...std.UnsafePointer) `ffi:"SDL_LogMessage"`
	SetAllPriority func(LogPriority)                                            `ffi:"SDL_LogSetAllPriority"`
	SetPriority    func(LogCategory, LogPriority)                               `ffi:"SDL_LogSetPriority"`

	Verbose  func(string, ...std.UnsafePointer) `ffi:"SDL_LogVerbose"`
	Debug    func(string, ...std.UnsafePointer) `ffi:"SDL_LogDebug"`
	Info     func(string, ...std.UnsafePointer) `ffi:"SDL_LogInfo"`
	Warn     func(string, ...std.UnsafePointer) `ffi:"SDL_LogWarn"`
	Error    func(string, ...std.UnsafePointer) `ffi:"SDL_LogError"`
	Critical func(string, ...std.UnsafePointer) `ffi:"SDL_LogCritical"`
}
