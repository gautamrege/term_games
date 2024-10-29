package main

import (
	"testing"
)

// TestOpponent tests the opponent function.
func TestOpponent(t *testing.T) {
	game := NewGame()

	// Test when it's BLACK's turn
	game.turn = BLACK
	if game.opponent() != WHITE {
		t.Errorf("Expected WHITE, got %d", game.opponent())
	}

	// Test when it's WHITE's turn
	game.turn = WHITE
	if game.opponent() != BLACK {
		t.Errorf("Expected BLACK, got %d", game.opponent())
	}
}

// TestApplyMove tests the ApplyMove function.
func TestApplyMove(t *testing.T) {
	game := NewGame()

	// Apply a valid move for BLACK
	game.ApplyMove(2, 3)
	if game.board[2][3] != BLACK {
		t.Errorf("Expected BLACK at (2, 3), got %d", game.board[2][3])
	}
	if game.board[3][3] != BLACK {
		t.Errorf("Expected BLACK at (3, 3), got %d", game.board[3][3])
	}

	// Apply a valid move for WHITE
	game.ApplyMove(2, 2)
	if game.board[2][2] != WHITE {
		t.Errorf("Expected WHITE at (2, 2), got %d", game.board[2][2])
	}
	if game.board[3][3] != WHITE {
		t.Errorf("Expected WHITE at (3, 3), got %d", game.board[3][3])
	}
}

// TestFlip tests the flip function.
func TestFlip(t *testing.T) {
	game := NewGame()

	// Apply a move that will flip pieces
	game.ApplyMove(2, 3)
	game.flip(2, 3, 1, 0)
	if game.board[3][3] != BLACK {
		t.Errorf("Expected BLACK at (3, 3), got %d", game.board[3][3])
	}
}

// TestHasValidMove tests the HasValidMove function.
func TestHasValidMove(t *testing.T) {
	game := NewGame()

	// Initially, there should be valid moves
	if !game.HasValidMove() {
		t.Errorf("Expected valid moves, but got none")
	}

	// Fill the board to simulate no valid moves
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			game.board[i][j] = BLACK
		}
	}
	if game.HasValidMove() {
		t.Errorf("Expected no valid moves, but got some")
	}
}
