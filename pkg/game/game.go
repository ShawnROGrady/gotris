package game

import (
	"log"
	"os"
)

// Game is responsible for handling the game state
type Game struct {
	term   *os.File
	canvas *canvas
}

// New returns a new game with the specified specifications
func New(term *os.File, width, height int) *Game {
	return &Game{
		term: term,
		canvas: newCanvas(
			"\u001b[32m", // Green
			width, height,
		),
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

				if err := g.canvas.render(g.term); err != nil {
					runErr <- err
				}
			}
		}
	}()
	return runErr
}

func (g *Game) handleDemoInput(input string) error {
	switch input {
	case "j":
		g.canvas.cells[0][0] = &cell{"\u001b[31m"} // Red
		g.canvas.cells[0][1] = &cell{"\u001b[32m"} // Green
	case "k":
		g.canvas.cells[0][1] = &cell{"\u001b[31m"} // Red
		g.canvas.cells[0][0] = &cell{"\u001b[32m"} // Green
	default:
		log.Printf("unhandled input: %s", input)
	}

	return nil
}
