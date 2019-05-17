package tetrimino

import (
	"fmt"

	"github.com/ShawnROGrady/gotris/pkg/canvas"
	"github.com/ShawnROGrady/gotris/pkg/game/board"
)

type iPiece struct {
	*tetriminoBase
}

func newIPiece(boardWidth, boardHeight int) Tetrimino {
	spawnOrientation := spawn

	piece := &iPiece{
		tetriminoBase: &tetriminoBase{
			orientation:     &spawnOrientation,
			prevOrientation: spawnOrientation,
		},
	}

	box := startingBox(boardWidth, boardHeight, piece)
	piece.box = box
	return piece
}

func (i *iPiece) YMax() Coordinates {
	var (
		boxBottomRight = i.box.BottomRight
		blocks         = i.Blocks()
	)

	yMax := findMaxY(blocks, boxBottomRight)

	return yMax
}

func (i *iPiece) YMin() Coordinates {
	var (
		boxTopLeft = i.box.TopLeft
		blocks     = i.Blocks()
	)

	yMin := findMinY(blocks, boxTopLeft)

	return yMin
}

func (i *iPiece) XMax() Coordinates {
	var (
		boxTopLeft = i.box.TopLeft
		blocks     = i.Blocks()
	)

	xMax := findMaxX(blocks, boxTopLeft)

	return xMax
}

func (i *iPiece) XMin() Coordinates {
	var (
		boxBottomRight = i.box.BottomRight
		blocks         = i.Blocks()
	)

	xMin := findMinX(blocks, boxBottomRight)

	return xMin
}

func (i *iPiece) Blocks() [][]*board.Block {
	switch *i.orientation {
	case clockwise:
		return [][]*board.Block{
			[]*board.Block{nil, nil, &board.Block{Color: canvas.Cyan}, nil},
			[]*board.Block{nil, nil, &board.Block{Color: canvas.Cyan}, nil},
			[]*board.Block{nil, nil, &board.Block{Color: canvas.Cyan}, nil},
			[]*board.Block{nil, nil, &board.Block{Color: canvas.Cyan}, nil},
		}
	case opposite:
		return [][]*board.Block{
			[]*board.Block{nil, nil, nil, nil},
			[]*board.Block{nil, nil, nil, nil},
			[]*board.Block{
				&board.Block{Color: canvas.Cyan},
				&board.Block{Color: canvas.Cyan},
				&board.Block{Color: canvas.Cyan},
				&board.Block{Color: canvas.Cyan},
			},
			[]*board.Block{nil, nil, nil, nil},
		}
	case counterclockwise:
		return [][]*board.Block{
			[]*board.Block{nil, &board.Block{Color: canvas.Cyan}, nil, nil},
			[]*board.Block{nil, &board.Block{Color: canvas.Cyan}, nil, nil},
			[]*board.Block{nil, &board.Block{Color: canvas.Cyan}, nil, nil},
			[]*board.Block{nil, &board.Block{Color: canvas.Cyan}, nil, nil},
		}
	case spawn:
		return [][]*board.Block{
			[]*board.Block{nil, nil, nil, nil},
			[]*board.Block{
				&board.Block{Color: canvas.Cyan},
				&board.Block{Color: canvas.Cyan},
				&board.Block{Color: canvas.Cyan},
				&board.Block{Color: canvas.Cyan},
			},
			[]*board.Block{nil, nil, nil, nil},
			[]*board.Block{nil, nil, nil, nil},
		}

	}
	return nil
}

func (i *iPiece) RotationTests() []RotationTest {
	return iPieceRotationTests(i, i.prevOrientation, i.pieceOrientation())
}

