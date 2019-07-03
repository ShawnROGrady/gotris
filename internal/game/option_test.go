package game

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/ShawnROGrady/gotris/internal/canvas"
)

var optionTests = map[string]struct {
	options []Option
	pass    []func(g *Game) error
}{
	"no options": {
		pass: []func(g *Game) error{
			checkControlScheme(HomeRow()),
			checkWithoutGhost(false),
			checkBackground(canvas.White),
			checkColor(canvas.White),
			checkDebugMode(false),
			checkWithoutSide(false),
			checkWidthScale(2),
			checkWidth(10),
			checkHeight(24), // includes hidden rows
			checkHiddenRows(4),
		},
	},
	"with arrowKeys scheme": {
		options: []Option{
			WithControlScheme(ArrowKeys()),
		},
		pass: []func(g *Game) error{
			checkControlScheme(ArrowKeys()),
			checkWithoutGhost(false),
			checkBackground(canvas.White),
			checkColor(canvas.White),
			checkDebugMode(false),
			checkWithoutSide(false),
			checkWidthScale(2),
			checkWidth(10),
			checkHeight(24), // includes hidden rows
			checkHiddenRows(4),
		},
	},
	"without ghost or side-bar": {
		options: []Option{
			WithoutGhost(),
			WithoutSide(),
		},
		pass: []func(g *Game) error{
			checkControlScheme(HomeRow()),
			checkWithoutGhost(true),
			checkBackground(canvas.White),
			checkColor(canvas.White),
			checkDebugMode(false),
			checkWithoutSide(true),
			checkWidthScale(2),
			checkWidth(10),
			checkHeight(24), // includes hidden rows
			checkHiddenRows(4),
		},
	},
	"with debug mode": {
		options: []Option{
			WithDebugMode(),
		},
		pass: []func(g *Game) error{
			checkControlScheme(HomeRow()),
			checkWithoutGhost(false),
			checkBackground(canvas.White),
			checkColor(canvas.White),
			checkDebugMode(true),
			checkWithoutSide(false),
			checkWidthScale(2),
			checkWidth(10),
			checkHeight(24), // includes hidden rows
			checkHiddenRows(4),
		},
	},
	"black background and color": {
		options: []Option{
			WithBackground(canvas.Black),
			WithColor(canvas.Black),
		},
		pass: []func(g *Game) error{
			checkControlScheme(HomeRow()),
			checkWithoutGhost(false),
			checkBackground(canvas.Black),
			checkColor(canvas.Black),
			checkDebugMode(false),
			checkWithoutSide(false),
			checkWidthScale(2),
			checkWidth(10),
			checkHeight(24), // includes hidden rows
			checkHiddenRows(4),
		},
	},
	"with width=height=40, widthScale=1": {
		options: []Option{
			WithDimensions(dimensions{
				width:      40,
				height:     40,
				widthScale: 1,
			}),
		},
		pass: []func(g *Game) error{
			checkControlScheme(HomeRow()),
			checkWithoutGhost(false),
			checkBackground(canvas.White),
			checkColor(canvas.White),
			checkDebugMode(false),
			checkWithoutSide(false),
			checkWidthScale(1),
			checkWidth(40),
			checkHeight(44), // includes hidden rows
			checkHiddenRows(4),
		},
	},
	"with 8 hiddenRows": {
		options: []Option{
			WithHiddenRows(8),
		},
		pass: []func(g *Game) error{
			checkControlScheme(HomeRow()),
			checkWithoutGhost(false),
			checkBackground(canvas.White),
			checkColor(canvas.White),
			checkDebugMode(false),
			checkWithoutSide(false),
			checkWidthScale(2),
			checkWidth(10),
			checkHeight(28), // includes hidden rows
			checkHiddenRows(8),
		},
	},
}

func TestOptions(t *testing.T) {
	for testName, test := range optionTests {
		var b bytes.Buffer

		g := New(&b, test.options...)
		for i := range test.pass {
			if err := test.pass[i](g); err != nil {
				t.Errorf("Test case '%s', Error: %s", testName, err)
			}
		}
	}
}

func checkControlScheme(expected ControlScheme) func(g *Game) error {
	return func(g *Game) error {
		if g.controlScheme.String() != expected.String() {
			return fmt.Errorf("unexpected controlScheme [expected = %d, actual = %d]", expected, g.controlScheme)
		}
		return nil
	}
}

func checkWithoutGhost(expected bool) func(g *Game) error {
	return func(g *Game) error {
		if g.disableGhost != expected {
			return fmt.Errorf("unexpected disableGhost [expected = %v, actual = %v]", expected, g.disableGhost)
		}
		return nil
	}
}

func checkBackground(expected canvas.Color) func(g *Game) error {
	return func(g *Game) error {
		if g.board.Background() != expected {
			return fmt.Errorf("unexpected board.Background() [expected = %#v, actual = %#v]", expected, g.board.Background())
		}
		return nil
	}
}

func checkColor(expected canvas.Color) func(g *Game) error {
	return func(g *Game) error {
		if g.color != expected {
			return fmt.Errorf("unexpected game color [expected = %#v, actual = %#v]", expected, g.color)
		}
		return nil
	}
}

func checkDebugMode(expected bool) func(g *Game) error {
	return func(g *Game) error {
		if g.debugMode != expected {
			return fmt.Errorf("unexpected game debugMode [expected = %v, actual = %v]", expected, g.debugMode)
		}
		return nil
	}
}

func checkWithoutSide(expected bool) func(g *Game) error {
	return func(g *Game) error {
		if g.disableSide != expected {
			return fmt.Errorf("unexpected game disableSide [expected = %v, actual = %v]", expected, g.disableSide)
		}
		return nil
	}
}

func checkWidthScale(expected int) func(g *Game) error {
	return func(g *Game) error {
		if g.widthScale != expected {
			return fmt.Errorf("unexpected widthScale [expected = %d, actual = %d]", expected, g.widthScale)
		}
		return nil
	}
}

func checkWidth(expected int) func(g *Game) error {
	return func(g *Game) error {
		boardWidth := boardWidth(g.board)
		if boardWidth != expected {
			return fmt.Errorf("unexpected board width [expected = %d, actual = %d]", expected, boardWidth)
		}
		return nil
	}
}

func checkHeight(expected int) func(g *Game) error {
	return func(g *Game) error {
		boardHeight := boardHeight(g.board)
		if boardHeight != expected {
			return fmt.Errorf("height unexpectedly not default [expected = %d, actual = %d]", expected, boardHeight)
		}
		return nil
	}
}

func checkHiddenRows(expected int) func(g *Game) error {
	return func(g *Game) error {
		boardHiddenRows := g.board.HiddenRows()
		if boardHiddenRows != expected {
			return fmt.Errorf("unexpected board hiddenRows [expected = %d, actual = %d]", expected, boardHiddenRows)
		}
		return nil
	}
}
