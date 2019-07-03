package tetrimino

import (
	"github.com/ShawnROGrady/gotris/internal/canvas"
	"github.com/ShawnROGrady/gotris/internal/game/board"
)

type oPiece struct {
	*tetriminoBase
}

func newOPiece(boardWidth, boardHeight int) Tetrimino {
	spawnOrientation := spawn

	piece := &oPiece{
		tetriminoBase: &tetriminoBase{
			orientation: &spawnOrientation,
			color:       canvas.Yellow,
		},
	}

	box := startingBox(boardWidth, boardHeight, piece)
	piece.box = box
	return piece
}

func (o *oPiece) YMax() Coordinates {
	var (
		boxBottomRight = o.box.BottomRight
		blocks         = o.Blocks()
	)

	yMax := findMaxY(blocks, boxBottomRight)

	return yMax
}

func (o *oPiece) YMin() Coordinates {
	var (
		boxTopLeft = o.box.TopLeft
		blocks     = o.Blocks()
	)

	yMin := findMinY(blocks, boxTopLeft)

	return yMin
}

func (o *oPiece) XMax() Coordinates {
	var (
		boxTopLeft = o.box.TopLeft
		blocks     = o.Blocks()
	)

	xMax := findMaxX(blocks, boxTopLeft)

	return xMax
}

func (o *oPiece) XMin() Coordinates {
	var (
		boxBottomRight = o.box.BottomRight
		blocks         = o.Blocks()
	)

	xMin := findMinX(blocks, boxBottomRight)

	return xMin
}

func (o *oPiece) Blocks() [][]*board.Block {
	switch *o.orientation {
	case spawn:
		return [][]*board.Block{
			[]*board.Block{nil, &board.Block{Color: o.color, Transparent: o.isGhost}, &board.Block{Color: o.color, Transparent: o.isGhost}, nil},
			[]*board.Block{nil, &board.Block{Color: o.color, Transparent: o.isGhost}, &board.Block{Color: o.color, Transparent: o.isGhost}, nil},
			[]*board.Block{nil, nil, nil, nil},
		}
	case clockwise:
		return [][]*board.Block{
			[]*board.Block{nil, &board.Block{Color: o.color, Transparent: o.isGhost}, &board.Block{Color: o.color, Transparent: o.isGhost}, nil},
			[]*board.Block{nil, &board.Block{Color: o.color, Transparent: o.isGhost}, &board.Block{Color: o.color, Transparent: o.isGhost}, nil},
			[]*board.Block{nil, nil, nil, nil},
		}
	case opposite:
		return [][]*board.Block{
			[]*board.Block{nil, &board.Block{Color: o.color, Transparent: o.isGhost}, &board.Block{Color: o.color, Transparent: o.isGhost}, nil},
			[]*board.Block{nil, &board.Block{Color: o.color, Transparent: o.isGhost}, &board.Block{Color: o.color, Transparent: o.isGhost}, nil},
			[]*board.Block{nil, nil, nil, nil},
		}
	case counterclockwise:
		return [][]*board.Block{
			[]*board.Block{nil, &board.Block{Color: o.color, Transparent: o.isGhost}, &board.Block{Color: o.color, Transparent: o.isGhost}, nil},
			[]*board.Block{nil, &board.Block{Color: o.color, Transparent: o.isGhost}, &board.Block{Color: o.color, Transparent: o.isGhost}, nil},
			[]*board.Block{nil, nil, nil, nil},
		}
	}
	return nil
}

func (o *oPiece) SpawnGhost() Tetrimino {
	copy := oPiece{
		tetriminoBase: &tetriminoBase{
			orientation:     o.orientation,
			prevOrientation: o.prevOrientation,
			color:           o.color, // TODO: make different color to distinguish
			box:             o.box,
			isGhost:         true,
		},
	}
	return &copy
}

func (o *oPiece) RotationTests() []RotationTest {
	// oPiece can't be resolved in case of conflicts
	return []RotationTest{}
}
