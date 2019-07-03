package canvas

import (
	"bytes"
	"testing"
)

var renderTests = map[string]struct {
	width            int
	height           int
	background       Color
	debugMode        bool
	cells            [][]Cell
	expectedContents string
}{
	"2x2 no update, debugMode": {
		width:            2,
		height:           2,
		background:       White,
		debugMode:        true,
		expectedContents: "\u001b[37m\u2588\u001b[37m\u2588\n\u001b[0m\u001b[37m\u2588\u001b[37m\u2588\n\u001b[0m\u001b[0m",
	},
	"2x2 update cells": {
		width:      2,
		height:     2,
		background: Blue,
		cells: [][]Cell{
			{&BlockCell{Color: White}, &BlockCell{Color: White}},
			{&BlockCell{Color: White}, &BlockCell{Color: White}},
		},
		expectedContents: "\x1b[0;0H\u001b[47m\u001b[37m\u2588\u001b[47m\u001b[37m\u2588\n\u001b[0m\u001b[47m\u001b[37m\u2588\u001b[47m\u001b[37m\u2588\n\u001b[0m\u001b[0m",
	},
}

func TestRender(t *testing.T) {
	for testName, test := range renderTests {
		var b bytes.Buffer
		opts := []Option{WithBackground(test.background), WithWidth(test.width), WithHeight(test.height)}
		if test.debugMode {
			opts = append(opts, WithDebugMode())
		}
		c := New(&b, opts...)

		if len(test.cells) != 0 {
			c.UpdateCells(test.cells)
		}

		err := c.Render()
		if err != nil {
			t.Fatalf("Unexpected error rendering canvas for test case '%s'", testName)
		}

		contents := b.String()
		if contents != test.expectedContents {
			t.Errorf("Unexpected contents for test case '%s' [expected = %#v, actual = %#v]", testName, test.expectedContents, contents)
		}
	}
}

var renderBenchmarks = []struct {
	name  string
	size  int
	cells [][]Cell
}{
	{
		name: "4x4 no updates",
		size: 4,
	},
	{
		name:  "4x4 update non-transparent",
		size:  4,
		cells: populateCells(4, false),
	},
	{
		name:  "4x4 update transparent",
		size:  4,
		cells: populateCells(4, true),
	},
	{
		name:  "10x10 update transparent",
		size:  10,
		cells: populateCells(10, true),
	},
	{
		name:  "20x20 update transparent",
		size:  20,
		cells: populateCells(20, true),
	},
	{
		name:  "50x50 update transparent",
		size:  50,
		cells: populateCells(50, true),
	},
	{
		name:  "100x100 update transparent",
		size:  100,
		cells: populateCells(100, true),
	},
}

var benchRes error

func BenchmarkRender(b *testing.B) {
	for _, benchmark := range renderBenchmarks {
		var (
			buf   bytes.Buffer
			size  = benchmark.size
			cells = benchmark.cells
			name  = benchmark.name
		)
		opts := []Option{WithWidth(size), WithHeight(size)}
		c := New(&buf, opts...)
		if len(cells) != 0 {
			c.UpdateCells(cells)
		}

		b.Run(name, func(b *testing.B) {
			var err error
			for n := 0; n < b.N; n++ {
				err = c.Render()
				if err != nil {
					b.Errorf("Error rendering canvas: %s", err)
				}
			}
			benchRes = err
		})
	}
}

// generates grid of cells for testing
func populateCells(size int, transparent bool) [][]Cell {
	cells := make([][]Cell, size)
	for i := range cells {
		row := make([]Cell, size)
		for j := range row {
			row[j] = &BlockCell{
				Color:       Blue,
				Transparent: transparent,
			}
		}
		cells[i] = row
	}
	return cells
}
