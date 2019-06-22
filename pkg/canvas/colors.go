package canvas

import (
	"strconv"
	"strings"
)

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

	// background colors
	BackgroundBlack
	BackgroundRed
	BackgroundGreen
	BackgroundYellow
	BackgroundBlue
	BackgroundMagenta
	BackgroundCyan
	BackgroundWhite

	// other colors
	Orange           Color = 208
	BackgroundOrange Color = -208
)

var resetControl = []byte{'\u001b', '[', '0', 'm'}

func (c Color) String() string {
	var b strings.Builder
	b.Grow(5)
	b.WriteString("\u001b[")
	switch c {
	case BrightBlack, BrightRed, BrightGreen, BrightYellow, BrightBlue, BrightMagenta, BrightCyan, BrightWhite:
		b.WriteString(strconv.Itoa(int(c) - 8))
		b.WriteString(";1m")
	case BackgroundBlack, BackgroundRed, BackgroundGreen, BackgroundYellow, BackgroundBlue, BackgroundMagenta, BackgroundCyan, BackgroundWhite:
		b.WriteString(strconv.Itoa(int(c) - 6))
		b.WriteString("m")
	case Reset, Black, Red, Green, Yellow, Blue, Magenta, Cyan, White:
		b.WriteString(strconv.Itoa(int(c)))
		b.WriteString("m")
	case BackgroundOrange:
		b.WriteString("48;5;")
		b.WriteString(strconv.Itoa(-1 * int(c)))
		b.WriteString("m")
	default:
		b.WriteString("38;5;")
		b.WriteString(strconv.Itoa(int(c)))
		b.WriteString("m")
	}
	return b.String()
}

func (c Color) description() string {
	colorDescriptions := map[Color]string{
		Black:             "black",
		Red:               "red",
		Green:             "green",
		Yellow:            "yellow",
		Blue:              "blue",
		Magenta:           "magenta",
		Cyan:              "cyan",
		White:             "white",
		BrightBlack:       "bright black",
		BrightRed:         "bright red",
		BrightGreen:       "bright green",
		BrightYellow:      "bright yellow",
		BrightBlue:        "bright blue",
		BrightMagenta:     "bright magenta",
		BrightCyan:        "bright cyan",
		BrightWhite:       "bright white",
		BackgroundBlack:   "background black",
		BackgroundRed:     "background red",
		BackgroundGreen:   "background green",
		BackgroundYellow:  "background yellow",
		BackgroundBlue:    "background blue",
		BackgroundMagenta: "background magenta",
		BackgroundCyan:    "background cyan",
		BackgroundWhite:   "background white",
		Orange:            "orange",
		Reset:             "reset",
	}

	return colorDescriptions[c]
}

// decorate formats the provided input so it can be printed in color
func (c Color) decorate(input string) string {
	var b strings.Builder
	b.Grow(13)
	b.WriteString(c.String())
	b.WriteString(input)
	return b.String()
}

func (c Color) background() Color {
	switch c {
	case Black, BrightBlack:
		return BackgroundBlack
	case Red, BrightRed:
		return BackgroundRed
	case Green, BrightGreen:
		return BackgroundGreen
	case Yellow, BrightYellow:
		return BackgroundYellow
	case Blue, BrightBlue:
		return BackgroundBlue
	case Magenta, BrightMagenta:
		return BackgroundMagenta
	case Cyan, BrightCyan:
		return BackgroundCyan
	case White, BrightWhite:
		return BackgroundWhite
	case Orange:
		return BackgroundOrange
	default:
		return c
	}
}
