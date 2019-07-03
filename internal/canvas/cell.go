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
	Type  PipeType
	Color Color
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
func Box(inner [][]Cell, caption string) [][]Cell {
	boxedCells := [][]Cell{}

	topRow := []Cell{&PipeCell{Type: TopLeft}}
	bottomRow := []Cell{&PipeCell{Type: BottomLeft}}

	var bar []Cell
	if len(caption) > len(inner[0]) {
		bar = make([]Cell, len(caption))
	} else {
		bar = make([]Cell, len(inner[0]))
	}
	for i := range bar {
		bar[i] = &PipeCell{Type: HorizontalBar}
	}

	bottomRow = append(bottomRow, bar...)
	bottomRow = append(bottomRow, &PipeCell{Type: BottomRight})

	// center the text on the top bar
	if caption != "" {
		var (
			textStart = (len(bar) - len(caption)) / 2
		)
		for i := 0; i < len(caption); i++ {
			bar[i+textStart] = &TextCell{
				Text: string(caption[i]),
			}
		}
	}

	topRow = append(topRow, bar...)
	topRow = append(topRow, &PipeCell{Type: TopRight})

	boxedCells = append(boxedCells, topRow)

	// add vertical bars on either side of the inner cells
	for i := range inner {
		var (
			innerRow   = inner[i]
			innerStart = (len(bar) + 2 - len(innerRow)) / 2
			row        = make([]Cell, len(bar)+2)
		)
		row[0] = &PipeCell{Type: VerticalBar}
		row[len(row)-1] = &PipeCell{Type: VerticalBar}
		for j := 1; j < len(row)-1; j++ {
			if j < innerStart || j > len(innerRow)+innerStart-1 {
				row[j] = &TextCell{
					Color: Reset,
					Text:  " ",
				}
				continue
			}
			row[j] = innerRow[j-innerStart]
		}
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
