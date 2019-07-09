package inputreader

import (
	"io"
	"testing"
	"time"
)

func TestInputReader(t *testing.T) {
	var (
		done     = make(chan bool)
		writeErr = make(chan error)
		recieved = []byte{}
	)

	tReader, tWriter := io.Pipe()

	reader := NewTermReader(tReader)

	testInputs := "hjkllkjh"

	inputs, readErr := reader.ReadInput(done)
	go func() {
		defer func() {
			tWriter.Close()
			time.Sleep(10 * time.Millisecond) // give ReadInput go routine enough time to read last input following 'done'
			close(done)
		}()
		for _, b := range []byte(testInputs) {
			_, err := tWriter.Write([]byte{b})
			if err != nil {
				writeErr <- err
				return
			}
			time.Sleep(1 * time.Millisecond) // 1 ms = 1/20 the gravity time at the fastest speed
		}
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
			if string(recieved) != testInputs {
				t.Errorf("Unexpected input read [expected = %s, actual = %s]", testInputs, string(recieved))
			}
			return
		}
	}
}
