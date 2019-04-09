package tetrimino

import (
	"github.com/ShawnROGrady/gotris/pkg/canvas"
	"github.com/ShawnROGrady/gotris/pkg/game/board"
)

type tPiece struct {
	box         Box
	orientation *orientation
}

func newTPiece(boardWidth, boardHeight int) Tetrimino {
	spawnOrientation := spawn
	return &tPiece{
		orientation: &spawnOrientation,
		box: Box{
			TopLeft: Coordinates{
				X: 0,
				Y: boardHeight - 1,
			},
			BottomRight: Coordinates{
				X: 2,
				Y: boardHeight - 3,
			},
		},
	}
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
	case spawn:
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
	case clockwise:
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
	case opposite:
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
	t.orientation.rotateClockwise()
}

func (t *tPiece) RotateCounter() {
	t.orientation.rotateCounter()
}
