package game

import (
	"fmt"
	"testing"

	"github.com/ShawnROGrady/gotris/pkg/canvas"
	"github.com/ShawnROGrady/gotris/pkg/game/board"
	"github.com/ShawnROGrady/gotris/pkg/game/tetrimino"
)

type tetriminoTestCase struct {
	expectedMaxY tetriminoCoordTest
	expectedMinY tetriminoCoordTest
	expectedMaxX tetriminoCoordTest
	expectedMinX tetriminoCoordTest
}

type tetriminoCoordTest struct {
	x       int
	ignoreX bool
	y       int
	ignoreY bool
}

func newTestGame(width, height, hiddenRows int, pieceConstructor tetrimino.PieceConstructor) *Game {
	piece := pieceConstructor(width, height+hiddenRows)
	return &Game{
		board: board.New(
			canvas.White,
			width, height,
			hiddenRows,
		),
		currentPiece: piece,
	}

}

var addPieceToBoardTests = map[string]struct {
	pieceConstructor tetrimino.PieceConstructor
	boardWidth       int
	boardHeight      int
	hiddenRows       int
	expectAtTop      bool
	expectedPosition tetriminoTestCase
}{
	"new i piece no hidden rows": {
		pieceConstructor: tetrimino.PieceConstructors[0],
		boardWidth:       10,
		boardHeight:      20,
		hiddenRows:       0,
		expectAtTop:      false,
		expectedPosition: tetriminoTestCase{
			expectedMaxY: tetriminoCoordTest{
				y:       18,
				ignoreX: true,
			},
			expectedMinY: tetriminoCoordTest{
				y:       18,
				ignoreX: true,
			},
			expectedMaxX: tetriminoCoordTest{
				y: 18,
				x: 6,
			},
			expectedMinX: tetriminoCoordTest{
				y: 18,
				x: 3,
			},
		},
	},
	"new i piece 4 hidden rows": {
		pieceConstructor: tetrimino.PieceConstructors[0],
		boardWidth:       10,
		boardHeight:      20,
		hiddenRows:       4,
		expectAtTop:      true,
		expectedPosition: tetriminoTestCase{
			expectedMaxY: tetriminoCoordTest{
				y:       22,
				ignoreX: true,
			},
			expectedMinY: tetriminoCoordTest{
				y:       22,
				ignoreX: true,
			},
			expectedMaxX: tetriminoCoordTest{
				y: 22,
				x: 6,
			},
			expectedMinX: tetriminoCoordTest{
				y: 22,
				x: 3,
			},
		},
	},
	"new j piece 4 hidden rows": {
		pieceConstructor: tetrimino.PieceConstructors[1],
		boardWidth:       10,
		boardHeight:      20,
		hiddenRows:       4,
		expectAtTop:      true,
		expectedPosition: tetriminoTestCase{
			expectedMaxY: tetriminoCoordTest{
				y: 23,
				x: 3,
			},
			expectedMinY: tetriminoCoordTest{
				y:       22,
				ignoreX: true,
			},
			expectedMaxX: tetriminoCoordTest{
				y: 22,
				x: 5,
			},
			expectedMinX: tetriminoCoordTest{
				ignoreY: true,
				x:       3,
			},
		},
	},
	"new l piece 4 hidden rows": {
		pieceConstructor: tetrimino.PieceConstructors[2],
		boardWidth:       10,
		boardHeight:      20,
		hiddenRows:       4,
		expectAtTop:      true,
		expectedPosition: tetriminoTestCase{
			expectedMaxY: tetriminoCoordTest{
				y: 23,
				x: 5,
			},
			expectedMinY: tetriminoCoordTest{
				y:       22,
				ignoreX: true,
			},
			expectedMaxX: tetriminoCoordTest{
				y: 23,
				x: 5,
			},
			expectedMinX: tetriminoCoordTest{
				ignoreY: true,
				x:       3,
			},
		},
	},
	"new o piece 4 hidden rows": {
		pieceConstructor: tetrimino.PieceConstructors[3],
		boardWidth:       10,
		boardHeight:      20,
		hiddenRows:       4,
		expectAtTop:      true,
		expectedPosition: tetriminoTestCase{
			expectedMaxY: tetriminoCoordTest{
				y:       23,
				ignoreX: true,
			},
			expectedMinY: tetriminoCoordTest{
				y:       22,
				ignoreX: true,
			},
			expectedMaxX: tetriminoCoordTest{
				ignoreY: true,
				x:       5,
			},
			expectedMinX: tetriminoCoordTest{
				ignoreY: true,
				x:       4,
			},
		},
	},
	"new s piece 4 hidden rows": {
		pieceConstructor: tetrimino.PieceConstructors[4],
		boardWidth:       10,
		boardHeight:      20,
		hiddenRows:       4,
		expectAtTop:      true,
		expectedPosition: tetriminoTestCase{
			expectedMaxY: tetriminoCoordTest{
				y:       23,
				ignoreX: true,
			},
			expectedMinY: tetriminoCoordTest{
				y:       22,
				ignoreX: true,
			},
			expectedMaxX: tetriminoCoordTest{
				y: 23,
				x: 5,
			},
			expectedMinX: tetriminoCoordTest{
				y: 22,
				x: 3,
			},
		},
	},
	"new t piece 4 hidden rows": {
		pieceConstructor: tetrimino.PieceConstructors[5],
		boardWidth:       10,
		boardHeight:      20,
		hiddenRows:       4,
		expectAtTop:      true,
		expectedPosition: tetriminoTestCase{
			expectedMaxY: tetriminoCoordTest{
				y: 23,
				x: 4,
			},
			expectedMinY: tetriminoCoordTest{
				y:       22,
				ignoreX: true,
			},
			expectedMaxX: tetriminoCoordTest{
				y: 22,
				x: 5,
			},
			expectedMinX: tetriminoCoordTest{
				y: 22,
				x: 3,
			},
		},
	},
	"new z piece 4 hidden rows": {
		pieceConstructor: tetrimino.PieceConstructors[6],
		boardWidth:       10,
		boardHeight:      20,
		hiddenRows:       4,
		expectAtTop:      true,
		expectedPosition: tetriminoTestCase{
			expectedMaxY: tetriminoCoordTest{
				y:       23,
				ignoreX: true,
			},
			expectedMinY: tetriminoCoordTest{
				y:       22,
				ignoreX: true,
			},
			expectedMaxX: tetriminoCoordTest{
				y: 22,
				x: 5,
			},
			expectedMinX: tetriminoCoordTest{
				y: 23,
				x: 3,
			},
		},
	},
}

