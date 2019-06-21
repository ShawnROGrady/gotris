package canvas

import (
	"bytes"
	"io"
	"runtime"
	"strconv"
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
	Background Color
	cells      [][]Cell
	debugMode  bool
}

// Config represents the configuration params for a terminal canvas
type Config struct {
	Term       io.Writer
	Width      int
	Height     int
	Background Color
	DebugMode  bool
}

// New returns a new canvas
func New(c Config) *TermCanvas {
	var cells = [][]Cell{}

	for i := 0; i < c.Height; i++ {
		row := make([]Cell, c.Width)
		cells = append(cells, row)
	}

	return &TermCanvas{
		dest:       c.Term,
		Background: c.Background,
		cells:      cells,
		debugMode:  c.DebugMode,
	}
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
		b.Write(c.setCursor(0, 0))
	}

	for _, row := range c.cells {
		for _, cell := range row {
			if cell == nil {
				b.WriteString(c.Background.String())
				b.WriteString(block)
				continue
			}
			b.WriteString(cell.String())
		}
		b.WriteString("\n")
		b.WriteString(Reset.String())
	}
	b.WriteString(Reset.String())

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

func (c *TermCanvas) setCursor(x, y int) []byte {
	var b bytes.Buffer
	b.WriteString("\033[")
	b.WriteString(strconv.Itoa(x))
	b.WriteString(";")
	b.WriteString(strconv.Itoa(y))
	b.WriteString("H")

	return b.Bytes()
}
