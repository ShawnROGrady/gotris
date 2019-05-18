package tetrimino

import (
	"github.com/ShawnROGrady/gotris/pkg/canvas"
	"github.com/ShawnROGrady/gotris/pkg/game/board"
)

type jPiece struct {
	*tetriminoBase
}

func newJPiece(boardWidth, boardHeight int) Tetrimino {
	spawnOrientation := spawn

	piece := &jPiece{
		tetriminoBase: &tetriminoBase{
			orientation:     &spawnOrientation,
			prevOrientation: spawnOrientation,
			color:           canvas.Blue,
		},
	}

	box := startingBox(boardWidth, boardHeight, piece)
	piece.box = box
	return piece
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
				&board.Block{Color: j.color},
				&board.Block{Color: j.color},
			},
			[]*board.Block{nil, &board.Block{Color: j.color}, nil},
			[]*board.Block{nil, &board.Block{Color: j.color}, nil},
		}
	case opposite:
		return [][]*board.Block{
			[]*board.Block{nil, nil, nil},
			[]*board.Block{
				&board.Block{Color: j.color},
				&board.Block{Color: j.color},
				&board.Block{Color: j.color},
			},
			[]*board.Block{
				nil,
				nil,
				&board.Block{Color: j.color},
			},
		}
	case counterclockwise:
		return [][]*board.Block{
			[]*board.Block{nil, &board.Block{Color: j.color}, nil},
			[]*board.Block{nil, &board.Block{Color: j.color}, nil},
			[]*board.Block{
				&board.Block{Color: j.color},
				&board.Block{Color: j.color},
				nil,
			},
		}
	case spawn:
		return [][]*board.Block{
			[]*board.Block{
				&board.Block{Color: j.color},
				nil,
				nil,
			},
			[]*board.Block{
				&board.Block{Color: j.color},
				&board.Block{Color: j.color},
				&board.Block{Color: j.color},
			},
			[]*board.Block{nil, nil, nil},
		}
	}
	return nil
}

func (j *jPiece) RotationTests() []RotationTest {
	return defaultRotationTests(j, j.prevOrientation, j.pieceOrientation())
}
