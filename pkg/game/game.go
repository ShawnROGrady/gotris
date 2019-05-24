package game

import (
	"fmt"
	"os"
	"time"

	"github.com/ShawnROGrady/gotris/pkg/canvas"
	"github.com/ShawnROGrady/gotris/pkg/game/board"
	"github.com/ShawnROGrady/gotris/pkg/game/tetrimino"
	"github.com/ShawnROGrady/gotris/pkg/inputreader"
)

// Game is responsible for handling the game state
type Game struct {
	inputreader   inputreader.InputReader
	canvas        canvas.Canvas
	board         *board.Board
	currentPiece  tetrimino.Tetrimino
	ghostPiece    tetrimino.Tetrimino
	newPieceSet   func(width, height int) []tetrimino.Tetrimino
	nextPieces    []tetrimino.Tetrimino
	level         level
	debugMode     bool
	disableGhost  bool
	controlScheme ControlScheme
}

// Config represents the configuration for a game
type Config struct {
	Term          *os.File
	Width         int
	Height        int
	HiddenRows    int
	DebugMode     bool
	DisableGhost  bool
	ControlScheme ControlScheme
}

// New returns a new game with the specified specifications
func New(c Config) *Game {
	initPieces := tetrimino.NewSet(c.Width, c.Height+c.HiddenRows)
	piece, pieceSet := initPieces[0], initPieces[1:]
	return &Game{
		inputreader: inputreader.NewTermReader(c.Term),
		canvas: canvas.New(canvas.Config{
			Term:       c.Term,
			Width:      c.Width,
			Height:     c.Height,
			Background: canvas.White,
			DebugMode:  c.DebugMode,
		}),
		board: board.New(
			canvas.White,
			c.Width, c.Height,
			c.HiddenRows,
		),
		currentPiece:  piece,
		newPieceSet:   tetrimino.NewSet,
		nextPieces:    pieceSet,
		level:         1,
		debugMode:     c.DebugMode,
		disableGhost:  c.DisableGhost,
		controlScheme: c.ControlScheme,
	}
}

// Run takes care of the core game functionality
func (g *Game) Run(done chan bool) (chan int, chan error) {
	var (
		runErr   = make(chan error)
		endScore = make(chan int)
	)
	controlMap, err := g.controlScheme.controlMap()
	if err != nil {
		runErr <- err
		return endScore, runErr
	}
	input, readErr := translateInput(done, g.inputreader, controlMap)

	// add initial piece to canvas
	g.addPieceToBoard(g.currentPiece)
	g.ghostPiece = g.findGhostPiece()

	g.canvas.UpdateCells(g.board.Cells())

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
			var gravity <-chan time.Time
			if !g.debugMode {
				gravity = time.After(g.level.gTime())
			}
			select {
			case err := <-readErr:
				runErr <- err
				return
			case <-done:
				return
			case <-gravity:
				if err := g.handleInput(moveDown, endScore); err != nil {
					runErr <- err
					return
				}
			case in := <-input:
				if g.debugMode {
					fmt.Printf("User input: %s\n", in)
				}

				if err := g.handleInput(in, endScore); err != nil {
					runErr <- err
					return
				}

			}
		}
	}()
	return endScore, runErr
}

