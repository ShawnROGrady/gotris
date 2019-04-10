package tetrimino

import "testing"

var tPieceTests = map[orientation]tetriminoTestCase{
	spawn: tetriminoTestCase{
		expectedMaxY: tetriminoCoordTest{
			y: 2,
			x: 1,
		},
		expectedMinY: tetriminoCoordTest{
			y: 0,
			x: 1,
		},
		expectedMaxX: tetriminoCoordTest{
			x: 2,
			y: 1,
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
			x: 1,
		},
		expectedMaxX: tetriminoCoordTest{
			x: 2,
			y: 1,
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
			y: 0,
			x: 1,
		},
		expectedMaxX: tetriminoCoordTest{
			x:       1,
			ignoreY: true,
		},
		expectedMinX: tetriminoCoordTest{
			x: 0,
			y: 1,
		},
	},
	counterclockwise: tetriminoTestCase{
		expectedMaxY: tetriminoCoordTest{
			y: 2,
			x: 1,
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
			x: 0,
			y: 1,
		},
	},
}

func TestTPiece(t *testing.T) {
	piece := newTPiece(3, 3)

	testPiece(t, piece, tPieceTests)
}
