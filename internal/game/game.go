package game

import (
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/ShawnROGrady/gotris/internal/canvas"
	"github.com/ShawnROGrady/gotris/internal/game/board"
	"github.com/ShawnROGrady/gotris/internal/game/tetrimino"
	"github.com/ShawnROGrady/gotris/internal/inputreader"
)

// Defaults for the game
const (
	defaultColor = canvas.White
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
	currentScore  int
	linesCleared  int
	debugMode     bool
	disableGhost  bool
	disableSide   bool
	controlScheme ControlScheme
	widthScale    int
	gameCells     gameCells
	color         canvas.Color
	mutex         *sync.Mutex
}

type gameCells struct {
	nextPiece [][]canvas.Cell
	score     [][]canvas.Cell
	controls  [][]canvas.Cell
}

// New returns a new game with the specified specifications
func New(termReader io.Reader, termWriter io.Writer, opts ...Option) *Game {
	g := &Game{
		inputreader:   inputreader.NewTermReader(termReader),
		newPieceSet:   tetrimino.NewSet,
		level:         0,
		currentScore:  0,
		linesCleared:  0,
		widthScale:    board.DefaultWidthScale,
		color:         defaultColor,
		controlScheme: HomeRow(),
		mutex:         &sync.Mutex{},
	}

	var (
		boardOpts  = []board.Option{}
		canvasOpts = []canvas.Option{}
	)

	for i := range opts {
		opts[i].Apply(g)
		if boardOpt, ok := opts[i].(board.Option); ok {
			boardOpts = append(boardOpts, boardOpt)
		}
		if canvasOpt, ok := opts[i].(canvas.Option); ok {
			canvasOpts = append(canvasOpts, canvasOpt)
		}
	}

	// initialize the games canvas (what's rendered)
	c := canvas.New(termWriter, canvasOpts...)
	gCanvas := &gCanvas{
		canvas:   c,
		newCells: make(chan [][]canvas.Cell),
		mut:      &sync.Mutex{},
	}
	g.canvas = gCanvas

	// initialize the games board (used for game logic)
	board := board.New(boardOpts...)
	g.board = board

	// initialize first pieces
	initPieces := tetrimino.NewSet(boardWidth(board), boardHeight(board))
	piece, pieceSet := initPieces[0], initPieces[1:]
	g.currentPiece = piece
	g.nextPieces = pieceSet

	return g
}

