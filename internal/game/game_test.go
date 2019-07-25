package game

import (
	"fmt"
	"io"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/ShawnROGrady/gotris/internal/canvas"
	"github.com/ShawnROGrady/gotris/internal/game/board"
	"github.com/ShawnROGrady/gotris/internal/game/tetrimino"
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

type testCanvas struct {
	cells [][]canvas.Cell
}

func (t *testCanvas) Init() error                          { return nil }
func (t *testCanvas) Render() error                        { return nil }
func (t *testCanvas) UpdateCells(newCells [][]canvas.Cell) { t.cells = newCells }

func newTestGame(width, height, hiddenRows int, pieceSetConstructor func(width, height int) []tetrimino.Tetrimino) *Game {
	initPieces := pieceSetConstructor(width, height+hiddenRows)
	piece, pieceSet := initPieces[0], initPieces[1:]
	opts := []board.Option{board.WithWidth(width), board.WithHeight(height), board.WithHiddenRows(hiddenRows)}
	return &Game{
		board:         board.New(opts...),
		currentPiece:  piece,
		nextPieces:    pieceSet,
		canvas:        &testCanvas{cells: [][]canvas.Cell{}},
		newPieceSet:   pieceSetConstructor,
		disableGhost:  false, // enabling ghost to catch potential nil-pointer/index-oob exceptions
		controlScheme: HomeRow(),
		mutex:         &sync.Mutex{},
	}
}

func testNewSet(pieceConstructor tetrimino.PieceConstructor) func(width, height int) []tetrimino.Tetrimino {
	return func(width, height int) []tetrimino.Tetrimino {
		pieceSet := []tetrimino.Tetrimino{}
		for i := 0; i < 7; i++ {
			pieceSet = append(pieceSet, pieceConstructor(width, height))
		}
		return pieceSet
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
		g := newTestGame(test.boardWidth, test.boardHeight, test.hiddenRows, testNewSet(test.pieceConstructor))

		// add piece to board
		g.addPieceToBoard(g.currentPiece)

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
		if g.pieceAtBottom(g.currentPiece) {
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

func fillInputSequence(input userInput, count int) []userInput {
	sequence := []userInput{}
	for i := 0; i <= count; i++ {
		sequence = append(sequence, input)
	}
	return sequence
}

func combineInputSequences(sequences ...[]userInput) []userInput {
	inputSequence := []userInput{}
	for _, sequence := range sequences {
		inputSequence = append(inputSequence, sequence...)
	}
	return inputSequence
}

var handleInputTests = map[string]struct {
	pieceConstructor      tetrimino.PieceConstructor
	boardWidth            int
	boardHeight           int
	hiddenRows            int
	inputSequence         []userInput
	expectAtTop           bool
	expectedPosition      tetriminoTestCase
	expectedGhostPosition tetriminoTestCase
	expectGameOver        bool
}{
	"move i piece to bottom once": {
		pieceConstructor: tetrimino.PieceConstructors[0],
		boardWidth:       10,
		boardHeight:      20,
		hiddenRows:       4,
		expectAtTop:      true,
		inputSequence:    fillInputSequence(moveDown, 22), // 21 to get to bottom, one to lock in place
		// new piece will also be an "I" piece
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
		// Ghost will be on top of previous piece
		expectedGhostPosition: tetriminoTestCase{
			expectedMaxY: tetriminoCoordTest{
				y:       1,
				ignoreX: true,
			},
			expectedMinY: tetriminoCoordTest{
				y:       1,
				ignoreX: true,
			},
			expectedMaxX: tetriminoCoordTest{
				y: 1,
				x: 6,
			},
			expectedMinX: tetriminoCoordTest{
				y: 1,
				x: 3,
			},
		},
	},
	"move i piece to bottom then horizontal conflict": {
		pieceConstructor: tetrimino.PieceConstructors[0],
		boardWidth:       10,
		boardHeight:      20,
		hiddenRows:       4,
		expectAtTop:      false,
		inputSequence: combineInputSequences(
			[]userInput{rotateLeft}, fillInputSequence(moveDown, 20), // rotate then move to bottom
			[]userInput{rotateLeft, moveRight}, fillInputSequence(moveDown, 17), []userInput{moveLeft}, // rotate, move right, move down then attempt to move left
		),
		// final move should fail
		expectedPosition: tetriminoTestCase{
			expectedMaxY: tetriminoCoordTest{
				y: 5,
				x: 5,
			},
			expectedMinY: tetriminoCoordTest{
				y: 2,
				x: 5,
			},
			expectedMaxX: tetriminoCoordTest{
				x:       5,
				ignoreY: true,
			},
			expectedMinX: tetriminoCoordTest{
				x:       5,
				ignoreY: true,
			},
		},
		expectedGhostPosition: tetriminoTestCase{
			expectedMaxY: tetriminoCoordTest{
				y: 3,
				x: 5,
			},
			expectedMinY: tetriminoCoordTest{
				y: 0,
				x: 5,
			},
			expectedMaxX: tetriminoCoordTest{
				x:       5,
				ignoreY: true,
			},
			expectedMinX: tetriminoCoordTest{
				x:       5,
				ignoreY: true,
			},
		},
	},
	"left rotate i piece, move to right then left rotate": {
		pieceConstructor: tetrimino.PieceConstructors[0],
		boardWidth:       10,
		boardHeight:      20,
		hiddenRows:       4,
		expectAtTop:      false,
		inputSequence: combineInputSequences(
			[]userInput{rotateLeft}, fillInputSequence(moveRight, 4), // rotate then move to right
			[]userInput{rotateLeft},
		),
		// rotation should succeed
		expectedPosition: tetriminoTestCase{
			expectedMaxY: tetriminoCoordTest{
				y:       21,
				ignoreX: true,
			},
			expectedMinY: tetriminoCoordTest{
				y:       21,
				ignoreX: true,
			},
			expectedMaxX: tetriminoCoordTest{
				y: 21,
				x: 9,
			},
			expectedMinX: tetriminoCoordTest{
				y: 21,
				x: 6,
			},
		},
		expectedGhostPosition: tetriminoTestCase{
			expectedMaxY: tetriminoCoordTest{
				y:       0,
				ignoreX: true,
			},
			expectedMinY: tetriminoCoordTest{
				y:       0,
				ignoreX: true,
			},
			expectedMaxX: tetriminoCoordTest{
				y: 0,
				x: 9,
			},
			expectedMinX: tetriminoCoordTest{
				y: 0,
				x: 6,
			},
		},
	},
	"left rotate i piece, move to right then right rotate": {
		pieceConstructor: tetrimino.PieceConstructors[0],
		boardWidth:       10,
		boardHeight:      20,
		hiddenRows:       4,
		expectAtTop:      false,
		inputSequence: combineInputSequences(
			[]userInput{rotateLeft}, fillInputSequence(moveRight, 4), // rotate then move to right
			[]userInput{rotateRight},
		),
		// rotation should succeed
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
				x: 9,
			},
			expectedMinX: tetriminoCoordTest{
				y: 22,
				x: 6,
			},
		},
		expectedGhostPosition: tetriminoTestCase{
			expectedMaxY: tetriminoCoordTest{
				y:       0,
				ignoreX: true,
			},
			expectedMinY: tetriminoCoordTest{
				y:       0,
				ignoreX: true,
			},
			expectedMaxX: tetriminoCoordTest{
				y: 0,
				x: 9,
			},
			expectedMinX: tetriminoCoordTest{
				y: 0,
				x: 6,
			},
		},
	},
	"right rotate i piece, move to right then right rotate": {
		pieceConstructor: tetrimino.PieceConstructors[0],
		boardWidth:       10,
		boardHeight:      20,
		hiddenRows:       4,
		expectAtTop:      false,
		inputSequence: combineInputSequences(
			[]userInput{rotateRight}, fillInputSequence(moveRight, 4), // rotate then move to right
			[]userInput{rotateRight},
		),
		// rotation should succeed
		expectedPosition: tetriminoTestCase{
			expectedMaxY: tetriminoCoordTest{
				y:       21,
				ignoreX: true,
			},
			expectedMinY: tetriminoCoordTest{
				y:       21,
				ignoreX: true,
			},
			expectedMaxX: tetriminoCoordTest{
				y: 21,
				x: 9,
			},
			expectedMinX: tetriminoCoordTest{
				y: 21,
				x: 6,
			},
		},
		expectedGhostPosition: tetriminoTestCase{
			expectedMaxY: tetriminoCoordTest{
				y:       0,
				ignoreX: true,
			},
			expectedMinY: tetriminoCoordTest{
				y:       0,
				ignoreX: true,
			},
			expectedMaxX: tetriminoCoordTest{
				y: 0,
				x: 9,
			},
			expectedMinX: tetriminoCoordTest{
				y: 0,
				x: 6,
			},
		},
	},
	"right rotate i piece, move to right then left rotate": {
		pieceConstructor: tetrimino.PieceConstructors[0],
		boardWidth:       10,
		boardHeight:      20,
		hiddenRows:       4,
		expectAtTop:      false,
		inputSequence: combineInputSequences(
			[]userInput{rotateRight}, fillInputSequence(moveRight, 4), // rotate then move to right
			[]userInput{rotateLeft},
		),
		// rotation should succeed
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
				x: 9,
			},
			expectedMinX: tetriminoCoordTest{
				y: 22,
				x: 6,
			},
		},
		expectedGhostPosition: tetriminoTestCase{
			expectedMaxY: tetriminoCoordTest{
				y:       0,
				ignoreX: true,
			},
			expectedMinY: tetriminoCoordTest{
				y:       0,
				ignoreX: true,
			},
			expectedMaxX: tetriminoCoordTest{
				y: 0,
				x: 9,
			},
			expectedMinX: tetriminoCoordTest{
				y: 0,
				x: 6,
			},
		},
	},
	"move i piece to bottom until end": {
		pieceConstructor: tetrimino.PieceConstructors[0],
		boardWidth:       10,
		boardHeight:      20,
		hiddenRows:       4,
		inputSequence:    append(fillInputSequence(moveDown, 273)), // sum(x, 3, 23)
		expectAtTop:      true,
		expectGameOver:   true,
	},
	"hard drop t piece until end": {
		pieceConstructor: tetrimino.PieceConstructors[5],
		boardWidth:       10,
		boardHeight:      20,
		hiddenRows:       4,
		inputSequence:    append(fillInputSequence(moveUp, 10)),
		expectAtTop:      true,
		expectGameOver:   true,
	},
	"new j piece move down then all the right": {
		pieceConstructor: tetrimino.PieceConstructors[1],
		boardWidth:       10,
		boardHeight:      20,
		hiddenRows:       4,
		expectAtTop:      false,
		inputSequence:    []userInput{moveDown, moveRight, moveRight, moveRight, moveRight, moveRight, moveRight},
		expectedPosition: tetriminoTestCase{
			expectedMaxY: tetriminoCoordTest{
				y: 22,
				x: 7,
			},
			expectedMinY: tetriminoCoordTest{
				y:       21,
				ignoreX: true,
			},
			expectedMaxX: tetriminoCoordTest{
				y: 21,
				x: 9,
			},
			expectedMinX: tetriminoCoordTest{
				ignoreY: true,
				x:       7,
			},
		},
		expectedGhostPosition: tetriminoTestCase{
			expectedMaxY: tetriminoCoordTest{
				y: 1,
				x: 7,
			},
			expectedMinY: tetriminoCoordTest{
				y:       0,
				ignoreX: true,
			},
			expectedMaxX: tetriminoCoordTest{
				y: 0,
				x: 9,
			},
			expectedMinX: tetriminoCoordTest{
				ignoreY: true,
				x:       7,
			},
		},
	},
}

