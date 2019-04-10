package tetrimino

import (
	"math/rand"
	"time"

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
	// primarily used for testing
	pieceOrientation() orientation
}

// Coordinates represent a blocks position on the board
type Coordinates struct {
	X int
	Y int
}

type pieceConstructor func(boardWidth, boardHeight int) Tetrimino

var pieceConstructors = []pieceConstructor{newIPiece, newJPiece, newLPiece, newOPiece, newSPiece, newTPiece, newZPiece}

// New generates a new tetrimino
func New(boardWidth, boardHeight int) Tetrimino {
	var (
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
		i = r.Intn(len(pieceConstructors))
	)

	return pieceConstructors[i](boardWidth, boardHeight)
}

// Box represents the box surrounding the current piece
// this way we don't have to track the coordinates of each block
type Box struct {
	TopLeft     Coordinates
	BottomRight Coordinates
}

func findMaxY(blocks [][]*board.Block, boxBottomRight Coordinates) Coordinates {
	var (
		yMax = boxBottomRight
	)
	for i, row := range blocks {
		for j, block := range row {
			if block == nil {
				continue
			}
			y := boxBottomRight.Y + (len(blocks) - 1) - i
			x := boxBottomRight.X - (len(row) - 1) + j

			if y > yMax.Y {
				yMax = Coordinates{
					X: x,
					Y: y,
				}
			}
		}
	}

	return yMax
}

func findMinY(blocks [][]*board.Block, boxTopLeft Coordinates) Coordinates {
	var (
		yMin = boxTopLeft
	)

	for i, row := range blocks {
		for j, block := range row {
			if block == nil {
				continue
			}
			y := boxTopLeft.Y - i
			x := boxTopLeft.X + j

			if y < yMin.Y {
				yMin = Coordinates{
					X: x,
					Y: y,
				}
			}
		}
	}

	return yMin
}

func findMaxX(blocks [][]*board.Block, boxTopLeft Coordinates) Coordinates {
	var (
		xMax = boxTopLeft
	)

	for i, row := range blocks {
		for j, block := range row {
			if block == nil {
				continue
			}
			y := boxTopLeft.Y - i
			x := boxTopLeft.X + j

			if x > xMax.X {
				xMax = Coordinates{
					X: x,
					Y: y,
				}
			}
		}
	}

	return xMax
}

func findMinX(blocks [][]*board.Block, boxBottomRight Coordinates) Coordinates {
	var (
		xMin = boxBottomRight
	)

	for i, row := range blocks {
		for j, block := range row {
			if block == nil {
				continue
			}
			y := boxBottomRight.Y + (len(blocks) - 1) - i
			x := boxBottomRight.X - (len(row) - 1) + j

			if x < xMin.X {
				xMin = Coordinates{
					X: x,
					Y: y,
				}
			}
		}
	}

	return xMin
}

type orientation int

const (
	spawn orientation = iota
	clockwise
	opposite
	counterclockwise
)

func (o *orientation) String() string {
	switch *o {
	case spawn:
		return "spawn"
	case clockwise:
		return "clockwise"
	case opposite:
		return "opposite"
	case counterclockwise:
		return "counterclockwise"
	default:
		return ""
	}
}

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
