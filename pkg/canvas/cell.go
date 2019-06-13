package canvas

import (
	"strings"
)

const (
	// block elements
	block                  = "\u2588"
	lightTransparentBlock  = "\u2591"
	mediumTransparentBlock = "\u2592"
	darkTransparentBlock   = "\u2593"

	// pipe elements
	horizontalBarPipe = "\u2550"
	verticalBarPipe   = "\u2551"
	topLeftPipe       = "\u2554"
	topRightPipe      = "\u2557"
	bottomLeftPipe    = "\u255A"
	bottomRightPipe   = "\u255D"
)

// Cell represents an item to be rendered on the canvas
type Cell interface {
	String() string
}

// BlockCell represents a single cell on the canvas
type BlockCell struct {
	Color       Color
	Background  Color
	Transparent bool
}

func (c *BlockCell) String() string {
	if c.Transparent {
		return c.Background.background().decorate(
			// TODO: allow for level of transparency to be modifiable
			c.Color.decorate(mediumTransparentBlock),
		)
	}
	return c.Color.background().decorate(
		c.Color.decorate(block),
	)
}

// PipeType represents a type of pipe cell
type PipeType int

// the available pipe cell types
const (
	HorizontalBar PipeType = iota
	VerticalBar
	TopLeft
	TopRight
	BottomLeft
	BottomRight
)

// PipeCell represents a single pipe element on the canvas
type PipeCell struct {
	Type       PipeType
	Color      Color
	Background Color
}

func (p *PipeCell) String() string {
	switch p.Type {
	case HorizontalBar:
		return p.Color.decorate(horizontalBarPipe)
	case VerticalBar:
		return p.Color.decorate(verticalBarPipe)
	case TopLeft:
		return p.Color.decorate(topLeftPipe)
	case TopRight:
		return p.Color.decorate(topRightPipe)
	case BottomLeft:
		return p.Color.decorate(bottomLeftPipe)
	case BottomRight:
		return p.Color.decorate(bottomRightPipe)
	default:
		return ""
	}
}

// Box wraps a block of cells in a piped box
func Box(inner [][]Cell) [][]Cell {
	boxedCells := [][]Cell{}
	topRow := []Cell{&PipeCell{Type: TopLeft}}
	bottomRow := []Cell{&PipeCell{Type: BottomLeft}}
	bar := []Cell{}
	for range inner[0] {
		bar = append(bar, &PipeCell{Type: HorizontalBar})
	}

	topRow = append(topRow, bar...)
	topRow = append(topRow, &PipeCell{Type: TopRight})

	bottomRow = append(bottomRow, bar...)
	bottomRow = append(bottomRow, &PipeCell{Type: BottomRight})

	boxedCells = append(boxedCells, topRow)

	for i := range inner {
		row := inner[i]
		row = append([]Cell{&PipeCell{Type: VerticalBar}}, row...)
		row = append(row, &PipeCell{Type: VerticalBar})
		boxedCells = append(boxedCells, row)
	}
	boxedCells = append(boxedCells, bottomRow)
	return boxedCells
}

// TextCell is a piece of text to be displayed
type TextCell struct {
	Text  string
	Color Color
}

func (t *TextCell) String() string {
	return t.Color.decorate(t.Text)
}

// CellsFromString constructs a grid of TextCells from a provided string
func CellsFromString(s string, color Color) [][]Cell {
	var (
		maxLine int
		lines   = strings.Split(s, "\n")
	)
	for _, line := range lines {
		if len(line) > maxLine {
			maxLine = len(line)
		}
	}

	cells := make([][]Cell, len(lines))

	for i, line := range lines {
		row := make([]Cell, maxLine)
		chars := strings.Split(line, "")
		for j := 0; j < maxLine; j++ {
			if j >= len(chars) {
				row[j] = &TextCell{
					Text:  " ",
					Color: Reset,
				}
				continue
			}
			row[j] = &TextCell{
				Text:  chars[j],
				Color: color,
			}
		}
		cells[i] = row
	}

	return cells
}
