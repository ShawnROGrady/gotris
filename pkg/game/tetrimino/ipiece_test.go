package tetrimino

import "testing"

var iPieceTests = map[orientation]tetriminoTestCase{
	spawn: tetriminoTestCase{
		expectedMaxY: tetriminoCoordTest{
			y: 3,
			x: 2,
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
			x:       2,
			ignoreY: true,
		},
	},
	clockwise: tetriminoTestCase{
		expectedMaxY: tetriminoCoordTest{
			y:       1,
			ignoreX: true,
		},
		expectedMinY: tetriminoCoordTest{
			y:       1,
			ignoreX: true,
		},
		expectedMaxX: tetriminoCoordTest{
			x: 3,
			y: 1,
		},
		expectedMinX: tetriminoCoordTest{
			x: 0,
			y: 1,
		},
	},
	opposite: tetriminoTestCase{
		expectedMaxY: tetriminoCoordTest{
			y: 3,
			x: 1,
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
			x:       1,
			ignoreY: true,
		},
	},
	counterclockwise: tetriminoTestCase{
		expectedMaxY: tetriminoCoordTest{
			y:       2,
			ignoreX: true,
		},
		expectedMinY: tetriminoCoordTest{
			y:       2,
			ignoreX: true,
		},
		expectedMaxX: tetriminoCoordTest{
			x: 3,
			y: 2,
		},
		expectedMinX: tetriminoCoordTest{
			x: 0,
			y: 2,
		},
	},
}

func TestIPiece(t *testing.T) {
	piece := newIPiece(4, 4)

	testPiece(t, piece, iPieceTests)
}
