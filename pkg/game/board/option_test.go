package board

import (
	"fmt"
	"testing"

	"github.com/ShawnROGrady/gotris/pkg/canvas"
)

var optionTests = map[string]struct {
	options []Option
	pass    []func(b *Board) error
}{
	"no options": {
		pass: []func(b *Board) error{
			checkDefaultBackground,
			checkDefaultHiddenRows,
			checkDefaultWidthScale,
			checkDefaultWidth,
			checkDefaultHeight,
		},
	},
	"with black background": {
		options: []Option{
			WithBackground(canvas.Black),
		},
		pass: []func(b *Board) error{
			func(b *Board) error {
				if b.Background() != canvas.Black {
					return fmt.Errorf("unexpected background [expected = %#v, actual = %#v]", canvas.Black, b.Background())
				}
				return nil
			},
			checkDefaultHiddenRows,
			checkDefaultWidthScale,
			checkDefaultWidth,
			checkDefaultHeight,
		},
	},
	"with single widthScale": {
		options: []Option{
			WithWidthScale(1),
		},
		pass: []func(b *Board) error{
			checkDefaultBackground,
			func(b *Board) error {
				if b.widthScale != 1 {
					return fmt.Errorf("unexpected widthScale [expected = %d, actual = %d]", 1, b.widthScale)
				}
				return nil
			},
			checkDefaultHiddenRows,
			checkDefaultWidth,
			checkDefaultHeight,
		},
	},
	"with no hidden rows": {
		options: []Option{
			WithHiddenRows(0),
		},
		pass: []func(b *Board) error{
			checkDefaultBackground,
			func(b *Board) error {
				if b.HiddenRows() != 0 {
					return fmt.Errorf("unexpected HiddenRows() [expected = %d, actual = %d]", 0, b.HiddenRows())
				}
				return nil
			},
			checkDefaultWidthScale,
			checkDefaultWidth,
			checkDefaultHeight,
		},
	},
	"with width = height = 40": {
		options: []Option{
			WithWidth(40),
			WithHeight(40),
		},
		pass: []func(b *Board) error{
			checkDefaultBackground,
			checkDefaultHiddenRows,
			checkDefaultWidthScale,
			func(b *Board) error {
				if b.width != 40 {
					return fmt.Errorf("unexpected width [expected = %d, actual = %d]", 40, b.width)
				}
				return nil
			},
			func(b *Board) error {
				if b.height != 40 {
					return fmt.Errorf("unexpected height [expected = %d, actual = %d]", 40, b.height)
				}
				return nil
			},
		},
	},
}

func TestOptions(t *testing.T) {
	for testName, test := range optionTests {
		b := New(test.options...)

		for i := range test.pass {
			if err := test.pass[i](b); err != nil {
				t.Errorf("Test case '%s', Error: %s", testName, err)
			}
		}
	}
}

func checkDefaultBackground(b *Board) error {
	if b.Background() != canvas.White {
		return fmt.Errorf("Background unexpectedly not default [expected = %#v, actual = %#v]", canvas.White, b.Background())
	}
	return nil
}

func checkDefaultHiddenRows(b *Board) error {
	if b.HiddenRows() != 4 {
		return fmt.Errorf("HiddenRows unexpectedly not default [expected = %d, actual = %d]", 4, b.HiddenRows())
	}
	return nil
}

func checkDefaultWidthScale(b *Board) error {
	if b.widthScale != 2 {
		return fmt.Errorf("widthScale unexpectedly not default [expected = %d, actual = %d]", 2, b.widthScale)
	}
	return nil
}

func checkDefaultWidth(b *Board) error {
	if b.width != 10 {
		return fmt.Errorf("width unexpectedly not default [expected = %d, actual = %d]", 10, b.width)
	}
	return nil
}

func checkDefaultHeight(b *Board) error {
	if b.height != 20 {
		return fmt.Errorf("height unexpectedly not default [expected = %d, actual = %d]", 20, b.height)
	}
	return nil
}
