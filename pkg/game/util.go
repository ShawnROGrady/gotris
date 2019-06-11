package game

import "github.com/ShawnROGrady/gotris/pkg/game/board"

// centerBlocks places the provided blocks in a grid of specified width and height
func centerBlocks(blocks [][]*board.Block, width, height int) [][]*board.Block {
	newBlocks := make([][]*board.Block, height)
	startingX := (width - len(blocks[0])) / 2
	startingY := (height - len(blocks)) / 2

	for i := range newBlocks {
		row := make([]*board.Block, width)
		if i < startingY || i > len(blocks)-1 {
			newBlocks[i] = row
			continue
		}
		for j := 0; j < width; j++ {
			if j < startingX || j > len(blocks[0])-1 {
				continue
			}
			row[j] = blocks[i-startingY][j-startingX]
		}
		newBlocks[i] = row
	}
	return newBlocks
}