func TestHandleInput(t *testing.T) {
	for testName, test := range handleInputTests {
		g := newTestGame(test.boardWidth, test.boardHeight, test.hiddenRows, testNewSet(test.pieceConstructor))
		// add piece to board
		g.addPieceToBoard(g.currentPiece)
		// have to initialize ghost piece
		g.ghostPiece = g.findGhostPiece()

		var (
			endScore  = make(chan int)
			handleErr = make(chan error)
			inputOver = make(chan bool)
			gameOver  = make(chan bool)
		)

		go func(gameOver chan bool) {
			for _, input := range test.inputSequence {
				select {
				case <-gameOver:
					return
				default:
					if err := g.handleInput(input, endScore); err != nil {
						handleErr <- err
					}
				}
			}
			inputOver <- true
		}(gameOver)

		select {
		case err := <-handleErr:
			t.Fatalf("Unexpected error handling user input for test case '%s': %s", testName, err)
		case score := <-endScore:
			if !test.expectGameOver {
				t.Fatalf("Game unexpectedly over for test case '%s' (final score = %d)", testName, score)
			}
		case <-inputOver:
			if test.expectGameOver {
				t.Fatalf("Game unexpectedly not over for test case '%s' (current piece maxY = %v, minY = %v)", testName, g.currentPiece.YMax(), g.currentPiece.YMin())
			}
			// verify current piece coordinates
			if err := testPieceCoords(g.currentPiece, testName, test.expectedPosition); err != nil {
				t.Errorf("Current Piece: %s", err)
			}
			// verify ghost piece coordinates
			if err := testPieceCoords(g.ghostPiece, testName, test.expectedGhostPosition); err != nil {
				t.Errorf("Ghost Piece: %s", err)
			}

			// check if piece at top
			if test.expectAtTop && !g.pieceAtTop() {
				t.Errorf("Piece unexpectedly not at top for test case '%s' (maxY = %v, minY = %v)", testName, g.currentPiece.YMax(), g.currentPiece.YMin())
			}
		}
	}
}

