package game

import (
	"time"
)

type level int

func (l level) gTime() time.Duration {
	// TODO: add durations based on level
	return 1 * time.Second
}
