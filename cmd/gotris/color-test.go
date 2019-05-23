package main

import (
	"fmt"

	"github.com/ShawnROGrady/gotris/pkg/canvas"
)

func printPotentialColors() {
	background := &canvas.Cell{Color: canvas.White}
	fmt.Printf("Background: %s%s\n", background, canvas.Reset)

	iCell := &canvas.Cell{Color: canvas.Cyan}
	fmt.Printf("I piece: %s%s\n", iCell, canvas.Reset)
	iCell = &canvas.Cell{Color: canvas.Cyan, Transparent: true, Background: canvas.White}
	fmt.Printf("I ghost: %s%s\n", iCell.String(), canvas.Reset)

	oCell := &canvas.Cell{Color: canvas.Yellow}
	fmt.Printf("O piece: %s%s\n", oCell, canvas.Reset)
	oCell = &canvas.Cell{Color: canvas.Yellow, Transparent: true, Background: canvas.White}
	fmt.Printf("O ghost: %s%s\n", oCell, canvas.Reset)

	tCell := &canvas.Cell{Color: canvas.Magenta}
	fmt.Printf("T piece: %s%s\n", tCell, canvas.Reset)
	tCell = &canvas.Cell{Color: canvas.Magenta, Transparent: true, Background: canvas.White}
	fmt.Printf("T ghost: %s%s\n", tCell, canvas.Reset)

	sCell := &canvas.Cell{Color: canvas.Green}
	fmt.Printf("S piece: %s%s\n", sCell, canvas.Reset)
	sCell = &canvas.Cell{Color: canvas.Green, Transparent: true, Background: canvas.White}
	fmt.Printf("S ghost: %s%s\n", sCell, canvas.Reset)

	zCell := &canvas.Cell{Color: canvas.Red}
	fmt.Printf("Z piece: %s%s\n", zCell, canvas.Reset)
	zCell = &canvas.Cell{Color: canvas.Red, Transparent: true, Background: canvas.White}
	fmt.Printf("Z ghost: %s%s\n", zCell, canvas.Reset)

	jCell := &canvas.Cell{Color: canvas.Blue}
	fmt.Printf("J piece: %s%s\n", jCell, canvas.Reset)
	jCell = &canvas.Cell{Color: canvas.Blue, Transparent: true, Background: canvas.White}
	fmt.Printf("J ghost: %s%s\n", jCell, canvas.Reset)

	lCell := &canvas.Cell{Color: canvas.Orange}
	fmt.Printf("L piece: %s%s\n", lCell, canvas.Reset)
	lCell = &canvas.Cell{Color: canvas.Orange, Transparent: true, Background: canvas.White}
	fmt.Printf("L ghost: %s%s\n", lCell, canvas.Reset)
}