func TestNextPiece(t *testing.T) {
	g := New(nil, nil)

	// verify correct number of pieces for new game (7 tetriminos - 1 currentPiece)
	if len(g.nextPieces) != 6 {
		t.Errorf("Unexpected number of next pieces for new game (expected = %d, actual = %d)", 6, len(g.nextPieces))
	}

	// go through all of the nextPieces slice
	for i := 0; i < 5; i++ {
		g.currentPiece = g.nextPiece()
		if len(g.nextPieces) != 5-i {
			t.Errorf("Unexpected number of next pieces after getting %d next piece(s) (expected = %d, actual = %d)", i, 5-i, len(g.nextPieces))
		}
	}

	// verify that after the next call we create a new slice of next pieces
	g.currentPiece = g.nextPiece()
	// should be 7 since this time the next piece came from the previous slice
	if len(g.nextPieces) != 7 {
		t.Errorf("Unexpected number of next pieces after going through initial nextPieces slice (expected = %d, actual = %d)", 7, len(g.nextPieces))
	}
}

var boardWithGhostTests = map[string]struct {
	pieceConstructor tetrimino.PieceConstructor
	boardWidth       int
	boardHeight      int
	hiddenRows       int
	inputSequence    []userInput
	overrideDiff     func(row, column int, block *board.Block) *board.Block
}{
	"new i piece no moves": {
		pieceConstructor: tetrimino.PieceConstructors[0],
		boardWidth:       10,
		boardHeight:      20,
		hiddenRows:       4,
		overrideDiff:     func(row, column int, block *board.Block) *board.Block { return block },
	},
	// the active and ghost piece will conflict on the board
	"new i piece rotate left then down 17": {
		pieceConstructor: tetrimino.PieceConstructors[0],
		boardWidth:       10,
		boardHeight:      20,
		hiddenRows:       4,
		inputSequence: combineInputSequences(
			[]userInput{rotateLeft}, fillInputSequence(moveDown, 17),
		),
		overrideDiff: func(row, column int, block *board.Block) *board.Block {
			// the third and fourth blocks in the fourth column will conflict
			if column == 4 {
				if row == 2 || row == 3 {
					// replace with ghost piece
					return &board.Block{
						Color:       canvas.Cyan,
						Transparent: true,
					}
				}
			}
			return block
		},
	},
}

