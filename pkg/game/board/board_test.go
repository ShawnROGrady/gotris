package board

import (
	"testing"

	"github.com/ShawnROGrady/gotris/pkg/canvas"
)

var clearRowsTests = map[string]struct {
	initialBlocks     [][]*Block
	expectedNewBlocks [][]*Block
	expectedFullRows  []int
	expectedNewCells  [][]*canvas.Cell
}{
	"only bottom row": {
		initialBlocks: [][]*Block{
			[]*Block{
				&Block{Color: canvas.Blue},
				&Block{Color: canvas.Blue},
				&Block{Color: canvas.Blue},
				&Block{Color: canvas.Blue},
			},
			[]*Block{nil, nil, nil, nil},
			[]*Block{nil, nil, nil, nil},
			[]*Block{nil, nil, nil, nil},
		},
		expectedNewBlocks: [][]*Block{
			[]*Block{nil, nil, nil, nil},
			[]*Block{nil, nil, nil, nil},
			[]*Block{nil, nil, nil, nil},
			[]*Block{nil, nil, nil, nil},
		},
		expectedFullRows: []int{0},
		expectedNewCells: [][]*canvas.Cell{
			[]*canvas.Cell{
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
			},
			[]*canvas.Cell{
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
			},
			[]*canvas.Cell{
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
			},
			[]*canvas.Cell{
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
			},
		},
	},
	"bottom row with other blocks": {
		initialBlocks: [][]*Block{
			[]*Block{
				&Block{Color: canvas.Blue},
				&Block{Color: canvas.Blue},
				&Block{Color: canvas.Blue},
				&Block{Color: canvas.Blue},
			},
			[]*Block{
				nil,
				&Block{Color: canvas.Blue},
				&Block{Color: canvas.Blue},
				nil,
			},
			[]*Block{nil, nil, nil, nil},
			[]*Block{nil, nil, nil, nil},
		},
		expectedNewBlocks: [][]*Block{
			[]*Block{
				nil,
				&Block{Color: canvas.Blue},
				&Block{Color: canvas.Blue},
				nil,
			},
			[]*Block{nil, nil, nil, nil},
			[]*Block{nil, nil, nil, nil},
			[]*Block{nil, nil, nil, nil},
		},
		expectedFullRows: []int{0},
		expectedNewCells: [][]*canvas.Cell{
			[]*canvas.Cell{
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
			},
			[]*canvas.Cell{
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
			},
			[]*canvas.Cell{
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
			},
			[]*canvas.Cell{
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Blue, Background: canvas.Green}, &canvas.Cell{Color: canvas.Blue, Background: canvas.Green},
				&canvas.Cell{Color: canvas.Blue, Background: canvas.Green}, &canvas.Cell{Color: canvas.Blue, Background: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
			},
		},
	},
	"second from bottom row with other blocks": {
		initialBlocks: [][]*Block{
			[]*Block{
				nil,
				&Block{Color: canvas.Blue},
				&Block{Color: canvas.Blue},
				nil,
			},
			[]*Block{
				&Block{Color: canvas.Blue},
				&Block{Color: canvas.Blue},
				&Block{Color: canvas.Blue},
				&Block{Color: canvas.Blue},
			},
			[]*Block{nil, nil, nil, nil},
			[]*Block{nil, nil, nil, nil},
		},
		expectedNewBlocks: [][]*Block{
			[]*Block{
				nil,
				&Block{Color: canvas.Blue},
				&Block{Color: canvas.Blue},
				nil,
			},
			[]*Block{nil, nil, nil, nil},
			[]*Block{nil, nil, nil, nil},
			[]*Block{nil, nil, nil, nil},
		},
		expectedFullRows: []int{1},
		expectedNewCells: [][]*canvas.Cell{
			[]*canvas.Cell{
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
			},
			[]*canvas.Cell{
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
			},
			[]*canvas.Cell{
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
			},
			[]*canvas.Cell{
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Blue, Background: canvas.Green}, &canvas.Cell{Color: canvas.Blue, Background: canvas.Green},
				&canvas.Cell{Color: canvas.Blue, Background: canvas.Green}, &canvas.Cell{Color: canvas.Blue, Background: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
			},
		},
	},
	"just bottom two": {
		initialBlocks: [][]*Block{
			[]*Block{
				&Block{Color: canvas.Blue},
				&Block{Color: canvas.Blue},
				&Block{Color: canvas.Blue},
				&Block{Color: canvas.Blue},
			},
			[]*Block{
				&Block{Color: canvas.Blue},
				&Block{Color: canvas.Blue},
				&Block{Color: canvas.Blue},
				&Block{Color: canvas.Blue},
			},
			[]*Block{nil, nil, nil, nil},
			[]*Block{nil, nil, nil, nil},
		},
		expectedNewBlocks: [][]*Block{
			[]*Block{nil, nil, nil, nil},
			[]*Block{nil, nil, nil, nil},
			[]*Block{nil, nil, nil, nil},
			[]*Block{nil, nil, nil, nil},
		},
		expectedFullRows: []int{0, 1},
		expectedNewCells: [][]*canvas.Cell{
			[]*canvas.Cell{
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
			},
			[]*canvas.Cell{
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
			},
			[]*canvas.Cell{
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
			},
			[]*canvas.Cell{
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
			},
		},
	},
	"bottom two rows with other blocks": {
		initialBlocks: [][]*Block{
			[]*Block{
				&Block{Color: canvas.Blue},
				&Block{Color: canvas.Blue},
				&Block{Color: canvas.Blue},
				&Block{Color: canvas.Blue},
			},
			[]*Block{
				&Block{Color: canvas.Blue},
				&Block{Color: canvas.Blue},
				&Block{Color: canvas.Blue},
				&Block{Color: canvas.Blue},
			},
			[]*Block{
				nil,
				&Block{Color: canvas.Blue},
				&Block{Color: canvas.Blue},
				nil,
			},
			[]*Block{nil, nil, nil, nil},
		},
		expectedNewBlocks: [][]*Block{
			[]*Block{
				nil,
				&Block{Color: canvas.Blue},
				&Block{Color: canvas.Blue},
				nil,
			},
			[]*Block{nil, nil, nil, nil},
			[]*Block{nil, nil, nil, nil},
			[]*Block{nil, nil, nil, nil},
		},
		expectedFullRows: []int{0, 1},
		expectedNewCells: [][]*canvas.Cell{
			[]*canvas.Cell{
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
			},
			[]*canvas.Cell{
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
			},
			[]*canvas.Cell{
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
			},
			[]*canvas.Cell{
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Blue, Background: canvas.Green}, &canvas.Cell{Color: canvas.Blue, Background: canvas.Green},
				&canvas.Cell{Color: canvas.Blue, Background: canvas.Green}, &canvas.Cell{Color: canvas.Blue, Background: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
			},
		},
	},
	"no clearing neccessary": {
		initialBlocks: [][]*Block{
			[]*Block{
				nil,
				&Block{Color: canvas.Blue},
				&Block{Color: canvas.Blue},
				nil,
			},
			[]*Block{
				&Block{Color: canvas.Blue},
				nil,
				&Block{Color: canvas.Blue},
				&Block{Color: canvas.Blue},
			},
			[]*Block{nil, nil, nil, nil},
			[]*Block{nil, nil, nil, nil},
		},
		expectedNewBlocks: [][]*Block{
			[]*Block{
				nil,
				&Block{Color: canvas.Blue},
				&Block{Color: canvas.Blue},
				nil,
			},
			[]*Block{
				&Block{Color: canvas.Blue},
				nil,
				&Block{Color: canvas.Blue},
				&Block{Color: canvas.Blue},
			},
			[]*Block{nil, nil, nil, nil},
			[]*Block{nil, nil, nil, nil},
		},
		expectedNewCells: [][]*canvas.Cell{
			[]*canvas.Cell{
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
			},
			[]*canvas.Cell{
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
			},
			[]*canvas.Cell{
				&canvas.Cell{Color: canvas.Blue, Background: canvas.Green}, &canvas.Cell{Color: canvas.Blue, Background: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Blue, Background: canvas.Green}, &canvas.Cell{Color: canvas.Blue, Background: canvas.Green},
				&canvas.Cell{Color: canvas.Blue, Background: canvas.Green}, &canvas.Cell{Color: canvas.Blue, Background: canvas.Green},
			},
			[]*canvas.Cell{
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
				&canvas.Cell{Color: canvas.Blue, Background: canvas.Green}, &canvas.Cell{Color: canvas.Blue, Background: canvas.Green},
				&canvas.Cell{Color: canvas.Blue, Background: canvas.Green}, &canvas.Cell{Color: canvas.Blue, Background: canvas.Green},
				&canvas.Cell{Color: canvas.Green}, &canvas.Cell{Color: canvas.Green},
			},
		},
	},
}