func TestAddPieceToBoard(t *testing.T) {
	for testName, test := range addPieceToBoardTests {
		g := newTestGame(test.boardWidth, test.boardHeight, test.hiddenRows, test.pieceConstructor)

		// add piece to board
		g.addPieceToBoard()

		// verify piece coordinates
		if err := testPieceCoords(g.currentPiece, testName, test.expectedPosition); err != nil {
			t.Errorf("%s", err)
		}

		// check if piece at top
		if test.expectAtTop && !g.pieceAtTop() {
			t.Errorf("Piece unexpectedly not at top for test case '%s'", testName)
		}

		if !test.expectAtTop && g.pieceAtTop() {
			t.Errorf("Piece unexpectedly at top for test case '%s'", testName)
		}

		// verify piece not at bottom
		if g.pieceAtBottom() {
			t.Errorf("New piece unexpectedly at bottom of board for test case '%s'", testName)
		}
	}
}

func testPieceCoords(piece tetrimino.Tetrimino, testName string, testCase tetriminoTestCase) error {
	// test XMax
	maxX := piece.XMax()
	if !testCase.expectedMaxX.ignoreX {
		if maxX.X != testCase.expectedMaxX.x {
			return fmt.Errorf("Unexpected xMax for test case %s [expected = %d, actual = %d]", testName, testCase.expectedMaxX.x, maxX.X)
		}
	}
	if !testCase.expectedMaxX.ignoreY {
		if maxX.Y != testCase.expectedMaxX.y {
			return fmt.Errorf("Unexpected xMax.y for test case '%s' [expected = %d, actual = %d]", testName, testCase.expectedMaxX.y, maxX.Y)
		}
	}

	// test XMin
	minX := piece.XMin()
	if !testCase.expectedMinX.ignoreX {
		if minX.X != testCase.expectedMinX.x {
			return fmt.Errorf("Unexpected xMin for test case '%s' [expected = %d, actual = %d]", testName, testCase.expectedMinX.x, minX.X)
		}
	}

	if !testCase.expectedMinX.ignoreY {
		if minX.Y != testCase.expectedMinX.y {
			return fmt.Errorf("Unexpected xMin.y for test case '%s' [expected = %d, actual = %d]", testName, testCase.expectedMinX.y, minX.Y)
		}
	}

	// test YMax
	maxY := piece.YMax()
	if !testCase.expectedMaxY.ignoreX {
		if maxY.X != testCase.expectedMaxY.x {
			return fmt.Errorf("Unexpected yMax.x for test case '%s' [expected = %d, actual = %d]", testName, testCase.expectedMaxY.x, maxY.X)
		}
	}

	if !testCase.expectedMaxY.ignoreY {
		if maxY.Y != testCase.expectedMaxY.y {
			return fmt.Errorf("Unexpected yMax for test case '%s' [expected = %d, actual = %d]", testName, testCase.expectedMaxY.y, maxY.Y)
		}
	}

	// test YMin
	minY := piece.YMin()
	if !testCase.expectedMinY.ignoreX {
		if minY.X != testCase.expectedMinY.x {
			return fmt.Errorf("Unexpected yMin.x for test case '%s' [expected = %d, actual = %d]", testName, testCase.expectedMinY.x, minY.X)
		}
	}

	if !testCase.expectedMinY.ignoreY {
		if minY.Y != testCase.expectedMinY.y {
			return fmt.Errorf("Unexpected yMin for test case '%s' [expected = %d, actual = %d]", testName, testCase.expectedMinY.y, minY.Y)
		}
	}
	return nil
}
