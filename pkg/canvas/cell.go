package canvas

const (
	block                  = "\u2588"
	lightTransparentBlock  = "\u2591"
	mediumTransparentBlock = "\u2592"
	darkTransparentBlock   = "\u2593"
)

// Cell represents a single cell on the canvas
type Cell struct {
	Color       Color
	Background  Color
	Transparent bool
}

func (c *Cell) String() string {
	if c.Transparent {
		return c.Background.background().decorate(
			// TODO: allow for level of transparency to be modifiable
			c.Color.decorate(mediumTransparentBlock),
		)
	}
	return c.Color.decorate(block)
}
