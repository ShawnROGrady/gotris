package canvas

import "testing"

var cellColorTests = map[Color]struct {
	ExpectedString string
}{
	Black: {
		ExpectedString: "\u001b[30m\u2588",
	},
	Red: {
		ExpectedString: "\u001b[31m\u2588",
	},
	Green: {
		ExpectedString: "\u001b[32m\u2588",
	},
	Yellow: {
		ExpectedString: "\u001b[33m\u2588",
	},
	Blue: {
		ExpectedString: "\u001b[34m\u2588",
	},
	Magenta: {
		ExpectedString: "\u001b[35m\u2588",
	},
	Cyan: {
		ExpectedString: "\u001b[36m\u2588",
	},
	White: {
		ExpectedString: "\u001b[37m\u2588",
	},
	BrightBlack: {
		ExpectedString: "\u001b[30;1m\u2588",
	},
	BrightRed: {
		ExpectedString: "\u001b[31;1m\u2588",
	},
	BrightGreen: {
		ExpectedString: "\u001b[32;1m\u2588",
	},
	BrightYellow: {
		ExpectedString: "\u001b[33;1m\u2588",
	},
	BrightBlue: {
		ExpectedString: "\u001b[34;1m\u2588",
	},
	BrightMagenta: {
		ExpectedString: "\u001b[35;1m\u2588",
	},
	BrightCyan: {
		ExpectedString: "\u001b[36;1m\u2588",
	},
	BrightWhite: {
		ExpectedString: "\u001b[37;1m\u2588",
	},
	Orange: {
		ExpectedString: "\u001b[38;5;208m\u2588",
	},
}

func TestCellColors(t *testing.T) {
	for color, test := range cellColorTests {
		cell := &Cell{
			Background: color,
		}

		cellString := cell.String()
		if cellString != test.ExpectedString {
			t.Errorf("Unexpected string for cell with %s background [expected=%s, actual=%s", color.description(), test.ExpectedString, cellString)
		}
	}
}
