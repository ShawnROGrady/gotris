package canvas

const block = "\u2588"

// Cell represents a single cell on the canvas
type Cell struct {
	Background Color
}

func (c *Cell) String() string {
	return c.Background.decorate(block)
}
