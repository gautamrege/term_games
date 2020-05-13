package main

import (
	tl "github.com/JoelOtter/termloop"
)

const (
	BOX_WIDTH  int = 5
	BOX_HEIGHT int = 3
	GRID       int = 6
)

type Location struct {
	X int
	Y int
}

var GAME Game = Game{}

func main() {
	game := tl.NewGame()
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorCyan,
	})

	// Where do we start drawing the grid from?
	origin := Location{75, 10}

	x := origin.X
	y := origin.Y

	for i := 1; i <= GRID*GRID; i++ {
		box := NewBox(x, y)
		level.AddEntity(&box)

		x += BOX_WIDTH + 1
		if i%GRID == 0 { // row is complete
			x = origin.X //reset position
			y += BOX_HEIGHT + 1
		}
	}

	player1 := NewPlayer("player6.png", "Gautam")
	player1.SetPosition(20, 5)
	level.AddEntity(&player1)

	player2 := NewPlayer("player7.png", "Anuj")
	player2.SetPosition(125, 5)
	level.AddEntity(&player2)

	game.Screen().SetLevel(level)
	game.Start()
}
