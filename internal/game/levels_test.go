package game

import (
	"testing"
	"time"
)

var gTimeTests = map[string]struct {
	level         level
	expectedGTime time.Duration
}{
	"level 0": {
		level:         0,
		expectedGTime: 1000 * time.Millisecond,
	},
	"level 1": {
		level:         1,
		expectedGTime: 966 * time.Millisecond,
	},
	"level 2": {
		level:         2,
		expectedGTime: 932 * time.Millisecond,
	},
	"1 before top speed level (28)": {
		level:         28,
		expectedGTime: 53 * time.Millisecond,
	},
	"top speed level (29)": {
		level:         29,
		expectedGTime: 20 * time.Millisecond,
	},
	"1 after top speed level (30)": {
		level:         30,
		expectedGTime: 20 * time.Millisecond,
	},
}

func TestGTime(t *testing.T) {
	for testName, test := range gTimeTests {
		gTime := test.level.gTime()
		if gTime != test.expectedGTime {
			t.Errorf("Unexpected gTime for test '%s' (expected = %s, actual = %s)", testName, test.expectedGTime, gTime)
		}
	}
}

var linePointsTests = map[string]struct {
	level          level
	linesCleared   int
	expectedPoints int
}{
	"level 0, 1 line cleared": {
		level:          0,
		linesCleared:   1,
		expectedPoints: 40,
	},
	"level 0, 2 lines cleared": {
		level:          0,
		linesCleared:   2,
		expectedPoints: 100,
	},
	"level 0, 3 lines cleared": {
		level:          0,
		linesCleared:   3,
		expectedPoints: 300,
	},
	"level 0, 4 lines cleared": {
		level:          0,
		linesCleared:   4,
		expectedPoints: 1200,
	},
	"level 9, 1 line cleared": {
		level:          9,
		linesCleared:   1,
		expectedPoints: 400,
	},
	"level 9, 2 lines cleared": {
		level:          9,
		linesCleared:   2,
		expectedPoints: 1000,
	},
	"level 9, 3 lines cleared": {
		level:          9,
		linesCleared:   3,
		expectedPoints: 3000,
	},
	"level 9, 4 lines cleared": {
		level:          9,
		linesCleared:   4,
		expectedPoints: 12000,
	},
	"level 9, 0 lines cleared": {
		level:          9,
		linesCleared:   0,
		expectedPoints: 0,
	},
}

func TestLevelLinePoints(t *testing.T) {
	for testName, test := range linePointsTests {
		points := test.level.linePoints(test.linesCleared)
		if points != test.expectedPoints {
			t.Errorf("Unexpected points for test case: '%s'", testName)
		}
	}
}

var updatedLevelTests = []struct {
	currentLevel     level
	linesCleared     int
	expectedNewLevel level
}{
	{
		currentLevel:     0,
		linesCleared:     9,
		expectedNewLevel: 0,
	},
	{
		currentLevel:     0,
		linesCleared:     10,
		expectedNewLevel: 1,
	},
	{
		currentLevel:     1,
		linesCleared:     10,
		expectedNewLevel: 1,
	},
	{
		currentLevel:     1,
		linesCleared:     19,
		expectedNewLevel: 1,
	},
	{
		currentLevel:     1,
		linesCleared:     20,
		expectedNewLevel: 2,
	},
}

func TestUpdatedLevel(t *testing.T) {
	for _, test := range updatedLevelTests {
		newLevel := test.currentLevel.updatedLevel(test.linesCleared)
		if newLevel != test.expectedNewLevel {
			t.Errorf("Unexpected new level for currentLevel=%d, linesCleared=%d (expected=%d, actual=%d)", test.currentLevel, test.linesCleared, test.expectedNewLevel, newLevel)
		}
	}
}
