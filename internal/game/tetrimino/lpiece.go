package tetrimino

import (
	"github.com/ShawnROGrady/gotris/internal/canvas"
	"github.com/ShawnROGrady/gotris/internal/game/board"
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
			color:           canvas.Orange,
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
			[]*board.Block{nil, &board.Block{Color: l.color, Transparent: l.isGhost}, nil},
			[]*board.Block{nil, &board.Block{Color: l.color, Transparent: l.isGhost}, nil},
			[]*board.Block{
				nil,
				&board.Block{Color: l.color, Transparent: l.isGhost},
				&board.Block{Color: l.color, Transparent: l.isGhost},
			},
		}
	case opposite:
		return [][]*board.Block{
			[]*board.Block{nil, nil, nil},
			[]*board.Block{
				&board.Block{Color: l.color, Transparent: l.isGhost},
				&board.Block{Color: l.color, Transparent: l.isGhost},
				&board.Block{Color: l.color, Transparent: l.isGhost},
			},
			[]*board.Block{
				&board.Block{Color: l.color, Transparent: l.isGhost},
				nil,
				nil,
			},
		}
	case counterclockwise:
		return [][]*board.Block{
			[]*board.Block{
				&board.Block{Color: l.color, Transparent: l.isGhost},
				&board.Block{Color: l.color, Transparent: l.isGhost},
				nil,
			},
			[]*board.Block{nil, &board.Block{Color: l.color, Transparent: l.isGhost}, nil},
			[]*board.Block{nil, &board.Block{Color: l.color, Transparent: l.isGhost}, nil},
		}
	case spawn:
		return [][]*board.Block{
			[]*board.Block{
				nil,
				nil,
				&board.Block{Color: l.color, Transparent: l.isGhost},
			},
			[]*board.Block{
				&board.Block{Color: l.color, Transparent: l.isGhost},
				&board.Block{Color: l.color, Transparent: l.isGhost},
				&board.Block{Color: l.color, Transparent: l.isGhost},
			},
			[]*board.Block{nil, nil, nil},
		}
	}
	return nil
}

func (l *lPiece) SpawnGhost() Tetrimino {
	copy := lPiece{
		tetriminoBase: &tetriminoBase{
			orientation:     l.orientation,
			prevOrientation: l.prevOrientation,
			color:           l.color, // TODO: make different color to distinguish
			box:             l.box,
			isGhost:         true,
		},
	}
	return &copy
}

func (l *lPiece) RotationTests() []RotationTest {
	return defaultRotationTests(l, l.prevOrientation, l.pieceOrientation())
}
