package game

import (
	"io"
	"testing"
	"time"

	"github.com/ShawnROGrady/gotris/internal/inputreader"
)

var translateInputTests = map[string]struct {
	scheme              ControlScheme
	inputs              []string
	expectedTranslation []userInput
}{
	"Home Row": {
		scheme: HomeRow(),
		inputs: []string{
			"h",
			"j",
			"k",
			"l",
			"l",
			"k",
			"j",
			"h",
		},
		expectedTranslation: []userInput{
			moveLeft,
			moveDown,
			moveUp,
			moveRight,
			moveRight,
			moveUp,
			moveDown,
			moveLeft,
		},
	},
	"Arrow Keys": {
		scheme: ArrowKeys(),
		inputs: []string{
			leftArrow().String(),
			downArrow().String(),
			upArrow().String(),
			rightArrow().String(),
			rightArrow().String(),
			upArrow().String(),
			downArrow().String(),
			leftArrow().String(),
		},
		expectedTranslation: []userInput{
			moveLeft,
			moveDown,
			moveUp,
			moveRight,
			moveRight,
			moveUp,
			moveDown,
			moveLeft,
		},
	},
	"Standard": {
		scheme: Standard(),
		inputs: []string{
			leftArrow().String(),
			downArrow().String(),
			spaceBar().String(),
			rightArrow().String(),
			rightArrow().String(),
			spaceBar().String(),
			downArrow().String(),
			leftArrow().String(),
		},
		expectedTranslation: []userInput{
			moveLeft,
			moveDown,
			moveUp,
			moveRight,
			moveRight,
			moveUp,
			moveDown,
			moveLeft,
		},
	},
	"Home Row+Arrow Keys": {
		scheme: ControlSchemes([]ControlScheme{HomeRow(), ArrowKeys()}),
		inputs: []string{
			"h",
			"j",
			"k",
			"l",
			rightArrow().String(),
			upArrow().String(),
			downArrow().String(),
			leftArrow().String(),
		},
		expectedTranslation: []userInput{
			moveLeft,
			moveDown,
			moveUp,
			moveRight,
			moveRight,
			moveUp,
			moveDown,
			moveLeft,
		},
	},
}

func TestTranslateInput(t *testing.T) {
	for testName, test := range translateInputTests {
		var (
			done     = make(chan bool)
			writeErr = make(chan error)
			recieved = []userInput{}
		)

		tReader, tWriter := io.Pipe()

		reader := inputreader.NewTermReader(tReader)

		inputs, readErr := translateInput(done, reader, test.scheme.controlMap())
		go func() {
			defer func() {
				tWriter.Close()
				time.Sleep(10 * time.Millisecond) // give ReadInput go routine enough time to read last input following 'done'
				close(done)
			}()
			for _, in := range test.inputs {
				_, err := tWriter.Write([]byte(in))
				if err != nil {
					writeErr <- err
					return
				}
				time.Sleep(1 * time.Millisecond) // 1 ms = 1/20 the gravity time at the fastest speed
			}
		}()

		var completed bool
		for !completed {
			select {
			case input := <-inputs:
				recieved = append(recieved, input)
			case err := <-writeErr:
				t.Fatalf("Unexpected error writing input for test case '%s': %s", testName, err)
			case err := <-readErr:
				if err != io.EOF {
					t.Fatalf("Unexpected error reading input for test case '%s' : %s", testName, err)
				}
			case <-time.After(500 * time.Millisecond):
				t.Fatalf("Timeout for test case '%s'", testName)
			case <-done:
				if len(test.expectedTranslation) != len(recieved) {
					t.Fatalf("Unexpected number of translated inputs received for test case '%s' [expected = %d, actual = %d]", testName, len(test.expectedTranslation), len(recieved))
				}
				for i := range recieved {
					if recieved[i] != test.expectedTranslation[i] {
						t.Errorf("Unexpected translations[%d] for test case '%s' [expected = %s, actual = %s]", i, testName, test.expectedTranslation[i], recieved[i])
					}
				}
				completed = true
			}
		}
	}
}
