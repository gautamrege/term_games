package main

import (
	"fmt"
	tl "github.com/JoelOtter/termloop"
)

type ScoreLabel struct {
	*tl.Text
}

func (s *ScoreLabel) Tick(e tl.Event) {
	// nothing to do here
}

func (s *ScoreLabel) Draw(screen *tl.Screen) {
	s.SetText(fmt.Sprintf("Score: %d", Score))
	s.Text.Draw(screen)
}

func NewScoreLabel() ScoreLabel {
	return ScoreLabel{tl.NewText(PLAY_WIDTH+1, 6, "Score: 0", tl.ColorWhite, tl.ColorBlack)}
}
