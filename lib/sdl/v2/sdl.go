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
	"unsafe"

	"runtime.link/std"
)

type location struct {
	linux  std.Location `libSDL2-2.0.so.0`
	darwin std.Location `libSDL2.dylib,/opt/homebrew/lib/libSDL2.dylib`
}

type Library struct {
	location

	Atomics      Atomics
	Audio        Audio
	AudioDevices AudioDevices
	AudioStreams AudioStreams
	Windows      Windows
	Draw         Draw
	Timer        Timer
	System       System
	Events       Events
	Errors       Errors
	Video        Video
	Log          Log
	Surfaces     Surfaces
}
type Module std.Uint32

// These are the flags which may be passed to System.Init(). You should
// specify the modules which you will be using in your application.
const (
	ModuleTimer          Module = 0x00000001 // timer module
	ModuleAudio          Module = 0x00000010 // audio module
	ModuleVideo          Module = 0x00000020 // video module; automatically initializes the events module
	ModuleJoystick       Module = 0x00000200 // joystick module; automatically initializes the events module
	ModuleHaptic         Module = 0x00001000 // haptic (force feedback) module
	ModuleGameController Module = 0x00002000 // controller module; automatically initializes the joystick module
	ModuleEvents         Module = 0x00004000 // events module
	ModuleSensor         Module = 0x00008000
	Modules              Module = ModuleTimer | ModuleAudio | ModuleVideo | ModuleJoystick | ModuleHaptic | ModuleGameController | ModuleEvents | ModuleSensor // all of the above modules                                                                                                  // compatibility; this flag is ignored
)

type System struct {
	location

	/*
		Init initialize the SDL library.

		The file I/O (for example: File.ReadWrite) and threading (Threads.Create)
		subsystems are initialized by default. Message boxes
		(GUI.ShowSimpleMessageBox) also attempt to work without initializing the
		video subsystem, in hopes of being useful in showing an error dialog when
		SDL_Init fails. You must specifically initialize other subsystems if you
		use them in your application.

		Logging (such as Log.Printf) works without initialization, too.

		* Subsystem initialization is ref-counted, you must call System.QuitSubSystem()
		* for each System.InitSubSystem() to correctly shutdown a subsystem manually (or
		* call SDL_Quit() to force shutdown). If a subsystem is already loaded then
		* this call will increase the ref-count and return.
	*/
	Init func(Module) std.Error `ffi:"SDL_Init"`
	/*
		Stop shuts down specific SDL subsystems.

		If you start a subsystem using a call to that subsystem's init function
		(for example SDL_VideoInit()) instead of SDL_Init() or SDL_InitSubSystem(),
		SDL_QuitSubSystem() and SDL_WasInit() will not work. You will need to use
		that subsystem's quit function (SDL_VideoQuit()) directly instead. But
		generally, you should not be using those functions directly anyhow; use
		SDL_Init() instead.

		You still need to call SDL_Quit() even if you close all open subsystems
		with SDL_QuitSubSystem().
	*/
	Stop   func(Module)        `ffi:"SDL_QuitSubSystem"`
	Loaded func(Module) Module `ffi:"SDL_WasInit"` // Loaded returns a mask of the specified subsystems which have previously been initialized.
	/*
		Quit cleans up all initialized subsystems.

		You should call this function even if you have already shutdown each
		initialized subsystem with SDL_QuitSubSystem(). It is safe to call this
		function even in the case of errors in initialization.

		If you start a subsystem using a call to that subsystem's init function
		(for example SDL_VideoInit()) instead of SDL_Init() or SDL_InitSubSystem(),
		then you must use that subsystem's quit function (SDL_VideoQuit()) to shut
		it down before calling SDL_Quit(). But generally, you should not be using
		those functions directly anyhow; use SDL_Init() instead.

		You can use this function with atexit() to ensure that it is run when your
		application is shutdown, but it is not wise to do this from a library or
		other dynamically loaded code.
	*/
	Quit func() `ffi:"SDL_Quit"` // Quit cleans up all initialized subsystems.

	Revision func() string  `ffi:"SDL_GetRevision"` // Revision returns the revision number of SDL that is linked against your program.
	Version  func(*Version) `ffi:"SDL_GetVersion"`  // Version returns the version of SDL that is linked against your program.

	DefaultAssertionHandler func() AssertionHandler                   `ffi:"SDL_GetAssertionHandler"`  // AssertionHandler returns the current assertion handler.
	SetAssertionHandler     func(AssertionHandler)                    `ffi:"SDL_SetAssertionHandler"`  // SetAssertionHandler sets a new assertion handler.
	GetAssertionHandler     func(*std.UnsafePointer) AssertionHandler `ffi:"SDL_GetAssertionHandler"`  // GetAssertionHandler returns the current assertion handler.
	GetAssertionReport      func() *AssertionData                     `ffi:"SDL_GetAssertionReport"`   // GetAssertionReport returns the last assertion reported, or nil if there weren't any.
	ResetAssertionReport    func()                                    `ffi:"SDL_ResetAssertionReport"` // ResetAssertionReport clears the list of all assertion failures.
}

