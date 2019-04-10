package tetrimino

import "testing"

type tetriminoTestCase struct {
	expectedMaxY tetriminoCoordTest
	expectedMinY tetriminoCoordTest
	expectedMaxX tetriminoCoordTest
	expectedMinX tetriminoCoordTest
}

type tetriminoCoordTest struct {
	x       int
	ignoreX bool
	y       int
	ignoreY bool
}

func testPiece(t *testing.T, piece Tetrimino, pieceTests map[orientation]tetriminoTestCase) {
	// test spawn orientation
	testCase, ok := pieceTests[piece.pieceOrientation()]
	if ok {
		testOrientation(t, piece, testCase)
	}

	// test right rotation
	piece.RotateClockwise()
	testCase, ok = pieceTests[piece.pieceOrientation()]
	if ok {
		// clockwise
		testOrientation(t, piece, testCase)
	}

	piece.RotateClockwise()
	testCase, ok = pieceTests[piece.pieceOrientation()]
	if ok {
		// opposite
		testOrientation(t, piece, testCase)
	}

	piece.RotateClockwise()
	testCase, ok = pieceTests[piece.pieceOrientation()]
	if ok {
		// counter
		testOrientation(t, piece, testCase)
	}

	piece.RotateClockwise()
	testCase, ok = pieceTests[piece.pieceOrientation()]
	if ok {
		// spawn
		testOrientation(t, piece, testCase)
	}

	// test left rotation
	piece.RotateCounter()
	testCase, ok = pieceTests[piece.pieceOrientation()]
	if ok {
		// counter
		testOrientation(t, piece, testCase)
	}

	piece.RotateCounter()
	testCase, ok = pieceTests[piece.pieceOrientation()]
	if ok {
		// opposite
		testOrientation(t, piece, testCase)
	}

	piece.RotateCounter()
	testCase, ok = pieceTests[piece.pieceOrientation()]
	if ok {
		// clockwise
		testOrientation(t, piece, testCase)
	}

	piece.RotateCounter()
	testCase, ok = pieceTests[piece.pieceOrientation()]
	if ok {
		// spawn
		testOrientation(t, piece, testCase)
	}

}

func testOrientation(t *testing.T, piece Tetrimino, testCase tetriminoTestCase) {
	orientation := piece.pieceOrientation()

	// test XMax
	maxX := piece.XMax()
	if !testCase.expectedMaxX.ignoreX {
		if maxX.X != testCase.expectedMaxX.x {
			t.Errorf("Unexpected xMax in %s orientation [expected = %d, actual = %d]", &orientation, testCase.expectedMaxX.x, maxX.X)
		}
	}

	if !testCase.expectedMaxX.ignoreY {
		if maxX.Y != testCase.expectedMaxX.y {
			t.Errorf("Unexpected xMax.y in %s orientation [expected = %d, actual = %d]", &orientation, testCase.expectedMaxX.y, maxX.Y)
		}
	}

	// test XMin
	minX := piece.XMin()
	if !testCase.expectedMinX.ignoreX {
		if minX.X != testCase.expectedMinX.x {
			t.Errorf("Unexpected xMin in %s orientation [expected = %d, actual = %d]", &orientation, testCase.expectedMinX.x, minX.X)
		}
	}

	if !testCase.expectedMinX.ignoreY {
		if minX.Y != testCase.expectedMinX.y {
			t.Errorf("Unexpected xMin.y in %s orientation [expected = %d, actual = %d]", &orientation, testCase.expectedMinX.y, minX.Y)
		}
	}

	// test YMax
	maxY := piece.YMax()
	if !testCase.expectedMaxY.ignoreX {
		if maxY.X != testCase.expectedMaxY.x {
			t.Errorf("Unexpected yMax.x in %s orientation [expected = %d, actual = %d]", &orientation, testCase.expectedMaxY.x, maxY.X)
		}
	}

	if !testCase.expectedMaxY.ignoreY {
		if maxY.Y != testCase.expectedMaxY.y {
			t.Errorf("Unexpected yMax in %s orientation [expected = %d, actual = %d]", &orientation, testCase.expectedMaxY.y, maxY.Y)
		}
	}

	// test YMin
	minY := piece.YMin()
	if !testCase.expectedMinY.ignoreX {
		if minY.X != testCase.expectedMinY.x {
			t.Errorf("Unexpected yMin.x in %s orientation [expected = %d, actual = %d]", &orientation, testCase.expectedMinY.x, minY.X)
		}
	}

	if !testCase.expectedMinY.ignoreY {
		if minY.Y != testCase.expectedMinY.y {
			t.Errorf("Unexpected yMin in %s orientation [expected = %d, actual = %d]", &orientation, testCase.expectedMinY.y, minY.Y)
		}
	}
}
