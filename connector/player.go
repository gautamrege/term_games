package main

import tl "github.com/JoelOtter/termloop"

type Player struct {
	*tl.Entity
	Color int
	Img   string
}

func NewPlayer(img string, color int) Player {
	p := Player{}
	p.Img = img
	p.Color = color
	p.Entity = tl.NewEntity(1, 1, 24, 24)
	p.Entity.ApplyCanvas(tl.BackgroundCanvasFromFile(img))

	return p
}
