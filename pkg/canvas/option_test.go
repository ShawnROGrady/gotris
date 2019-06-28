package canvas

import (
	"bytes"
	"fmt"
	"testing"
)

var optionTests = map[string]struct {
	options []Option
	pass    []func(c *TermCanvas) error
}{
	"no options": {
		pass: []func(c *TermCanvas) error{
			checkDefaultBackground,
			checkDefaultDebugMode,
			checkDefaultWidth,
			checkDefaultHeight,
		},
	},
	"with debug mode": {
		options: []Option{WithDebugMode()},
		pass: []func(c *TermCanvas) error{
			checkDefaultBackground,
			func(c *TermCanvas) error {
				if c.debugMode != true {
					return fmt.Errorf("unexpected debugMode [expected = %#v, actual = %#v]", true, c.debugMode)
				}
				return nil
			},
			checkDefaultWidth,
			checkDefaultHeight,
		},
	},
	"with black background": {
		options: []Option{
			WithBackground(Black),
		},
		pass: []func(c *TermCanvas) error{
			func(c *TermCanvas) error {
				if c.background != Black {
					return fmt.Errorf("unexpected background [expected = %#v, actual = %#v]", Black, c.background)
				}
				return nil
			},
			checkDefaultDebugMode,
			checkDefaultWidth,
			checkDefaultHeight,
		},
	},
	"with width = height = 40": {
		options: []Option{
			WithWidth(40),
			WithHeight(40),
		},
		pass: []func(c *TermCanvas) error{
			checkDefaultBackground,
			checkDefaultDebugMode,
			func(c *TermCanvas) error {
				if c.width != 40 {
					return fmt.Errorf("unexpected width [expected = %d, actual = %d]", 40, c.width)
				}
				return nil
			},
			func(c *TermCanvas) error {
				if c.height != 40 {
					return fmt.Errorf("unexpected height [expected = %d, actual = %d]", 40, c.height)
				}
				return nil
			},
		},
	},
}

func TestOptions(t *testing.T) {
	for testName, test := range optionTests {
		var b bytes.Buffer
		c := New(&b, test.options...)

		for i := range test.pass {
			if err := test.pass[i](c); err != nil {
				t.Errorf("Test case '%s', Error: %s", testName, err)
			}
		}
	}
}

func checkDefaultBackground(c *TermCanvas) error {
	if c.background != White {
		return fmt.Errorf("background unexpectedly not default [expected = %#v, actual = %#v]", White, c.background)
	}
	return nil
}

func checkDefaultDebugMode(c *TermCanvas) error {
	if c.debugMode != false {
		return fmt.Errorf("debugMode unexpectedly not default [expected = %#v, actual = %#v]", false, c.debugMode)
	}
	return nil
}

func checkDefaultWidth(c *TermCanvas) error {
	if c.width != 20 {
		return fmt.Errorf("width unexpectedly not default [expected = %d, actual = %d]", 10, c.width)
	}
	return nil
}

func checkDefaultHeight(c *TermCanvas) error {
	if c.height != 20 {
		return fmt.Errorf("height unexpectedly not default [expected = %d, actual = %d]", 20, c.height)
	}
	return nil
}
