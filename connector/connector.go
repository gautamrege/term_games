package main

import (
	tl "github.com/JoelOtter/termloop"
)

const (
	BOX_WIDTH  int = 6
	BOX_HEIGHT int = 3
	GRID       int = 6
)

type Location struct {
	X int
	Y int
}

var GAME Game = Game{}
var Matrix [][]*Box

var renderBall chan Ball
var renderWinner chan Player

func main() {
	game := tl.NewGame()
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorWhite,
	})

	// Where do we start drawing the grid from?
	origin := Location{60, 10}

	x := origin.X
	y := origin.Y

	Matrix = [][]*Box{}
	row := []*Box{}
	for r := 0; r < GRID; r++ {
		for c := 0; c < GRID; c++ {
			box := NewBox(x, y, r, c)
			level.AddEntity(&box)

			row = append(row, &box)
			x += BOX_WIDTH + 1
		}

		// append row to grid and reset it
		Matrix = append(Matrix, row)
		row = []*Box{}

		x = origin.X //reset position
		y += BOX_HEIGHT + 1
	}

	renderWinner = make(chan Player)
	go func() {
		for p := range renderWinner {
			if p.Color == RED {
				player1 := NewPlayer("RED.png", RED)
				player1.SetPosition(20, 10)
				level.AddEntity(&player1)
			} else {
				player2 := NewPlayer("BLUE.png", BLUE)
				player2.SetPosition(115, 10)
				level.AddEntity(&player2)
			}
		}
	}()

	// Render Red/Blue Ball
	renderBall = make(chan Ball)
	go func() {
		for b := range renderBall {
			tmp := NewBall(b.Color, b.X, b.Y)
			tmp.SetPosition(tmp.X, tmp.Y)
			level.AddEntity(&tmp)
		}
	}()

	game.Screen().SetLevel(level)
	game.Start()
}
