package game

import (
	"testing"

	"github.com/ShawnROGrady/gotris/pkg/canvas"
)

var checkRowsTests = map[string]struct {
	initialBlocks     [][]*block
	expectedNewBlocks [][]*block
}{
	"only bottom row": {
		initialBlocks: [][]*block{
			[]*block{nil, nil, nil, nil},
			[]*block{nil, nil, nil, nil},
			[]*block{nil, nil, nil, nil},
			[]*block{
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
			},
		},
		expectedNewBlocks: [][]*block{
			[]*block{nil, nil, nil, nil},
			[]*block{nil, nil, nil, nil},
			[]*block{nil, nil, nil, nil},
			[]*block{nil, nil, nil, nil},
		},
	},
	"bottom row with other blocks": {
		initialBlocks: [][]*block{
			[]*block{nil, nil, nil, nil},
			[]*block{nil, nil, nil, nil},
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
		},
		expectedNewBlocks: [][]*block{
			[]*block{nil, nil, nil, nil},
			[]*block{nil, nil, nil, nil},
			[]*block{nil, nil, nil, nil},
			[]*block{
				nil,
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
				nil,
			},
		},
	},
	"second from bottom row with other blocks": {
		initialBlocks: [][]*block{
			[]*block{nil, nil, nil, nil},
			[]*block{nil, nil, nil, nil},
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
		},
		expectedNewBlocks: [][]*block{
			[]*block{nil, nil, nil, nil},
			[]*block{nil, nil, nil, nil},
			[]*block{nil, nil, nil, nil},
			[]*block{
				nil,
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
				nil,
			},
		},
	},
	"no clearing neccessary": {
		initialBlocks: [][]*block{
			[]*block{nil, nil, nil, nil},
			[]*block{nil, nil, nil, nil},
			[]*block{
				&block{color: canvas.Blue},
				nil,
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
			},
			[]*block{
				nil,
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
				nil,
			},
		},
		expectedNewBlocks: [][]*block{
			[]*block{nil, nil, nil, nil},
			[]*block{nil, nil, nil, nil},
			[]*block{
				&block{color: canvas.Blue},
				nil,
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
			},
			[]*block{
				nil,
				&block{color: canvas.Blue},
				&block{color: canvas.Blue},
				nil,
			},
		},
	},
}

func TestCheckRows(t *testing.T) {
	for testName, test := range checkRowsTests {
		b := &board{
			blocks: test.initialBlocks,
		}

		b.checkRows()

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

	}
}
