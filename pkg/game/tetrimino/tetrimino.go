package tetrimino

import (
	"fmt"
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
	previousOrientation() orientation
	// for resolving rotation conflicts
	RotationTests() []RotationTest
}

// Coordinates represent a blocks position on the board
type Coordinates struct {
	X int
	Y int
}

// PieceConstructor represents the function signature of each pieces constructor
// Currently only exporting this to test game logic
// TODO: figure out better way to enable testing
type PieceConstructor func(boardWidth, boardHeight int) Tetrimino

// PieceConstructors represent all constructors for all possible pieces
// Currently only exporting this to test game logic
// TODO: figure out better way to enable testing
var PieceConstructors = []PieceConstructor{newIPiece, newJPiece, newLPiece, newOPiece, newSPiece, newTPiece, newZPiece}

// New generates a new tetrimino
func New(boardWidth, boardHeight int) Tetrimino {
	var (
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
		i = r.Intn(len(PieceConstructors))
	)

	return PieceConstructors[i](boardWidth, boardHeight)
}

// Box represents the box surrounding the current piece
// this way we don't have to track the coordinates of each block
type Box struct {
	TopLeft     Coordinates
	BottomRight Coordinates
}

func startingBox(boardWidth, boardHeight int, piece Tetrimino) Box {
	midpoint := (boardWidth / 2) - 1
	switch piece.(type) {
	case *iPiece:
		return Box{
			TopLeft: Coordinates{
				X: midpoint - 1,
				Y: boardHeight - 1,
			},
			BottomRight: Coordinates{
				X: midpoint + 2,
				Y: boardHeight - 4,
			},
		}
	case *oPiece:
		return Box{
			TopLeft: Coordinates{
				X: midpoint - 1,
				Y: boardHeight - 1,
			},
			BottomRight: Coordinates{
				X: midpoint + 2,
				Y: boardHeight - 3,
			},
		}
	default:
		return Box{
			TopLeft: Coordinates{
				X: midpoint - 1,
				Y: boardHeight - 1,
			},
			BottomRight: Coordinates{
				X: midpoint + 1,
				Y: boardHeight - 3,
			},
		}
	}
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

// RotationTest is used to attempt to resolve rotation conflicts then revert those changes on failure
type RotationTest struct {
	ApplyTest  func(xmax, ymax int)
	RevertTest func(xmax, ymax int)
}

func defaultRotationTests(t Tetrimino, prevOrientation, newOrientation orientation) []RotationTest {
	switch prevOrientation {
	case spawn:
		switch newOrientation {
		case clockwise:
			return []RotationTest{
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveLeft() },
					RevertTest: func(xmax, ymax int) { t.MoveRight(xmax) },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveLeft(); t.MoveDown() },
					RevertTest: func(xmax, ymax int) { t.MoveRight(xmax); t.MoveUp(ymax) },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveDown(); t.MoveDown() },
					RevertTest: func(xmax, ymax int) { t.MoveUp(ymax); t.MoveUp(ymax) },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveLeft(); t.MoveDown(); t.MoveDown() },
					RevertTest: func(xmax, ymax int) { t.MoveRight(xmax); t.MoveUp(ymax); t.MoveUp(ymax) },
				},
			}
		case counterclockwise:
			return []RotationTest{
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveRight(xmax) },
					RevertTest: func(xmax, ymax int) { t.MoveLeft() },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveRight(xmax); t.MoveUp(ymax) },
					RevertTest: func(xmax, ymax int) { t.MoveLeft(); t.MoveDown() },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveDown(); t.MoveDown() },
					RevertTest: func(xmax, ymax int) { t.MoveUp(ymax); t.MoveUp(ymax) },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveRight(xmax); t.MoveDown(); t.MoveDown() },
					RevertTest: func(xmax, ymax int) { t.MoveLeft(); t.MoveUp(ymax); t.MoveUp(ymax) },
				},
			}
		}
	case clockwise:
		switch newOrientation {
		case spawn:
			return []RotationTest{
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveRight(xmax) },
					RevertTest: func(xmax, ymax int) { t.MoveLeft() },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveRight(xmax); t.MoveDown() },
					RevertTest: func(xmax, ymax int) { t.MoveLeft(); t.MoveUp(ymax) },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveUp(ymax); t.MoveUp(ymax) },
					RevertTest: func(xmax, ymax int) { t.MoveDown(); t.MoveDown() },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveRight(xmax); t.MoveUp(ymax); t.MoveUp(ymax) },
					RevertTest: func(xmax, ymax int) { t.MoveLeft(); t.MoveDown(); t.MoveDown() },
				},
			}
		case opposite:
			return []RotationTest{
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveRight(xmax) },
					RevertTest: func(xmax, ymax int) { t.MoveLeft() },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveRight(xmax); t.MoveDown() },
					RevertTest: func(xmax, ymax int) { t.MoveLeft(); t.MoveUp(ymax) },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveUp(ymax); t.MoveUp(ymax) },
					RevertTest: func(xmax, ymax int) { t.MoveDown(); t.MoveDown() },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveRight(xmax); t.MoveUp(ymax); t.MoveUp(ymax) },
					RevertTest: func(xmax, ymax int) { t.MoveLeft(); t.MoveDown(); t.MoveDown() },
				},
			}
		}
	case opposite:
		switch newOrientation {
		case clockwise:
			return []RotationTest{
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveLeft() },
					RevertTest: func(xmax, ymax int) { t.MoveRight(xmax) },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveLeft(); t.MoveUp(ymax) },
					RevertTest: func(xmax, ymax int) { t.MoveRight(xmax); t.MoveDown() },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveDown(); t.MoveDown() },
					RevertTest: func(xmax, ymax int) { t.MoveUp(ymax); t.MoveUp(ymax) },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveLeft(); t.MoveDown(); t.MoveDown() },
					RevertTest: func(xmax, ymax int) { t.MoveRight(xmax); t.MoveUp(ymax); t.MoveUp(ymax) },
				},
			}
		case counterclockwise:
			return []RotationTest{
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveRight(xmax) },
					RevertTest: func(xmax, ymax int) { t.MoveLeft() },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveRight(xmax); t.MoveUp(ymax) },
					RevertTest: func(xmax, ymax int) { t.MoveLeft(); t.MoveDown() },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveDown(); t.MoveDown() },
					RevertTest: func(xmax, ymax int) { t.MoveUp(ymax); t.MoveUp(ymax) },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveRight(xmax); t.MoveDown(); t.MoveDown() },
					RevertTest: func(xmax, ymax int) { t.MoveLeft(); t.MoveUp(ymax); t.MoveUp(ymax) },
				},
			}
		}
	case counterclockwise:
		switch newOrientation {
		case opposite:
			return []RotationTest{
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveLeft() },
					RevertTest: func(xmax, ymax int) { t.MoveRight(xmax) },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveLeft(); t.MoveDown() },
					RevertTest: func(xmax, ymax int) { t.MoveRight(xmax); t.MoveUp(ymax) },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveUp(ymax); t.MoveUp(ymax) },
					RevertTest: func(xmax, ymax int) { t.MoveDown(); t.MoveDown() },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveLeft(); t.MoveUp(ymax); t.MoveUp(ymax) },
					RevertTest: func(xmax, ymax int) { t.MoveRight(xmax); t.MoveDown(); t.MoveDown() },
				},
			}
		case spawn:
			return []RotationTest{
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveLeft() },
					RevertTest: func(xmax, ymax int) { t.MoveRight(xmax) },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveLeft(); t.MoveDown() },
					RevertTest: func(xmax, ymax int) { t.MoveRight(xmax); t.MoveUp(ymax) },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveUp(ymax); t.MoveUp(ymax) },
					RevertTest: func(xmax, ymax int) { t.MoveDown(); t.MoveDown() },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveLeft(); t.MoveUp(ymax); t.MoveUp(ymax) },
					RevertTest: func(xmax, ymax int) { t.MoveRight(xmax); t.MoveDown(); t.MoveDown() },
				},
			}
		}
	}
	fmt.Printf("Unhandled orientation combo (prev, new) = (%s, %s)\n", &prevOrientation, &newOrientation)
	return []RotationTest{}
}
