package tetrimino

import "testing"

var tPieceTests = map[orientation]tetriminoTestCase{
	clockwise: tetriminoTestCase{
		expectedMaxY: tetriminoCoordTest{
			y: 3,
			x: 1,
		},
		expectedMinY: tetriminoCoordTest{
			y: 1,
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
	opposite: tetriminoTestCase{
		expectedMaxY: tetriminoCoordTest{
			y:       2,
			ignoreX: true,
		},
		expectedMinY: tetriminoCoordTest{
			y: 1,
			x: 1,
		},
		expectedMaxX: tetriminoCoordTest{
			x: 2,
			y: 2,
		},
		expectedMinX: tetriminoCoordTest{
			x: 0,
			y: 2,
		},
	},
	counterclockwise: tetriminoTestCase{
		expectedMaxY: tetriminoCoordTest{
			y: 3,
			x: 1,
		},
		expectedMinY: tetriminoCoordTest{
			y: 1,
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
	spawn: tetriminoTestCase{
		expectedMaxY: tetriminoCoordTest{
			y: 3,
			x: 1,
		},
		expectedMinY: tetriminoCoordTest{
			y:       2,
			ignoreX: true,
		},
		expectedMaxX: tetriminoCoordTest{
			x: 2,
			y: 2,
		},
		expectedMinX: tetriminoCoordTest{
			x: 0,
			y: 2,
		},
	},
}

func TestTPiece(t *testing.T) {
	piece := newTPiece(4, 4)

	testPiece(t, piece, tPieceTests)

	piece = newTPiece(10, 24)
	if err := testRotationTests(piece, defaultWallKickTests()); err != nil {
		t.Errorf("%s", err)
	}
}
