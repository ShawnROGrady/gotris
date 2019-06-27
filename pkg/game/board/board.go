package board

import (
	"github.com/ShawnROGrady/gotris/pkg/canvas"
)

// Defaults for board
const (
	DefaultWidthScale = 2
	defaultHiddenRows = 4
)

// Board represents the game board
type Board struct {
	Background canvas.Color
	Blocks     [][]*Block
	HiddenRows int
	widthScale int
	Width      int
	Height     int
}

// New creates a new board
func New(opts ...Option) *Board {
	var blocks = [][]*Block{}

	b := &Board{
		Background: canvas.DefaultBackground,
		HiddenRows: defaultHiddenRows,
		widthScale: DefaultWidthScale,
		Width:      canvas.DefaultWidth / DefaultWidthScale,
		Height:     canvas.DefaultHeight,
	}

	for i := range opts {
		opts[i].ApplyToBoard(b)
	}

	for i := 0; i < b.Height+b.HiddenRows; i++ {
		row := make([]*Block, b.Width)
		blocks = append(blocks, row)
	}

	b.Blocks = blocks

	return b
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

	for i := 0; i < len(b); i++ {
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
	cells := [][]canvas.Cell{}
	// reverse the rows
	for i := len(activeBlocks) - 1; i >= 0; i-- {
		row := []canvas.Cell{}
		for _, block := range activeBlocks[i] {
			if block == nil {
				for i := 0; i < b.widthScale; i++ {
					row = append(row, &canvas.BlockCell{
						Color:      b.Background,
						Background: b.Background,
					})
				}
				continue
			}
			blockCell := block.cell()
			blockCell.Background = b.Background
			for i := 0; i < b.widthScale; i++ {
				row = append(row, blockCell)
			}
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
