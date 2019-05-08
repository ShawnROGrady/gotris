package tetrimino

import (
	"github.com/ShawnROGrady/gotris/pkg/canvas"
	"github.com/ShawnROGrady/gotris/pkg/game/board"
)

type lPiece struct {
	box             Box
	orientation     *orientation
	prevOrientation orientation
}

func newLPiece(boardWidth, boardHeight int) Tetrimino {
	spawnOrientation := spawn

	piece := &lPiece{
		orientation:     &spawnOrientation,
		prevOrientation: spawnOrientation,
	}

	box := startingBox(boardWidth, boardHeight, piece)
	piece.box = box
	return piece
}

func (l lPiece) pieceOrientation() orientation {
	return *l.orientation
}

func (l *lPiece) previousOrientation() orientation {
	return l.prevOrientation
}

func (l *lPiece) ContainingBox() Box {
	return l.box
}

func (l *lPiece) YMax() Coordinates {
	var (
		boxBottomRight = l.box.BottomRight
		blocks         = l.Blocks()
	)

	yMax := findMaxY(blocks, boxBottomRight)

	return yMax
}

func (l *lPiece) YMin() Coordinates {
	var (
		boxTopLeft = l.box.TopLeft
		blocks     = l.Blocks()
	)

	yMin := findMinY(blocks, boxTopLeft)

	return yMin
}

func (l *lPiece) XMax() Coordinates {
	var (
		boxTopLeft = l.box.TopLeft
		blocks     = l.Blocks()
	)

	xMax := findMaxX(blocks, boxTopLeft)

	return xMax
}

func (l *lPiece) XMin() Coordinates {
	var (
		boxBottomRight = l.box.BottomRight
		blocks         = l.Blocks()
	)

	xMin := findMinX(blocks, boxBottomRight)

	return xMin
}

func (l *lPiece) Blocks() [][]*board.Block {
	switch *l.orientation {
	case clockwise:
		return [][]*board.Block{
			[]*board.Block{nil, &board.Block{Color: canvas.Orange}, nil},
			[]*board.Block{nil, &board.Block{Color: canvas.Orange}, nil},
			[]*board.Block{
				nil,
				&board.Block{Color: canvas.Orange},
				&board.Block{Color: canvas.Orange},
			},
		}
	case opposite:
		return [][]*board.Block{
			[]*board.Block{nil, nil, nil},
			[]*board.Block{
				&board.Block{Color: canvas.Orange},
				&board.Block{Color: canvas.Orange},
				&board.Block{Color: canvas.Orange},
			},
			[]*board.Block{
				&board.Block{Color: canvas.Orange},
				nil,
				nil,
			},
		}
	case counterclockwise:
		return [][]*board.Block{
			[]*board.Block{
				&board.Block{Color: canvas.Orange},
				&board.Block{Color: canvas.Orange},
				nil,
			},
			[]*board.Block{nil, &board.Block{Color: canvas.Orange}, nil},
			[]*board.Block{nil, &board.Block{Color: canvas.Orange}, nil},
		}
	case spawn:
		return [][]*board.Block{
			[]*board.Block{
				nil,
				nil,
				&board.Block{Color: canvas.Orange},
			},
			[]*board.Block{
				&board.Block{Color: canvas.Orange},
				&board.Block{Color: canvas.Orange},
				&board.Block{Color: canvas.Orange},
			},
			[]*board.Block{nil, nil, nil},
		}
	}
	return nil
}

func (l *lPiece) MoveUp(ymax int) {
	if l.YMax().Y < ymax {
		l.box.TopLeft.Y++
		l.box.BottomRight.Y++
	}
}

func (l *lPiece) MoveDown() {
	if l.YMin().Y > 0 {
		l.box.BottomRight.Y--
		l.box.TopLeft.Y--
	}
}

func (l *lPiece) MoveLeft() {
	if l.XMin().X > 0 {
		l.box.TopLeft.X--
		l.box.BottomRight.X--
	}
}

func (l *lPiece) MoveRight(xmax int) {
	if l.XMax().X < xmax {
		l.box.BottomRight.X++
		l.box.TopLeft.X++
	}
}

func (l *lPiece) RotateClockwise() {
	l.prevOrientation = *l.orientation
	l.orientation.rotateClockwise()
}

func (l *lPiece) RotateCounter() {
	l.prevOrientation = *l.orientation
	l.orientation.rotateCounter()
}

func (l *lPiece) RotationTests() []RotationTest {
	return defaultRotationTests(l, l.prevOrientation, l.pieceOrientation())
}
