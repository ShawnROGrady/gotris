package game

// ControlScheme represents a mapping of keys to user input
type ControlScheme interface {
	controlMap() map[string]userInput
}

// HomeRow represents the control scheme which focuses on the home row
type HomeRow struct{}

func (h HomeRow) controlMap() map[string]userInput {
	return map[string]userInput{
		"h": moveLeft,
		"j": moveDown,
		"k": moveUp,
		"l": moveRight,
		"a": rotateLeft,
		"d": rotateRight,
	}
}

// ArrowKeys represents the control scheme which utilizes the arrow keys
type ArrowKeys struct{}

func (a ArrowKeys) controlMap() map[string]userInput {
	return map[string]userInput{
		"\u001b[A": moveUp,
		"\u001b[B": moveDown,
		"\u001b[C": moveRight,
		"\u001b[D": moveLeft,
		"z":        rotateLeft,
		"x":        rotateRight,
	}
}
