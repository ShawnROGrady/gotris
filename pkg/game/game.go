package game

import (
	"log"
	"os"
)

// Game is responsible for handling the game state
type Game struct {
	term         *os.File
	canvas       *canvas
	currentPiece *piece
}

// New returns a new game with the specified specifications
func New(term *os.File, width, height int) *Game {
	return &Game{
		term: term,
		canvas: newCanvas(
			"\u001b[32m", // Green
			width, height,
		),
		currentPiece: &piece{
			color: "\u001b[34m", //Blue
		},
	}
}

// RunDemo is a placeholder function to test core functionality
func (g *Game) RunDemo(done chan bool) chan error {
	input, readErr := readInput(done, g.term)
	runErr := make(chan error)

	// render initial canvas
	if err := g.canvas.render(g.term); err != nil {
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
				if err := g.handleDemoInput(string(in)); err != nil {
					runErr <- err
				}

				coords := g.currentPiece.coordinates

				g.canvas.cells[coords.y][coords.x] = &cell{
					background: g.currentPiece.color,
				}

				if err := g.canvas.render(g.term); err != nil {
					runErr <- err
				}
			}
		}
	}()
	return runErr
}

func (g *Game) handleDemoInput(input string) error {
	g.currentPiece.move(input,
		len(g.canvas.cells[0])-1,
		len(g.canvas.cells)-1,
	)
	return nil
}
