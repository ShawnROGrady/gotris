package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ShawnROGrady/gotris/pkg/game"
)

func main() {
	schemeArgs := &stringArrayFlag{}
	colorTest := flag.Bool("colors", false, "Display the colors that will be used throughout the game then exit")
	debugMode := flag.Bool("debug", false, "Run the game in debug mode. This disables gravity as well as canvas clearing")
	disableGhost := flag.Bool("disable-ghost", false, "Don't show the 'ghost' of the current piece")
	flag.Var(schemeArgs, "scheme", fmt.Sprintf("The control scheme to use, multiple may be specified (default: %s)", game.HomeRowName))
	describeScheme := flag.Bool("describe-scheme", false, "Prints the specified control scheme then exits. If none specified then all available schemes are described")

	flag.Parse()

	if colorTest != nil && *colorTest {
		printPotentialColors()
		os.Exit(0)
	}

	// set the specified control scheme
	var scheme game.ControlSchemes
	for _, arg := range *schemeArgs {
		s, err := game.SchemeFromName(arg)
		if err != nil {
			log.Fatalf("%s", err)
			os.Exit(1)
		}
		scheme = append(scheme, s)
	}
	if describeScheme != nil && *describeScheme {
		describeSchemes(scheme)
		os.Exit(0)
	}

	if scheme == nil {
		// default to home row if no scheme provided
		scheme = game.ControlSchemes{game.HomeRow()}
	}

	undoSetup, err := setupTerm()
	if err != nil {
		log.Fatalf("error setting up terminal: %s", err)
		os.Exit(1)
	}
	defer undoSetup()

	f, err := os.OpenFile("/dev/tty", os.O_RDWR, 0755)
	if err != nil {
		log.Fatalf("Error opening controlling terminal: %s", err)
	}
	defer f.Close()

	conf := game.Config{
		Term:  f,
		Width: 10, Height: 20,
		HiddenRows:    4,
		DebugMode:     *debugMode,
		DisableGhost:  *disableGhost,
		ControlScheme: scheme,
		WidthScale:    2,
	}

	g := game.New(conf)

	done := make(chan bool)
	defer close(done)

	endScore, runErr := g.Run(done)

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
