package game

import (
	"fmt"
	"sort"
	"strings"
)

// the available schemes
const (
	HomeRowName   = "home-row"
	ArrowKeysName = "arrow-keys"
)

// ControlScheme represents a mapping of keys to user input
type ControlScheme interface {
	controlMap() map[string]userInput
	keyMap() map[key]userInput
	Description() string
	String() string
}

// SchemeFromName retrieves the scheme associated with the specified name
func SchemeFromName(name string) (ControlScheme, error) {
	switch name {
	case HomeRowName:
		return HomeRow(), nil
	case ArrowKeysName:
		return ArrowKeys(), nil
	default:
		return nil, fmt.Errorf("Unrecognized control scheme '%s'", name)
	}
}

// AvailableSchemes represents the set of available control schemes
func AvailableSchemes() []ControlScheme {
	return []ControlScheme{HomeRow(), ArrowKeys()}
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
		name:  "\u2193",
		value: "\u001b[B",
	}
}

func rightArrow() key {
	return key{
		name:  "\u2192",
		value: "\u001b[C",
	}
}

func leftArrow() key {
	return key{
		name:  "\u2190",
		value: "\u001b[D",
	}
}

type keyMapping struct {
	name    string
	mapping func() map[key]userInput
}

func (k keyMapping) keyMap() map[key]userInput {
	return k.mapping()
}

func (k keyMapping) controlMap() map[string]userInput {
	return ctrlMap(k)
}

func (k keyMapping) Description() string {
	return schemeDescription(k)
}

func (k keyMapping) String() string {
	return k.name
}

// HomeRow represents the control scheme which focuses on the home row
func HomeRow() ControlScheme {
	return keyMapping{
		name: HomeRowName,
		mapping: func() map[key]userInput {
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
		},
	}
}

// ArrowKeys represents the control scheme which utilizes the arrow keys
func ArrowKeys() ControlScheme {
	return keyMapping{
		name: ArrowKeysName,
		mapping: func() map[key]userInput {
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
		},
	}
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

// Description is the combined descriptions of multiple schemes
func (c ControlSchemes) Description() string {
	return schemeDescription(c)
}

func (c ControlSchemes) String() string {
	names := []string{}
	for _, scheme := range c {
		names = append(names, scheme.String())
	}
	return strings.Join(names, ", ")
}

func schemeDescription(c ControlScheme) string {
	var (
		inputMap            = inputMap(c)
		mappingDescriptions = []string{}
		inputs              = []userInput{}
	)

	for input := range inputMap {
		inputs = append(inputs, input)
	}

	// sort to allow consistent ordering
	sort.Slice(inputs, func(i, j int) bool {
		return inputs[i] < inputs[j]
	})

	for _, input := range inputs {
		mappings := inputMap[input]
		mappingDescriptions = append(
			mappingDescriptions,
			fmt.Sprintf("%s: %s", input, strings.Join(mappings, ", ")),
		)
	}

	return strings.Join(mappingDescriptions, "\n")
}

func inputMap(c ControlScheme) map[userInput][]string {
	var (
		keyMap   = c.keyMap()
		mappings = make(map[userInput][]string)
	)

	for key, input := range keyMap {
		if synonyms, ok := mappings[input]; ok {
			synonyms = append(synonyms, key.displayKey())
			mappings[input] = synonyms
		} else {
			mappings[input] = []string{key.displayKey()}
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
