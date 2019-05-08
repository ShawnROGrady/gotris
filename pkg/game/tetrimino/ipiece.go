package tetrimino

import (
	"fmt"

	"github.com/ShawnROGrady/gotris/pkg/canvas"
	"github.com/ShawnROGrady/gotris/pkg/game/board"
)

type iPiece struct {
	box             Box
	orientation     *orientation
	prevOrientation orientation
}

func newIPiece(boardWidth, boardHeight int) Tetrimino {
	spawnOrientation := spawn

	piece := &iPiece{
		orientation:     &spawnOrientation,
		prevOrientation: spawnOrientation,
	}

	box := startingBox(boardWidth, boardHeight, piece)
	piece.box = box
	return piece
}

func (i *iPiece) pieceOrientation() orientation {
	return *i.orientation
}

func (i *iPiece) previousOrientation() orientation {
	return i.prevOrientation
}

func (i *iPiece) ContainingBox() Box {
	return i.box
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

func (i *iPiece) MoveUp(ymax int) {
	if i.YMax().Y < ymax {
		i.box.TopLeft.Y++
		i.box.BottomRight.Y++
	}
}

func (i *iPiece) MoveDown() {
	if i.YMin().Y > 0 {
		i.box.BottomRight.Y--
		i.box.TopLeft.Y--
	}
}

func (i *iPiece) MoveLeft() {
	if i.XMin().X > 0 {
		i.box.TopLeft.X--
		i.box.BottomRight.X--
	}
}

func (i *iPiece) MoveRight(xmax int) {
	if i.XMax().X < xmax {
		i.box.BottomRight.X++
		i.box.TopLeft.X++
	}
}

func (i *iPiece) RotateClockwise() {
	i.prevOrientation = *i.orientation
	i.orientation.rotateClockwise()
}

func (i *iPiece) RotateCounter() {
	i.prevOrientation = *i.orientation
	i.orientation.rotateCounter()
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
					ApplyTest:  func(xmax, ymax int) { t.MoveLeft(); t.MoveLeft() },
					RevertTest: func(xmax, ymax int) { t.MoveRight(xmax); t.MoveRight(xmax) },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveRight(xmax) },
					RevertTest: func(xmax, ymax int) { t.MoveLeft() },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveLeft(); t.MoveLeft(); t.MoveDown() },
					RevertTest: func(xmax, ymax int) { t.MoveRight(xmax); t.MoveRight(xmax); t.MoveUp(ymax) },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveRight(xmax); t.MoveUp(ymax); t.MoveUp(ymax) },
					RevertTest: func(xmax, ymax int) { t.MoveLeft(); t.MoveDown(); t.MoveDown() },
				},
			}
		case counterclockwise:
			return []RotationTest{
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveLeft() },
					RevertTest: func(xmax, ymax int) { t.MoveRight(xmax) },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveRight(xmax); t.MoveRight(xmax) },
					RevertTest: func(xmax, ymax int) { t.MoveLeft(); t.MoveLeft() },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveLeft(); t.MoveUp(ymax); t.MoveUp(ymax) },
					RevertTest: func(xmax, ymax int) { t.MoveRight(xmax); t.MoveDown(); t.MoveDown() },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveRight(xmax); t.MoveRight(xmax); t.MoveDown() },
					RevertTest: func(xmax, ymax int) { t.MoveLeft(); t.MoveLeft(); t.MoveUp(ymax) },
				},
			}
		}
	case clockwise:
		switch newOrientation {
		case spawn:
			return []RotationTest{
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveRight(xmax); t.MoveRight(xmax) },
					RevertTest: func(xmax, ymax int) { t.MoveLeft(); t.MoveLeft() },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveLeft() },
					RevertTest: func(xmax, ymax int) { t.MoveRight(xmax) },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveRight(xmax); t.MoveRight(xmax); t.MoveUp(ymax) },
					RevertTest: func(xmax, ymax int) { t.MoveLeft(); t.MoveLeft(); t.MoveDown() },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveLeft(); t.MoveDown(); t.MoveDown() },
					RevertTest: func(xmax, ymax int) { t.MoveRight(xmax); t.MoveUp(ymax); t.MoveUp(ymax) },
				},
			}
		case opposite:
			return []RotationTest{
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveLeft() },
					RevertTest: func(xmax, ymax int) { t.MoveRight(xmax) },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveRight(xmax); t.MoveRight(xmax) },
					RevertTest: func(xmax, ymax int) { t.MoveLeft(); t.MoveLeft() },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveLeft(); t.MoveUp(ymax); t.MoveUp(ymax) },
					RevertTest: func(xmax, ymax int) { t.MoveRight(xmax); t.MoveDown(); t.MoveDown() },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveRight(xmax); t.MoveRight(xmax); t.MoveDown() },
					RevertTest: func(xmax, ymax int) { t.MoveLeft(); t.MoveLeft(); t.MoveUp(ymax) },
				},
			}
		}
	case opposite:
		switch newOrientation {
		case clockwise:
			return []RotationTest{
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveRight(xmax) },
					RevertTest: func(xmax, ymax int) { t.MoveLeft() },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveLeft(); t.MoveLeft() },
					RevertTest: func(xmax, ymax int) { t.MoveRight(xmax); t.MoveRight(xmax) },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveRight(xmax); t.MoveDown(); t.MoveDown() },
					RevertTest: func(xmax, ymax int) { t.MoveLeft(); t.MoveUp(ymax); t.MoveUp(ymax) },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveLeft(); t.MoveLeft(); t.MoveUp(xmax) },
					RevertTest: func(xmax, ymax int) { t.MoveRight(xmax); t.MoveRight(xmax); t.MoveDown() },
				},
			}
		case counterclockwise:
			return []RotationTest{
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveRight(xmax); t.MoveRight(xmax) },
					RevertTest: func(xmax, ymax int) { t.MoveLeft(); t.MoveLeft() },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveLeft() },
					RevertTest: func(xmax, ymax int) { t.MoveRight(xmax) },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveRight(xmax); t.MoveRight(xmax); t.MoveUp(xmax) },
					RevertTest: func(xmax, ymax int) { t.MoveLeft(); t.MoveLeft(); t.MoveDown() },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveLeft(); t.MoveDown(); t.MoveDown() },
					RevertTest: func(xmax, ymax int) { t.MoveRight(xmax); t.MoveUp(ymax); t.MoveUp(ymax) },
				},
			}
		}
	case counterclockwise:
		switch newOrientation {
		case opposite:
			return []RotationTest{
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveLeft(); t.MoveLeft() },
					RevertTest: func(xmax, ymax int) { t.MoveRight(xmax); t.MoveRight(xmax) },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveRight(xmax) },
					RevertTest: func(xmax, ymax int) { t.MoveLeft() },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveLeft(); t.MoveLeft(); t.MoveDown() },
					RevertTest: func(xmax, ymax int) { t.MoveRight(xmax); t.MoveRight(xmax); t.MoveUp(xmax) },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveRight(xmax); t.MoveUp(ymax); t.MoveUp(ymax) },
					RevertTest: func(xmax, ymax int) { t.MoveLeft(); t.MoveDown(); t.MoveDown() },
				},
			}
		case spawn:
			return []RotationTest{
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveRight(xmax) },
					RevertTest: func(xmax, ymax int) { t.MoveLeft() },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveLeft(); t.MoveLeft() },
					RevertTest: func(xmax, ymax int) { t.MoveRight(xmax); t.MoveRight(xmax) },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveRight(xmax); t.MoveDown(); t.MoveDown() },
					RevertTest: func(xmax, ymax int) { t.MoveLeft(); t.MoveUp(ymax); t.MoveUp(ymax) },
				},
				{
					ApplyTest:  func(xmax, ymax int) { t.MoveLeft(); t.MoveLeft(); t.MoveUp(xmax) },
					RevertTest: func(xmax, ymax int) { t.MoveRight(xmax); t.MoveRight(xmax); t.MoveDown() },
				},
			}
		}
	}
	fmt.Printf("Unhandled orientation combo (prev, new) = (%s, %s)\n", &prevOrientation, &newOrientation)
	return []RotationTest{}
}
