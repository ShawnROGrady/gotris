package game

import "testing"

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