func TestBoardWithGhost(t *testing.T) {
	for testName, test := range boardWithGhostTests {
		g := newTestGame(test.boardWidth, test.boardHeight, test.hiddenRows, testNewSet(test.pieceConstructor))
		// add piece to board
		g.addPieceToBoard(g.currentPiece)
		// have to initialize ghost piece
		g.ghostPiece = g.findGhostPiece()

		var (
			endScore = make(chan int)
		)

		for _, input := range test.inputSequence {
			if err := g.handleInput(input, endScore); err != nil {
				t.Fatalf("Unexpected error handling input for test case '%s'", testName)
			}
		}

		newBoard := g.boardWithGhost()
		// check that board's height hasn't changed
		if len(newBoard.Blocks) != len(g.board.Blocks) {
			t.Fatalf("Board height modified by adding ghost piece for test case '%s' (original=%d, new=%d)", testName, len(g.board.Blocks), len(newBoard.Blocks))
		}

		diffBlocks := [][]*board.Block{}

		// determine the blocks that are different
		for i := range g.board.Blocks {
			row := g.board.Blocks[i]
			newRow := newBoard.Blocks[i]
			// make sure width hasn't changed
			if len(row) != len(newRow) {
				t.Fatalf("Board width (row %d) modified by adding ghost piece for test case '%s' (original=%d, new=%d)", i, testName, len(row), len(newRow))
			}
			diffRow := make([]*board.Block, len(row))
			for j := range row {
				if row[j] == nil && newRow[j] == nil {
					diffRow[j] = nil
					continue
				}
				if row[j] == nil && newRow[j] != nil {
					diffRow[j] = newRow[j]
					continue
				}
				if row[j] != nil && newRow[j] == nil {
					diffRow[j] = row[j]
					continue
				}

				if *row[j] != *newRow[j] {
					diffRow[j] = row[j]
				} else {
					diffRow[j] = nil
				}
			}
			diffBlocks = append(diffBlocks, diffRow)
		}

		// create a brand new board and add the ghost piece
		opts := []board.Option{board.WithWidth(test.boardWidth), board.WithHeight(test.boardHeight), board.WithHiddenRows(test.hiddenRows)}
		board2 := board.New(opts...)
		g.board = board2
		g.addPieceToBoard(g.ghostPiece)

		// we expect the difference between the board w/ and w/o the ghost to be
		// the same as a new board with just the ghost piece
		for i := range board2.Blocks {
			diffRow := diffBlocks[i]
			newRow := board2.Blocks[i]
			for j := range newRow {
				diffRow[j] = test.overrideDiff(i, j, diffRow[j])

				if diffRow[j] == nil && newRow[j] == nil {
					continue
				}
				if diffRow[j] == nil && newRow[j] != nil {
					t.Errorf("Unexpected block at board.Blocks[%d][%d] for test case '%s' (expected = %v, actual = %v)", i, j, testName, newRow[j], diffRow[j])
					continue
				}
				if diffRow[j] != nil && newRow[j] == nil {
					t.Errorf("Unexpected block at board.Blocks[%d][%d] for test case '%s' (expected = %v, actual = %v)", i, j, testName, newRow[j], diffRow[j])
					continue
				}

				if *diffRow[j] != *newRow[j] {
					t.Errorf("Unexpected block at board.Blocks[%d][%d] for test case '%s' (expected = %+v, actual = %+v)", i, j, testName, *newRow[j], *diffRow[j])
				}
			}
		}
	}
}

