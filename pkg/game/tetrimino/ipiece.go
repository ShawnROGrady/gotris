package tetrimino

import (
	"github.com/ShawnROGrady/gotris/pkg/canvas"
	"github.com/ShawnROGrady/gotris/pkg/game/board"
)

type iPiece struct {
	box Box
}

func (i *iPiece) ContainingBox() Box {
	return i.box
}

func (i *iPiece) YMax() Coordinates {
	boxBottomRight := i.box.BottomRight

	var yMax = boxBottomRight

	for i, row := range i.Blocks() {
		for j, block := range row {
			if block == nil {
				continue
			}
			y := boxBottomRight.Y + i
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

	var yMin = boxTopLeft

	for i, row := range i.Blocks() {
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

	var xMax = boxTopLeft

	for i, row := range i.Blocks() {
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

	var xMin = boxBottomRight

	for i, row := range i.Blocks() {
		for j, block := range row {
			if block == nil {
				continue
			}
			y := boxBottomRight.Y + i
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
	return iBlocks
}

var iBlocks = [][]*board.Block{
	[]*board.Block{nil, nil, &board.Block{Color: canvas.Cyan}, nil},
	[]*board.Block{nil, nil, &board.Block{Color: canvas.Cyan}, nil},
	[]*board.Block{nil, nil, &board.Block{Color: canvas.Cyan}, nil},
	[]*board.Block{nil, nil, &board.Block{Color: canvas.Cyan}, nil},
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
