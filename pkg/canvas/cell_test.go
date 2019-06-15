package canvas

import "testing"

var cellColorTests = map[string]struct {
	Color          Color
	Background     Color
	Transparent    bool
	ExpectedString string
}{
	"Black solid": {
		Color:          Black,
		ExpectedString: "\u001b[40m\u001b[30m\u2588",
	},
	"Red solid": {
		Color:          Red,
		ExpectedString: "\u001b[41m\u001b[31m\u2588",
	},
	"Green solid": {
		Color:          Green,
		ExpectedString: "\u001b[42m\u001b[32m\u2588",
	},
	"Yellow solid": {
		Color:          Yellow,
		ExpectedString: "\u001b[43m\u001b[33m\u2588",
	},
	"Blue solid": {
		Color:          Blue,
		ExpectedString: "\u001b[44m\u001b[34m\u2588",
	},
	"Magenta solid": {
		Color:          Magenta,
		ExpectedString: "\u001b[45m\u001b[35m\u2588",
	},
	"Cyan solid": {
		Color:          Cyan,
		ExpectedString: "\u001b[46m\u001b[36m\u2588",
	},
	"White solid": {
		Color:          White,
		ExpectedString: "\u001b[47m\u001b[37m\u2588",
	},
	"BrightBlack solid": {
		Color:          BrightBlack,
		ExpectedString: "\u001b[40m\u001b[30;1m\u2588",
	},
	"BrightRed solid": {
		Color:          BrightRed,
		ExpectedString: "\u001b[41m\u001b[31;1m\u2588",
	},
	"BrightGreen solid": {
		Color:          BrightGreen,
		ExpectedString: "\u001b[42m\u001b[32;1m\u2588",
	},
	"BrightYellow solid": {
		Color:          BrightYellow,
		ExpectedString: "\u001b[43m\u001b[33;1m\u2588",
	},
	"BrightBlue solid": {
		Color:          BrightBlue,
		ExpectedString: "\u001b[44m\u001b[34;1m\u2588",
	},
	"BrightMagenta solid": {
		Color:          BrightMagenta,
		ExpectedString: "\u001b[45m\u001b[35;1m\u2588",
	},
	"BrightCyan solid": {
		Color:          BrightCyan,
		ExpectedString: "\u001b[46m\u001b[36;1m\u2588",
	},
	"BrightWhite solid": {
		Color:          BrightWhite,
		ExpectedString: "\u001b[47m\u001b[37;1m\u2588",
	},
	"Orange solid": {
		Color:          Orange,
		ExpectedString: "\u001b[48;5;208m\u001b[38;5;208m\u2588",
	},
	"Cyan transparent white background": {
		Color:          Cyan,
		Transparent:    true,
		Background:     White,
		ExpectedString: "\u001b[47m\u001b[36m\u2592",
	},
}

func TestCellColors(t *testing.T) {
	for testName, test := range cellColorTests {
		cell := &BlockCell{
			Color:       test.Color,
			Transparent: test.Transparent,
			Background:  test.Background,
		}

		cellString := cell.String()
		if cellString != test.ExpectedString {
			t.Errorf("Unexpected string for test case '%s' background [expected=%s,%s actual=%s]%s", testName, test.ExpectedString, Reset, cellString, Reset)
		}
	}
}

