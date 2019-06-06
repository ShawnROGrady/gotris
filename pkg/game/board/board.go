package board

import (
	"github.com/ShawnROGrady/gotris/pkg/canvas"
)

// Board represents the game board
type Board struct {
	background canvas.Color
	Blocks     [][]*Block
	HiddenRows int
}

// New creates a new board
func New(background canvas.Color, width, height, hiddenRows int) *Board {
	var blocks = [][]*Block{}

	for i := 0; i < height+hiddenRows; i++ {
		row := make([]*Block, width)
		blocks = append(blocks, row)
	}

	return &Board{
		background: background,
		Blocks:     blocks,
		HiddenRows: hiddenRows,
	}
}

// Block represents a single block on the board
type Block struct {
	Color       canvas.Color
	Transparent bool
}

func (b *Block) cell() *canvas.Cell {
	return &canvas.Cell{
		Color:       b.Color,
		Transparent: b.Transparent,
	}
}

// Cells generates a visual representation of the board
func (b *Board) Cells() [][]*canvas.Cell {
	cells := [][]*canvas.Cell{}

	// reverse the rows
	for i := len(b.Blocks) - b.HiddenRows - 1; i >= 0; i-- {
		row := []*canvas.Cell{}
		for _, block := range b.Blocks[i] {
			if block == nil {
				row = append(row, []*canvas.Cell{
					{
						Color:      b.background,
						Background: b.background,
					},
					{
						Color:      b.background,
						Background: b.background,
					},
				}...)
				continue
			}
			blockCell := block.cell()
			blockCell.Background = b.background
			row = append(row, []*canvas.Cell{
				blockCell,
				blockCell,
			}...)
		}
		cells = append(cells, row)
	}
	return cells
}

// ClearFullRows checks if any rows are full and clears them if so
func (b *Board) ClearFullRows() int {
	fullRows := b.CheckRows()

	blocksPerRow := len(b.Blocks[0])
	for i, fullRow := range fullRows {
		// all blocks in row non-nil->clear
		// subtract i to account for rows already cleared
		copy(b.Blocks[fullRow-i:], b.Blocks[fullRow+1-i:])

		// remove full row
		b.Blocks = b.Blocks[:len(b.Blocks)-1]

		// insert empty row at top
		b.Blocks = append(b.Blocks, [][]*Block{make([]*Block, blocksPerRow)}...)
	}
	return len(fullRows)
}

// CheckRows checks if there are any full rows
func (b *Board) CheckRows() []int {
	fullRows := []int{}

	for i, row := range b.Blocks {
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
