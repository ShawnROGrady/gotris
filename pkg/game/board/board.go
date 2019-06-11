package board

import (
	"github.com/ShawnROGrady/gotris/pkg/canvas"
)

// Board represents the game board
type Board struct {
	Background canvas.Color
	Blocks     [][]*Block
	HiddenRows int
	widthScale int
}

// New creates a new board
func New(background canvas.Color, width, height, hiddenRows, widthScale int) *Board {
	var blocks = [][]*Block{}

	for i := 0; i < height+hiddenRows; i++ {
		row := make([]*Block, width)
		blocks = append(blocks, row)
	}

	return &Board{
		Background: background,
		Blocks:     blocks,
		HiddenRows: hiddenRows,
		widthScale: widthScale,
	}
}

// Block represents a single block on the board
type Block struct {
	Color       canvas.Color
	Transparent bool
}

func (b *Block) cell() *canvas.BlockCell {
	return &canvas.BlockCell{
		Color:       b.Color,
		Transparent: b.Transparent,
	}
}

// BlockGridCells generates the visual representation of a grid of cells
func BlockGridCells(b [][]*Block, background canvas.Color, widthScale int) [][]canvas.Cell {
	cells := [][]canvas.Cell{}

	// reverse the rows
	for i := len(b) - 1; i >= 0; i-- {
		row := []canvas.Cell{}
		for _, block := range b[i] {
			if block == nil {
				for i := 0; i < widthScale; i++ {
					row = append(row, &canvas.BlockCell{
						Color:      background,
						Background: background,
					})
				}
				continue
			}
			blockCell := block.cell()
			blockCell.Background = background
			for i := 0; i < widthScale; i++ {
				row = append(row, blockCell)
			}
		}
		cells = append(cells, row)
	}
	return cells
}

// Cells generates a visual representation of the board
func (b *Board) Cells() [][]canvas.Cell {
	activeBlocks := b.Blocks[:len(b.Blocks)-b.HiddenRows]
	return BlockGridCells(activeBlocks, b.Background, b.widthScale)
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
