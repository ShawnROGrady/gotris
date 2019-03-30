package game

import "fmt"

const block = "\u2588"

type cell struct {
	background string
}

func (c *cell) String() string {
	return fmt.Sprintf("%s%s", c.background, block)
}
