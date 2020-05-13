package main

import (
	tl "github.com/JoelOtter/termloop"
	"math/rand"
	"time"
)

// Inital size of brick, which reduces over time.
const BRICK_SIZE_DEFAULT = 4

type Brick struct {
	*tl.Rectangle
	size int
}

func (brick *Brick) Tick(event tl.Event) {
	// We don't need to handle events
	return
}

func (brick *Brick) Move() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	x := r.Int31() % (PLAY_WIDTH - BRICK_SIZE_DEFAULT)
	y := r.Int31() % (PLAY_HEIGHT - BRICK_SIZE_DEFAULT)

	// reduce the size of the brick as the score increases.
	if Score == 4 || Score == 8 {
		brick.size -= 1
	}
	s := brick.size
	brick.SetSize(s+s, s)
	brick.SetPosition(int(x), int(y))
}

func NewBrick() Brick {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	x := r.Int31() % (PLAY_WIDTH - BRICK_SIZE_DEFAULT)
	y := r.Int31() % (PLAY_HEIGHT - BRICK_SIZE_DEFAULT)
	size := BRICK_SIZE_DEFAULT

	// set at least at Left offet of 6
	if x < 6 {
		x = 6
	}

	// set at least at Top offet of 6
	if y < 6 {
		y = 6
	}
	return Brick{tl.NewRectangle(int(x), int(y), size+size, size, tl.ColorMagenta), int(size)}
}
