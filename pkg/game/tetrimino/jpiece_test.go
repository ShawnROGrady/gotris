package tetrimino

import "testing"

var jPieceTests = map[orientation]tetriminoTestCase{
	spawn: tetriminoTestCase{
		expectedMaxY: tetriminoCoordTest{
			y:       2,
			ignoreX: true,
		},
		expectedMinY: tetriminoCoordTest{
			y: 0,
			x: 1,
		},
		expectedMaxX: tetriminoCoordTest{
			x: 2,
			y: 2,
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
			x: 2,
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
	opposite: tetriminoTestCase{
		expectedMaxY: tetriminoCoordTest{
			y: 2,
			x: 1,
		},
		expectedMinY: tetriminoCoordTest{
			y:       0,
			ignoreX: true,
		},
		expectedMaxX: tetriminoCoordTest{
			x:       1,
			ignoreY: true,
		},
		expectedMinX: tetriminoCoordTest{
			x: 0,
			y: 0,
		},
	},
	counterclockwise: tetriminoTestCase{
		expectedMaxY: tetriminoCoordTest{
			y: 2,
			x: 0,
		},
		expectedMinY: tetriminoCoordTest{
			y:       1,
			ignoreX: true,
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
}

func TestJPiece(t *testing.T) {
	piece := newJPiece(3, 3)

	testPiece(t, piece, jPieceTests)
}
