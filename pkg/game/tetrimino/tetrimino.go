package tetrimino

import (
	"github.com/ShawnROGrady/gotris/pkg/game/board"
)

// Tetrimino represents an active game piece
type Tetrimino interface {
	Blocks() [][]*board.Block
	MoveUp(ymax int)
	MoveDown()
	MoveLeft()
	MoveRight(xmax int)
	RotateClockwise()
	RotateCounter()
	ContainingBox() Box
	// used for detecting collisions
	YMax() Coordinates
	YMin() Coordinates
	XMax() Coordinates
	XMin() Coordinates
}

// Coordinates represent a blocks position on the board
type Coordinates struct {
	X int
	Y int
}

// New generates a new tetrimino
func New(boardWidth, boardHeight int) Tetrimino {
	// TODO: should randomly generate piece type
	spawnOrientation := spawn
	return &iPiece{
		orientation: &spawnOrientation,
		box: Box{
			TopLeft: Coordinates{
				X: 0,
				Y: boardHeight - 1,
			},
			BottomRight: Coordinates{
				X: 3,
				Y: boardHeight - 4,
			},
		},
	}
}

// Box represents the box surrounding the current piece
// this way we don't have to track the coordinates of each block
type Box struct {
	TopLeft     Coordinates
	BottomRight Coordinates
}

type orientation int

const (
	spawn orientation = iota
	clockwise
	opposite
	counterclockwise
)

func (o *orientation) rotateClockwise() {
	switch *o {
	case spawn:
		*o = clockwise
	case clockwise:
		*o = opposite
	case opposite:
		*o = counterclockwise
	case counterclockwise:
		*o = spawn
	}
}

func (o *orientation) rotateCounter() {
	switch *o {
	case spawn:
		*o = counterclockwise
	case clockwise:
		*o = spawn
	case opposite:
		*o = clockwise
	case counterclockwise:
		*o = opposite
	}
}
