package tetrimino

import (
	"github.com/ShawnROGrady/gotris/pkg/canvas"
	"github.com/ShawnROGrady/gotris/pkg/game/board"
)

type tPiece struct {
	box             Box
	orientation     *orientation
	prevOrientation orientation
}

func newTPiece(boardWidth, boardHeight int) Tetrimino {
	spawnOrientation := spawn

	piece := &tPiece{
		orientation:     &spawnOrientation,
		prevOrientation: spawnOrientation,
	}

	box := startingBox(boardWidth, boardHeight, piece)
	piece.box = box
	return piece
}

func (t tPiece) pieceOrientation() orientation {
	return *t.orientation
}

func (t *tPiece) previousOrientation() orientation {
	return t.prevOrientation
}

func (t *tPiece) ContainingBox() Box {
	return t.box
}

func (t *tPiece) YMax() Coordinates {
	var (
		boxBottomRight = t.box.BottomRight
		blocks         = t.Blocks()
	)

	yMax := findMaxY(blocks, boxBottomRight)

	return yMax
}

func (t *tPiece) YMin() Coordinates {
	var (
		boxTopLeft = t.box.TopLeft
		blocks     = t.Blocks()
	)

	yMin := findMinY(blocks, boxTopLeft)

	return yMin
}

func (t *tPiece) XMax() Coordinates {
	var (
		boxTopLeft = t.box.TopLeft
		blocks     = t.Blocks()
	)

	xMax := findMaxX(blocks, boxTopLeft)

	return xMax
}

func (t *tPiece) XMin() Coordinates {
	var (
		boxBottomRight = t.box.BottomRight
		blocks         = t.Blocks()
	)

	xMin := findMinX(blocks, boxBottomRight)

	return xMin
}

func (t *tPiece) Blocks() [][]*board.Block {
	switch *t.orientation {
	case clockwise:
		return [][]*board.Block{
			[]*board.Block{
				nil,
				&board.Block{Color: canvas.Magenta},
				nil,
			},
			[]*board.Block{
				nil,
				&board.Block{Color: canvas.Magenta},
				&board.Block{Color: canvas.Magenta},
			},
			[]*board.Block{
				nil,
				&board.Block{Color: canvas.Magenta},
				nil,
			},
		}
	case opposite:
		return [][]*board.Block{
			[]*board.Block{nil, nil, nil},
			[]*board.Block{
				&board.Block{Color: canvas.Magenta},
				&board.Block{Color: canvas.Magenta},
				&board.Block{Color: canvas.Magenta},
			},
			[]*board.Block{
				nil,
				&board.Block{Color: canvas.Magenta},
				nil,
			},
		}
	case counterclockwise:
		return [][]*board.Block{
			[]*board.Block{
				nil,
				&board.Block{Color: canvas.Magenta},
				nil,
			},
			[]*board.Block{
				&board.Block{Color: canvas.Magenta},
				&board.Block{Color: canvas.Magenta},
				nil,
			},
			[]*board.Block{
				nil,
				&board.Block{Color: canvas.Magenta},
				nil,
			},
		}
	case spawn:
		return [][]*board.Block{
			[]*board.Block{
				nil,
				&board.Block{Color: canvas.Magenta},
				nil,
			},
			[]*board.Block{
				&board.Block{Color: canvas.Magenta},
				&board.Block{Color: canvas.Magenta},
				&board.Block{Color: canvas.Magenta},
			},
			[]*board.Block{nil, nil, nil},
		}
	}
	return nil
}

func (t *tPiece) MoveUp(ymax int) {
	if t.YMax().Y < ymax {
		t.box.TopLeft.Y++
		t.box.BottomRight.Y++
	}
}

func (t *tPiece) MoveDown() {
	if t.YMin().Y > 0 {
		t.box.BottomRight.Y--
		t.box.TopLeft.Y--
	}
}

func (t *tPiece) MoveLeft() {
	if t.XMin().X > 0 {
		t.box.TopLeft.X--
		t.box.BottomRight.X--
	}
}

func (t *tPiece) MoveRight(xmax int) {
	if t.XMax().X < xmax {
		t.box.BottomRight.X++
		t.box.TopLeft.X++
	}
}

func (t *tPiece) RotateClockwise() {
	t.prevOrientation = *t.orientation
	t.orientation.rotateClockwise()
}

func (t *tPiece) RotateCounter() {
	t.prevOrientation = *t.orientation
	t.orientation.rotateCounter()
}

func (t *tPiece) RotationTests() []RotationTest {
	return defaultRotationTests(t, t.prevOrientation, t.pieceOrientation())
}
