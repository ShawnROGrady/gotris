package tetrimino

import (
	"github.com/ShawnROGrady/gotris/pkg/canvas"
	"github.com/ShawnROGrady/gotris/pkg/game/board"
)

type lPiece struct {
	*tetriminoBase
}

func newLPiece(boardWidth, boardHeight int) Tetrimino {
	spawnOrientation := spawn

	piece := &lPiece{
		tetriminoBase: &tetriminoBase{
			orientation:     &spawnOrientation,
			prevOrientation: spawnOrientation,
		},
	}

	box := startingBox(boardWidth, boardHeight, piece)
	piece.box = box
	return piece
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

func (l *lPiece) RotationTests() []RotationTest {
	return defaultRotationTests(l, l.prevOrientation, l.pieceOrientation())
}
