package tetrimino

import (
	"github.com/ShawnROGrady/gotris/pkg/canvas"
	"github.com/ShawnROGrady/gotris/pkg/game/board"
)

type jPiece struct {
	box         Box
	orientation *orientation
}

func newJPiece(boardWidth, boardHeight int) Tetrimino {
	spawnOrientation := spawn
	return &jPiece{
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

func (j jPiece) pieceOrientation() orientation {
	return *j.orientation
}

func (j *jPiece) ContainingBox() Box {
	return j.box
}

func (j *jPiece) YMax() Coordinates {
	var (
		boxBottomRight = j.box.BottomRight
		blocks         = j.Blocks()
	)

	yMax := findMaxY(blocks, boxBottomRight)

	return yMax
}

func (j *jPiece) YMin() Coordinates {
	var (
		boxTopLeft = j.box.TopLeft
		blocks     = j.Blocks()
	)

	yMin := findMinY(blocks, boxTopLeft)

	return yMin
}

func (j *jPiece) XMax() Coordinates {
	var (
		boxTopLeft = j.box.TopLeft
		blocks     = j.Blocks()
	)

	xMax := findMaxX(blocks, boxTopLeft)

	return xMax
}

func (j *jPiece) XMin() Coordinates {
	var (
		boxBottomRight = j.box.BottomRight
		blocks         = j.Blocks()
	)

	xMin := findMinX(blocks, boxBottomRight)

	return xMin
}

func (j *jPiece) Blocks() [][]*board.Block {
	switch *j.orientation {
	case spawn:
		return [][]*board.Block{
			[]*board.Block{
				nil,
				&board.Block{Color: canvas.Blue},
				&board.Block{Color: canvas.Blue},
			},
			[]*board.Block{nil, &board.Block{Color: canvas.Blue}, nil},
			[]*board.Block{nil, &board.Block{Color: canvas.Blue}, nil},
		}
	case clockwise:
		return [][]*board.Block{
			[]*board.Block{nil, nil, nil},
			[]*board.Block{
				&board.Block{Color: canvas.Blue},
				&board.Block{Color: canvas.Blue},
				&board.Block{Color: canvas.Blue},
			},
			[]*board.Block{
				nil,
				nil,
				&board.Block{Color: canvas.Blue},
			},
		}
	case opposite:
		return [][]*board.Block{
			[]*board.Block{nil, &board.Block{Color: canvas.Blue}, nil},
			[]*board.Block{nil, &board.Block{Color: canvas.Blue}, nil},
			[]*board.Block{
				&board.Block{Color: canvas.Blue},
				&board.Block{Color: canvas.Blue},
				nil,
			},
		}
	case counterclockwise:
		return [][]*board.Block{
			[]*board.Block{
				&board.Block{Color: canvas.Blue},
				nil,
				nil,
			},
			[]*board.Block{
				&board.Block{Color: canvas.Blue},
				&board.Block{Color: canvas.Blue},
				&board.Block{Color: canvas.Blue},
			},
			[]*board.Block{nil, nil, nil},
		}
	}
	return nil
}

func (j *jPiece) MoveUp(ymax int) {
	if j.YMax().Y < ymax {
		j.box.TopLeft.Y++
		j.box.BottomRight.Y++
	}
}

func (j *jPiece) MoveDown() {
	if j.YMin().Y > 0 {
		j.box.BottomRight.Y--
		j.box.TopLeft.Y--
	}
}

func (j *jPiece) MoveLeft() {
	if j.XMin().X > 0 {
		j.box.TopLeft.X--
		j.box.BottomRight.X--
	}
}

func (j *jPiece) MoveRight(xmax int) {
	if j.XMax().X < xmax {
		j.box.BottomRight.X++
		j.box.TopLeft.X++
	}
}

func (j *jPiece) RotateClockwise() {
	j.orientation.rotateClockwise()
}

func (j *jPiece) RotateCounter() {
	j.orientation.rotateCounter()
}
