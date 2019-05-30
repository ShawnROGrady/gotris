package game

import (
	"fmt"
	"strings"
)

// ControlScheme represents a mapping of keys to user input
type ControlScheme interface {
	controlMap() map[string]userInput
	keyMap() map[key]userInput
}

// this approach allows us to more easily change the actual value based on GOOS
type key struct {
	name  string
	value string
}

func (k key) String() string     { return k.value }
func (k key) displayKey() string { return k.name }

func upArrow() key {
	return key{
		name:  "\u2191",
		value: "\u001b[A",
	}
}

func downArrow() key {
	return key{
		name:  "\u2192",
		value: "\u001b[B",
	}
}

func rightArrow() key {
	return key{
		name:  "\u2193",
		value: "\u001b[C",
	}
}

func leftArrow() key {
	return key{
		name:  "\u2190",
		value: "\u001b[D",
	}
}

// HomeRow represents the control scheme which focuses on the home row
type HomeRow struct{}

func (h HomeRow) keyMap() map[key]userInput {
	var (
		upKey          = key{name: "k", value: "k"}
		downKey        = key{name: "j", value: "j"}
		rightKey       = key{name: "l", value: "l"}
		leftKey        = key{name: "h", value: "h"}
		rotateLeftKey  = key{name: "a", value: "a"}
		rotateRightKey = key{name: "d", value: "d"}
	)

	return map[key]userInput{
		upKey:          moveUp,
		downKey:        moveDown,
		rightKey:       moveRight,
		leftKey:        moveLeft,
		rotateLeftKey:  rotateLeft,
		rotateRightKey: rotateRight,
	}
}

func (h HomeRow) controlMap() map[string]userInput {
	return ctrlMap(h)
}

func (h HomeRow) String() string {
	return schemeDescription(h)
}

// ArrowKeys represents the control scheme which utilizes the arrow keys
type ArrowKeys struct{}

func (a ArrowKeys) keyMap() map[key]userInput {
	var (
		upKey          = upArrow()
		downKey        = downArrow()
		rightKey       = rightArrow()
		leftKey        = leftArrow()
		rotateLeftKey  = key{name: "z", value: "z"}
		rotateRightKey = key{name: "x", value: "x"}
	)

	return map[key]userInput{
		upKey:          moveUp,
		downKey:        moveDown,
		rightKey:       moveRight,
		leftKey:        moveLeft,
		rotateLeftKey:  rotateLeft,
		rotateRightKey: rotateRight,
	}
}

func (a ArrowKeys) controlMap() map[string]userInput {
	return ctrlMap(a)
}

func (a ArrowKeys) String() string {
	return schemeDescription(a)
}

// ControlSchemes is a union of one or more mappings of keys to user input
type ControlSchemes []ControlScheme

func (c ControlSchemes) keyMap() map[key]userInput {
	keyMap := make(map[key]userInput)
	for _, scheme := range c {
		kMap := scheme.keyMap()
		for k, v := range kMap {
			keyMap[k] = v
		}
	}

	return keyMap
}

func (c ControlSchemes) controlMap() map[string]userInput {
	return ctrlMap(c)
}

func (c ControlSchemes) String() string {
	return schemeDescription(c)
}

func schemeDescription(c ControlScheme) string {
	var (
		inputMap            = inputMap(c)
		mappingDescriptions = []string{}
	)

	for input, mappings := range inputMap {
		mappingDescriptions = append(
			mappingDescriptions,
			fmt.Sprintf("%s: %s", input, strings.Join(mappings, ", ")),
		)
	}

	return strings.Join(mappingDescriptions, "\n")
}

func inputMap(c ControlScheme) map[string][]string {
	var (
		keyMap   = c.keyMap()
		mappings = make(map[string][]string)
	)

	for key, input := range keyMap {
		if synonyms, ok := mappings[input.String()]; ok {
			synonyms = append(synonyms, key.displayKey())
			mappings[input.String()] = synonyms
		} else {
			mappings[input.String()] = []string{key.displayKey()}
		}
	}

	return mappings
}

func ctrlMap(c ControlScheme) map[string]userInput {
	var (
		keyMap     = c.keyMap()
		controlMap = make(map[string]userInput)
	)
	for key, input := range keyMap {
		controlMap[key.String()] = input
	}
	return controlMap
}
