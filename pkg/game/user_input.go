package game

import "github.com/ShawnROGrady/gotris/pkg/inputreader"

type userInput int

const (
	moveLeft userInput = iota
	moveDown
	moveUp
	moveRight
)

func (u userInput) String() string {
	inputDescriptions := map[userInput]string{
		moveLeft:  "move left",
		moveDown:  "move down",
		moveUp:    "move up",
		moveRight: "move right",
	}

	return inputDescriptions[u]
}

func translateInput(done chan bool, inputreader inputreader.InputReader) (chan userInput, chan error) {
	rawInput, readErr := inputreader.ReadInput(done)

	translatedInput := make(chan userInput)

	// TODO allow this to be configurable
	controlMap := map[string]userInput{
		"h": moveLeft,
		"j": moveDown,
		"k": moveUp,
		"l": moveRight,
	}

	go func() {
		for {
			select {
			case <-done:
				return
			case input := <-rawInput:
				if translated, ok := controlMap[string(input)]; ok {
					translatedInput <- translated
				}
			}
		}
	}()

	return translatedInput, readErr
}
