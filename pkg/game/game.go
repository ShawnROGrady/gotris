package game

import (
	"os"

	"github.com/ShawnROGrady/gotris/pkg/canvas"
	"github.com/ShawnROGrady/gotris/pkg/inputreader"
)

// Game is responsible for handling the game state
type Game struct {
	inputreader  inputreader.InputReader
	canvas       *canvas.Canvas
	board        *board
	currentPiece tetrimino
}

// New returns a new game with the specified specifications
func New(term *os.File, width, height int) *Game {
	return &Game{
		inputreader: inputreader.NewTermReader(term),
		canvas: canvas.New(
			term,
			canvas.Green,
			width, height,
		),
		board: newBoard(
			canvas.Green,
			width, height,
		),
		currentPiece: &iPiece{
			box: box{
				topLeft: coordinates{
					x: 0,
					y: height - 1,
				},
				bottomRight: coordinates{
					x: 3,
					y: height - 4,
				},
			},
		},
	}
}

// RunDemo is a placeholder function to test core functionality
func (g *Game) RunDemo(done chan bool) chan error {
	input, readErr := translateInput(done, g.inputreader)
	runErr := make(chan error)

	blocks := g.currentPiece.blocks()
	topL := g.currentPiece.containingBox().topLeft

	// add initial piece to canvas
	for i, row := range blocks {
		for j, block := range row {
			if block == nil {
				continue
			}
			x := topL.x + j
			y := topL.y - i

			g.board.blocks[y][x] = block
		}
	}

	g.canvas.Cells = g.board.cells()

	// render initial canvas
	if err := g.canvas.Render(); err != nil {
		runErr <- err
		return runErr
	}

	go func() {
		for {
			select {
			case err := <-readErr:
				runErr <- err
				return
			case <-done:
				return
			case in := <-input:
				// TODO: print if in debug mode
				//log.Printf("User input: %s", in)

				topL := g.currentPiece.containingBox().topLeft

				// add initial piece to canvas

				if err := g.handleDemoInput(in); err != nil {
					runErr <- err
					return
				}

				newTopL := g.currentPiece.containingBox().topLeft
				newBottomR := g.currentPiece.containingBox().bottomRight

				// clear cell where piece was
				for i, row := range blocks {
					for j, block := range row {
						if block == nil {
							continue
						}
						x := topL.x + j
						y := topL.y - i

						g.board.blocks[y][x] = nil
					}
				}

				topL = newTopL
				bottomR := newBottomR

				// update cell at pieces new position
				blocks = g.currentPiece.blocks()
				for i, row := range blocks {
					for j, block := range row {
						if block == nil {
							continue
						}
						x := topL.x + j
						y := topL.y - i

						g.board.blocks[y][x] = block
					}
				}

				// generate new current piece if at bottom or on top of another piece
				if bottomR.y == 0 || g.board.blocks[g.currentPiece.yMin().y-1][g.currentPiece.yMin().x] != nil {
					// check if any rows can be cleared
					// TODO: add scoring
					g.board.clearFullRows()

					g.currentPiece = &iPiece{
						box: box{
							topLeft: coordinates{
								x: 0,
								y: len(g.board.blocks) - 1,
							},
							bottomRight: coordinates{
								x: 3,
								y: len(g.board.blocks) - 4,
							},
						},
					}

					topL = newTopL
					blocks = g.currentPiece.blocks()

					// add new piece to canvas
					for i, row := range blocks {
						for j, block := range row {
							if block == nil {
								continue
							}
							x := topL.x + j
							y := topL.y - i

							g.board.blocks[y][x] = block
						}
					}
				}

				g.canvas.Cells = g.board.cells()

				if err := g.canvas.Render(); err != nil {
					runErr <- err
					return
				}
			}
		}
	}()
	return runErr
}

func (g *Game) handleDemoInput(input userInput) error {
	g.currentPiece.move(input,
		len(g.board.blocks[0])-1,
		len(g.board.blocks)-1,
	)
	return nil
}
