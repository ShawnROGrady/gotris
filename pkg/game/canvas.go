package game

import "os"

type canvas struct {
	width      int
	height     int
	background string
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

	for i := 0; i < c.height; i++ {
		var buf = []byte{}
		for j := 0; j < c.width; j++ {
			buf = append(buf, []byte(c.background)...)
		}
		buf = append(buf, '\n')

		_, err := dest.Write(buf)
		if err != nil {
			return err
		}
	}

	return nil
}