var boxTests = map[string]struct {
	inner         [][]Cell
	caption       string
	expectedCells [][]Cell
}{
	"4x4 blocks no caption": {
		inner: [][]Cell{
			{&BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}},
			{&BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}},
			{&BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}},
			{&BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}},
		},
		caption: "",
		expectedCells: [][]Cell{
			{&PipeCell{Type: TopLeft}, &PipeCell{Type: HorizontalBar}, &PipeCell{Type: HorizontalBar}, &PipeCell{Type: HorizontalBar}, &PipeCell{Type: HorizontalBar}, &PipeCell{Type: TopRight}},
			{&PipeCell{Type: VerticalBar}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &PipeCell{Type: VerticalBar}},
			{&PipeCell{Type: VerticalBar}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &PipeCell{Type: VerticalBar}},
			{&PipeCell{Type: VerticalBar}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &PipeCell{Type: VerticalBar}},
			{&PipeCell{Type: VerticalBar}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &PipeCell{Type: VerticalBar}},
			{&PipeCell{Type: BottomLeft}, &PipeCell{Type: HorizontalBar}, &PipeCell{Type: HorizontalBar}, &PipeCell{Type: HorizontalBar}, &PipeCell{Type: HorizontalBar}, &PipeCell{Type: BottomRight}},
		},
	},
	"4x4 blocks 2 letter caption": {
		inner: [][]Cell{
			{&BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}},
			{&BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}},
			{&BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}},
			{&BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}},
		},
		caption: "HI",
		expectedCells: [][]Cell{
			{&PipeCell{Type: TopLeft}, &PipeCell{Type: HorizontalBar}, &TextCell{Text: "H"}, &TextCell{Text: "I"}, &PipeCell{Type: HorizontalBar}, &PipeCell{Type: TopRight}},
			{&PipeCell{Type: VerticalBar}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &PipeCell{Type: VerticalBar}},
			{&PipeCell{Type: VerticalBar}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &PipeCell{Type: VerticalBar}},
			{&PipeCell{Type: VerticalBar}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &PipeCell{Type: VerticalBar}},
			{&PipeCell{Type: VerticalBar}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &PipeCell{Type: VerticalBar}},
			{&PipeCell{Type: BottomLeft}, &PipeCell{Type: HorizontalBar}, &PipeCell{Type: HorizontalBar}, &PipeCell{Type: HorizontalBar}, &PipeCell{Type: HorizontalBar}, &PipeCell{Type: BottomRight}},
		},
	},
	"4x4 blocks 3 letter caption": {
		inner: [][]Cell{
			{&BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}},
			{&BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}},
			{&BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}},
			{&BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}},
		},
		caption: "HEY",
		expectedCells: [][]Cell{
			{&PipeCell{Type: TopLeft}, &TextCell{Text: "H"}, &TextCell{Text: "E"}, &TextCell{Text: "Y"}, &PipeCell{Type: HorizontalBar}, &PipeCell{Type: TopRight}},
			{&PipeCell{Type: VerticalBar}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &PipeCell{Type: VerticalBar}},
			{&PipeCell{Type: VerticalBar}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &PipeCell{Type: VerticalBar}},
			{&PipeCell{Type: VerticalBar}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &PipeCell{Type: VerticalBar}},
			{&PipeCell{Type: VerticalBar}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &PipeCell{Type: VerticalBar}},
			{&PipeCell{Type: BottomLeft}, &PipeCell{Type: HorizontalBar}, &PipeCell{Type: HorizontalBar}, &PipeCell{Type: HorizontalBar}, &PipeCell{Type: HorizontalBar}, &PipeCell{Type: BottomRight}},
		},
	},
	"4x4 blocks 4 letter caption": {
		inner: [][]Cell{
			{&BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}},
			{&BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}},
			{&BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}},
			{&BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}},
		},
		caption: "TEST",
		expectedCells: [][]Cell{
			{&PipeCell{Type: TopLeft}, &TextCell{Text: "T"}, &TextCell{Text: "E"}, &TextCell{Text: "S"}, &TextCell{Text: "T"}, &PipeCell{Type: TopRight}},
			{&PipeCell{Type: VerticalBar}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &PipeCell{Type: VerticalBar}},
			{&PipeCell{Type: VerticalBar}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &PipeCell{Type: VerticalBar}},
			{&PipeCell{Type: VerticalBar}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &PipeCell{Type: VerticalBar}},
			{&PipeCell{Type: VerticalBar}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &PipeCell{Type: VerticalBar}},
			{&PipeCell{Type: BottomLeft}, &PipeCell{Type: HorizontalBar}, &PipeCell{Type: HorizontalBar}, &PipeCell{Type: HorizontalBar}, &PipeCell{Type: HorizontalBar}, &PipeCell{Type: BottomRight}},
		},
	},
	"4x4 blocks 5 letter caption": {
		inner: [][]Cell{
			{&BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}},
			{&BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}},
			{&BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}},
			{&BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}},
		},
		caption: "TEST2",
		expectedCells: [][]Cell{
			{&PipeCell{Type: TopLeft}, &TextCell{Text: "T"}, &TextCell{Text: "E"}, &TextCell{Text: "S"}, &TextCell{Text: "T"}, &TextCell{Text: "2"}, &PipeCell{Type: TopRight}},
			{&PipeCell{Type: VerticalBar}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &TextCell{Text: " ", Color: Reset}, &PipeCell{Type: VerticalBar}},
			{&PipeCell{Type: VerticalBar}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &TextCell{Text: " ", Color: Reset}, &PipeCell{Type: VerticalBar}},
			{&PipeCell{Type: VerticalBar}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &TextCell{Text: " ", Color: Reset}, &PipeCell{Type: VerticalBar}},
			{&PipeCell{Type: VerticalBar}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &TextCell{Text: " ", Color: Reset}, &PipeCell{Type: VerticalBar}},
			{&PipeCell{Type: BottomLeft}, &PipeCell{Type: HorizontalBar}, &PipeCell{Type: HorizontalBar}, &PipeCell{Type: HorizontalBar}, &PipeCell{Type: HorizontalBar}, &PipeCell{Type: HorizontalBar}, &PipeCell{Type: BottomRight}},
		},
	},
	"4x4 blocks 6 letter caption": {
		inner: [][]Cell{
			{&BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}},
			{&BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}},
			{&BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}},
			{&BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}},
		},
		caption: "TEST00",
		expectedCells: [][]Cell{
			{&PipeCell{Type: TopLeft}, &TextCell{Text: "T"}, &TextCell{Text: "E"}, &TextCell{Text: "S"}, &TextCell{Text: "T"}, &TextCell{Text: "0"}, &TextCell{Text: "0"}, &PipeCell{Type: TopRight}},
			{&PipeCell{Type: VerticalBar}, &TextCell{Text: " ", Color: Reset}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &TextCell{Text: " ", Color: Reset}, &PipeCell{Type: VerticalBar}},
			{&PipeCell{Type: VerticalBar}, &TextCell{Text: " ", Color: Reset}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &TextCell{Text: " ", Color: Reset}, &PipeCell{Type: VerticalBar}},
			{&PipeCell{Type: VerticalBar}, &TextCell{Text: " ", Color: Reset}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &TextCell{Text: " ", Color: Reset}, &PipeCell{Type: VerticalBar}},
			{&PipeCell{Type: VerticalBar}, &TextCell{Text: " ", Color: Reset}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &BlockCell{Color: Blue, Background: Blue}, &TextCell{Text: " ", Color: Reset}, &PipeCell{Type: VerticalBar}},
			{&PipeCell{Type: BottomLeft}, &PipeCell{Type: HorizontalBar}, &PipeCell{Type: HorizontalBar}, &PipeCell{Type: HorizontalBar}, &PipeCell{Type: HorizontalBar}, &PipeCell{Type: HorizontalBar}, &PipeCell{Type: HorizontalBar}, &PipeCell{Type: BottomRight}},
		},
	},
}

