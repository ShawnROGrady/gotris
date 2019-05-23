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
				&board.Block{Color: j.color, Transparent: j.isGhost},
				&board.Block{Color: j.color, Transparent: j.isGhost},
			},
			[]*board.Block{nil, &board.Block{Color: j.color, Transparent: j.isGhost}, nil},
			[]*board.Block{nil, &board.Block{Color: j.color, Transparent: j.isGhost}, nil},
		}
	case opposite:
		return [][]*board.Block{
			[]*board.Block{nil, nil, nil},
			[]*board.Block{
				&board.Block{Color: j.color, Transparent: j.isGhost},
				&board.Block{Color: j.color, Transparent: j.isGhost},
				&board.Block{Color: j.color, Transparent: j.isGhost},
			},
			[]*board.Block{
				nil,
				nil,
				&board.Block{Color: j.color, Transparent: j.isGhost},
			},
		}
	case counterclockwise:
		return [][]*board.Block{
			[]*board.Block{nil, &board.Block{Color: j.color, Transparent: j.isGhost}, nil},
			[]*board.Block{nil, &board.Block{Color: j.color, Transparent: j.isGhost}, nil},
			[]*board.Block{
				&board.Block{Color: j.color, Transparent: j.isGhost},
				&board.Block{Color: j.color, Transparent: j.isGhost},
				nil,
			},
		}
	case spawn:
		return [][]*board.Block{
			[]*board.Block{
				&board.Block{Color: j.color, Transparent: j.isGhost},
				nil,
				nil,
			},
			[]*board.Block{
				&board.Block{Color: j.color, Transparent: j.isGhost},
				&board.Block{Color: j.color, Transparent: j.isGhost},
				&board.Block{Color: j.color, Transparent: j.isGhost},
			},
			[]*board.Block{nil, nil, nil},
		}
	}
	return nil
}

func (j *jPiece) SpawnGhost() Tetrimino {
	copy := jPiece{
		tetriminoBase: &tetriminoBase{
			orientation:     j.orientation,
			prevOrientation: j.prevOrientation,
			color:           j.color, // TODO: make different color to distinguish
			box:             j.box,
			isGhost:         true,
		},
	}
	return &copy
}

func (j *jPiece) RotationTests() []RotationTest {
	return defaultRotationTests(j, j.prevOrientation, j.pieceOrientation())
}
