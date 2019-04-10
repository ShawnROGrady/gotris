package tetrimino

import "testing"

var lPieceTests = map[orientation]tetriminoTestCase{
	spawn: tetriminoTestCase{
		expectedMaxY: tetriminoCoordTest{
			y: 2,
			x: 1,
		},
		expectedMinY: tetriminoCoordTest{
			y:       0,
			ignoreX: true,
		},
		expectedMaxX: tetriminoCoordTest{
			x: 2,
			y: 0,
		},
		expectedMinX: tetriminoCoordTest{
			x:       1,
			ignoreY: true,
		},
	},
	clockwise: tetriminoTestCase{
		expectedMaxY: tetriminoCoordTest{
			y:       1,
			ignoreX: true,
		},
		expectedMinY: tetriminoCoordTest{
			y: 0,
			x: 0,
		},
		expectedMaxX: tetriminoCoordTest{
			x: 2,
			y: 1,
		},
		expectedMinX: tetriminoCoordTest{
			x:       0,
			ignoreY: true,
		},
	},
	opposite: tetriminoTestCase{
		expectedMaxY: tetriminoCoordTest{
			y:       2,
			ignoreX: true,
		},
		expectedMinY: tetriminoCoordTest{
			y: 0,
			x: 1,
		},
		expectedMaxX: tetriminoCoordTest{
			x:       1,
			ignoreY: true,
		},
		expectedMinX: tetriminoCoordTest{
			x: 0,
			y: 2,
		},
	},
	counterclockwise: tetriminoTestCase{
		expectedMaxY: tetriminoCoordTest{
			y: 2,
			x: 2,
		},
		expectedMinY: tetriminoCoordTest{
			y:       1,
			ignoreX: true,
		},
		expectedMaxX: tetriminoCoordTest{
			x:       2,
			ignoreY: true,
		},
		expectedMinX: tetriminoCoordTest{
			x: 0,
			y: 1,
		},
	},
}

func TestLPiece(t *testing.T) {
	piece := newLPiece(3, 3)

	testPiece(t, piece, lPieceTests)
}
