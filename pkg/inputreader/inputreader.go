package inputreader

import (
	"io"
)

// InputReader represents a way to read user input
type InputReader interface {
	ReadInput(done chan bool) (chan []byte, chan error)
}

// TermReader reads user input from the supplied terminal
type TermReader struct {
	term io.Reader
}

// NewTermReader returns a new terminal reader
func NewTermReader(term io.Reader) *TermReader {
	return &TermReader{
		term: term,
	}
}

// ReadInput reads the user input from the supplied terminal
func (t *TermReader) ReadInput(done chan bool) (chan []byte, chan error) {
	var (
		input   = make(chan []byte)
		readErr = make(chan error)
	)

	go func() {
		for {
			select {
			case <-done:
				return
			default:
				buf := make([]byte, 128)
				n, err := t.term.Read(buf)
				if err != nil {
					readErr <- err
				}
				if len(buf) != 0 {
					input <- buf[:n]
				}
			}
		}
	}()

	return input, readErr
}
