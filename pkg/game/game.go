package game

import (
	"log"
	"os"

	"github.com/ShawnROGrady/gotris/pkg/canvas"
	"github.com/ShawnROGrady/gotris/pkg/inputreader"
)

// Game is responsible for handling the game state
type Game struct {
	inputreader  inputreader.InputReader
	canvas       *canvas.Canvas
	board        *board
	currentPiece *piece
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
		currentPiece: &piece{
			color: canvas.Blue,
			coordinates: coordinates{
				x: 0,
				y: height - 1,
			},
		},
	}
}

// RunDemo is a placeholder function to test core functionality
func (g *Game) RunDemo(done chan bool) chan error {
	input, readErr := g.inputreader.ReadInput(done)
	runErr := make(chan error)

	coords := g.currentPiece.coordinates
	// add initial piece to canvas
	g.board.blocks[coords.y][coords.x] = &block{
		color: g.currentPiece.color,
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
			case <-done:
				return
			case in := <-input:
				log.Printf("User input: %s", in)

				coords := g.currentPiece.coordinates

				// clear cell where piece was
				g.board.blocks[coords.y][coords.x] = nil

				if err := g.handleDemoInput(string(in)); err != nil {
					runErr <- err
				}

				coords = g.currentPiece.coordinates

				// update cell at pieces new position
				g.board.blocks[coords.y][coords.x] = &block{
					color: g.currentPiece.color,
				}

				// generate new current piece if at bottom
				if coords.y == 0 {
					g.currentPiece = &piece{
						color: canvas.Blue,
						coordinates: coordinates{
							x: 0,
							y: len(g.board.blocks) - 1,
						},
					}

					coords = g.currentPiece.coordinates

					// add new piece to canvas
					g.board.blocks[coords.y][coords.x] = &block{
						color: g.currentPiece.color,
					}
				}

				g.canvas.Cells = g.board.cells()

				if err := g.canvas.Render(); err != nil {
					runErr <- err
				}
			}
		}
	}()
	return runErr
}

func (g *Game) handleDemoInput(input string) error {
	g.currentPiece.move(input,
		len(g.board.blocks[0])-1,
		len(g.board.blocks)-1,
	)
	return nil
}
