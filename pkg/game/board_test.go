package game

import (
	"testing"

	"github.com/ShawnROGrady/gotris/pkg/canvas"
)

var clearRowsTests = map[string]struct {
	initialBlocks     [][]*block
	expectedNewBlocks [][]*block
	expectedFullRows  []int
	expectedNewCells  [][]*canvas.Cell
}{
	"only bottom row": {
		initialBlocks: [][]*block{
			[]*block{
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
			},
			[]*block{nil, nil, nil, nil},
			[]*block{nil, nil, nil, nil},
			[]*block{nil, nil, nil, nil},
		},
		expectedNewBlocks: [][]*block{
			[]*block{nil, nil, nil, nil},
			[]*block{nil, nil, nil, nil},
			[]*block{nil, nil, nil, nil},
			[]*block{nil, nil, nil, nil},
		},
		expectedFullRows: []int{0},
		expectedNewCells: [][]*canvas.Cell{
			[]*canvas.Cell{
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
			},
			[]*canvas.Cell{
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
			},
			[]*canvas.Cell{
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
			},
			[]*canvas.Cell{
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
			},
		},
	},
	"bottom row with other blocks": {
		initialBlocks: [][]*block{
			[]*block{
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
			},
			[]*block{
				nil,
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
				nil,
			},
			[]*block{nil, nil, nil, nil},
			[]*block{nil, nil, nil, nil},
		},
		expectedNewBlocks: [][]*block{
			[]*block{
				nil,
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
				nil,
			},
			[]*block{nil, nil, nil, nil},
			[]*block{nil, nil, nil, nil},
			[]*block{nil, nil, nil, nil},
		},
		expectedFullRows: []int{0},
		expectedNewCells: [][]*canvas.Cell{
			[]*canvas.Cell{
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
			},
			[]*canvas.Cell{
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
			},
			[]*canvas.Cell{
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
			},
			[]*canvas.Cell{
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Blue},
				&canvas.Cell{Background: canvas.Blue},
				&canvas.Cell{Background: canvas.Green},
			},
		},
	},
	"second from bottom row with other blocks": {
		initialBlocks: [][]*block{
			[]*block{
				nil,
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
				nil,
			},
			[]*block{
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
			},
			[]*block{nil, nil, nil, nil},
			[]*block{nil, nil, nil, nil},
		},
		expectedNewBlocks: [][]*block{
			[]*block{
				nil,
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
				nil,
			},
			[]*block{nil, nil, nil, nil},
			[]*block{nil, nil, nil, nil},
			[]*block{nil, nil, nil, nil},
		},
		expectedFullRows: []int{1},
		expectedNewCells: [][]*canvas.Cell{
			[]*canvas.Cell{
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
			},
			[]*canvas.Cell{
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
			},
			[]*canvas.Cell{
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
			},
			[]*canvas.Cell{
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Blue},
				&canvas.Cell{Background: canvas.Blue},
				&canvas.Cell{Background: canvas.Green},
			},
		},
	},
	"just bottom two": {
		initialBlocks: [][]*block{
			[]*block{
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
			},
			[]*block{
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
			},
			[]*block{nil, nil, nil, nil},
			[]*block{nil, nil, nil, nil},
		},
		expectedNewBlocks: [][]*block{
			[]*block{nil, nil, nil, nil},
			[]*block{nil, nil, nil, nil},
			[]*block{nil, nil, nil, nil},
			[]*block{nil, nil, nil, nil},
		},
		expectedFullRows: []int{0, 1},
		expectedNewCells: [][]*canvas.Cell{
			[]*canvas.Cell{
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
			},
			[]*canvas.Cell{
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
			},
			[]*canvas.Cell{
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
			},
			[]*canvas.Cell{
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
			},
		},
	},
	"bottom two rows with other blocks": {
		initialBlocks: [][]*block{
			[]*block{
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
			},
			[]*block{
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
			},
			[]*block{
				nil,
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
				nil,
			},
			[]*block{nil, nil, nil, nil},
		},
		expectedNewBlocks: [][]*block{
			[]*block{
				nil,
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
				nil,
			},
			[]*block{nil, nil, nil, nil},
			[]*block{nil, nil, nil, nil},
			[]*block{nil, nil, nil, nil},
		},
		expectedFullRows: []int{0, 1},
		expectedNewCells: [][]*canvas.Cell{
			[]*canvas.Cell{
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
			},
			[]*canvas.Cell{
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
			},
			[]*canvas.Cell{
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
			},
			[]*canvas.Cell{
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Blue},
				&canvas.Cell{Background: canvas.Blue},
				&canvas.Cell{Background: canvas.Green},
			},
		},
	},
	"no clearing neccessary": {
		initialBlocks: [][]*block{
			[]*block{
				nil,
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
				nil,
			},
			[]*block{
				&block{color: canvas.Blue},
				nil,
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
			},
			[]*block{nil, nil, nil, nil},
			[]*block{nil, nil, nil, nil},
		},
		expectedNewBlocks: [][]*block{
			[]*block{
				nil,
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
				nil,
			},
			[]*block{
				&block{color: canvas.Blue},
				nil,
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
			},
			[]*block{nil, nil, nil, nil},
			[]*block{nil, nil, nil, nil},
		},
		expectedNewCells: [][]*canvas.Cell{
			[]*canvas.Cell{
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
			},
			[]*canvas.Cell{
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Green},
			},
			[]*canvas.Cell{
				&canvas.Cell{Background: canvas.Blue},
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Blue},
				&canvas.Cell{Background: canvas.Blue},
			},
			[]*canvas.Cell{
				&canvas.Cell{Background: canvas.Green},
				&canvas.Cell{Background: canvas.Blue},
				&canvas.Cell{Background: canvas.Blue},
				&canvas.Cell{Background: canvas.Green},
			},
		},
	},
}

