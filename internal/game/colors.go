package game

import (
	"github.com/ShawnROGrady/gotris/internal/canvas"
	"github.com/ShawnROGrady/gotris/internal/game/board"
	"github.com/ShawnROGrady/gotris/internal/game/tetrimino"
)

// DisplayPotentialColors prints a demo board with the potential pieces and colors
// This is used to assist the user in selecting options
// e.g. WithoutGhost() should be applied if ghost pieces render oddly in the canvas
func (g *Game) DisplayPotentialColors() error {
	var can canvas.Canvas
	if gCanvas, ok := g.canvas.(*gCanvas); ok {
		can = gCanvas.canvas
	} else {
		can = g.canvas
	}

	if err := can.Init(); err != nil {
		return err
	}

	boardWidth := boardWidth(g.board)
	boardBlocks := [][]*board.Block{}

	for i := range tetrimino.PieceConstructors {
		var (
			piece       = tetrimino.PieceConstructors[i](boardWidth, 4)
			pieceBlocks = piece.Blocks()
			ghostBlocks [][]*board.Block
		)
		if !g.disableGhost {
			ghost := piece.SpawnGhost()
			ghostBlocks = ghost.Blocks()
		}
		for i := range pieceBlocks {
			row := make([]*board.Block, boardWidth)
			for j := range pieceBlocks[i] {
				row[j] = pieceBlocks[i][j]
				if !g.disableGhost {
					row[j+(boardWidth/2)] = ghostBlocks[i][j]
				}
			}
			boardBlocks = append(boardBlocks, row)
		}
	}
	g.board.Blocks = boardBlocks
	g.updateCells(g.board.Background())
	cells := g.cells(g.board)
	can.UpdateCells(cells)

	return can.Render()
}
