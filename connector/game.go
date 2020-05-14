package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

// Which color? (Red or Blue)
// What move? (mark or eliminate)
// Is Game over?

type Game struct {
	prevBox *Box
}

const (
	MARK      bool = true
	ELIMINATE bool = false
)

var player int = RED
var move bool = MARK // MARK or ELIMINATE

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

// Return RED or BLUE
func (g *Game) GetPlayer() int {
	return player
}

func (g *Game) GetMove() bool {
	return move
}

// toggle the move (internal function)
func (g *Game) NextMove() bool {
	move = !move

	// Toggle player for next mark
	if move == MARK && player == RED {
		player = BLUE
	} else if move == MARK && player == BLUE {
		player = RED
	}
	return move
}

func (g *Game) ValidateMove(box *Box) bool {
	if move == MARK {
		// TODO: Add a validation that you MUST have at least 1 empty neighbour to mark the square!
		if !hasEmptyNeighbour(box.Row, box.Col) {
			return false
		}

		//  Mark and save
		g.prevBox = box
		return true
	}

	// If next move is eliminate, then we need to verify that box is a neighbour
	x, y := g.prevBox.Position()
	newX, newY := box.Position()

	// check if box is a neighbour
	if (newX >= (x-BOX_WIDTH-1) && newX <= (x+BOX_WIDTH+1)) &&
		(newY >= (y-BOX_HEIGHT-1) && newY <= (y+BOX_HEIGHT+1)) {
		return true
	}

	// this is not a neighbour
	return false
}

func hasEmptyNeighbour(r, c int) bool {
	// box with all 8 neighbours
	if (c+1 < GRID && Matrix[r][c+1].Status == EMPTY) || // right
		(c-1 >= 0 && Matrix[r][c-1].Status == EMPTY) || // left
		(r-1 >= 0 && Matrix[r-1][c].Status == EMPTY) || // top
		(r+1 < GRID && Matrix[r+1][c].Status == EMPTY) || // bottom
		(r+1 < GRID && c+1 < GRID && Matrix[r+1][c+1].Status == EMPTY) || // bottom-right corner
		(r+1 < GRID && c-1 >= 0 && Matrix[r+1][c-1].Status == EMPTY) || // top-right corner
		(r-1 >= 0 && c-1 >= 0 && Matrix[r-1][c-1].Status == EMPTY) || // top-left corner
		(r-1 >= 0 && c+1 < GRID && Matrix[r-1][c+1].Status == EMPTY) { // bottom-left corner
		return true
	}

	// has no empty neighbour
	return false
}

func getKey(r, c int) string {
	return fmt.Sprintf("%d%d", r, c)
}

func checkRightBox(r, c int, status int) bool {
	if c+1 < GRID && Matrix[r][c+1].Status == status {
		return true
	}
	return false
}

