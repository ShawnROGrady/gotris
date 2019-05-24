package game

import "fmt"

// ControlScheme represents a mapping of keys to user input
type ControlScheme int

// the available control schemes
const (
	HomeRow ControlScheme = iota
	ArrowKeys
)

func (c ControlScheme) controlMap() (map[string]userInput, error) {
	switch c {
	case HomeRow:
		return map[string]userInput{
			"h": moveLeft,
			"j": moveDown,
			"k": moveUp,
			"l": moveRight,
			"a": rotateLeft,
			"d": rotateRight,
		}, nil
	case ArrowKeys:
		return map[string]userInput{
			"\u001b[A": moveUp,
			"\u001b[B": moveDown,
			"\u001b[C": moveRight,
			"\u001b[D": moveLeft,
			"z":        rotateLeft,
			"x":        rotateRight,
		}, nil
	default:
		return nil, fmt.Errorf("Invalid control scheme: %d", c)
	}
}