func TestBox(t *testing.T) {
	for testName, test := range boxTests {
		boxedCells := Box(test.inner, test.caption)
		if len(boxedCells) != len(test.expectedCells) {
			t.Fatalf("Unexpected number of rows for test case '%s' [expected=%d, actual=%d]", testName, len(test.expectedCells), len(boxedCells))
		}

		for i := range boxedCells {
			newRow := boxedCells[i]
			expectedRow := test.expectedCells[i]
			if len(newRow) != len(expectedRow) {
				t.Fatalf("Unexpected number of cells in row %d for test case '%s' [expected=%d, actual=%d]", i, testName, len(expectedRow), len(newRow))
			}

			for j := range newRow {
				if newRow[j] == nil && expectedRow[j] == nil {
					continue
				}
				if newRow[j] == nil && expectedRow[j] != nil {
					t.Errorf("cells[%d][%d] unexpectedly nil for test case '%s'", i, j, testName)
					continue
				}
				if newRow[j] != nil && expectedRow[j] == nil {
					t.Errorf("cells[%d][%d] unexpectedly non-nil for test case '%s'", i, j, testName)
					continue
				}
				if newRow[j].String() != expectedRow[j].String() {
					t.Errorf("Unexpected cells[%d][%d] value for test case '%s' [expected=%#v, actual=%#v]", i, j, testName, expectedRow[j], newRow[j])
				}
			}
		}
	}
}

