package board

import "github.com/ShawnROGrady/gotris/pkg/canvas"

// Option represents a configuration option for the board
type Option interface {
	ApplyToBoard(b *Board)
}

// WithBackground returns an option specifies the background for the canvas and board
func WithBackground(c canvas.Color) Option {
	return withBackground(c)
}

type withBackground canvas.Color

func (w withBackground) ApplyToBoard(b *Board) {
	b.background = canvas.Color(w)
}

// WithWidthScale returns an option that specifies how blocks should be scaled to cells
func WithWidthScale(scale int) Option {
	return withWidthScale(scale)
}

type withWidthScale int

func (w withWidthScale) ApplyToBoard(b *Board) {
	b.widthScale = int(w)
}

// WithHiddenRows returns an option that specifies how many rows of the board shouldn't be rendered
func WithHiddenRows(rows int) Option {
	return withHiddenRows(rows)
}

type withHiddenRows int

func (w withHiddenRows) ApplyToBoard(b *Board) {
	b.hiddenRows = int(w)
}

// WithWidth returns an option the specifies the width of the board
func WithWidth(width int) Option {
	return withWidth(width)
}

type withWidth int

func (w withWidth) ApplyToBoard(b *Board) {
	b.width = int(w)
}

// WithHeight returns an option the specifies the height of the board
func WithHeight(height int) Option {
	return withHeight(height)
}

type withHeight int

func (w withHeight) ApplyToBoard(b *Board) {
	b.height = int(w)
}
