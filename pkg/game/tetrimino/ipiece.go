package tetrimino

import (
	"github.com/ShawnROGrady/gotris/pkg/canvas"
	"github.com/ShawnROGrady/gotris/pkg/game/board"
)

type iPiece struct {
	box         Box
	orientation *orientation
}

func (i *iPiece) ContainingBox() Box {
	return i.box
}

func (i *iPiece) YMax() Coordinates {
	boxBottomRight := i.box.BottomRight

	var (
		yMax   = boxBottomRight
		blocks = i.Blocks()
	)

	for i, row := range blocks {
		for j, block := range row {
			if block == nil {
				continue
			}
			y := boxBottomRight.Y + (len(blocks) - 1) - i
			x := boxBottomRight.X - (len(row) - 1) + j

			if y > yMax.Y {
				yMax = Coordinates{
					X: x,
					Y: y,
				}
			}
		}
	}

	return yMax
}

func (i *iPiece) YMin() Coordinates {
	boxTopLeft := i.box.TopLeft

	var (
		yMin   = boxTopLeft
		blocks = i.Blocks()
	)

	for i, row := range blocks {
		for j, block := range row {
			if block == nil {
				continue
			}
			y := boxTopLeft.Y - i
			x := boxTopLeft.X + j

			if y < yMin.Y {
				yMin = Coordinates{
					X: x,
					Y: y,
				}
			}
		}
	}

	return yMin
}

func (i *iPiece) XMax() Coordinates {
	boxTopLeft := i.box.TopLeft

	var (
		xMax   = boxTopLeft
		blocks = i.Blocks()
	)

	for i, row := range blocks {
		for j, block := range row {
			if block == nil {
				continue
			}
			y := boxTopLeft.Y - i
			x := boxTopLeft.X + j

			if x > xMax.X {
				xMax = Coordinates{
					X: x,
					Y: y,
				}
			}
		}
	}

	return xMax
}

func (i *iPiece) XMin() Coordinates {
	boxBottomRight := i.box.BottomRight

	var (
		xMin   = boxBottomRight
		blocks = i.Blocks()
	)

	for i, row := range blocks {
		for j, block := range row {
			if block == nil {
				continue
			}
			y := boxBottomRight.Y + (len(blocks) - 1) - i
			x := boxBottomRight.X - (len(row) - 1) + j

			if x < xMin.X {
				xMin = Coordinates{
					X: x,
					Y: y,
				}
			}
		}
	}

	return xMin
}

func (i *iPiece) Blocks() [][]*board.Block {
	switch *i.orientation {
	case spawn:
		return [][]*board.Block{
			[]*board.Block{nil, nil, &board.Block{Color: canvas.Cyan}, nil},
			[]*board.Block{nil, nil, &board.Block{Color: canvas.Cyan}, nil},
			[]*board.Block{nil, nil, &board.Block{Color: canvas.Cyan}, nil},
			[]*board.Block{nil, nil, &board.Block{Color: canvas.Cyan}, nil},
		}
	case clockwise:
		return [][]*board.Block{
			[]*board.Block{nil, nil, nil, nil},
			[]*board.Block{nil, nil, nil, nil},
			[]*board.Block{
				&board.Block{Color: canvas.Cyan},
				&board.Block{Color: canvas.Cyan},
				&board.Block{Color: canvas.Cyan},
				&board.Block{Color: canvas.Cyan},
			},
			[]*board.Block{nil, nil, nil, nil},
		}
	case opposite:
		return [][]*board.Block{
			[]*board.Block{nil, &board.Block{Color: canvas.Cyan}, nil, nil},
			[]*board.Block{nil, &board.Block{Color: canvas.Cyan}, nil, nil},
			[]*board.Block{nil, &board.Block{Color: canvas.Cyan}, nil, nil},
			[]*board.Block{nil, &board.Block{Color: canvas.Cyan}, nil, nil},
		}
	case counterclockwise:
		return [][]*board.Block{
			[]*board.Block{nil, nil, nil, nil},
			[]*board.Block{
				&board.Block{Color: canvas.Cyan},
				&board.Block{Color: canvas.Cyan},
				&board.Block{Color: canvas.Cyan},
				&board.Block{Color: canvas.Cyan},
			},
			[]*board.Block{nil, nil, nil, nil},
			[]*board.Block{nil, nil, nil, nil},
		}

	}
	return nil
}

func (i *iPiece) MoveUp(ymax int) {
	if i.YMax().Y < ymax {
		i.box.TopLeft.Y++
		i.box.BottomRight.Y++
	}
}

func (i *iPiece) MoveDown() {
	if i.YMin().Y > 0 {
		i.box.BottomRight.Y--
		i.box.TopLeft.Y--
	}
}

func (i *iPiece) MoveLeft() {
	if i.XMin().X > 0 {
		i.box.TopLeft.X--
		i.box.BottomRight.X--
	}
}

func (i *iPiece) MoveRight(xmax int) {
	if i.XMax().X < xmax {
		i.box.BottomRight.X++
		i.box.TopLeft.X++
	}
}

func (i *iPiece) RotateClockwise() {
	i.orientation.rotateClockwise()
}

func (i *iPiece) RotateCounter() {
	i.orientation.rotateCounter()
}
