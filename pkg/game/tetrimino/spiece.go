package tetrimino

import (
	"github.com/ShawnROGrady/gotris/pkg/canvas"
	"github.com/ShawnROGrady/gotris/pkg/game/board"
)

type sPiece struct {
	box         Box
	orientation *orientation
}

func newSPiece(boardWidth, boardHeight int) Tetrimino {
	spawnOrientation := spawn

	piece := &sPiece{
		orientation: &spawnOrientation,
	}

	box := startingBox(boardWidth, boardHeight, piece)
	piece.box = box
	return piece
}

func (s sPiece) pieceOrientation() orientation {
	return *s.orientation
}

func (s *sPiece) ContainingBox() Box {
	return s.box
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
				&board.Block{Color: canvas.Green},
				nil,
			},
			[]*board.Block{
				nil,
				&board.Block{Color: canvas.Green},
				&board.Block{Color: canvas.Green},
			},
			[]*board.Block{
				nil,
				nil,
				&board.Block{Color: canvas.Green},
			},
		}
	case opposite:
		return [][]*board.Block{
			[]*board.Block{nil, nil, nil},
			[]*board.Block{
				nil,
				&board.Block{Color: canvas.Green},
				&board.Block{Color: canvas.Green},
			},
			[]*board.Block{
				&board.Block{Color: canvas.Green},
				&board.Block{Color: canvas.Green},
				nil,
			},
		}
	case counterclockwise:
		return [][]*board.Block{
			[]*board.Block{
				&board.Block{Color: canvas.Green},
				nil,
				nil,
			},
			[]*board.Block{
				&board.Block{Color: canvas.Green},
				&board.Block{Color: canvas.Green},
				nil,
			},
			[]*board.Block{
				nil,
				&board.Block{Color: canvas.Green},
				nil,
			},
		}
	case spawn:
		return [][]*board.Block{
			[]*board.Block{
				nil,
				&board.Block{Color: canvas.Green},
				&board.Block{Color: canvas.Green},
			},
			[]*board.Block{
				&board.Block{Color: canvas.Green},
				&board.Block{Color: canvas.Green},
				nil,
			},
			[]*board.Block{nil, nil, nil},
		}
	}
	return nil
}

func (s *sPiece) MoveUp(ymax int) {
	if s.YMax().Y < ymax {
		s.box.TopLeft.Y++
		s.box.BottomRight.Y++
	}
}

func (s *sPiece) MoveDown() {
	if s.YMin().Y > 0 {
		s.box.BottomRight.Y--
		s.box.TopLeft.Y--
	}
}

func (s *sPiece) MoveLeft() {
	if s.XMin().X > 0 {
		s.box.TopLeft.X--
		s.box.BottomRight.X--
	}
}

func (s *sPiece) MoveRight(xmax int) {
	if s.XMax().X < xmax {
		s.box.BottomRight.X++
		s.box.TopLeft.X++
	}
}

func (s *sPiece) RotateClockwise() {
	s.orientation.rotateClockwise()
}

func (s *sPiece) RotateCounter() {
	s.orientation.rotateCounter()
}
