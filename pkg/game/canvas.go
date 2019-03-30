package game

import "os"

type canvas struct {
	background string
	cells      [][]*cell
}

func newCanvas(background string, width, height int) *canvas {
	var cells = [][]*cell{}

	for i := 0; i < height; i++ {
		row := []*cell{}
		for j := 0; j < width; j++ {
			row = append(row, &cell{
				background: background,
			})
		}
		cells = append(cells, row)
	}

	return &canvas{
		background: background,
		cells:      cells,
	}
}

func (c *canvas) render(dest *os.File) error {
	// clear the canvas
	_, err := dest.WriteString("\033[2J")
	if err != nil {
		return err
	}

	_, err = dest.Seek(0, 0)
	if err != nil {
		return err
	}

	for _, row := range c.cells {
		var buf = []byte{}
		for _, cell := range row {
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
