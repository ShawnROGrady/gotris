package canvas

import (
	"fmt"
	"os"
)

// Canvas represents what is actually rendered to the user
type Canvas struct {
	Background string
	Cells      [][]*Cell
}

// New returns a new canvas
func New(background string, width, height int) *Canvas {
	var cells = [][]*Cell{}

	for i := 0; i < height; i++ {
		row := make([]*Cell, width)
		cells = append(cells, row)
	}

	return &Canvas{
		Background: background,
		Cells:      cells,
	}
}

// Render renders the current canvas
func (c *Canvas) Render(dest *os.File) error {
	// clear the canvas
	_, err := dest.WriteString("\033[2J")
	if err != nil {
		return err
	}

	_, err = dest.Seek(0, 0)
	if err != nil {
		return err
	}

	for _, row := range c.Cells {
		var buf = []byte{}
		for _, cell := range row {
			if cell == nil {
				buf = append(buf,
					[]byte(fmt.Sprintf("%s%s", c.Background, block))...,
				)
				continue
			}
			buf = append(buf, []byte(cell.String())...)
		}
		buf = append(buf, '\n')

		_, err := dest.Write(buf)
		if err != nil {
			return err
		}
	}

	return nil
}
