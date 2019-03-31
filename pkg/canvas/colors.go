package canvas

import "fmt"

// Color is used to alter the color of cells on the canvas
type Color int

// the available colors
const (
	Black Color = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

func (c Color) String() string {
	return fmt.Sprintf("\u001b[3%dm", c)
}

// Decorate formats the provided input so it can be printed in color
func (c Color) Decorate(input string) string {
	return fmt.Sprintf("%s%s", c, input)
}