func TestCheckRows(t *testing.T) {
	for testName, test := range clearRowsTests {
		b := &Board{
			background: canvas.Green,
			Blocks:     test.initialBlocks,
		}

		fullRows := b.CheckRows()

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
		b := &Board{
			background: canvas.Green,
			Blocks:     test.initialBlocks,
		}

		clearedRows := b.ClearFullRows()
		if clearedRows != len(test.expectedFullRows) {
			t.Fatalf("Unexpected number of cleared rows for test case '%s' [expected = %v, actual = %v", testName, len(test.expectedFullRows), clearedRows)
		}

		if len(b.Blocks) != len(test.initialBlocks) {
			t.Fatalf("checking rows resulted in new row count for test case '%s' [expected = %d, actual = %d]", testName, len(test.initialBlocks), len(b.Blocks))
			return
		}

		if len(b.Blocks) != len(test.expectedNewBlocks) {
			t.Fatalf("unexpected number of columns for test case '%s' [expected = %d, actual = %d]", testName, len(test.expectedNewBlocks), len(b.Blocks))
		}

		for i := range b.Blocks {
			if len(b.Blocks[i]) != len(test.expectedNewBlocks[i]) {
				t.Fatalf("unexpected number of blocks in updated row %d for test case '%s' [expected = %d, actual = %d]", i, testName, len(test.expectedNewBlocks[i]), len(b.Blocks[i]))
			}
			for j := range b.Blocks[i] {
				if b.Blocks[i][j] == nil && test.expectedNewBlocks[i][j] == nil {
					continue
				}

				if b.Blocks[i][j] == nil && test.expectedNewBlocks[i][j] != nil {
					t.Fatalf("board.blocks[%d][%d] unexpectedly nil for test case '%s'", i, j, testName)
					return
				}

				if b.Blocks[i][j] != nil && test.expectedNewBlocks[i][j] == nil {
					t.Fatalf("board.blocks[%d][%d] unexpectedly non-nil for test case '%s'", i, j, testName)
					return
				}

				if *b.Blocks[i][j] != *test.expectedNewBlocks[i][j] {
					t.Fatalf("unexpected new board.blocks[%d][%d] for test case '%s' [expected = %v, actual = %v]", i, j, testName, *test.expectedNewBlocks[i][j], *b.Blocks[i][j])
					return
				}
			}
		}

		cells := b.Cells()
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
					t.Fatalf("unexpected new cells[%d][%d] for test case '%s' [expected = %#v, actual = %#v]", i, j, testName, *test.expectedNewCells[i][j], *cells[i][j])
					return
				}
			}
		}

	}
}
