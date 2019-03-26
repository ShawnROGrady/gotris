package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

type canvas struct {
	width      int
	height     int
	background string
}

func (c *canvas) render(dest *os.File) error {
	// clear the canvas
	_, err := dest.WriteString("\033[2J")
	if err != nil {
		return err
	}

	_, err = dest.Seek(0, 0)
	if err != nil {
		return err
	}

	for i := 0; i < c.height; i++ {
		var buf = []byte{}
		for j := 0; j < c.width; j++ {
			buf = append(buf, []byte(c.background)...)
		}
		buf = append(buf, '\n')

		_, err := dest.Write(buf)
		if err != nil {
			return err
		}
	}
	//return dest.Sync()
	return nil
}

func main() {
	// set min number of characters for reading to 1
	if err := exec.Command("stty", "-f", "/dev/tty", "-icanon", "min", "1").Run(); err != nil {
		log.Fatalf("error limiting input minimum: %s", err)
		os.Exit(1)
	}
	// do not echo user input
	if err := exec.Command("stty", "-f", "/dev/tty", "-echo").Run(); err != nil {
		log.Fatalf("error disabling user input echoing: %s", err)
		os.Exit(1)
	}
	// re-enable echoing user input
	defer exec.Command("stty", "-f", "echo")

	f, err := os.OpenFile("/dev/tty", os.O_RDWR, 0755)
	if err != nil {
		log.Fatalf("Error opening controlling terminal: %s", err)
	}
	defer f.Close()

	c := &canvas{
		width:      8,
		height:     20,
		background: "\u001b[31m\u2588", // Red
	}

	if err := c.render(f); err != nil {
		log.Fatalf("Error rendering canvas: %s", err)
	}

	time.Sleep(1 * time.Second)
	c.background = "\u001b[32m\u2588" // Green

	if err := c.render(f); err != nil {
		log.Fatalf("Error rendering canvas: %s", err)
	}

	var input = make(chan []byte)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for {
			buf := make([]byte, 128)
			_, err := f.Read(buf)
			if err != nil {
				log.Fatalf("Error reading user input: %s", err)
				os.Exit(1)
			}
			if len(buf) != 0 {
				input <- buf
			}
		}
	}()
	for {
		select {
		case userIn := <-input:
			fmt.Printf("user input: %s\n", userIn)
		case sig := <-sigs:
			fmt.Printf("received signal: %s\n", sig)
			return
		}
	}
}
