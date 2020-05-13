package main

import (
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/mock"
	"testing"
)

func TestSomething(t *testing.T) {
	assert.True(t, true, "True is true!")
}

/*
type MockHangman struct {
	mock.Mock
}

func (h *MockHangman) Render_game(p []string, e map[string]bool) error {
	args := h.Called()
	return args.Error(0)
}

func (h *MockHangman) Get_entry() string {
	args := h.Called()
	return args.String(0)
}

func TestPlayWithWin(t *testing.T) {
	game := new(MockHangman)

	game.On("Render_game").Return(nil)
	game.On("Get_entry").Return("e")

	play(game, "elephant")
}

*/
/*
func TestPlayWithLoss(t *testing.T) {
}

func TestPlayWinWithWord(t *testing.T) {
}
*/
