package game

import (
	"os"

	"github.com/ShawnROGrady/gotris/pkg/canvas"
	"github.com/ShawnROGrady/gotris/pkg/game/board"
	"github.com/ShawnROGrady/gotris/pkg/game/tetrimino"
	"github.com/ShawnROGrady/gotris/pkg/inputreader"
)

// Game is responsible for handling the game state
type Game struct {
	inputreader  inputreader.InputReader
	canvas       *canvas.Canvas
	board        *board.Board
	currentPiece tetrimino.Tetrimino
}

// New returns a new game with the specified specifications
func New(term *os.File, width, height, hiddenRows int) *Game {
	piece := tetrimino.New(width, height)
	return &Game{
		inputreader: inputreader.NewTermReader(term),
		canvas: canvas.New(
			term,
			canvas.White,
			width, height,
		),
		board: board.New(
			canvas.White,
			width, height,
			hiddenRows,
		),
		currentPiece: piece,
	}
}

// RunDemo is a placeholder function to test core functionality
func (g *Game) RunDemo(done chan bool) (chan int, chan error) {
	input, readErr := translateInput(done, g.inputreader)
	var (
		runErr   = make(chan error)
		endScore = make(chan int)
	)

	// add initial piece to canvas
	g.addPieceToBoard()

	g.canvas.Cells = g.board.Cells()

	// initialize the canvas
	if err := g.canvas.Init(); err != nil {
		runErr <- err
		return endScore, runErr
	}

	// render initial canvas
	if err := g.canvas.Render(); err != nil {
		runErr <- err
		return endScore, runErr
	}

	go func() {
		for {
			select {
			case err := <-readErr:
				runErr <- err
				return
			case <-done:
				return
			case in := <-input:
				// TODO: print if in debug mode
				//log.Printf("User input: %s", in)

				if err := g.handleInput(in, endScore); err != nil {
					runErr <- err
					return
				}

			}
		}
	}()
	return endScore, runErr
}

func (g *Game) addPieceToBoard() {
	var (
		piece  = g.currentPiece
		topL   = piece.ContainingBox().TopLeft
		blocks = piece.Blocks()
	)

	for i, row := range blocks {
		for j, block := range row {
			if block == nil {
				continue
			}
			x := topL.X + j
			y := topL.Y - i

			g.board.Blocks[y][x] = block
		}
	}
}

// pieceConflicts checks if the current piece is in an occupied space
func (g *Game) pieceConflicts(oldTopL tetrimino.Coordinates, oldBlocks [][]*board.Block) bool {
	var (
		topL       = g.currentPiece.ContainingBox().TopLeft
		blocks     = g.currentPiece.Blocks()
		prevCoords = make(map[tetrimino.Coordinates]bool)
	)

	if oldTopL == topL {
		// piece didn't move
		return false
	}

	for i, row := range oldBlocks {
		for j, block := range row {
			if block == nil {
				continue
			}
			x := oldTopL.X + j
			y := oldTopL.Y - i
			prevCoords[tetrimino.Coordinates{
				X: x,
				Y: y,
			}] = true
		}
	}

	for i, row := range blocks {
		for j, block := range row {
			if block == nil {
				continue
			}
			x := topL.X + j
			y := topL.Y - i

			if g.board.Blocks[y][x] != nil && !prevCoords[tetrimino.Coordinates{X: x, Y: y}] {
				return true
			}
		}
	}

	return false
}

// pieceOutOfBounds checks if the piece is no longer in the bounds of the board
// this can happen after a rotation
func (g *Game) pieceOutOfBounds() bool {
	var (
		topL   = g.currentPiece.ContainingBox().TopLeft
		blocks = g.currentPiece.Blocks()
	)
	for i, row := range blocks {
		for j, block := range row {
			if block == nil {
				continue
			}
			x := topL.X + j
			y := topL.Y - i

			// can't rotate due to horizontal constraints
			if x < 0 || x > len(g.board.Blocks[0])-1 {
				return true
			}

			if y < 0 || y > len(g.board.Blocks)-1 {
				return true
			}
		}
	}
	return false
}

// checks if the current piece is in the hidden row(s)
func (g *Game) pieceAtTop() bool {
	return g.currentPiece.YMax().Y > len(g.board.Blocks)-g.board.HiddenRows-1
}

// current piece is at minimum vertical position
// either at bottom or on top of another piece
func (g *Game) pieceAtBottom() bool {
	var (
		topL        = g.currentPiece.ContainingBox().TopLeft
		blocks      = g.currentPiece.Blocks()
		pieceCoords = make(map[tetrimino.Coordinates]bool)
	)

	if g.currentPiece.YMin().Y == 0 {
		return true
	}

	for i, row := range blocks {
		for j, block := range row {
			if block == nil {
				continue
			}
			x := topL.X + j
			y := topL.Y - i
			pieceCoords[tetrimino.Coordinates{
				X: x,
				Y: y,
			}] = true
		}
	}

	for i, row := range blocks {
		for j, block := range row {
			if block == nil {
				continue
			}
			x := topL.X + j
			y := topL.Y - i

			if g.board.Blocks[y-1][x] != nil && !pieceCoords[tetrimino.Coordinates{X: x, Y: y - 1}] {
				return true
			}
		}
	}

	return false
}

func (g *Game) handleInput(input userInput, endScore chan int) error {
	topL := g.currentPiece.ContainingBox().TopLeft
	blocks := g.currentPiece.Blocks()

	g.movePiece(input)
	// new space already occupied
	if (input == moveLeft || input == moveRight || input == rotateLeft || input == rotateRight) && g.pieceConflicts(topL, blocks) {
		// move back to original spot
		if opposite := input.opposite(); opposite != ignore {
			g.movePiece(opposite)
			return nil
		}
	}
	if (input == rotateLeft || input == rotateRight) && g.pieceOutOfBounds() {
		// TODO: allow for 'kickback' off wall
		// move back to original spot
		if opposite := input.opposite(); opposite != ignore {
			g.movePiece(opposite)
			return nil
		}
	}

	newTopL := g.currentPiece.ContainingBox().TopLeft

	// clear cell where piece was
	for i, row := range blocks {
		for j, block := range row {
			if block == nil {
				continue
			}
			x := topL.X + j
			y := topL.Y - i

			g.board.Blocks[y][x] = nil
		}
	}

	topL = newTopL

	// update cell at pieces new position
	g.addPieceToBoard()

	// generate new current piece if at bottom or on top of another piece
	if g.pieceAtBottom() {
		// check if any rows can be cleared
		// TODO: add scoring
		g.board.ClearFullRows()

		if g.pieceAtTop() {
			// still render game-over state
			g.canvas.Cells = g.board.Cells()
			if err := g.canvas.Render(); err != nil {
				return err
			}
			endScore <- 0
			return nil
		}

		g.currentPiece = tetrimino.New(len(g.board.Blocks[0]), len(g.board.Blocks))

		// add new piece to canvas
		g.addPieceToBoard()
	}

	g.canvas.Cells = g.board.Cells()

	return g.canvas.Render()
}

func (g *Game) movePiece(input userInput) {
	var (
		xmax  = len(g.board.Blocks[0]) - 1
		ymax  = len(g.board.Blocks) - 1
		piece = g.currentPiece
	)

	switch input {
	case moveLeft:
		piece.MoveLeft()
	case moveDown:
		piece.MoveDown()
	case moveUp:
		piece.MoveUp(ymax)
	case moveRight:
		piece.MoveRight(xmax)
	case rotateLeft:
		piece.RotateCounter()
	case rotateRight:
		piece.RotateClockwise()
	}
}
