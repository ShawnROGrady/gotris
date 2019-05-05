package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/ShawnROGrady/gotris/pkg/game"
)

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

	g := game.New(f, 8, 20, 4)

	done := make(chan bool)
	defer close(done)

	endScore, runErr := g.RunDemo(done)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-done:
		return
	case err := <-runErr:
		log.Fatalf("Error running game: %s", err)
		os.Exit(1)
	case score := <-endScore:
		fmt.Printf("GAME OVER (score = %d)\n", score)
		return
	case sig := <-sigs:
		fmt.Printf("received signal: %s\n", sig)
		return
	}

}
