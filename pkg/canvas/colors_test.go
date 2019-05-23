package canvas

import "testing"

var colorTests = map[Color]struct {
	ExpectedDescription    string
	ExpectedString         string
	ExpectedDecoratedInput string
	ExpectedBackground     Color
}{
	Black: {
		ExpectedDescription:    "black",
		ExpectedString:         "\u001b[30m",
		ExpectedDecoratedInput: "\u001b[30minput",
		ExpectedBackground:     BackgroundBlack,
	},
	Red: {
		ExpectedDescription:    "red",
		ExpectedString:         "\u001b[31m",
		ExpectedDecoratedInput: "\u001b[31minput",
		ExpectedBackground:     BackgroundRed,
	},
	Green: {
		ExpectedDescription:    "green",
		ExpectedString:         "\u001b[32m",
		ExpectedDecoratedInput: "\u001b[32minput",
		ExpectedBackground:     BackgroundGreen,
	},
	Yellow: {
		ExpectedDescription:    "yellow",
		ExpectedString:         "\u001b[33m",
		ExpectedDecoratedInput: "\u001b[33minput",
		ExpectedBackground:     BackgroundYellow,
	},
	Blue: {
		ExpectedDescription:    "blue",
		ExpectedString:         "\u001b[34m",
		ExpectedDecoratedInput: "\u001b[34minput",
		ExpectedBackground:     BackgroundBlue,
	},
	Magenta: {
		ExpectedDescription:    "magenta",
		ExpectedString:         "\u001b[35m",
		ExpectedDecoratedInput: "\u001b[35minput",
		ExpectedBackground:     BackgroundMagenta,
	},
	Cyan: {
		ExpectedDescription:    "cyan",
		ExpectedString:         "\u001b[36m",
		ExpectedDecoratedInput: "\u001b[36minput",
		ExpectedBackground:     BackgroundCyan,
	},
	White: {
		ExpectedDescription:    "white",
		ExpectedString:         "\u001b[37m",
		ExpectedDecoratedInput: "\u001b[37minput",
		ExpectedBackground:     BackgroundWhite,
	},
	BrightBlack: {
		ExpectedDescription:    "bright black",
		ExpectedString:         "\u001b[30;1m",
		ExpectedDecoratedInput: "\u001b[30;1minput",
		ExpectedBackground:     BrightBlack,
	},
	BrightRed: {
		ExpectedDescription:    "bright red",
		ExpectedString:         "\u001b[31;1m",
		ExpectedDecoratedInput: "\u001b[31;1minput",
		ExpectedBackground:     BrightRed,
	},
	BrightGreen: {
		ExpectedDescription:    "bright green",
		ExpectedString:         "\u001b[32;1m",
		ExpectedDecoratedInput: "\u001b[32;1minput",
		ExpectedBackground:     BrightGreen,
	},
	BrightYellow: {
		ExpectedDescription:    "bright yellow",
		ExpectedString:         "\u001b[33;1m",
		ExpectedDecoratedInput: "\u001b[33;1minput",
		ExpectedBackground:     BrightYellow,
	},
	BrightBlue: {
		ExpectedDescription:    "bright blue",
		ExpectedString:         "\u001b[34;1m",
		ExpectedDecoratedInput: "\u001b[34;1minput",
		ExpectedBackground:     BrightBlue,
	},
	BrightMagenta: {
		ExpectedDescription:    "bright magenta",
		ExpectedString:         "\u001b[35;1m",
		ExpectedDecoratedInput: "\u001b[35;1minput",
		ExpectedBackground:     BrightMagenta,
	},
	BrightCyan: {
		ExpectedDescription:    "bright cyan",
		ExpectedString:         "\u001b[36;1m",
		ExpectedDecoratedInput: "\u001b[36;1minput",
		ExpectedBackground:     BrightCyan,
	},
	BrightWhite: {
		ExpectedDescription:    "bright white",
		ExpectedString:         "\u001b[37;1m",
		ExpectedDecoratedInput: "\u001b[37;1minput",
		ExpectedBackground:     BrightWhite,
	},
	BackgroundBlack: {
		ExpectedDescription:    "background black",
		ExpectedString:         "\u001b[40m",
		ExpectedDecoratedInput: "\u001b[40minput",
		ExpectedBackground:     BackgroundBlack,
	},
	BackgroundRed: {
		ExpectedDescription:    "background red",
		ExpectedString:         "\u001b[41m",
		ExpectedDecoratedInput: "\u001b[41minput",
		ExpectedBackground:     BackgroundRed,
	},
	BackgroundGreen: {
		ExpectedDescription:    "background green",
		ExpectedString:         "\u001b[42m",
		ExpectedDecoratedInput: "\u001b[42minput",
		ExpectedBackground:     BackgroundGreen,
	},
	BackgroundYellow: {
		ExpectedDescription:    "background yellow",
		ExpectedString:         "\u001b[43m",
		ExpectedDecoratedInput: "\u001b[43minput",
		ExpectedBackground:     BackgroundYellow,
	},
	BackgroundBlue: {
		ExpectedDescription:    "background blue",
		ExpectedString:         "\u001b[44m",
		ExpectedDecoratedInput: "\u001b[44minput",
		ExpectedBackground:     BackgroundBlue,
	},
	BackgroundMagenta: {
		ExpectedDescription:    "background magenta",
		ExpectedString:         "\u001b[45m",
		ExpectedDecoratedInput: "\u001b[45minput",
		ExpectedBackground:     BackgroundMagenta,
	},
	BackgroundCyan: {
		ExpectedDescription:    "background cyan",
		ExpectedString:         "\u001b[46m",
		ExpectedDecoratedInput: "\u001b[46minput",
		ExpectedBackground:     BackgroundCyan,
	},
	BackgroundWhite: {
		ExpectedDescription:    "background white",
		ExpectedString:         "\u001b[47m",
		ExpectedDecoratedInput: "\u001b[47minput",
		ExpectedBackground:     BackgroundWhite,
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

		background := color.background()
		if background != test.ExpectedBackground {
			t.Errorf("Unexpected background for color '%s' [expected=%s, actual=%s]", color, test.ExpectedBackground, background)
		}
	}
}
