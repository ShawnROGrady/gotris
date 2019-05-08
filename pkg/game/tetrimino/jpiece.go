package tetrimino

import (
	"github.com/ShawnROGrady/gotris/pkg/canvas"
	"github.com/ShawnROGrady/gotris/pkg/game/board"
)

type jPiece struct {
	box             Box
	orientation     *orientation
	prevOrientation orientation
}

func newJPiece(boardWidth, boardHeight int) Tetrimino {
	spawnOrientation := spawn

	piece := &jPiece{
		orientation:     &spawnOrientation,
		prevOrientation: spawnOrientation,
	}

	box := startingBox(boardWidth, boardHeight, piece)
	piece.box = box
	return piece
}

func (j jPiece) pieceOrientation() orientation {
	return *j.orientation
}

func (j *jPiece) previousOrientation() orientation {
	return j.prevOrientation
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
	case clockwise:
		return [][]*board.Block{
			[]*board.Block{
				nil,
				&board.Block{Color: canvas.Blue},
				&board.Block{Color: canvas.Blue},
			},
			[]*board.Block{nil, &board.Block{Color: canvas.Blue}, nil},
			[]*board.Block{nil, &board.Block{Color: canvas.Blue}, nil},
		}
	case opposite:
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
	case counterclockwise:
		return [][]*board.Block{
			[]*board.Block{nil, &board.Block{Color: canvas.Blue}, nil},
			[]*board.Block{nil, &board.Block{Color: canvas.Blue}, nil},
			[]*board.Block{
				&board.Block{Color: canvas.Blue},
				&board.Block{Color: canvas.Blue},
				nil,
			},
		}
	case spawn:
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

func (j *jPiece) MoveUp() {
	j.box.TopLeft.Y++
	j.box.BottomRight.Y++
}

func (j *jPiece) MoveDown() {
	j.box.BottomRight.Y--
	j.box.TopLeft.Y--
}

func (j *jPiece) MoveLeft() {
	j.box.TopLeft.X--
	j.box.BottomRight.X--
}

func (j *jPiece) MoveRight() {
	j.box.BottomRight.X++
	j.box.TopLeft.X++
}

func (j *jPiece) RotateClockwise() {
	j.prevOrientation = *j.orientation
	j.orientation.rotateClockwise()
}

func (j *jPiece) RotateCounter() {
	j.prevOrientation = *j.orientation
	j.orientation.rotateCounter()
}

func (j *jPiece) RotationTests() []RotationTest {
	return defaultRotationTests(j, j.prevOrientation, j.pieceOrientation())
}
