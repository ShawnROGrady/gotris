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
		canvas: &canvas{
			width:      width,
			height:     height,
			background: "\u001b[32m\u2588", // Green
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
		g.canvas.background = "\u001b[32m\u2588" // Green
	case "k":
		g.canvas.background = "\u001b[31m\u2588" // Red
	default:
		log.Printf("unhandled input: %s", input)
	}

	return nil
}
