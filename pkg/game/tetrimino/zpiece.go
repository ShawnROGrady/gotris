package tetrimino

import (
	"github.com/ShawnROGrady/gotris/pkg/canvas"
	"github.com/ShawnROGrady/gotris/pkg/game/board"
)

type zPiece struct {
	box         Box
	orientation *orientation
}

func newZPiece(boardWidth, boardHeight int) Tetrimino {
	spawnOrientation := spawn
	return &zPiece{
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

func (z zPiece) pieceOrientation() orientation {
	return *z.orientation
}

func (z *zPiece) ContainingBox() Box {
	return z.box
}

func (z *zPiece) YMax() Coordinates {
	var (
		boxBottomRight = z.box.BottomRight
		blocks         = z.Blocks()
	)

	yMax := findMaxY(blocks, boxBottomRight)

	return yMax
}

func (z *zPiece) YMin() Coordinates {
	var (
		boxTopLeft = z.box.TopLeft
		blocks     = z.Blocks()
	)

	yMin := findMinY(blocks, boxTopLeft)

	return yMin
}

func (z *zPiece) XMax() Coordinates {
	var (
		boxTopLeft = z.box.TopLeft
		blocks     = z.Blocks()
	)

	xMax := findMaxX(blocks, boxTopLeft)

	return xMax
}

func (z *zPiece) XMin() Coordinates {
	var (
		boxBottomRight = z.box.BottomRight
		blocks         = z.Blocks()
	)

	xMin := findMinX(blocks, boxBottomRight)

	return xMin
}

func (z *zPiece) Blocks() [][]*board.Block {
	switch *z.orientation {
	case spawn:
		return [][]*board.Block{
			[]*board.Block{
				nil,
				nil,
				&board.Block{Color: canvas.Red},
			},
			[]*board.Block{
				nil,
				&board.Block{Color: canvas.Red},
				&board.Block{Color: canvas.Red},
			},
			[]*board.Block{
				nil,
				&board.Block{Color: canvas.Red},
				nil,
			},
		}
	case clockwise:
		return [][]*board.Block{
			[]*board.Block{nil, nil, nil},
			[]*board.Block{
				&board.Block{Color: canvas.Red},
				&board.Block{Color: canvas.Red},
				nil,
			},
			[]*board.Block{
				nil,
				&board.Block{Color: canvas.Red},
				&board.Block{Color: canvas.Red},
			},
		}
	case opposite:
		return [][]*board.Block{
			[]*board.Block{
				nil,
				&board.Block{Color: canvas.Red},
				nil,
			},
			[]*board.Block{
				&board.Block{Color: canvas.Red},
				&board.Block{Color: canvas.Red},
				nil,
			},
			[]*board.Block{
				&board.Block{Color: canvas.Red},
				nil,
				nil,
			},
		}
	case counterclockwise:
		return [][]*board.Block{
			[]*board.Block{
				&board.Block{Color: canvas.Red},
				&board.Block{Color: canvas.Red},
				nil,
			},
			[]*board.Block{
				nil,
				&board.Block{Color: canvas.Red},
				&board.Block{Color: canvas.Red},
			},
			[]*board.Block{nil, nil, nil},
		}
	}
	return nil
}

func (z *zPiece) MoveUp(ymax int) {
	if z.YMax().Y < ymax {
		z.box.TopLeft.Y++
		z.box.BottomRight.Y++
	}
}

func (z *zPiece) MoveDown() {
	if z.YMin().Y > 0 {
		z.box.BottomRight.Y--
		z.box.TopLeft.Y--
	}
}

func (z *zPiece) MoveLeft() {
	if z.XMin().X > 0 {
		z.box.TopLeft.X--
		z.box.BottomRight.X--
	}
}

func (z *zPiece) MoveRight(xmax int) {
	if z.XMax().X < xmax {
		z.box.BottomRight.X++
		z.box.TopLeft.X++
	}
}

func (z *zPiece) RotateClockwise() {
	z.orientation.rotateClockwise()
}

func (z *zPiece) RotateCounter() {
	z.orientation.rotateCounter()
}