var runTests = map[string]struct {
	currentLevel   level
	inputs         []string
	inputDelay     time.Duration
	expectedScore  int
	expectGameOver bool
}{
	"hard drop until end": {
		inputs:         fillStringSlice("k", 20),
		expectGameOver: true,
		inputDelay:     1 * time.Millisecond,
	},
	"max level, unhandled input until end": {
		currentLevel:   29,
		inputs:         fillStringSlice("r", 2*273), // 2 * sum(x, 3, 23), need to double to account for 'sliding' piece
		expectGameOver: true,
		inputDelay:     22 * time.Millisecond,
	},
	"clear one line, lvl 0": {
		currentLevel: 0,
		inputs: combineStringSlice(
			append(fillStringSlice("l", 4), "k"),
			append(fillStringSlice("h", 4), "k"),
			[]string{"a", "k"},
			[]string{"a", "l", "k"},
			[]string{" "},
		),
		expectGameOver: false,
		expectedScore:  40,
		inputDelay:     1 * time.Millisecond,
	},
}

func TestRun(t *testing.T) {
	for testName, test := range runTests {
		var (
			done                 = make(chan bool)
			inReader, inWriter   = io.Pipe()
			outReader, outWriter = io.Pipe()
		)

		// Need to read output written otherwise writes block
		go func() {
			defer outReader.Close()
			for {
				select {
				case <-done:
					return
				default:
					// TODO: verify written output matches what we expect
					buf := make([]byte, 128)
					_, err := outReader.Read(buf)
					if err != nil && err != io.EOF {
						log.Panicf("Error reading from out for test case '%s': %s", testName, err)
						return
					}
				}
			}
		}()

		g := New(inReader, outWriter, WithControlScheme(HomeRow()))
		g.level = test.currentLevel

		// Using exclusively 'I' pieces for easy testing
		pieceSetConstructor := testNewSet(tetrimino.PieceConstructors[0])
		initPieces := pieceSetConstructor(boardWidth(g.board), boardHeight(g.board))
		piece, pieceSet := initPieces[0], initPieces[1:]

		g.currentPiece, g.nextPieces = piece, pieceSet
		g.newPieceSet = pieceSetConstructor

		finalScore, runErr := g.Run(done)

		go func() {
			defer func() {
				inWriter.Close()
				time.Sleep(10 * time.Millisecond) // give ReadInput go routine enough time to read last input following 'done'
				close(done)
			}()
			for _, in := range test.inputs {
				_, err := inWriter.Write([]byte(in))
				if err != nil {
					log.Panicf("Error writing input for test case '%s': %s", testName, err)
					return
				}
				time.Sleep(test.inputDelay)
			}
		}()

		select {
		case score := <-finalScore:
			if !test.expectGameOver {
				t.Errorf("Game unexpectedly over after handling inputs for test case '%s'", testName)
			}
			if score != test.expectedScore {
				t.Errorf("Unexpected final score for test case '%s' [expected = %d, actual = %d]", testName, test.expectedScore, score)
			}

			// wait for goroutine writing inputs to complete (might still be running due to input delay)
			<-done
		case <-done:
			if test.expectGameOver {
				t.Errorf("Game unexpectedly not over after handling inputs for test case '%s'", testName)
			}
			g.mutex.Lock()
			score := g.currentScore
			g.mutex.Unlock()

			if score != test.expectedScore {
				t.Errorf("Unexpected current score for test case '%s' [expected = %d, actual = %d]", testName, test.expectedScore, score)
			}
		case err := <-runErr:
			t.Errorf("Unexpected error running game for test case '%s: %s'", testName, err)
		}
		// NOTE: uncomment the below to print the state of the game at the end of the test
		/*
			g.canvas = canvas.New(os.Stdout)
			g.canvas.UpdateCells(g.cells(g.board))
			g.canvas.Render()
		*/

	}
}

