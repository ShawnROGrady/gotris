package game

import "github.com/ShawnROGrady/gotris/pkg/canvas"

type board struct {
	background string
	blocks     [][]*block
}

func newBoard(background string, width, height int) *board {
	var blocks = [][]*block{}

	for i := 0; i < height; i++ {
		row := make([]*block, width)
		blocks = append(blocks, row)
	}

	return &board{
		background: background,
		blocks:     blocks,
	}
}

type block struct {
	color string
}

func (b *block) cell() *canvas.Cell {
	return &canvas.Cell{
		Background: b.color,
	}
}

func (b *board) canvas() *canvas.Canvas {
	canv := canvas.Canvas{}

	// reverse the rows
	for i := len(b.blocks) - 1; i >= 0; i-- {
		row := []*canvas.Cell{}
		for _, block := range b.blocks[i] {
			if block == nil {
				row = append(row, &canvas.Cell{
					Background: b.background,
				})
				continue
			}
			row = append(row, block.cell())
		}
		canv.Cells = append(canv.Cells, row)
	}
	return &canv
}
