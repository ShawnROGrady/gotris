package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"

	"github.com/ShawnROGrady/gotris/internal/canvas"
	"github.com/ShawnROGrady/gotris/internal/game"
)

func main() {
	schemeArgs := &stringArrayFlag{}
	colorTest := flag.Bool("colors", false, "Display the colors that will be used throughout the game then exit")
	debugMode := flag.Bool("debug", false, "Run the game in debug mode. This disables gravity as well as canvas clearing")
	disableGhost := flag.Bool("disable-ghost", false, "Don't show the 'ghost' of the current piece")
	disableSide := flag.Bool("disable-side", false, "Don't show the side bar (next piece, current score, and controls)")
	flag.Var(schemeArgs, "scheme", fmt.Sprintf("The control scheme to use, multiple may be specified (default: %s)", game.HomeRowName))
	describeScheme := flag.Bool("describe-scheme", false, "Prints the specified control scheme then exits. If none specified then all available schemes are described")
	lightMode := flag.Bool("light-mode", false, "Update colors to work for light color schemes")
	lowContrastMode := flag.Bool("low-contrast", false, "Update colors to use lower contrast (updates background to white for 'light-mode', black otherwise)")
	cpuprofile := flag.String("cpuprofile", "", "write cpu profile to specified file")

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

	opts := []game.Option{}

	if lightMode != nil && *lightMode {
		background := canvas.Black
		if lowContrastMode != nil && *lowContrastMode {
			background = canvas.White
		}
		opts = append(opts, game.WithBackground(background))
		opts = append(opts, game.WithColor(canvas.Black))
	} else {
		if lowContrastMode != nil && *lowContrastMode {
			opts = append(opts, game.WithBackground(canvas.Black))
		}
	}

	if disableGhost != nil && *disableGhost {
		opts = append(opts, game.WithoutGhost())
	}

	if debugMode != nil && *debugMode {
		opts = append(opts, game.WithDebugMode())
	}

	if disableSide != nil && *disableSide {
		opts = append(opts, game.WithoutSide())
	}

	if describeScheme != nil && *describeScheme {
		describeSchemes(scheme)
		os.Exit(0)
	}

	if scheme == nil {
		// default to home row if no scheme provided
		scheme = game.ControlSchemes{game.HomeRow()}
	}
	opts = append(opts, game.WithControlScheme(scheme))

	if cpuprofile != nil && *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	undoSetup, err := setupTerm()
	if err != nil {
		log.Fatalf("error setting up terminal: %s", err)
		os.Exit(1)
	}
	defer func() {
		if err := undoSetup(); err != nil {
			log.Fatalf("Error undoing set up: %s", err)
		}
	}()

	f, err := os.OpenFile("/dev/tty", os.O_RDWR, 0755)
	if err != nil {
		log.Fatalf("Error opening controlling terminal: %s", err)
	}
	defer f.Close()

	g := game.New(f, opts...)

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