var cellFromStringTests = map[string]struct {
	inputText     string
	inputColor    Color
	expectedCells [][]Cell
}{
	"3 lines equal length": {
		inputText:  "A: 1\nB: 2\nC: 3",
		inputColor: White,
		expectedCells: [][]Cell{
			{&TextCell{Text: "A", Color: White}, &TextCell{Text: ":", Color: White}, &TextCell{Text: " ", Color: White}, &TextCell{Text: "1", Color: White}},
			{&TextCell{Text: "B", Color: White}, &TextCell{Text: ":", Color: White}, &TextCell{Text: " ", Color: White}, &TextCell{Text: "2", Color: White}},
			{&TextCell{Text: "C", Color: White}, &TextCell{Text: ":", Color: White}, &TextCell{Text: " ", Color: White}, &TextCell{Text: "3", Color: White}},
		},
	},
	"3 lines un-equal length": {
		inputText:  "A:1\nB: 2\nC: 3",
		inputColor: White,
		expectedCells: [][]Cell{
			{&TextCell{Text: "A", Color: White}, &TextCell{Text: ":", Color: White}, &TextCell{Text: "1", Color: White}, &TextCell{Text: " ", Color: Reset}},
			{&TextCell{Text: "B", Color: White}, &TextCell{Text: ":", Color: White}, &TextCell{Text: " ", Color: White}, &TextCell{Text: "2", Color: White}},
			{&TextCell{Text: "C", Color: White}, &TextCell{Text: ":", Color: White}, &TextCell{Text: " ", Color: White}, &TextCell{Text: "3", Color: White}},
		},
	},
	"1 line": {
		inputText:  "A: 1",
		inputColor: White,
		expectedCells: [][]Cell{
			{&TextCell{Text: "A", Color: White}, &TextCell{Text: ":", Color: White}, &TextCell{Text: " ", Color: White}, &TextCell{Text: "1", Color: White}},
		},
	},
	"no text": {
		inputText:     "",
		inputColor:    White,
		expectedCells: [][]Cell{{}},
	},
}

func TestCellFromString(t *testing.T) {
	for testName, test := range cellFromStringTests {
		cells := CellsFromString(test.inputText, test.inputColor)

		if len(cells) != len(test.expectedCells) {
			t.Fatalf("Unexpected number of rows for test case '%s' [expected=%d, actual=%d]", testName, len(test.expectedCells), len(cells))
		}

		for i := range cells {
			newRow := cells[i]
			expectedRow := test.expectedCells[i]
			if len(newRow) != len(expectedRow) {
				t.Fatalf("Unexpected number of cells in row %d for test case '%s' [expected=%d, actual=%d]", i, testName, len(expectedRow), len(newRow))
			}

			for j := range newRow {
				if newRow[j] == nil && expectedRow[j] == nil {
					continue
				}
				if newRow[j] == nil && expectedRow[j] != nil {
					t.Errorf("cells[%d][%d] unexpectedly nil for test case '%s'", i, j, testName)
					continue
				}
				if newRow[j] != nil && expectedRow[j] == nil {
					t.Errorf("cells[%d][%d] unexpectedly non-nil for test case '%s'", i, j, testName)
					continue
				}
				if newRow[j].String() != expectedRow[j].String() {
					t.Errorf("Unexpected cells[%d][%d] value for test case '%s' [expected=%#v, actual=%#v]", i, j, testName, expectedRow[j], newRow[j])
				}
			}
		}
	}
}
