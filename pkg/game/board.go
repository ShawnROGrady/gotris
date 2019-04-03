package game

import "github.com/ShawnROGrady/gotris/pkg/canvas"

type board struct {
	background canvas.Color
	blocks     [][]*block
}

func newBoard(background canvas.Color, width, height int) *board {
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
	color canvas.Color
}

func (b *block) cell() *canvas.Cell {
	return &canvas.Cell{
		Background: b.color,
	}
}

func (b *board) cells() [][]*canvas.Cell {
	cells := [][]*canvas.Cell{}

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
		cells = append(cells, row)
	}
	return cells
}

func (b *board) checkRows() {
	for i, row := range b.blocks {
		for j, cell := range row {
			if cell == nil {
				break
			}

			if j == len(row)-1 {
				// all blocks in row non-nil->clear
				copy(b.blocks[i:], b.blocks[i+1:])

				// remove full row
				b.blocks = b.blocks[:len(b.blocks)-1]

				// insert empty row at top
				b.blocks = append(b.blocks, [][]*block{make([]*block, j+1)}...)
			}
		}
	}

}