// Run takes care of the core game functionality
func (g *Game) Run(done chan bool) (chan int, chan error) {
	var (
		runErr   = make(chan error)
		endScore = make(chan int)
	)
	controlMap := g.controlScheme.controlMap()

	input, readErr := translateInput(done, g.inputreader, controlMap)

	// add initial piece to canvas
	g.addPieceToBoard(g.currentPiece)
	g.ghostPiece = g.findGhostPiece()

	if !g.disableSide {
		// add initial sidebar cells
		g.updateCells(g.board.Background())
	}

	// initialize the canvas
	if err := g.canvas.Init(); err != nil {
		runErr <- err
		return endScore, runErr
	}

	if gCanvas, ok := g.canvas.(*gCanvas); ok {
		gCanvas.run(done)
	}

	// render initial canvas
	g.canvas.UpdateCells(g.cells(g.board))
	if err := g.canvas.Render(); err != nil {
		runErr <- err
		return endScore, runErr
	}

	go func() {
		var gravity <-chan time.Time
		if !g.debugMode {
			// set initial gravity
			gravity = time.After(g.level.gTime())
		}
		for {
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
				if !g.debugMode {
					gravity = time.After(g.level.gTime())
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
			if x < 0 || x > boardWidth(g.board)-1 {
				return true
			}

			if y < 0 || y > boardHeight(g.board)-1 {
				return true
			}
		}
	}
	return false
}

// checks if the current piece is in the hidden row(s)
func (g *Game) pieceAtTop() bool {
	return g.currentPiece.YMax().Y > len(g.board.Blocks)-g.board.HiddenRows()-1
}

// current piece is at minimum vertical position
// either at bottom or on top of another piece
func (g *Game) pieceAtBottom(piece tetrimino.Tetrimino) bool {
	var (
		topL        = piece.ContainingBox().TopLeft
		blocks      = piece.Blocks()
		pieceCoords = make(map[tetrimino.Coordinates]bool)
	)

	if piece.YMin().Y <= 0 {
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
	var (
		topL     = g.currentPiece.ContainingBox().TopLeft
		blocks   = g.currentPiece.Blocks()
		canSlide = true
	)

	// should expect exclusive access when handling input
	g.mutex.Lock()
	defer g.mutex.Unlock()

	g.movePiece(input)

	if input == rotateLeft || input == rotateRight {
		if g.pieceOutOfBounds() || g.pieceConflicts(topL, blocks) {
			if !g.resolveRotation() {
				// move back to original spot
				if opposite := input.opposite(); opposite != ignore {
					g.movePiece(opposite)
					g.ghostPiece = g.findGhostPiece()
					return nil
				}
			}
		}
	}

	// new space already occupied
	if (input == moveLeft || input == moveRight || (input == moveUp && g.debugMode)) && (g.pieceOutOfBounds() || g.pieceConflicts(topL, blocks)) {
		// move back to original spot
		if opposite := input.opposite(); opposite != ignore {
			g.movePiece(opposite)
			g.ghostPiece = g.findGhostPiece()
			return nil
		}
	}

	// piece was already at the bottom
	if (!g.debugMode && input == moveUp) || g.pieceOutOfBounds() || g.pieceConflicts(topL, blocks) {
		canSlide = false
		if g.pieceOutOfBounds() || g.pieceConflicts(topL, blocks) {
			if opposite := input.opposite(); opposite != ignore {
				g.movePiece(opposite)
				g.ghostPiece = g.findGhostPiece()
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
	if g.pieceAtBottom(g.currentPiece) && !canSlide {
		// check if any rows can be cleared
		linesCleared := g.board.ClearFullRows()
		if linesCleared != 0 {
			g.linesCleared += linesCleared
			g.currentScore += g.level.linePoints(linesCleared)
			newLevel := g.level.updatedLevel(g.linesCleared)
			g.level = newLevel
		}

		if g.pieceAtTop() {
			// still render game-over state
			g.canvas.UpdateCells(g.cells(g.board))
			if err := g.canvas.Render(); err != nil {
				return err
			}
			endScore <- g.currentScore
			return nil
		}

		g.currentPiece = g.nextPiece()
		g.ghostPiece = g.findGhostPiece()

		if !g.disableSide {
			// update cells to include new next + updated score
			g.updateCells(g.board.Background())
		}

		// add new piece to canvas
		g.addPieceToBoard(g.currentPiece)
		if g.pieceAtBottom(g.currentPiece) {
			// new piece already at bottom -> game over
			g.canvas.UpdateCells(g.cells(g.board))
			if err := g.canvas.Render(); err != nil {
				return err
			}
			endScore <- g.currentScore
			return nil
		}
	}

	if !g.disableGhost {
		newBoard := g.boardWithGhost()
		g.canvas.UpdateCells(g.cells(newBoard))
	} else {
		g.canvas.UpdateCells(g.cells(g.board))
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
			boardWidth  = boardWidth(g.board)
			boardHeight = boardHeight(g.board)
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

func (g *Game) updateCells(background canvas.Color) {
	nextPiece := g.nextPieces[0]
	formattedBlocks := centerBlocks(nextPiece.Blocks(), tetrimino.MaxWidth, tetrimino.MaxHeight)
	nextPieceCells := canvas.Box(board.BlockGridCells(formattedBlocks, background, g.widthScale), "NEXT")

	currentScore := fmt.Sprintf("Score: %d\nLevel: %d", g.currentScore, g.level)
	scoreCells := canvas.Box(canvas.CellsFromString(currentScore, g.color), "")

	schemeCells := canvas.Box(canvas.CellsFromString(g.controlScheme.Description(), g.color), "CONTROLS")

	g.gameCells = gameCells{
		nextPiece: nextPieceCells,
		score:     scoreCells,
		controls:  schemeCells,
	}
}

func (g *Game) cells(b *board.Board) [][]canvas.Cell {
	gameCells := canvas.Box(b.Cells(), "GAME")

	if !g.disableSide {
		nextPieceCells := g.gameCells.nextPiece
		for i := range nextPieceCells {
			gameCells[i] = append(gameCells[i], nextPieceCells[i]...)
		}

		scoreCells := g.gameCells.score
		for i := range scoreCells {
			gameCells[i+len(nextPieceCells)] = append(gameCells[i+len(nextPieceCells)], scoreCells[i]...)
		}

		schemeCells := g.gameCells.controls
		for i := range schemeCells {
			gameCells[i+len(nextPieceCells)+len(scoreCells)] = append(gameCells[i+len(nextPieceCells)+len(scoreCells)], schemeCells[i]...)
		}
	}

	return gameCells
}