func iPieceRotationTests(t Tetrimino, prevOrientation, newOrientation orientation) []RotationTest {
	switch prevOrientation {
	case spawn:
		switch newOrientation {
		case clockwise:
			return []RotationTest{
				{
					ApplyTest:  func() { t.MoveLeft(); t.MoveLeft() },
					RevertTest: func() { t.MoveRight(); t.MoveRight() },
				},
				{
					ApplyTest:  func() { t.MoveRight() },
					RevertTest: func() { t.MoveLeft() },
				},
				{
					ApplyTest:  func() { t.MoveLeft(); t.MoveLeft(); t.MoveDown() },
					RevertTest: func() { t.MoveRight(); t.MoveRight(); t.MoveUp() },
				},
				{
					ApplyTest:  func() { t.MoveRight(); t.MoveUp(); t.MoveUp() },
					RevertTest: func() { t.MoveLeft(); t.MoveDown(); t.MoveDown() },
				},
			}
		case counterclockwise:
			return []RotationTest{
				{
					ApplyTest:  func() { t.MoveLeft() },
					RevertTest: func() { t.MoveRight() },
				},
				{
					ApplyTest:  func() { t.MoveRight(); t.MoveRight() },
					RevertTest: func() { t.MoveLeft(); t.MoveLeft() },
				},
				{
					ApplyTest:  func() { t.MoveLeft(); t.MoveUp(); t.MoveUp() },
					RevertTest: func() { t.MoveRight(); t.MoveDown(); t.MoveDown() },
				},
				{
					ApplyTest:  func() { t.MoveRight(); t.MoveRight(); t.MoveDown() },
					RevertTest: func() { t.MoveLeft(); t.MoveLeft(); t.MoveUp() },
				},
			}
		}
	case clockwise:
		switch newOrientation {
		case spawn:
			return []RotationTest{
				{
					ApplyTest:  func() { t.MoveRight(); t.MoveRight() },
					RevertTest: func() { t.MoveLeft(); t.MoveLeft() },
				},
				{
					ApplyTest:  func() { t.MoveLeft() },
					RevertTest: func() { t.MoveRight() },
				},
				{
					ApplyTest:  func() { t.MoveRight(); t.MoveRight(); t.MoveUp() },
					RevertTest: func() { t.MoveLeft(); t.MoveLeft(); t.MoveDown() },
				},
				{
					ApplyTest:  func() { t.MoveLeft(); t.MoveDown(); t.MoveDown() },
					RevertTest: func() { t.MoveRight(); t.MoveUp(); t.MoveUp() },
				},
			}
		case opposite:
			return []RotationTest{
				{
					ApplyTest:  func() { t.MoveLeft() },
					RevertTest: func() { t.MoveRight() },
				},
				{
					ApplyTest:  func() { t.MoveRight(); t.MoveRight() },
					RevertTest: func() { t.MoveLeft(); t.MoveLeft() },
				},
				{
					ApplyTest:  func() { t.MoveLeft(); t.MoveUp(); t.MoveUp() },
					RevertTest: func() { t.MoveRight(); t.MoveDown(); t.MoveDown() },
				},
				{
					ApplyTest:  func() { t.MoveRight(); t.MoveRight(); t.MoveDown() },
					RevertTest: func() { t.MoveLeft(); t.MoveLeft(); t.MoveUp() },
				},
			}
		}
	case opposite:
		switch newOrientation {
		case clockwise:
			return []RotationTest{
				{
					ApplyTest:  func() { t.MoveRight() },
					RevertTest: func() { t.MoveLeft() },
				},
				{
					ApplyTest:  func() { t.MoveLeft(); t.MoveLeft() },
					RevertTest: func() { t.MoveRight(); t.MoveRight() },
				},
				{
					ApplyTest:  func() { t.MoveRight(); t.MoveDown(); t.MoveDown() },
					RevertTest: func() { t.MoveLeft(); t.MoveUp(); t.MoveUp() },
				},
				{
					ApplyTest:  func() { t.MoveLeft(); t.MoveLeft(); t.MoveUp() },
					RevertTest: func() { t.MoveRight(); t.MoveRight(); t.MoveDown() },
				},
			}
		case counterclockwise:
			return []RotationTest{
				{
					ApplyTest:  func() { t.MoveRight(); t.MoveRight() },
					RevertTest: func() { t.MoveLeft(); t.MoveLeft() },
				},
				{
					ApplyTest:  func() { t.MoveLeft() },
					RevertTest: func() { t.MoveRight() },
				},
				{
					ApplyTest:  func() { t.MoveRight(); t.MoveRight(); t.MoveUp() },
					RevertTest: func() { t.MoveLeft(); t.MoveLeft(); t.MoveDown() },
				},
				{
					ApplyTest:  func() { t.MoveLeft(); t.MoveDown(); t.MoveDown() },
					RevertTest: func() { t.MoveRight(); t.MoveUp(); t.MoveUp() },
				},
			}
		}
	case counterclockwise:
		switch newOrientation {
		case opposite:
			return []RotationTest{
				{
					ApplyTest:  func() { t.MoveLeft(); t.MoveLeft() },
					RevertTest: func() { t.MoveRight(); t.MoveRight() },
				},
				{
					ApplyTest:  func() { t.MoveRight() },
					RevertTest: func() { t.MoveLeft() },
				},
				{
					ApplyTest:  func() { t.MoveLeft(); t.MoveLeft(); t.MoveDown() },
					RevertTest: func() { t.MoveRight(); t.MoveRight(); t.MoveUp() },
				},
				{
					ApplyTest:  func() { t.MoveRight(); t.MoveUp(); t.MoveUp() },
					RevertTest: func() { t.MoveLeft(); t.MoveDown(); t.MoveDown() },
				},
			}
		case spawn:
			return []RotationTest{
				{
					ApplyTest:  func() { t.MoveRight() },
					RevertTest: func() { t.MoveLeft() },
				},
				{
					ApplyTest:  func() { t.MoveLeft(); t.MoveLeft() },
					RevertTest: func() { t.MoveRight(); t.MoveRight() },
				},
				{
					ApplyTest:  func() { t.MoveRight(); t.MoveDown(); t.MoveDown() },
					RevertTest: func() { t.MoveLeft(); t.MoveUp(); t.MoveUp() },
				},
				{
					ApplyTest:  func() { t.MoveLeft(); t.MoveLeft(); t.MoveUp() },
					RevertTest: func() { t.MoveRight(); t.MoveRight(); t.MoveDown() },
				},
			}
		}
	}
	fmt.Printf("Unhandled orientation combo (prev, new) = (%s, %s)\n", &prevOrientation, &newOrientation)
	return []RotationTest{}
}
