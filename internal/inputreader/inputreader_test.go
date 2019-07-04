package inputreader

import (
	"bytes"
	"io"
	"testing"
	"time"
)

func TestInputReader(t *testing.T) {
	var (
		term     bytes.Buffer
		done     = make(chan bool)
		writeErr = make(chan error)
		recieved = []byte{}
	)

	reader := NewTermReader(&term)

	testInputs := "hjkllkjh"

	inputs, readErr := reader.ReadInput(done)
	go func() {
		for _, b := range []byte(testInputs) {
			if err := term.WriteByte(b); err != nil {
				writeErr <- err
			}
			time.Sleep(10 * time.Millisecond) // 10 ms = 1/2 the gravity time at the fastest speed
		}
		close(done)
	}()
	for {
		select {
		case input := <-inputs:
			recieved = append(recieved, input...)
		case err := <-writeErr:
			t.Fatalf("Unexpected error writing input: %s", err)
		case err := <-readErr:
			if err != io.EOF {
				t.Fatalf("Unexpected error reading input: %s", err)
			}
		case <-done:
			time.Sleep(10 * time.Millisecond) // give ReadInput go routine enough time to read last input following 'done'
			if string(recieved) != testInputs {
				t.Errorf("Unexpected input read [expected = %s, actual = %s]", testInputs, string(recieved))
			}
			return
		}
	}
}
