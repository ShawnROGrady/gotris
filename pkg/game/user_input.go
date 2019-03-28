package game

import (
	"os"
)

func readInput(done chan bool, term *os.File) (chan []byte, chan error) {
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
				n, err := term.Read(buf)
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
