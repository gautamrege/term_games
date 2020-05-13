package main

import (
	"time"

	tl "github.com/JoelOtter/termloop"
)

// Inital size of ball, which reduces over time.
const (
	BALL_FONT_SIZE = 4
	RIGHT          = 1
	DOWN           = 1
	LEFT           = -1
	UP             = -1
)

type Ball struct {
	*tl.Entity
	DirHorizontal int // 1: right, -1: left
	DirVertical   int // 1: down, -1: up
	speed         int
}

func (ball *Ball) new_position(x, y, xEdge, yEdge int) (nx, ny int) {
	if x <= xEdge { // left edge
		ball.DirHorizontal = RIGHT
	} else if x > PLAY_WIDTH+3 { // right edge
		ball.DirHorizontal = LEFT
	}

	if y <= yEdge { // top edge
		ball.DirVertical = DOWN
	} else if y > PLAY_HEIGHT+4 {
		ball.DirVertical = UP
	}

	nx = x + (ball.speed * ball.DirHorizontal)
	ny = y + (ball.speed * ball.DirVertical)

	return
}

func (ball *Ball) Draw(screen *tl.Screen) {
	x, y := ball.Position()

	nx, ny := ball.new_position(x, y, 6, 6)
	ball.SetPosition(nx, ny)

	ball.Entity.Draw(screen)
	time.Sleep(50 * time.Millisecond)
}

func (ball *Ball) IncSpeed() {
	ball.speed += 1
}

type BallPosition struct {
	x int
	y int
}

var playCh chan BallPosition

func ManageBrickHit(playCh chan BallPosition) {
	for data := range playCh {
		bx, by := data.x, data.y // ball position
		for _, b := range Bricks {
			// Check if the ball crashes into this brick
			x, y := b.Position()

			// Is ball{x,y} between the brick{x,y}?
			if bx > x && bx < (x+b.size) && by > y && by < (y+b.size) {
				Score += 1 // increment score
				if Score == 4 {
					GameBall.IncSpeed() // speed up the game
				}
				b.Move()
				break
			}
		}
	}
}

func (ball *Ball) Collide(collision tl.Physical) {
	bx, by := ball.Position()
	// Check if it's a Rectangle we're colliding with
	if _, ok := collision.(*Brick); ok { // we don't get it exact
		playCh <- BallPosition{bx, by}
	} else if paddle, ok := collision.(*Paddle); ok {
		x, y := paddle.Position()

		if x-bx < 10 && y-by < 2 {
			// toggle direction
			ball.DirVertical *= -1
			ball.SetPosition(bx+(ball.speed*ball.DirHorizontal), by+(ball.speed*ball.DirVertical))
		}
	}
}

func NewBall() *Ball {
	ball := Ball{tl.NewEntity(90, 24, PLAY_WIDTH, PLAY_HEIGHT), 1, 1, 1}
	ball.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '🔴'})

	playCh = make(chan BallPosition)
	go ManageBrickHit(playCh)

	return &ball
}
