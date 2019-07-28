package game

import (
	"fmt"
	"time"
)

type level int

// the available starting difficulties
const (
	BeginnerDifficulty = "beginner"
	NoviceDifficulty   = "novice"
	ProDifficulty      = "pro"
	ExpertDifficulty   = "expert"
)

// LevelFromDifficulty retrieves the starting level associated with a specified difficulty
func LevelFromDifficulty(difficulty string) (int, error) {
	switch difficulty {
	case BeginnerDifficulty:
		return 0, nil
	case NoviceDifficulty:
		return 5, nil
	case ProDifficulty:
		return 10, nil
	case ExpertDifficulty:
		return 15, nil
	default:
		return 0, fmt.Errorf("unrecognized difficulty: '%s'", difficulty)
	}
}

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
