package game

import (
	"github.com/ShawnROGrady/gotris/internal/canvas"
	"github.com/ShawnROGrady/gotris/internal/game/board"
)

// Option represents a game option
type Option interface {
	Apply(g *Game)
}

// WithControlScheme returns an Option which specifies the control scheme
func WithControlScheme(scheme ControlScheme) Option {
	return withControlScheme{scheme: scheme}
}

type withControlScheme struct {
	scheme ControlScheme
}

func (w withControlScheme) Apply(g *Game) {
	g.controlScheme = w.scheme
}

// WithoutGhost returns an option that disables the ghost piece
func WithoutGhost() Option {
	return withoutGhost{}
}

type withoutGhost struct{}

func (w withoutGhost) Apply(g *Game) {
	g.disableGhost = true
}

// WithBackground returns an option specifies the background for the canvas and board
func WithBackground(c canvas.Color) Option {
	return withBackground(c)
}

type withBackground canvas.Color

func (w withBackground) Apply(g *Game) {}

func (w withBackground) ApplyToBoard(b *board.Board) {
	board.WithBackground(canvas.Color(w)).ApplyToBoard(b)
}

func (w withBackground) ApplyToCanvas(c *canvas.TermCanvas) {
	canvas.WithBackground(canvas.Color(w)).ApplyToCanvas(c)
}

// WithColor returns an option specifies the primary color for the game
func WithColor(c canvas.Color) Option {
	return withColor(c)
}

type withColor canvas.Color

func (w withColor) Apply(g *Game) {
	g.color = canvas.Color(w)
}

// WithDebugMode returns an option that specifies that the game and canvas should use debug behaviour
func WithDebugMode() Option {
	return withDebugMode{}
}

type withDebugMode struct{}

func (w withDebugMode) Apply(g *Game) {
	g.debugMode = true
}

func (w withDebugMode) ApplyToCanvas(c *canvas.TermCanvas) {
	canvas.WithDebugMode().ApplyToCanvas(c)
}

// WithoutSide returns an option that disables the side bar
func WithoutSide() Option {
	return withoutSide{}
}

type withoutSide struct{}

func (w withoutSide) Apply(g *Game) {
	g.disableSide = true
}

// WithDimensions returns an option that specifies the dimensions of the board and canvas
func WithDimensions(d dimensions) Option {
	return d
}

type dimensions struct {
	width      int
	height     int
	widthScale int
}

//type withDimensions dimensions

func (w dimensions) Apply(g *Game) {
	g.widthScale = int(w.widthScale)
}

func (w dimensions) ApplyToBoard(b *board.Board) {
	board.WithWidthScale(int(w.widthScale)).ApplyToBoard(b)
	board.WithWidth(int(w.width)).ApplyToBoard(b)
	board.WithHeight(int(w.height)).ApplyToBoard(b)
}

func (w dimensions) ApplyToCanvas(c *canvas.TermCanvas) {
	canvas.WithWidth(int(w.width) * int(w.widthScale)).ApplyToCanvas(c)
	canvas.WithHeight(int(w.height)).ApplyToCanvas(c)
}

// WithHiddenRows returns an option that specifies how many rows of the board shouldn't be rendered
func WithHiddenRows(rows int) Option {
	return withHiddenRows(rows)
}

type withHiddenRows int

func (w withHiddenRows) Apply(g *Game) {}

func (w withHiddenRows) ApplyToBoard(b *board.Board) {
	board.WithHiddenRows(int(w)).ApplyToBoard(b)
}

// WithInitialLevel returns an option that specifies the initial level
func WithInitialLevel(level int) Option {
	return withInitialLevel(level)
}

type withInitialLevel int

func (w withInitialLevel) Apply(g *Game) {
	g.level = level(w)
}