type AssertionHandler std.Func[func(*AssertionData, std.UnsafePointer) AssertionState]

type AssertionState std.Enum

const (
	AssertionRetry        AssertionState = iota // Retry the assert immediately.
	AssertionBreak                              // Make the debugger trigger a breakpoint.
	AssertionAbort                              // Terminate the program.
	AssertionIgnore                             // Ignore the assert and continue execution.
	AssertionAlwaysIgnore                       // Ignore the assert from now on.
)

type AssertionData struct {
	AlwaysIgnore std.Int
	TriggerCount std.UnsignedInt
	Condition    std.String
	Filename     std.String
	LineNumber   std.Int
	Function     std.String
	Next         *AssertionData
}

type Window std.Handle[Window]

type WindowFlags std.Uint32

const (
	WindowOpenGL = 0x00000002
	WindowShown  = 0x00000004
)

const (
	WindowCentered = 0x2FFF0000
)

type Windows struct {
	location

	Error func() string `ffi:"SDL_GetError"`

	Create func(title string, x, y, w, h std.Int, flags WindowFlags) (Window, error) `ffi:"SDL_CreateWindow"`

	GetSurface    func(Window) (Surface, error) `ffi:"SDL_GetWindowSurface"`
	UpdateSurface func(Window) std.Error        `ffi:"SDL_UpdateWindowSurface"`
	Destroy       func(Window)                  `ffi:"SDL_DestroyWindow"`
}

type Surface std.Handle[Surface]

type Color std.Uint32

type Rect struct {
	X, Y, W, H std.Int
}

type Draw struct {
	location

	FilledRect func(Surface, *Rect, Color) `ffi:"SDL_FillRect"`
}

type Timer struct {
	location

	Delay func(ms std.Uint32) `ffi:"SDL_Delay"`
}

type MainFunc std.Func[func(std.Int, *std.String) std.Int]

type Version struct {
	Major std.Uint8
	Minor std.Uint8
	Patch std.Uint8
}

type Userdata std.Handle[Userdata]

type Bool std.Enum

const (
	False Bool = iota
	True
)

type eventType std.Uint32

const (
	eventQuit eventType = 0x100
)

type Event struct {
	etype eventType
	data  [max(
		unsafe.Sizeof(Quit{}),
		60,
	) - unsafe.Sizeof(std.Uint32(0))]byte
}

func (ev *Event) Data() any {
	switch ev.etype {
	case eventQuit:
		return (*Quit)(unsafe.Pointer(ev))
	default:
		return nil
	}
}

type Quit struct {
	_         eventType
	Timestamp std.Uint32
}

type Events struct {
	location

	Poll func(*Event) std.Int `ffi:"SDL_PollEvent"`
}

type Errors struct {
	location

	Clear           func()                                           `ffi:"SDL_ClearError"`
	Get             func() string                                    `ffi:"SDL_GetError"`
	GetErrorMessage func(std.String, std.Int) std.String             `ffi:"SDL_GetErrorMsg"`
	SetError        func(std.String, ...std.UnsafePointer) std.Error `ffi:"SDL_SetError"`
}