func TestCheckRows(t *testing.T) {
	for testName, test := range clearRowsTests {
		b := &board{
			background: canvas.Green,
			blocks:     test.initialBlocks,
		}

		fullRows := b.checkRows()

		if len(fullRows) != len(test.expectedFullRows) {
			t.Fatalf("Unexpected full rows detected for test case '%s' [expected = %v, actual = %v", testName, test.expectedFullRows, fullRows)
		}

		for i := range fullRows {
			if fullRows[i] != test.expectedFullRows[i] {
				t.Fatalf("Unexpected full rows detected for test case '%s' [expected = %v, actual = %v", testName, test.expectedFullRows, fullRows)
				return
			}
		}
	}
}

func TestClearFullRows(t *testing.T) {
	for testName, test := range clearRowsTests {
		b := &board{
			background: canvas.Green,
			blocks:     test.initialBlocks,
		}

		b.clearFullRows()

		if len(b.blocks) != len(test.initialBlocks) {
			t.Fatalf("checking rows resulted in new row count for test case '%s' [expected = %d, actual = %d]", testName, len(test.initialBlocks), len(b.blocks))
			return
		}

		if len(b.blocks) != len(test.expectedNewBlocks) {
			t.Fatalf("unexpected number of columns for test case '%s' [expected = %d, actual = %d]", testName, len(test.expectedNewBlocks), len(b.blocks))
		}

		for i := range b.blocks {
			if len(b.blocks[i]) != len(test.expectedNewBlocks[i]) {
				t.Fatalf("unexpected number of blocks in updated row %d for test case '%s' [expected = %d, actual = %d]", i, testName, len(test.expectedNewBlocks[i]), len(b.blocks[i]))
			}
			for j := range b.blocks[i] {
				if b.blocks[i][j] == nil && test.expectedNewBlocks[i][j] == nil {
					continue
				}

				if b.blocks[i][j] == nil && test.expectedNewBlocks[i][j] != nil {
					t.Fatalf("board.blocks[%d][%d] unexpectedly nil for test case '%s'", i, j, testName)
					return
				}

				if b.blocks[i][j] != nil && test.expectedNewBlocks[i][j] == nil {
					t.Fatalf("board.blocks[%d][%d] unexpectedly non-nil for test case '%s'", i, j, testName)
					return
				}

				if *b.blocks[i][j] != *test.expectedNewBlocks[i][j] {
					t.Fatalf("unexpected new board.blocks[%d][%d] for test case '%s' [expected = %v, actual = %v]", i, j, testName, *test.expectedNewBlocks[i][j], *b.blocks[i][j])
					return
				}
			}
		}

		cells := b.cells()
		if len(cells) != len(test.expectedNewCells) {
			t.Fatalf("unexpected number of cells for test case '%s' [expected = %d, actual = %d]", testName, len(test.expectedNewCells), len(cells))
		}

		if len(cells) != len(test.initialBlocks) {
			t.Fatalf("checking rows resulted in new cell count for test case '%s' [expected = %d, actual = %d]", testName, len(test.initialBlocks), len(cells))
			return
		}

		for i := range cells {
			if len(cells[i]) != len(test.expectedNewCells[i]) {
				t.Fatalf("unexpected number of cells in updated row %d for test case '%s' [expected = %d, actual = %d]", i, testName, len(test.expectedNewCells[i]), len(cells))
			}
			for j := range cells[i] {
				if cells[i][j] == nil && test.expectedNewCells[i][j] == nil {
					continue
				}

				if cells[i][j] == nil && test.expectedNewCells[i][j] != nil {
					t.Fatalf("cells[%d][%d] unexpectedly nil for test case '%s'", i, j, testName)
					return
				}

				if cells[i][j] != nil && test.expectedNewCells[i][j] == nil {
					t.Fatalf("cells[%d][%d] unexpectedly non-nil for test case '%s'", i, j, testName)
					return
				}

				if *cells[i][j] != *test.expectedNewCells[i][j] {
					t.Fatalf("unexpected new cells[%d][%d] for test case '%s' [expected = %v, actual = %v]", i, j, testName, *test.expectedNewCells[i][j], *cells[i][j])
					return
				}
			}
		}

	}
}
