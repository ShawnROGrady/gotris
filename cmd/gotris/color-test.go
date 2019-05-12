package main

import (
	"fmt"

	"github.com/ShawnROGrady/gotris/pkg/canvas"
)

func printPotentialColors() {
	background := &canvas.Cell{Background: canvas.White}
	fmt.Printf("Background: %s%s\n", background, canvas.Reset)

	iCell := &canvas.Cell{Background: canvas.Cyan}
	fmt.Printf("I piece: %s%s\n", iCell, canvas.Reset)

	oCell := &canvas.Cell{Background: canvas.Yellow}
	fmt.Printf("O piece: %s%s\n", oCell, canvas.Reset)

	tCell := &canvas.Cell{Background: canvas.Magenta}
	fmt.Printf("T piece: %s%s\n", tCell, canvas.Reset)

	sCell := &canvas.Cell{Background: canvas.Green}
	fmt.Printf("S piece: %s%s\n", sCell, canvas.Reset)

	zCell := &canvas.Cell{Background: canvas.Red}
	fmt.Printf("Z piece: %s%s\n", zCell, canvas.Reset)

	jCell := &canvas.Cell{Background: canvas.Blue}
	fmt.Printf("J piece: %s%s\n", jCell, canvas.Reset)

	lCell := &canvas.Cell{Background: canvas.Orange}
	fmt.Printf("L piece: %s%s\n", lCell, canvas.Reset)
}
