package canvas

import (
	"bytes"
	"io"
	"runtime"
	"strconv"
)

// Defaults for canvas
const (
	DefaultBackground = White
	DefaultWidth      = 20
	DefaultHeight     = 20
)

// Canvas represents a way to render to game to the user
type Canvas interface {
	Init() error
	Render() error
	UpdateCells(newCells [][]Cell)
}

// TermCanvas represents what is actually rendered to the user via the terminal
type TermCanvas struct {
	dest       io.Writer
	background Color
	cells      [][]Cell
	debugMode  bool
	width      int
	height     int
}

// New returns a new canvas
func New(term io.Writer, opts ...Option) *TermCanvas {
	var cells = [][]Cell{}

	t := &TermCanvas{
		dest:       term,
		background: DefaultBackground,
		width:      DefaultWidth,
		height:     DefaultHeight,
	}

	for i := range opts {
		opts[i].ApplyToCanvas(t)
	}

	for i := 0; i < t.height; i++ {
		row := make([]Cell, t.width)
		cells = append(cells, row)
	}
	t.cells = cells

	return t
}

// Init sets up the canvas in order to be written to
// for now this just clears the entire screen
func (c *TermCanvas) Init() error {
	return c.clear()
}

// Render renders the current canvas
func (c *TermCanvas) Render() error {
	var b bytes.Buffer
	if !c.debugMode {
		// reset cursor
		b.Write(resetTermCursor)
	}

	for _, row := range c.cells {
		for _, cell := range row {
			if cell == nil {
				b.WriteString(c.background.String())
				b.WriteString(block)
				continue
			}
			b.WriteString(cell.String())
		}
		b.WriteByte('\n')
		b.Write(resetControl)
	}
	b.Write(resetControl)

	// clear any potential formatting
	_, err := c.dest.Write(b.Bytes())
	return err
}

// UpdateCells updates the cells to be rendered
func (c *TermCanvas) UpdateCells(newCells [][]Cell) {
	c.cells = newCells
}

func resetString() string {
	switch runtime.GOOS {
	case "linux":
		return "\033c"
	case "darwin":
		return "\033[2J"
	default:
		return "\033[2J"
	}
}

func (c *TermCanvas) clear() error {
	_, err := c.dest.Write([]byte(resetString()))
	if err != nil {
		return err
	}
	return nil
}

var resetTermCursor = []byte{'\x1b', '[', '0', ';', '0', 'H'}

func (c *TermCanvas) setCursor(x, y int) []byte {
	var b bytes.Buffer
	b.WriteString("\033[")
	b.WriteString(strconv.Itoa(x))
	b.WriteString(";")
	b.WriteString(strconv.Itoa(y))
	b.WriteString("H")

	return b.Bytes()
}
