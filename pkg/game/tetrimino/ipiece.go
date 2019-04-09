package tetrimino

import (
	"github.com/ShawnROGrady/gotris/pkg/canvas"
	"github.com/ShawnROGrady/gotris/pkg/game/board"
)

type iPiece struct {
	box         Box
	orientation *orientation
}

func newIPiece(boardWidth, boardHeight int) Tetrimino {
	spawnOrientation := spawn
	return &iPiece{
		orientation: &spawnOrientation,
		box: Box{
			TopLeft: Coordinates{
				X: 0,
				Y: boardHeight - 1,
			},
			BottomRight: Coordinates{
				X: 3,
				Y: boardHeight - 4,
			},
		},
	}
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
	case spawn:
		return [][]*board.Block{
			[]*board.Block{nil, nil, &board.Block{Color: canvas.Cyan}, nil},
			[]*board.Block{nil, nil, &board.Block{Color: canvas.Cyan}, nil},
			[]*board.Block{nil, nil, &board.Block{Color: canvas.Cyan}, nil},
			[]*board.Block{nil, nil, &board.Block{Color: canvas.Cyan}, nil},
		}
	case clockwise:
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
	case opposite:
		return [][]*board.Block{
			[]*board.Block{nil, &board.Block{Color: canvas.Cyan}, nil, nil},
			[]*board.Block{nil, &board.Block{Color: canvas.Cyan}, nil, nil},
			[]*board.Block{nil, &board.Block{Color: canvas.Cyan}, nil, nil},
			[]*board.Block{nil, &board.Block{Color: canvas.Cyan}, nil, nil},
		}
	case counterclockwise:
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
	i.orientation.rotateClockwise()
}

func (i *iPiece) RotateCounter() {
	i.orientation.rotateCounter()
}
