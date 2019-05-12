package canvas

import "fmt"

// Color is used to alter the color of cells on the canvas
type Color int

// the available colors
const (
	// normal colors
	Reset Color = iota
	Black Color = iota + 29
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White

	// bright colors
	BrightBlack
	BrightRed
	BrightGreen
	BrightYellow
	BrightBlue
	BrightMagenta
	BrightCyan
	BrightWhite

	// other colors
	Orange Color = 208
)

func (c Color) String() string {
	switch c {
	case BrightBlack, BrightRed, BrightGreen, BrightYellow, BrightBlue, BrightMagenta, BrightCyan, BrightWhite:
		return fmt.Sprintf("\u001b[%d;1m", c-8)
	case Orange:
		return fmt.Sprintf("\u001b[38;5;%dm", c)
	default:
		return fmt.Sprintf("\u001b[%dm", c)
	}
}

func (c Color) description() string {
	colorDescriptions := map[Color]string{
		Black:         "black",
		Red:           "red",
		Green:         "green",
		Yellow:        "yellow",
		Blue:          "blue",
		Magenta:       "magenta",
		Cyan:          "cyan",
		White:         "white",
		BrightBlack:   "bright black",
		BrightRed:     "bright red",
		BrightGreen:   "bright green",
		BrightYellow:  "bright yellow",
		BrightBlue:    "bright blue",
		BrightMagenta: "bright magenta",
		BrightCyan:    "bright cyan",
		BrightWhite:   "bright white",
		Orange:        "orange",
		Reset:         "reset",
	}

	return colorDescriptions[c]
}

// decorate formats the provided input so it can be printed in color
func (c Color) decorate(input string) string {
	return fmt.Sprintf("%s%s", c, input)
}
