package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const winWidth, winHeight int = 800, 600

type color struct {
	r, g, b byte
}

func setPixel(x, y int, c color, pixels []byte) {
	index := (y*winWidth + x) * 4

	if index < len(pixels)-4 && index >= 0 {
		pixels[index] = c.r
		pixels[index+1] = c.g
		pixels[index+2] = c.b
	}
}
func main() {

	sdl.Init(sdl.INIT_EVERYTHING)
	defer sdl.Quit()

	// window, err := sdl.CreateWindow("testing SDL2", 100, 100, 800, 600, sdl.WINDOW_SHOWN)

	// create a window and defer its destroy
	window, err := sdl.CreateWindow("testing SDL2", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(winWidth), int32(winHeight), sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer window.Destroy()

	// create a renderer for the window and defer its destroy
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer renderer.Destroy()

	// create a texture on the renderer with a format and defer its destroy
	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(winWidth), int32(winHeight))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer tex.Destroy()

	// byte array of pixels
	// 0th through 8th byte in the array can contain one color
	// screen height times width give the size and times 4 is for each byte in the pixel format
	pixels := make([]byte, winWidth*winHeight*4)

	for y := 0; y < winHeight; y++ {
		for x := 0; x < winWidth; x++ {
			// setPixel(x, y, color{255, 0, 0}, pixels)
			setPixel(x, y, color{byte(x % 255), byte(y % 255), 0}, pixels)
		}
	}

	// update the texture with our pixels
	tex.Update(nil, pixels, winWidth*4)

	// copy the texture to the renderer and present the renderer
	renderer.Copy(tex, nil, nil)
	renderer.Present()

	// found this on the web
	// running := true
	// for running {
	// 	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
	// 		switch event.(type) {
	// 		case *sdl.QuitEvent:
	// 			println("Quit")
	// 			running = false
	// 			break
	// 		}
	// 	}
	// 	sdl.Delay(16)
	// }

	// Jack provided this (and I found it after working thru all this already! ;) )
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
	}
}
