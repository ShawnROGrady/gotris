package tetrimino

import (
	"github.com/ShawnROGrady/gotris/pkg/canvas"
	"github.com/ShawnROGrady/gotris/pkg/game/board"
)

type oPiece struct {
	box         Box
	orientation *orientation
}

func newOPiece(boardWidth, boardHeight int) Tetrimino {
	spawnOrientation := spawn

	piece := &oPiece{
		orientation: &spawnOrientation,
	}

	box := startingBox(boardWidth, boardHeight, piece)
	piece.box = box
	return piece
}

func (o oPiece) pieceOrientation() orientation {
	return *o.orientation
}

func (o *oPiece) ContainingBox() Box {
	return o.box
}

func (o *oPiece) YMax() Coordinates {
	var (
		boxBottomRight = o.box.BottomRight
		blocks         = o.Blocks()
	)

	yMax := findMaxY(blocks, boxBottomRight)

	return yMax
}

func (o *oPiece) YMin() Coordinates {
	var (
		boxTopLeft = o.box.TopLeft
		blocks     = o.Blocks()
	)

	yMin := findMinY(blocks, boxTopLeft)

	return yMin
}

func (o *oPiece) XMax() Coordinates {
	var (
		boxTopLeft = o.box.TopLeft
		blocks     = o.Blocks()
	)

	xMax := findMaxX(blocks, boxTopLeft)

	return xMax
}

func (o *oPiece) XMin() Coordinates {
	var (
		boxBottomRight = o.box.BottomRight
		blocks         = o.Blocks()
	)

	xMin := findMinX(blocks, boxBottomRight)

	return xMin
}

func (o *oPiece) Blocks() [][]*board.Block {
	switch *o.orientation {
	case spawn:
		return [][]*board.Block{
			[]*board.Block{nil, &board.Block{Color: canvas.Yellow}, &board.Block{Color: canvas.Yellow}, nil},
			[]*board.Block{nil, &board.Block{Color: canvas.Yellow}, &board.Block{Color: canvas.Yellow}, nil},
			[]*board.Block{nil, nil, nil, nil},
		}
	case clockwise:
		return [][]*board.Block{
			[]*board.Block{nil, &board.Block{Color: canvas.Yellow}, &board.Block{Color: canvas.Yellow}, nil},
			[]*board.Block{nil, &board.Block{Color: canvas.Yellow}, &board.Block{Color: canvas.Yellow}, nil},
			[]*board.Block{nil, nil, nil, nil},
		}
	case opposite:
		return [][]*board.Block{
			[]*board.Block{nil, &board.Block{Color: canvas.Yellow}, &board.Block{Color: canvas.Yellow}, nil},
			[]*board.Block{nil, &board.Block{Color: canvas.Yellow}, &board.Block{Color: canvas.Yellow}, nil},
			[]*board.Block{nil, nil, nil, nil},
		}
	case counterclockwise:
		return [][]*board.Block{
			[]*board.Block{nil, &board.Block{Color: canvas.Yellow}, &board.Block{Color: canvas.Yellow}, nil},
			[]*board.Block{nil, &board.Block{Color: canvas.Yellow}, &board.Block{Color: canvas.Yellow}, nil},
			[]*board.Block{nil, nil, nil, nil},
		}
	}
	return nil
}

func (o *oPiece) MoveUp(ymax int) {
	if o.YMax().Y < ymax {
		o.box.TopLeft.Y++
		o.box.BottomRight.Y++
	}
}

func (o *oPiece) MoveDown() {
	if o.YMin().Y > 0 {
		o.box.BottomRight.Y--
		o.box.TopLeft.Y--
	}
}

func (o *oPiece) MoveLeft() {
	if o.XMin().X > 0 {
		o.box.TopLeft.X--
		o.box.BottomRight.X--
	}
}

func (o *oPiece) MoveRight(xmax int) {
	if o.XMax().X < xmax {
		o.box.BottomRight.X++
		o.box.TopLeft.X++
	}
}

func (o *oPiece) RotateClockwise() {
	o.orientation.rotateClockwise()
}

func (o *oPiece) RotateCounter() {
	o.orientation.rotateCounter()
}
