package main

import (
	"fmt"
)

const (
	EMPTY = 0
	BLACK = 1
	WHITE = 2
)

var directions = [8][2]int{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
}

type Game struct {
	board [8][8]int
	turn  int
}

func NewGame() *Game {
	g := &Game{turn: BLACK}
	g.board[3][3], g.board[4][4] = WHITE, WHITE
	g.board[3][4], g.board[4][3] = BLACK, BLACK
	return g
}

func (g *Game) DisplayBoard() {
	fmt.Println("  a b c d e f g h")
	for i := 0; i < 8; i++ {
		fmt.Printf("%d ", i+1)
		for j := 0; j < 8; j++ {
			switch g.board[i][j] {
			case EMPTY:
				fmt.Print(". ")
			case BLACK:
				fmt.Print("B ")
			case WHITE:
				fmt.Print("W ")
			}
		}
		fmt.Println()
	}
}

func (g *Game) IsValidMove(x, y int) bool {
	if g.board[x][y] != EMPTY {
		return false
	}
	for _, d := range directions {
		if g.canFlip(x, y, d[0], d[1]) {
			return true
		}
	}
	return false
}

func (g *Game) canFlip(x, y, dx, dy int) bool {
	x += dx
	y += dy
	count := 0
	for x >= 0 && x < 8 && y >= 0 && y < 8 && g.board[x][y] == g.opponent() {
		x += dx
		y += dy
		count++
	}
	if count > 0 && x >= 0 && x < 8 && y >= 0 && y < 8 && g.board[x][y] == g.turn {
		return true
	}
	return false
}

func (g *Game) opponent() int {
	if g.turn == BLACK {
		return WHITE
	}
	return BLACK
}

func (g *Game) ApplyMove(x, y int) {
	g.board[x][y] = g.turn
	for _, d := range directions {
		if g.canFlip(x, y, d[0], d[1]) {
			g.flip(x, y, d[0], d[1])
		}
	}
	g.turn = g.opponent()
}

func (g *Game) flip(x, y, dx, dy int) {
	x += dx
	y += dy
	for g.board[x][y] == g.opponent() {
		g.board[x][y] = g.turn
		x += dx
		y += dy
	}
}

func (g *Game) HasValidMove() bool {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if g.IsValidMove(i, j) {
				return true
			}
		}
	}
	return false
}

func (g *Game) CountPieces() (black, white int) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			switch g.board[i][j] {
			case BLACK:
				black++
			case WHITE:
				white++
			}
		}
	}
	return
}

func main() {
	game := NewGame()
	for {
		game.DisplayBoard()
		if !game.HasValidMove() {
			break
		}
		var move string
		fmt.Printf("Player %d, enter your move (e.g., d3): ", game.turn)
		fmt.Scan(&move)
		x, y := int(move[1]-'1'), int(move[0]-'a')
		if game.IsValidMove(x, y) {
			game.ApplyMove(x, y)
		} else {
			fmt.Println("Invalid move. Try again.")
		}
	}
	black, white := game.CountPieces()
	fmt.Printf("Game over! Black: %d, White: %d\n", black, white)
	if black > white {
		fmt.Println("Black wins!")
	} else if white > black {
		fmt.Println("White wins!")
	} else {
		fmt.Println("It's a tie!")
	}
}
