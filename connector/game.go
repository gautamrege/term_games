package main

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
