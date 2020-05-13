package main

import tl "github.com/JoelOtter/termloop"

type Player struct {
	*tl.Entity
	Name string
	Img  string
}

func NewPlayer(img, name string) Player {
	p := Player{}
	p.Img = img
	p.Name = name
	p.Entity = tl.NewEntity(1, 1, 48, 40)
	p.Entity.ApplyCanvas(tl.BackgroundCanvasFromFile(img))

	return p
}
