package main

import tl "github.com/JoelOtter/termloop"

type Ball struct {
	*tl.Entity
	Color int
	X     int
	Y     int
}

func NewBall(color, x, y int) Ball {
	ball := Ball{tl.NewEntity(1, 1, 5, 3), color, x, y}

	if color == RED {
		ball.SetCell(2, 1, &tl.Cell{Fg: tl.ColorRed, Ch: '🔴'})
	} else {
		ball.SetCell(2, 1, &tl.Cell{Fg: tl.ColorBlue, Ch: '🔵'})
	}

	return ball
}
