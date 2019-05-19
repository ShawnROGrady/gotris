package tetrimino

import (
	"github.com/ShawnROGrady/gotris/pkg/canvas"
	"github.com/ShawnROGrady/gotris/pkg/game/board"
)

type zPiece struct {
	*tetriminoBase
}

func newZPiece(boardWidth, boardHeight int) Tetrimino {
	spawnOrientation := spawn

	piece := &zPiece{
		tetriminoBase: &tetriminoBase{
			orientation:     &spawnOrientation,
			prevOrientation: spawnOrientation,
			color:           canvas.Red,
		},
	}

	box := startingBox(boardWidth, boardHeight, piece)
	piece.box = box
	return piece
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
	case clockwise:
		return [][]*board.Block{
			[]*board.Block{
				nil,
				nil,
				&board.Block{Color: z.color},
			},
			[]*board.Block{
				nil,
				&board.Block{Color: z.color},
				&board.Block{Color: z.color},
			},
			[]*board.Block{
				nil,
				&board.Block{Color: z.color},
				nil,
			},
		}
	case opposite:
		return [][]*board.Block{
			[]*board.Block{nil, nil, nil},
			[]*board.Block{
				&board.Block{Color: z.color},
				&board.Block{Color: z.color},
				nil,
			},
			[]*board.Block{
				nil,
				&board.Block{Color: z.color},
				&board.Block{Color: z.color},
			},
		}
	case counterclockwise:
		return [][]*board.Block{
			[]*board.Block{
				nil,
				&board.Block{Color: z.color},
				nil,
			},
			[]*board.Block{
				&board.Block{Color: z.color},
				&board.Block{Color: z.color},
				nil,
			},
			[]*board.Block{
				&board.Block{Color: z.color},
				nil,
				nil,
			},
		}
	case spawn:
		return [][]*board.Block{
			[]*board.Block{
				&board.Block{Color: z.color},
				&board.Block{Color: z.color},
				nil,
			},
			[]*board.Block{
				nil,
				&board.Block{Color: z.color},
				&board.Block{Color: z.color},
			},
			[]*board.Block{nil, nil, nil},
		}
	}
	return nil
}

func (z *zPiece) SpawnGhost() Tetrimino {
	copy := zPiece{
		tetriminoBase: &tetriminoBase{
			orientation:     z.orientation,
			prevOrientation: z.prevOrientation,
			color:           z.color, // TODO: make different color to distinguish
			box:             z.box,
		},
	}
	return &copy
}

func (z *zPiece) RotationTests() []RotationTest {
	return defaultRotationTests(z, z.prevOrientation, z.pieceOrientation())
}
