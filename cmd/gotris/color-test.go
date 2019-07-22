package main

import (
	"log"
	"os"

	"github.com/ShawnROGrady/gotris/internal/game"
)

func printPotentialColors(opts ...game.Option) {
	demo := game.New(nil, os.Stdout, opts...)
	if err := demo.DisplayPotentialColors(); err != nil {
		log.Fatalf("Error displaying potential colors: %s", err)
	}
}
