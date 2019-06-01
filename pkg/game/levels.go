package game

import (
	"time"
)

type level int

func (l level) gTime() time.Duration {
	// TODO: add durations based on level
	return 1 * time.Second
}

func (l level) linePoints(linesCleared int) int {
	// using NES scoring: https://tetris.wiki/Scoring#Original_Nintendo_scoring_system
	lineMultiplier := map[int]int{
		1: 40,
		2: 100,
		3: 300,
		4: 1200,
	}
	return (int(l) + 1) * lineMultiplier[linesCleared]
}
