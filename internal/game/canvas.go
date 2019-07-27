package game

import (
	"sync"

	"github.com/ShawnROGrady/gotris/internal/canvas"
)

// gCanvas wraps a standard canvas
type gCanvas struct {
	canvas    canvas.Canvas
	newCells  chan [][]canvas.Cell
	mut       *sync.Mutex
	renderErr error
}

func (g *gCanvas) run(done <-chan bool) {
	go func() {
		for {
			select {
			case <-done:
				return
			case cells := <-g.newCells:
				g.mut.Lock()
				g.canvas.UpdateCells(cells)
				if err := g.canvas.Render(); err != nil {
					g.renderErr = err
				}
				g.mut.Unlock()
			}
		}
	}()
}

func (g *gCanvas) Init() error {
	return g.canvas.Init()
}

func (g *gCanvas) UpdateCells(newCells [][]canvas.Cell) {
	g.newCells <- newCells
}

func (g *gCanvas) Render() error {
	g.mut.Lock()
	defer g.mut.Unlock()
	return g.renderErr
}