func depth_first(r, c int, chain *[]*Box, traversed map[string]bool, player int) (ok bool) {
	//fmt.Printf("ENTRY: rc: %d%d, len(chain): %d, len(traversed): %d, player: %d\n", r, c, len(chain), len(traversed), player)
	log.WithFields(log.Fields{
		"r":              r,
		"c":              c,
		"len(chain)":     len(*chain),
		"len(traversed)": len(traversed),
		"player":         player,
	}).Info("ENTRY")
	if c+1 < GRID && Matrix[r][c+1].Status == player && !traversed[getKey(r, c+1)] { // right
		*chain = append(*chain, Matrix[r][c+1])
		traversed[getKey(r, c+1)] = true
		return depth_first(r, c+1, chain, traversed, player)
	} else if c-1 >= 0 && Matrix[r][c-1].Status == player && !traversed[getKey(r, c-1)] { // left
		*chain = append(*chain, Matrix[r][c-1])
		traversed[getKey(r, c-1)] = true
		return depth_first(r, c-1, chain, traversed, player)
	} else if r-1 >= 0 && Matrix[r-1][c].Status == player && !traversed[getKey(r-1, c)] { // top
		*chain = append(*chain, Matrix[r-1][c])
		traversed[getKey(r-1, c)] = true
		return depth_first(r-1, c, chain, traversed, player)
	} else if r+1 < GRID && Matrix[r+1][c].Status == player && !traversed[getKey(r+1, c)] { // bottom
		*chain = append(*chain, Matrix[r+1][c])
		traversed[getKey(r+1, c)] = true
		return depth_first(r+1, c, chain, traversed, player)
	} else if r+1 < GRID && c+1 < GRID && Matrix[r+1][c+1].Status == player && !traversed[getKey(r+1, c+1)] { // bottom-right corner
		*chain = append(*chain, Matrix[r+1][c+1])
		traversed[getKey(r+1, c+1)] = true
		return depth_first(r+1, c+1, chain, traversed, player)
	} else if r+1 < GRID && c-1 >= 0 && Matrix[r+1][c-1].Status == player && !traversed[getKey(r+1, c-1)] { // top-right corner
		*chain = append(*chain, Matrix[r+1][c-1])
		traversed[getKey(r+1, c-1)] = true
		return depth_first(r+1, c-1, chain, traversed, player)
	} else if r-1 >= 0 && c-1 >= 0 && Matrix[r-1][c-1].Status == player && !traversed[getKey(r-1, c-1)] { // top-left corner
		*chain = append(*chain, Matrix[r-1][c-1])
		traversed[getKey(r-1, c-1)] = true
		return depth_first(r-1, c-1, chain, traversed, player)
	} else if r-1 >= 0 && c+1 < GRID && Matrix[r-1][c+1].Status == player && !traversed[getKey(r-1, c+1)] { // bottom-left corner
		*chain = append(*chain, Matrix[r-1][c+1])
		traversed[getKey(r-1, c+1)] = true
		return depth_first(r-1, c+1, chain, traversed, player)
	}

	// no where else to go!
	//fmt.Printf("EXIT: rc: %d%d, len(chain): %d, len(traversed): %d, player: %d\n", r, c, len(chain), len(traversed), player)
	log.WithFields(log.Fields{
		"r":              r,
		"c":              c,
		"len(chain)":     len(*chain),
		"len(traversed)": len(traversed),
		"player":         player,
	}).Info("EXIT")
	return true
}

func (g *Game) EndGame() (ok bool, winner []*Box) {
	// Iterate through the Matrix and implement a Depth-First search to find the longest chain!

	// If there are 2 EMPTY neighbours, game is not over.
	for r := 0; r < GRID; r++ {
		for c := 0; c < GRID; c++ {
			box := Matrix[r][c]
			if box.Status == EMPTY {
				// Get empty neighbours
				if hasEmptyNeighbour(r, c) {
					// game is not over yet, return!
					ok = false
					return
				}
			}
		}
	}

	log.Warn("Starting .. brute force check")
	// Use brute force and move the root from 00 till 66 to find longest chain
	for r := 0; r < GRID; r++ {
		for c := 0; c < GRID; c++ {
			// root for the chain should always be for a player!
			if Matrix[r][c].Status == RED || Matrix[r][c].Status == BLUE {
				traversed := map[string]bool{} // 11 => true means box with row-1, col-1 has been traversed!
				player := Matrix[r][c].Status  // RED or BLUE
				traversed[getKey(r, c)] = true
				chain := []*Box{Matrix[r][c]}

				ok = depth_first(r, c, &chain, traversed, player)

				// TODO: How do we handle a Draw? o_O
				// If the chain is longer than "winner", update winner
				if ok && len(chain) > len(winner) {
					log.Info("Updating winner chain", "FROM: ", len(winner), "TO: ", len(chain))
					winner = chain
				}
			}

		}
	}

	log.Info("WINNING CHAIN...")
	for i, b := range winner {
		log.WithFields(log.Fields{
			"i":   i,
			"Row": b.Row,
			"Col": b.Col,
		}).Info("..")
	}
	return
}
