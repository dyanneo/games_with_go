package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const winWidth, winHeight int = 800, 600

type color struct {
	r, g, b byte
}
type pos struct {
	x, y float32
}
type ball struct {
	pos
	radius int
	xv     float32 // velocity in direction of x
	yv     float32 // velocity in direction of y
	color  color
}

func (ball *ball) draw(pixels []byte) {
	for y := -ball.radius; y < ball.radius; y++ {
		for x := -ball.radius; x < ball.radius; x++ {
			// per Jack, may not be super efficient, but compares each pixel position to whether it is within the ball's radius
			// to determine if it should be colored white
			if x*x+y*y < ball.radius*ball.radius {
				setPixel(int(ball.x)+x, int(ball.y)+y, ball.color, pixels)
			}
		}
	}
}

func (ball *ball) update(leftPaddle *paddle, rightPaddle *paddle) {
	ball.x += ball.xv
	ball.y += ball.yv

	// if ball goes to top or bottom edge, bounce it back
	if int(ball.y)-ball.radius < 0 || int(ball.y)+ball.radius > winHeight {
		ball.yv = -ball.yv
	}

	// if ball goes past the sides, put it back to where it started
	if int(ball.x)-ball.radius < 0 || int(ball.x)+ball.radius > winWidth {
		ball.x = 400
		ball.y = 300
	}

	if int(ball.x)-ball.radius < int(leftPaddle.x)+leftPaddle.w/2 {
		if int(ball.y) > int(leftPaddle.y)-leftPaddle.h/2 && int(ball.y) < int(leftPaddle.y)+leftPaddle.h/2 {
			ball.xv = -ball.xv
		}
	}
	if int(ball.x)+ball.radius > int(rightPaddle.x)-rightPaddle.w/2 {
		if int(ball.y) > int(rightPaddle.y)-rightPaddle.h/2 && int(ball.y) < int(rightPaddle.y)+rightPaddle.h/2 {
			ball.xv = -ball.xv
		}
	}
}

// note how pos is NOT named in paddle - go will 'bring along' the pos struct within the paddle struct
// and can be referenced like paddle.x

type paddle struct {
	pos
	w     int
	h     int
	color color
}

func (paddle *paddle) draw(pixels []byte) {
	startX := int(paddle.x) - paddle.w/2
	startY := int(paddle.y) - paddle.h/2
	// starting with y because memory is loaded at least 64 bytes at a time, (as opposed to the one byte you want)
	// and the more often you can move through it sequentially, the faster it will go
	// the way a square screen is loaded is sequential starting with y, not x

	for y := 0; y < paddle.h; y++ {
		for x := 0; x < paddle.w; x++ {
			setPixel(startX+x, startY+y, paddle.color, pixels)
		}
	}
}

func (paddle *paddle) update(keyState []uint8) {
	if keyState[sdl.SCANCODE_UP] != 0 {
		paddle.y -= 5
	}
	if keyState[sdl.SCANCODE_DOWN] != 0 {
		paddle.y += 5
	}
}

func (paddle *paddle) aiUpdate(ball *ball) {
	paddle.y = ball.y
}

func clear(pixels []byte) {
	for i := range pixels {
		pixels[i] = 0
	}
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

	window, err := sdl.CreateWindow("testing SDL2", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(winWidth), int32(winHeight), sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer renderer.Destroy()

	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(winWidth), int32(winHeight))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer tex.Destroy()

	pixels := make([]byte, winWidth*winHeight*4)

	player1 := paddle{pos{50, 100}, 20, 100, color{255, 255, 255}}
	player2 := paddle{pos{float32(winWidth) - 50, 100}, 20, 100, color{255, 0, 0}}
	ball := ball{pos{400, 300}, 20, 5, 5, color{255, 255, 255}}

	keyState := sdl.GetKeyboardState()

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		clear(pixels)

		player1.update(keyState)
		player2.aiUpdate(&ball)
		ball.update(&player1, &player2)

		player1.draw(pixels)
		player2.draw(pixels)
		ball.draw(pixels)

		tex.Update(nil, pixels, winWidth*4)

		renderer.Copy(tex, nil, nil)
		renderer.Present()

		sdl.Delay(16)
	}
}
