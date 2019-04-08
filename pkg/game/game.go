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
func New(term *os.File, width, height int) *Game {
	piece := tetrimino.New(width, height)
	return &Game{
		inputreader: inputreader.NewTermReader(term),
		canvas: canvas.New(
			term,
			canvas.Green,
			width, height,
		),
		board: board.New(
			canvas.Green,
			width, height,
		),
		currentPiece: piece,
	}
}

// RunDemo is a placeholder function to test core functionality
func (g *Game) RunDemo(done chan bool) chan error {
	input, readErr := translateInput(done, g.inputreader)
	runErr := make(chan error)

	blocks := g.currentPiece.Blocks()

	// add initial piece to canvas
	g.addPieceToBoard()

	g.canvas.Cells = g.board.Cells()

	// render initial canvas
	if err := g.canvas.Render(); err != nil {
		runErr <- err
		return runErr
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

				topL := g.currentPiece.ContainingBox().TopLeft
				blocks = g.currentPiece.Blocks()

				if err := g.handleDemoInput(in); err != nil {
					runErr <- err
					return
				}

				// new space already occupied
				if (in == moveLeft || in == moveRight) && g.pieceConflicts(topL, blocks) {
					// move back to original spot
					if opposite := in.opposite(); opposite != ignore {
						if err := g.handleDemoInput(opposite); err != nil {
							runErr <- err
							return
						}
						continue
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
				if g.currentPiece.YMin().Y == 0 || g.board.Blocks[g.currentPiece.YMin().Y-1][g.currentPiece.YMin().X] != nil {
					// check if any rows can be cleared
					// TODO: add scoring
					g.board.ClearFullRows()

					g.currentPiece = tetrimino.New(len(g.board.Blocks[0]), len(g.board.Blocks))

					// add new piece to canvas
					g.addPieceToBoard()
				}

				g.canvas.Cells = g.board.Cells()

				if err := g.canvas.Render(); err != nil {
					runErr <- err
					return
				}
			}
		}
	}()
	return runErr
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

func (g *Game) handleDemoInput(input userInput) error {
	g.movePiece(input)
	return nil
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
