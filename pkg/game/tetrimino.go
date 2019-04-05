package game

import "github.com/ShawnROGrady/gotris/pkg/canvas"

type tetrimino interface {
	blocks() [][]*block
	move(direction userInput, xmax, ymax int)
	containingBox() box
	// used for detecting collisions
	yMax() coordinates
	yMin() coordinates
	xMax() coordinates
	xMin() coordinates
}

// NewPiece generates a new tetrimino
func NewPiece(boardWidth, boardHeight int) tetrimino {
	// TODO: should randomly generate piece type
	return &iPiece{
		box: box{
			topLeft: coordinates{
				x: 0,
				y: boardHeight - 1,
			},
			bottomRight: coordinates{
				x: 3,
				y: boardHeight - 4,
			},
		},
	}
}

// each piece can be thought of as being contained in a box
// this way we don't have to track the coordinates of each block
type box struct {
	topLeft     coordinates
	bottomRight coordinates
}

type iPiece struct {
	box box
}

func (i *iPiece) containingBox() box {
	return i.box
}

func (i *iPiece) yMax() coordinates {
	boxBottomRight := i.box.bottomRight

	var yMax = boxBottomRight

	for i, row := range i.blocks() {
		for j, block := range row {
			if block == nil {
				continue
			}
			y := boxBottomRight.y + i
			x := boxBottomRight.x - (len(row) - 1) + j

			if y > yMax.y {
				yMax = coordinates{
					x: x,
					y: y,
				}
			}
		}
	}

	return yMax
}

func (i *iPiece) yMin() coordinates {
	boxTopLeft := i.box.topLeft

	var yMin = boxTopLeft

	for i, row := range i.blocks() {
		for j, block := range row {
			if block == nil {
				continue
			}
			y := boxTopLeft.y - i
			x := boxTopLeft.x + j

			if y < yMin.y {
				yMin = coordinates{
					x: x,
					y: y,
				}
			}
		}
	}

	return yMin
}

func (i *iPiece) xMax() coordinates {
	boxTopLeft := i.box.topLeft

	var xMax = boxTopLeft

	for i, row := range i.blocks() {
		for j, block := range row {
			if block == nil {
				continue
			}
			y := boxTopLeft.y - i
			x := boxTopLeft.x + j

			if x > xMax.x {
				xMax = coordinates{
					x: x,
					y: y,
				}
			}
		}
	}

	return xMax
}

func (i *iPiece) xMin() coordinates {
	boxBottomRight := i.box.bottomRight

	var xMin = boxBottomRight

	for i, row := range i.blocks() {
		for j, block := range row {
			if block == nil {
				continue
			}
			y := boxBottomRight.y + i
			x := boxBottomRight.x - (len(row) - 1) + j

			if x < xMin.x {
				xMin = coordinates{
					x: x,
					y: y,
				}
			}
		}
	}

	return xMin
}

func (i *iPiece) blocks() [][]*block {
	return iBlocks
}

var iBlocks = [][]*block{
	[]*block{nil, nil, &block{canvas.Cyan}, nil},
	[]*block{nil, nil, &block{canvas.Cyan}, nil},
	[]*block{nil, nil, &block{canvas.Cyan}, nil},
	[]*block{nil, nil, &block{canvas.Cyan}, nil},
}

func (i *iPiece) move(direction userInput, xmax, ymax int) {
	switch direction {
	case moveLeft:
		// left
		if i.xMin().x > 0 {
			i.box.topLeft.x -= 1
			i.box.bottomRight.x -= 1
		}
	case moveDown:
		// down
		if i.yMin().y > 0 {
			i.box.bottomRight.y -= 1
			i.box.topLeft.y -= 1
		}
	case moveUp:
		// up
		if i.yMax().y < ymax {
			i.box.topLeft.y += 1
			i.box.bottomRight.y += 1
		}
	case moveRight:
		// right
		if i.xMax().x < xmax {
			i.box.bottomRight.x += 1
			i.box.topLeft.x += 1
		}
	}
}
