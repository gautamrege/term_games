package main

import (
	tl "github.com/JoelOtter/termloop"
)

const (
	EMPTY int = iota
	RED
	BLUE
	ELIMINATED
)

type Box struct {
	*tl.Rectangle
	Status int
	Row    int
	Col    int
}

func (box *Box) validateClick(ev tl.Event) bool {
	// check if box is empty
	if box.Status != EMPTY {
		return false
	}

	// check the game play
	return GAME.ValidateMove(box)
}

func (box *Box) Tick(ev tl.Event) {
	if ev.Type != tl.EventMouse {
		return
	}

	x, y := box.Position()
	switch ev.Key {
	case tl.MouseLeft:
		if ev.MouseX > x && ev.MouseX < (x+BOX_WIDTH) && ev.MouseY < (y+BOX_HEIGHT) && ev.MouseY > y {
			if box.validateClick(ev) {
				// paint the box!
				move := GAME.GetMove()
				if move == ELIMINATE {
					box.SetColor(tl.ColorBlack)
					box.Status = ELIMINATED
				} else {
					player := GAME.GetPlayer()
					if player == RED {
						//box.SetColor(tl.ColorRed)
						renderBall <- Ball{nil, RED, x, y}
					} else {
						//box.SetColor(tl.ColorBlue)
						renderBall <- Ball{nil, BLUE, x, y}
					}
					box.Status = player
				}
				GAME.NextMove()
				end, winner := GAME.EndGame()
				if end {
					renderWinner <- Player{
						Color: winner[0].Status,
					}
				}
			}
		}
	}
}

func NewBox(x, y, row, column int) Box {
	return Box{tl.NewRectangle(x, y, BOX_WIDTH, BOX_HEIGHT, tl.ColorGreen), EMPTY, row, column}
}
