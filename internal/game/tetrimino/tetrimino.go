package tetrimino

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ShawnROGrady/gotris/internal/canvas"
	"github.com/ShawnROGrady/gotris/internal/game/board"
)

// the maximum size of a tetrimino
const (
	MaxWidth  = 4
	MaxHeight = 4
)

// Tetrimino represents an active game piece
type Tetrimino interface {
	Blocks() [][]*board.Block
	MoveUp()
	MoveDown()
	MoveLeft()
	MoveRight()
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
	SpawnGhost() Tetrimino
	ToggleGhost()
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

// NewSet generates a new set of tetriminos
// this set is a random permutation of all piece types: https://harddrop.com/wiki/Random_Generator
func NewSet(boardWidth, boardHeight int) []Tetrimino {
	var (
		r        = rand.New(rand.NewSource(time.Now().UnixNano()))
		perm     = r.Perm(len(PieceConstructors))
		pieceSet = []Tetrimino{}
	)

	for i := range perm {
		pieceSet = append(pieceSet, PieceConstructors[perm[i]](boardWidth, boardHeight))
	}

	return pieceSet
}

type tetriminoBase struct {
	box             Box
	orientation     *orientation
	prevOrientation orientation
	color           canvas.Color
	isGhost         bool
}

func (t *tetriminoBase) pieceOrientation() orientation {
	return *t.orientation
}

func (t *tetriminoBase) previousOrientation() orientation {
	return t.prevOrientation
}

func (t *tetriminoBase) ContainingBox() Box {
	return t.box
}

func (t *tetriminoBase) MoveUp() {
	t.box.TopLeft.Y++
	t.box.BottomRight.Y++
}

func (t *tetriminoBase) MoveDown() {
	t.box.BottomRight.Y--
	t.box.TopLeft.Y--
}

func (t *tetriminoBase) MoveLeft() {
	t.box.TopLeft.X--
	t.box.BottomRight.X--
}

func (t *tetriminoBase) MoveRight() {
	t.box.BottomRight.X++
	t.box.TopLeft.X++
}

func (t *tetriminoBase) RotateClockwise() {
	t.prevOrientation = *t.orientation
	t.orientation.rotateClockwise()
}

func (t *tetriminoBase) RotateCounter() {
	t.prevOrientation = *t.orientation
	t.orientation.rotateCounter()
}

func (t *tetriminoBase) ToggleGhost() {
	t.isGhost = !t.isGhost
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
	ApplyTest  func()
	RevertTest func()
}

// https://harddrop.com/wiki/SRS#Wall_Kicks
func defaultRotationTests(t Tetrimino, prevOrientation, newOrientation orientation) []RotationTest {
	switch prevOrientation {
	case spawn:
		switch newOrientation {
		case clockwise:
			return []RotationTest{
				{
					ApplyTest:  func() { t.MoveLeft() },
					RevertTest: func() { t.MoveRight() },
				},
				{
					ApplyTest:  func() { t.MoveLeft(); t.MoveUp() },
					RevertTest: func() { t.MoveRight(); t.MoveDown() },
				},
				{
					ApplyTest:  func() { t.MoveDown(); t.MoveDown() },
					RevertTest: func() { t.MoveUp(); t.MoveUp() },
				},
				{
					ApplyTest:  func() { t.MoveLeft(); t.MoveDown(); t.MoveDown() },
					RevertTest: func() { t.MoveRight(); t.MoveUp(); t.MoveUp() },
				},
			}
		case counterclockwise:
			return []RotationTest{
				{
					ApplyTest:  func() { t.MoveRight() },
					RevertTest: func() { t.MoveLeft() },
				},
				{
					ApplyTest:  func() { t.MoveRight(); t.MoveUp() },
					RevertTest: func() { t.MoveLeft(); t.MoveDown() },
				},
				{
					ApplyTest:  func() { t.MoveDown(); t.MoveDown() },
					RevertTest: func() { t.MoveUp(); t.MoveUp() },
				},
				{
					ApplyTest:  func() { t.MoveRight(); t.MoveDown(); t.MoveDown() },
					RevertTest: func() { t.MoveLeft(); t.MoveUp(); t.MoveUp() },
				},
			}
		}
	case clockwise:
		switch newOrientation {
		case spawn:
			return []RotationTest{
				{
					ApplyTest:  func() { t.MoveRight() },
					RevertTest: func() { t.MoveLeft() },
				},
				{
					ApplyTest:  func() { t.MoveRight(); t.MoveDown() },
					RevertTest: func() { t.MoveLeft(); t.MoveUp() },
				},
				{
					ApplyTest:  func() { t.MoveUp(); t.MoveUp() },
					RevertTest: func() { t.MoveDown(); t.MoveDown() },
				},
				{
					ApplyTest:  func() { t.MoveRight(); t.MoveUp(); t.MoveUp() },
					RevertTest: func() { t.MoveLeft(); t.MoveDown(); t.MoveDown() },
				},
			}
		case opposite:
			return []RotationTest{
				{
					ApplyTest:  func() { t.MoveRight() },
					RevertTest: func() { t.MoveLeft() },
				},
				{
					ApplyTest:  func() { t.MoveRight(); t.MoveDown() },
					RevertTest: func() { t.MoveLeft(); t.MoveUp() },
				},
				{
					ApplyTest:  func() { t.MoveUp(); t.MoveUp() },
					RevertTest: func() { t.MoveDown(); t.MoveDown() },
				},
				{
					ApplyTest:  func() { t.MoveRight(); t.MoveUp(); t.MoveUp() },
					RevertTest: func() { t.MoveLeft(); t.MoveDown(); t.MoveDown() },
				},
			}
		}
	case opposite:
		switch newOrientation {
		case clockwise:
			return []RotationTest{
				{
					ApplyTest:  func() { t.MoveLeft() },
					RevertTest: func() { t.MoveRight() },
				},
				{
					ApplyTest:  func() { t.MoveLeft(); t.MoveUp() },
					RevertTest: func() { t.MoveRight(); t.MoveDown() },
				},
				{
					ApplyTest:  func() { t.MoveDown(); t.MoveDown() },
					RevertTest: func() { t.MoveUp(); t.MoveUp() },
				},
				{
					ApplyTest:  func() { t.MoveLeft(); t.MoveDown(); t.MoveDown() },
					RevertTest: func() { t.MoveRight(); t.MoveUp(); t.MoveUp() },
				},
			}
		case counterclockwise:
			return []RotationTest{
				{
					ApplyTest:  func() { t.MoveRight() },
					RevertTest: func() { t.MoveLeft() },
				},
				{
					ApplyTest:  func() { t.MoveRight(); t.MoveUp() },
					RevertTest: func() { t.MoveLeft(); t.MoveDown() },
				},
				{
					ApplyTest:  func() { t.MoveDown(); t.MoveDown() },
					RevertTest: func() { t.MoveUp(); t.MoveUp() },
				},
				{
					ApplyTest:  func() { t.MoveRight(); t.MoveDown(); t.MoveDown() },
					RevertTest: func() { t.MoveLeft(); t.MoveUp(); t.MoveUp() },
				},
			}
		}
	case counterclockwise:
		switch newOrientation {
		case opposite:
			return []RotationTest{
				{
					ApplyTest:  func() { t.MoveLeft() },
					RevertTest: func() { t.MoveRight() },
				},
				{
					ApplyTest:  func() { t.MoveLeft(); t.MoveDown() },
					RevertTest: func() { t.MoveRight(); t.MoveUp() },
				},
				{
					ApplyTest:  func() { t.MoveUp(); t.MoveUp() },
					RevertTest: func() { t.MoveDown(); t.MoveDown() },
				},
				{
					ApplyTest:  func() { t.MoveLeft(); t.MoveUp(); t.MoveUp() },
					RevertTest: func() { t.MoveRight(); t.MoveDown(); t.MoveDown() },
				},
			}
		case spawn:
			return []RotationTest{
				{
					ApplyTest:  func() { t.MoveLeft() },
					RevertTest: func() { t.MoveRight() },
				},
				{
					ApplyTest:  func() { t.MoveLeft(); t.MoveDown() },
					RevertTest: func() { t.MoveRight(); t.MoveUp() },
				},
				{
					ApplyTest:  func() { t.MoveUp(); t.MoveUp() },
					RevertTest: func() { t.MoveDown(); t.MoveDown() },
				},
				{
					ApplyTest:  func() { t.MoveLeft(); t.MoveUp(); t.MoveUp() },
					RevertTest: func() { t.MoveRight(); t.MoveDown(); t.MoveDown() },
				},
			}
		}
	}
	fmt.Printf("Unhandled orientation combo (prev, new) = (%s, %s)\n", &prevOrientation, &newOrientation)
	return []RotationTest{}
}
