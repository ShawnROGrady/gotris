package game

import "github.com/ShawnROGrady/gotris/pkg/canvas"

type coordinates struct {
	x int
	y int
}

type piece struct {
	color       canvas.Color
	coordinates coordinates
}

func (p *piece) move(direction userInput, xmax, ymax int) {
	// TODO: make direction its own type
	switch direction {
	case moveLeft:
		// left
		if p.coordinates.x > 0 {
			p.coordinates.x -= 1
		}
	case moveDown:
		// down
		if p.coordinates.y > 0 {
			p.coordinates.y -= 1
		}
	case moveUp:
		// up
		if p.coordinates.y < ymax {
			p.coordinates.y += 1
		}
	case moveRight:
		// right
		if p.coordinates.x < xmax {
			p.coordinates.x += 1
		}
	}
}
