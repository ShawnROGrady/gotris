package board

import "github.com/ShawnROGrady/gotris/pkg/canvas"

// Board represents the game board
type Board struct {
	background canvas.Color
	Blocks     [][]*Block
}

// New creates a new board
func New(background canvas.Color, width, height int) *Board {
	var blocks = [][]*Block{}

	for i := 0; i < height; i++ {
		row := make([]*Block, width)
		blocks = append(blocks, row)
	}

	return &Board{
		background: background,
		Blocks:     blocks,
	}
}

// Block represents a single block on the board
type Block struct {
	Color canvas.Color
}

func (b *Block) cell() *canvas.Cell {
	return &canvas.Cell{
		Background: b.Color,
	}
}

// Cells generates a visual representation of the board
func (b *Board) Cells() [][]*canvas.Cell {
	cells := [][]*canvas.Cell{}

	// reverse the rows
	for i := len(b.Blocks) - 1; i >= 0; i-- {
		row := []*canvas.Cell{}
		for _, block := range b.Blocks[i] {
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

// ClearFullRows checks if any rows are full and clears them if so
func (b *Board) ClearFullRows() {
	// TODO this should return a score
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