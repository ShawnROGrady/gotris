package tetrimino

import (
	"github.com/ShawnROGrady/gotris/pkg/canvas"
	"github.com/ShawnROGrady/gotris/pkg/game/board"
)

type sPiece struct {
	*tetriminoBase
}

func newSPiece(boardWidth, boardHeight int) Tetrimino {
	spawnOrientation := spawn

	piece := &sPiece{
		tetriminoBase: &tetriminoBase{
			orientation:     &spawnOrientation,
			prevOrientation: spawnOrientation,
			color:           canvas.Green,
		},
	}

	box := startingBox(boardWidth, boardHeight, piece)
	piece.box = box
	return piece
}

func (s *sPiece) YMax() Coordinates {
	var (
		boxBottomRight = s.box.BottomRight
		blocks         = s.Blocks()
	)

	yMax := findMaxY(blocks, boxBottomRight)

	return yMax
}

func (s *sPiece) YMin() Coordinates {
	var (
		boxTopLeft = s.box.TopLeft
		blocks     = s.Blocks()
	)

	yMin := findMinY(blocks, boxTopLeft)

	return yMin
}

func (s *sPiece) XMax() Coordinates {
	var (
		boxTopLeft = s.box.TopLeft
		blocks     = s.Blocks()
	)

	xMax := findMaxX(blocks, boxTopLeft)

	return xMax
}

func (s *sPiece) XMin() Coordinates {
	var (
		boxBottomRight = s.box.BottomRight
		blocks         = s.Blocks()
	)

	xMin := findMinX(blocks, boxBottomRight)

	return xMin
}

func (s *sPiece) Blocks() [][]*board.Block {
	switch *s.orientation {
	case clockwise:
		return [][]*board.Block{
			[]*board.Block{
				nil,
				&board.Block{Color: s.color},
				nil,
			},
			[]*board.Block{
				nil,
				&board.Block{Color: s.color},
				&board.Block{Color: s.color},
			},
			[]*board.Block{
				nil,
				nil,
				&board.Block{Color: s.color},
			},
		}
	case opposite:
		return [][]*board.Block{
			[]*board.Block{nil, nil, nil},
			[]*board.Block{
				nil,
				&board.Block{Color: s.color},
				&board.Block{Color: s.color},
			},
			[]*board.Block{
				&board.Block{Color: s.color},
				&board.Block{Color: s.color},
				nil,
			},
		}
	case counterclockwise:
		return [][]*board.Block{
			[]*board.Block{
				&board.Block{Color: s.color},
				nil,
				nil,
			},
			[]*board.Block{
				&board.Block{Color: s.color},
				&board.Block{Color: s.color},
				nil,
			},
			[]*board.Block{
				nil,
				&board.Block{Color: s.color},
				nil,
			},
		}
	case spawn:
		return [][]*board.Block{
			[]*board.Block{
				nil,
				&board.Block{Color: s.color},
				&board.Block{Color: s.color},
			},
			[]*board.Block{
				&board.Block{Color: s.color},
				&board.Block{Color: s.color},
				nil,
			},
			[]*board.Block{nil, nil, nil},
		}
	}
	return nil
}

func (s *sPiece) SpawnGhost() Tetrimino {
	copy := sPiece{
		tetriminoBase: &tetriminoBase{
			orientation:     s.orientation,
			prevOrientation: s.prevOrientation,
			color:           s.color, // TODO: make different color to distinguish
			box:             s.box,
		},
	}
	return &copy
}

func (s *sPiece) RotationTests() []RotationTest {
	return defaultRotationTests(s, s.prevOrientation, s.pieceOrientation())
}
