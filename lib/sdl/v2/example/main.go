package main

import (
	"fmt"

	"runtime.link/dll"
	"runtime.link/lib/sdl/v2"
)

var SDL = dll.Import[sdl.Functions]()

func main() {
	if err := SDL.System.Init(sdl.Modules); err != 0 {
		panic(SDL.Errors.Get())
	}

	window, err := SDL.Windows.Create("Hello Square", sdl.WindowCentered,
		sdl.WindowCentered, 640, 480, sdl.WindowOpenGL|sdl.WindowShown)
	if err != nil {
		panic(err)
	}
	defer SDL.Windows.Destroy(window)

	surface, err := SDL.Windows.GetSurface(window)
	if err != nil {
		panic(err)
	}
	SDL.Draw.FilledRect(surface, nil, 0xFFFFFF)
	SDL.Draw.FilledRect(surface, &sdl.Rect{
		X: 640/2 - 50, Y: 480/2 - 50, W: 100, H: 100,
	}, 0xFF0000)
	SDL.Windows.UpdateSurface(window)

	fmt.Println(SDL.Audio.Driver())

	var event sdl.Event
	for ; true; SDL.Events.Poll(&event) {
		switch data := event.Data().(type) {
		case *sdl.Quit:
			fmt.Println(data.Timestamp)
			SDL.System.Quit()
			return
		}
	}
}
