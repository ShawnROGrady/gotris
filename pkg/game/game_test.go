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
	piece := pieceConstructor(width, height)
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
	expectedPosition tetriminoTestCase
}{
	"new i piece no hidden rows": {
		pieceConstructor: tetrimino.PieceConstructors[0],
		//pieceConstructor: tetrimino.NewTestIPiece,
		boardWidth:  10,
		boardHeight: 20,
		hiddenRows:  0,
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
