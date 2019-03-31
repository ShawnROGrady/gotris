package game

type coordinates struct {
	x int
	y int
}

type piece struct {
	color       string
	coordinates coordinates
}

func (p *piece) move(direction string, xmax, ymax int) {
	// TODO: make direction its own type
	switch direction {
	case "h":
		// left
		if p.coordinates.x > 0 {
			p.coordinates.x -= 1
		}
	case "j":
		// down
		if p.coordinates.y > 0 {
			p.coordinates.y -= 1
		}
	case "k":
		// up
		if p.coordinates.y < ymax {
			p.coordinates.y += 1
		}
	case "l":
		// right
		if p.coordinates.x < xmax {
			p.coordinates.x += 1
		}
	}
}
