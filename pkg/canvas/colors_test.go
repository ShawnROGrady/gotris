package canvas

import "testing"

var colorTests = map[Color]struct {
	ExpectedDescription    string
	ExpectedString         string
	ExpectedDecoratedInput string
}{
	Black: {
		ExpectedDescription:    "black",
		ExpectedString:         "\u001b[30m",
		ExpectedDecoratedInput: "\u001b[30minput",
	},
	Red: {
		ExpectedDescription:    "red",
		ExpectedString:         "\u001b[31m",
		ExpectedDecoratedInput: "\u001b[31minput",
	},
	Green: {
		ExpectedDescription:    "green",
		ExpectedString:         "\u001b[32m",
		ExpectedDecoratedInput: "\u001b[32minput",
	},
	Yellow: {
		ExpectedDescription:    "yellow",
		ExpectedString:         "\u001b[33m",
		ExpectedDecoratedInput: "\u001b[33minput",
	},
	Blue: {
		ExpectedDescription:    "blue",
		ExpectedString:         "\u001b[34m",
		ExpectedDecoratedInput: "\u001b[34minput",
	},
	Magenta: {
		ExpectedDescription:    "magenta",
		ExpectedString:         "\u001b[35m",
		ExpectedDecoratedInput: "\u001b[35minput",
	},
	Cyan: {
		ExpectedDescription:    "cyan",
		ExpectedString:         "\u001b[36m",
		ExpectedDecoratedInput: "\u001b[36minput",
	},
	White: {
		ExpectedDescription:    "white",
		ExpectedString:         "\u001b[37m",
		ExpectedDecoratedInput: "\u001b[37minput",
	},
	BrightBlack: {
		ExpectedDescription:    "bright black",
		ExpectedString:         "\u001b[30;1m",
		ExpectedDecoratedInput: "\u001b[30;1minput",
	},
	BrightRed: {
		ExpectedDescription:    "bright red",
		ExpectedString:         "\u001b[31;1m",
		ExpectedDecoratedInput: "\u001b[31;1minput",
	},
	BrightGreen: {
		ExpectedDescription:    "bright green",
		ExpectedString:         "\u001b[32;1m",
		ExpectedDecoratedInput: "\u001b[32;1minput",
	},
	BrightYellow: {
		ExpectedDescription:    "bright yellow",
		ExpectedString:         "\u001b[33;1m",
		ExpectedDecoratedInput: "\u001b[33;1minput",
	},
	BrightBlue: {
		ExpectedDescription:    "bright blue",
		ExpectedString:         "\u001b[34;1m",
		ExpectedDecoratedInput: "\u001b[34;1minput",
	},
	BrightMagenta: {
		ExpectedDescription:    "bright magenta",
		ExpectedString:         "\u001b[35;1m",
		ExpectedDecoratedInput: "\u001b[35;1minput",
	},
	BrightCyan: {
		ExpectedDescription:    "bright cyan",
		ExpectedString:         "\u001b[36;1m",
		ExpectedDecoratedInput: "\u001b[36;1minput",
	},
	BrightWhite: {
		ExpectedDescription:    "bright white",
		ExpectedString:         "\u001b[37;1m",
		ExpectedDecoratedInput: "\u001b[37;1minput",
	},
}

func TestColors(t *testing.T) {
	for color, test := range colorTests {
		colorDescription := color.description()
		if colorDescription != test.ExpectedDescription {
			t.Errorf("Unexpected description for color %d [expected=%s, actual=%s", color, test.ExpectedDescription, colorDescription)
		}

		colorString := color.String()
		if colorString != test.ExpectedString {
			t.Errorf("Unexpected string for color %s [expected=%s, actual=%s", color.description(), test.ExpectedString, colorString)
		}

		decorated := color.decorate("input")
		if decorated != test.ExpectedDecoratedInput {
			t.Errorf("Unexpected decorated 'input' for color %s [expected=%s, actual=%s", color.description(), test.ExpectedDecoratedInput, decorated)
		}

	}
}
