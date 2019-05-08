package tetrimino

import "testing"

var zPieceTests = map[orientation]tetriminoTestCase{
	clockwise: tetriminoTestCase{
		expectedMaxY: tetriminoCoordTest{
			y: 3,
			x: 2,
		},
		expectedMinY: tetriminoCoordTest{
			y: 1,
			x: 1,
		},
		expectedMaxX: tetriminoCoordTest{
			x:       2,
			ignoreY: true,
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
			y:       1,
			ignoreX: true,
		},
		expectedMaxX: tetriminoCoordTest{
			x: 2,
			y: 1,
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
			x: 0,
		},
		expectedMaxX: tetriminoCoordTest{
			x:       1,
			ignoreY: true,
		},
		expectedMinX: tetriminoCoordTest{
			x:       0,
			ignoreY: true,
		},
	},
	spawn: tetriminoTestCase{
		expectedMaxY: tetriminoCoordTest{
			y:       3,
			ignoreX: true,
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
			y: 3,
		},
	},
}

func TestZPiece(t *testing.T) {
	piece := newZPiece(4, 4)

	testPiece(t, piece, zPieceTests)

	piece = newZPiece(10, 24)
	if err := testRotationTests(piece); err != nil {
		t.Errorf("%s", err)
	}
}
