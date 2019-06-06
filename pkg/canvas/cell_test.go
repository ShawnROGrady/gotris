package canvas

import "testing"

var cellColorTests = map[string]struct {
	Color          Color
	Background     Color
	Transparent    bool
	ExpectedString string
}{
	"Black solid": {
		Color:          Black,
		ExpectedString: "\u001b[40m\u001b[30m\u2588",
	},
	"Red solid": {
		Color:          Red,
		ExpectedString: "\u001b[41m\u001b[31m\u2588",
	},
	"Green solid": {
		Color:          Green,
		ExpectedString: "\u001b[42m\u001b[32m\u2588",
	},
	"Yellow solid": {
		Color:          Yellow,
		ExpectedString: "\u001b[43m\u001b[33m\u2588",
	},
	"Blue solid": {
		Color:          Blue,
		ExpectedString: "\u001b[44m\u001b[34m\u2588",
	},
	"Magenta solid": {
		Color:          Magenta,
		ExpectedString: "\u001b[45m\u001b[35m\u2588",
	},
	"Cyan solid": {
		Color:          Cyan,
		ExpectedString: "\u001b[46m\u001b[36m\u2588",
	},
	"White solid": {
		Color:          White,
		ExpectedString: "\u001b[47m\u001b[37m\u2588",
	},
	"BrightBlack solid": {
		Color:          BrightBlack,
		ExpectedString: "\u001b[40m\u001b[30;1m\u2588",
	},
	"BrightRed solid": {
		Color:          BrightRed,
		ExpectedString: "\u001b[41m\u001b[31;1m\u2588",
	},
	"BrightGreen solid": {
		Color:          BrightGreen,
		ExpectedString: "\u001b[42m\u001b[32;1m\u2588",
	},
	"BrightYellow solid": {
		Color:          BrightYellow,
		ExpectedString: "\u001b[43m\u001b[33;1m\u2588",
	},
	"BrightBlue solid": {
		Color:          BrightBlue,
		ExpectedString: "\u001b[44m\u001b[34;1m\u2588",
	},
	"BrightMagenta solid": {
		Color:          BrightMagenta,
		ExpectedString: "\u001b[45m\u001b[35;1m\u2588",
	},
	"BrightCyan solid": {
		Color:          BrightCyan,
		ExpectedString: "\u001b[46m\u001b[36;1m\u2588",
	},
	"BrightWhite solid": {
		Color:          BrightWhite,
		ExpectedString: "\u001b[47m\u001b[37;1m\u2588",
	},
	"Orange solid": {
		Color:          Orange,
		ExpectedString: "\u001b[48;5;208m\u001b[38;5;208m\u2588",
	},
	"Cyan transparent white background": {
		Color:          Cyan,
		Transparent:    true,
		Background:     White,
		ExpectedString: "\u001b[47m\u001b[36m\u2592",
	},
}

func TestCellColors(t *testing.T) {
	for testName, test := range cellColorTests {
		cell := &Cell{
			Color:       test.Color,
			Transparent: test.Transparent,
			Background:  test.Background,
		}

		cellString := cell.String()
		if cellString != test.ExpectedString {
			t.Errorf("Unexpected string for test case '%s' background [expected=%s,%s actual=%s]%s", testName, test.ExpectedString, Reset, cellString, Reset)
		}
	}
}
