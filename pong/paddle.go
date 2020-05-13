package main

import (
	tl "github.com/JoelOtter/termloop"
)

const (
	PADDLE_WIDTH  = 11
	PADDLE_HEIGHT = 1
)

type Paddle struct {
	*tl.Rectangle
}

func (paddle *Paddle) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		x, y := paddle.Position()
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyArrowRight:
			if x < PLAY_WIDTH-6 {
				paddle.SetPosition(x+5, y)
			}
		case tl.KeyArrowLeft:
			if x > 6 {
				paddle.SetPosition(x-5, y)
			}
		case tl.KeyArrowUp:
			if y > 7 {
				paddle.SetPosition(x, y-2)
			}
		case tl.KeyArrowDown:
			if y < PLAY_HEIGHT+5 {
				paddle.SetPosition(x, y+2)
			}
		}
	}
}

func NewPaddle() Paddle {
	return Paddle{tl.NewRectangle(86, 25, PADDLE_WIDTH, PADDLE_HEIGHT, tl.ColorBlack)}
}