func (g *Game) addPieceToBoard(piece tetrimino.Tetrimino) {
	var (
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
func (g *Game) pieceAtBottom(piece tetrimino.Tetrimino) bool {
	var (
		topL        = piece.ContainingBox().TopLeft
		blocks      = piece.Blocks()
		pieceCoords = make(map[tetrimino.Coordinates]bool)
	)

	if piece.YMin().Y == 0 {
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
	if (input == moveLeft || input == moveRight || (input == moveUp && g.debugMode)) && (g.pieceOutOfBounds() || g.pieceConflicts(topL, blocks)) {
		// move back to original spot
		if opposite := input.opposite(); opposite != ignore {
			g.movePiece(opposite)
			g.ghostPiece = g.findGhostPiece()
			return nil
		}
	}
	if (input == rotateLeft || input == rotateRight) && (g.pieceOutOfBounds() || g.pieceConflicts(topL, blocks)) {
		if !g.resolveRotation() {
			// move back to original spot
			if opposite := input.opposite(); opposite != ignore {
				g.movePiece(opposite)
				g.ghostPiece = g.findGhostPiece()
				return nil
			}
		}
	}
	g.ghostPiece = g.findGhostPiece()

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

	// update cell at pieces new position
	g.addPieceToBoard(g.currentPiece)
	g.ghostPiece = g.findGhostPiece()

	// generate new current piece if at bottom or on top of another piece
	if g.pieceAtBottom(g.currentPiece) {
		// check if any rows can be cleared
		// TODO: add scoring
		g.board.ClearFullRows()

		if g.pieceAtTop() {
			// still render game-over state
			g.canvas.UpdateCells(g.board.Cells())
			if err := g.canvas.Render(); err != nil {
				return err
			}
			endScore <- 0
			return nil
		}

		g.currentPiece = g.nextPiece()
		g.ghostPiece = g.findGhostPiece()

		// add new piece to canvas
		g.addPieceToBoard(g.currentPiece)
		if g.pieceAtBottom(g.currentPiece) {
			// new piece already at bottom -> game over
			g.canvas.UpdateCells(g.board.Cells())
			if err := g.canvas.Render(); err != nil {
				return err
			}
			endScore <- 0
			return nil
		}
	}

	if !g.disableGhost {
		newBoard := g.boardWithGhost()
		g.canvas.UpdateCells(newBoard.Cells())
	} else {
		g.canvas.UpdateCells(g.board.Cells())
	}

	return g.canvas.Render()
}

func (g *Game) movePiece(input userInput) {
	var (
		piece = g.currentPiece
	)

	switch input {
	case moveLeft:
		piece.MoveLeft()
	case moveDown:
		piece.MoveDown()
	case moveUp:
		if g.debugMode {
			piece.MoveUp()
		} else {
			// hard drop
			g.ghostPiece.ToggleGhost()
			g.currentPiece = g.ghostPiece
			g.ghostPiece = nil
		}
	case moveRight:
		piece.MoveRight()
	case rotateLeft:
		piece.RotateCounter()
	case rotateRight:
		piece.RotateClockwise()
	}
}

func (g *Game) nextPiece() tetrimino.Tetrimino {
	var nextPiece tetrimino.Tetrimino
	if len(g.nextPieces) == 1 {
		var (
			boardWidth  = len(g.board.Blocks[0])
			boardHeight = len(g.board.Blocks)
		)
		nextPiece = g.nextPieces[0]
		g.nextPieces = g.newPieceSet(boardWidth, boardHeight)
		return nextPiece
	}
	nextPiece, g.nextPieces = g.nextPieces[0], g.nextPieces[1:]
	return nextPiece
}

func (g *Game) resolveRotation() bool {
	var (
		piece  = g.currentPiece
		topL   = g.currentPiece.ContainingBox().TopLeft
		blocks = g.currentPiece.Blocks()
	)

	for _, rotationTest := range piece.RotationTests() {
		rotationTest.ApplyTest()
		if !(g.pieceOutOfBounds() || g.pieceConflicts(topL, blocks)) {
			return true
		}

		rotationTest.RevertTest()
	}

	return false
}

func (g *Game) findGhostPiece() tetrimino.Tetrimino {
	var (
		ghost = g.currentPiece.SpawnGhost()
	)

	for !g.pieceAtBottom(ghost) {
		ghost.MoveDown()
	}

	return ghost
}

func pieceCoords(piece tetrimino.Tetrimino, boardBlocks [][]*board.Block) map[tetrimino.Coordinates]bool {
	var (
		topL        = piece.ContainingBox().TopLeft
		blocks      = piece.Blocks()
		pieceCoords = make(map[tetrimino.Coordinates]bool)
	)

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
	return pieceCoords
}

// the board with the ghost piece included
// should NOT modify actual game board since ghost piece is irrelevant to game logic
func (g *Game) boardWithGhost() *board.Board {
	var (
		ghost       = g.ghostPiece
		ghostTopL   = ghost.ContainingBox().TopLeft
		ghostBlocks = ghost.Blocks()
		g2          = *g
		newBoard    = *g2.board
		newBlocks   [][]*board.Block
	)

	currentPieceCoords := pieceCoords(g.currentPiece, g.board.Blocks)

	for i := range g.board.Blocks {
		row := []*board.Block{}
		for j := range g.board.Blocks[i] {
			row = append(row, g.board.Blocks[i][j])
		}
		newBlocks = append(newBlocks, row)
	}
	newBoard.Blocks = newBlocks

	for i := range ghostBlocks {
		for j, block := range ghostBlocks[i] {
			if block == nil {
				continue
			}
			x := ghostTopL.X + j
			y := ghostTopL.Y - i

			if currentPieceCoords[tetrimino.Coordinates{X: x, Y: y}] {
				// active piece should be displayed in case of conflict with ghost
				continue
			}
			newBoard.Blocks[y][x] = block
		}
	}
	return &newBoard
}
