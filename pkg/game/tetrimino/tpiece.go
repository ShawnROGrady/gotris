package tetrimino

import (
	"github.com/ShawnROGrady/gotris/pkg/canvas"
	"github.com/ShawnROGrady/gotris/pkg/game/board"
)

type tPiece struct {
	*tetriminoBase
}

func newTPiece(boardWidth, boardHeight int) Tetrimino {
	spawnOrientation := spawn

	piece := &tPiece{
		tetriminoBase: &tetriminoBase{
			orientation:     &spawnOrientation,
			prevOrientation: spawnOrientation,
			color:           canvas.Magenta,
		},
	}

	box := startingBox(boardWidth, boardHeight, piece)
	piece.box = box
	return piece
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
				&board.Block{Color: t.color},
				nil,
			},
			[]*board.Block{
				nil,
				&board.Block{Color: t.color},
				&board.Block{Color: t.color},
			},
			[]*board.Block{
				nil,
				&board.Block{Color: t.color},
				nil,
			},
		}
	case opposite:
		return [][]*board.Block{
			[]*board.Block{nil, nil, nil},
			[]*board.Block{
				&board.Block{Color: t.color},
				&board.Block{Color: t.color},
				&board.Block{Color: t.color},
			},
			[]*board.Block{
				nil,
				&board.Block{Color: t.color},
				nil,
			},
		}
	case counterclockwise:
		return [][]*board.Block{
			[]*board.Block{
				nil,
				&board.Block{Color: t.color},
				nil,
			},
			[]*board.Block{
				&board.Block{Color: t.color},
				&board.Block{Color: t.color},
				nil,
			},
			[]*board.Block{
				nil,
				&board.Block{Color: t.color},
				nil,
			},
		}
	case spawn:
		return [][]*board.Block{
			[]*board.Block{
				nil,
				&board.Block{Color: t.color},
				nil,
			},
			[]*board.Block{
				&board.Block{Color: t.color},
				&board.Block{Color: t.color},
				&board.Block{Color: t.color},
			},
			[]*board.Block{nil, nil, nil},
		}
	}
	return nil
}

func (t *tPiece) RotationTests() []RotationTest {
	return defaultRotationTests(t, t.prevOrientation, t.pieceOrientation())
}
