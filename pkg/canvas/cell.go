package canvas

import "fmt"

const block = "\u2588"

// Cell represents a single cell on the canvas
type Cell struct {
	Background string
}

func (c *Cell) String() string {
	return fmt.Sprintf("%s%s", c.Background, block)
}
