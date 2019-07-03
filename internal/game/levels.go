package game

import (
	"time"
)

type level int

func (l level) gTime() time.Duration {
	var (
		maxMilliseconds float64 = 1000
		minMilliseconds float64 = 20
		topSpeedLevel   float64 = 29 // the level after which speed stops increasing
	)

	// decrease linearly
	gMilliseconds := maxMilliseconds * (1 - (((maxMilliseconds - minMilliseconds) / maxMilliseconds) * (float64(l) / topSpeedLevel)))
	if gMilliseconds < minMilliseconds {
		return time.Duration(minMilliseconds) * time.Millisecond
	}
	return time.Duration(gMilliseconds) * time.Millisecond
}

func (l level) linePoints(linesCleared int) int {
	if linesCleared == 0 {
		return 0
	}
	// using NES scoring: https://tetris.wiki/Scoring#Original_Nintendo_scoring_system
	lineMultipliers := []int{
		40,
		100,
		300,
		1200,
	}
	return (int(l) + 1) * lineMultipliers[linesCleared-1]
}

func (l level) updatedLevel(linesCleared int) level {
	// using 10 lines before level increase
	if linesCleared >= (int(l)+1)*10 {
		return level(int(l) + 1)
	}
	return l
}
