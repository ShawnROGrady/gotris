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

func (b *board) clearFullRows() {
	// TODO this should return a score
	fullRows := b.checkRows()

	blocksPerRow := len(b.blocks[0])
	for i, fullRow := range fullRows {
		// all blocks in row non-nil->clear
		// subtract i to account for rows already cleared
		copy(b.blocks[fullRow-i:], b.blocks[fullRow+1-i:])

		// remove full row
		b.blocks = b.blocks[:len(b.blocks)-1]

		// insert empty row at top
		b.blocks = append(b.blocks, [][]*block{make([]*block, blocksPerRow)}...)
	}
}

func (b *board) checkRows() []int {
	fullRows := []int{}

	for i, row := range b.blocks {
		for j, cell := range row {
			if cell == nil {
				break
			}

			if j == len(row)-1 {
				fullRows = append(fullRows, i)
			}
		}
	}

	return fullRows
}
