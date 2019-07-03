package main

import (
	"fmt"

	"github.com/ShawnROGrady/gotris/internal/game"
)

func describeSchemes(scheme game.ControlSchemes) {
	if scheme != nil && len(scheme) != 0 {
		fmt.Printf("Selected scheme: %s\nControls:\n%s\n", scheme, scheme.Description())
		return
	}
	allSchemes := game.AvailableSchemes()
	fmt.Println("AvailableSchemes:")
	for _, s := range allSchemes {
		fmt.Printf("Name: %s\nControls:\n%s\n\n", s, s.Description())
	}
}