func BenchmarkRun(b *testing.B) {
	var (
		done                 = make(chan bool)
		inReader, inWriter   = io.Pipe()
		outReader, outWriter = io.Pipe()
	)

	// Need to read output written otherwise writes block
	go func() {
		defer outReader.Close()
		for {
			select {
			case <-done:
				return
			default:
				buf := make([]byte, 128)
				_, err := outReader.Read(buf)
				if err != nil && err != io.EOF {
					log.Panicf("Error reading from out for test case : %s", err)
					return
				}
			}
		}
	}()

	g := New(inReader, outWriter, WithControlScheme(HomeRow()))

	// Using exclusively 'I' pieces for easy testing
	pieceSetConstructor := testNewSet(tetrimino.PieceConstructors[0])
	initPieces := pieceSetConstructor(boardWidth(g.board), boardHeight(g.board))
	piece, pieceSet := initPieces[0], initPieces[1:]

	g.currentPiece, g.nextPieces = piece, pieceSet
	g.newPieceSet = pieceSetConstructor

	endScore, runErr := g.Run(done)

	for n := 0; n < b.N; n++ {
		var (
			writeErr = make(chan error)
			written  = make(chan struct{})
		)
		go func() {
			// always reset the blocks prior to writing input
			g.mutex.Lock()
			for i := range g.board.Blocks {
				for j := range g.board.Blocks[i] {
					g.board.Blocks[i][j] = nil
				}
			}
			g.mutex.Unlock()
			_, err := inWriter.Write([]byte("j"))
			if err != nil {
				writeErr <- err
			}
			written <- struct{}{}
		}()

		select {
		case err := <-runErr:
			b.Fatalf("Error running game: %s", err)
		case <-endScore:
			log.Printf("game over: %d", n)
		case err := <-writeErr:
			b.Errorf("Error writing input: %s", err)
		case <-written:
		}
	}
	close(done)
	inWriter.Close()
}

func fillStringSlice(input string, count int) []string {
	sequence := []string{}
	for i := 0; i <= count; i++ {
		sequence = append(sequence, input)
	}
	return sequence
}

func combineStringSlice(sequences ...[]string) []string {
	inputSequence := []string{}
	for _, sequence := range sequences {
		inputSequence = append(inputSequence, sequence...)
	}
	return inputSequence
}
