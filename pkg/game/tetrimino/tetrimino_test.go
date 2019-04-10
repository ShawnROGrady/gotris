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
	// test in default location
	testPieceRotations(t, piece, pieceTests, Coordinates{
		X: 0,
		Y: 0,
	})

	// move piece up (within allowed range)
	piece.MoveUp(20)
	testPieceRotations(t, piece, pieceTests, Coordinates{
		X: 0,
		Y: 1,
	})

	// move piece right (within allowed range)
	piece.MoveRight(20)
	testPieceRotations(t, piece, pieceTests, Coordinates{
		X: 1,
		Y: 1,
	})

	// move piece down
	piece.MoveDown()
	testPieceRotations(t, piece, pieceTests, Coordinates{
		X: 1,
		Y: 0,
	})

	// move piece left
	piece.MoveLeft()
	testPieceRotations(t, piece, pieceTests, Coordinates{
		X: 0,
		Y: 0,
	})
}

func testPieceRotations(t *testing.T, piece Tetrimino, pieceTests map[orientation]tetriminoTestCase, offset Coordinates) {
	// test spawn orientation
	testCase, ok := pieceTests[piece.pieceOrientation()]
	if ok {
		testOrientation(t, piece, testCase, offset)
	}

	// test right rotation
	piece.RotateClockwise()
	testCase, ok = pieceTests[piece.pieceOrientation()]
	if ok {
		// clockwise
		testOrientation(t, piece, testCase, offset)
	}

	piece.RotateClockwise()
	testCase, ok = pieceTests[piece.pieceOrientation()]
	if ok {
		// opposite
		testOrientation(t, piece, testCase, offset)
	}

	piece.RotateClockwise()
	testCase, ok = pieceTests[piece.pieceOrientation()]
	if ok {
		// counter
		testOrientation(t, piece, testCase, offset)
	}

	piece.RotateClockwise()
	testCase, ok = pieceTests[piece.pieceOrientation()]
	if ok {
		// spawn
		testOrientation(t, piece, testCase, offset)
	}

	// test left rotation
	piece.RotateCounter()
	testCase, ok = pieceTests[piece.pieceOrientation()]
	if ok {
		// counter
		testOrientation(t, piece, testCase, offset)
	}

	piece.RotateCounter()
	testCase, ok = pieceTests[piece.pieceOrientation()]
	if ok {
		// opposite
		testOrientation(t, piece, testCase, offset)
	}

	piece.RotateCounter()
	testCase, ok = pieceTests[piece.pieceOrientation()]
	if ok {
		// clockwise
		testOrientation(t, piece, testCase, offset)
	}

	piece.RotateCounter()
	testCase, ok = pieceTests[piece.pieceOrientation()]
	if ok {
		// spawn
		testOrientation(t, piece, testCase, offset)
	}
}

func testOrientation(t *testing.T, piece Tetrimino, testCase tetriminoTestCase, offset Coordinates) {
	orientation := piece.pieceOrientation()

	// test XMax
	maxX := piece.XMax()
	if !testCase.expectedMaxX.ignoreX {
		if maxX.X != testCase.expectedMaxX.x+offset.X {
			t.Errorf("Unexpected xMax in %s orientation (offset = %v)[expected = %d, actual = %d]", &orientation, offset, testCase.expectedMaxX.x, maxX.X+offset.X)
		}
	}

	if !testCase.expectedMaxX.ignoreY {
		if maxX.Y != testCase.expectedMaxX.y+offset.Y {
			t.Errorf("Unexpected xMax.y in %s orientation (offset = %v) [expected = %d, actual = %d]", &orientation, offset, testCase.expectedMaxX.y, maxX.Y+offset.Y)
		}
	}

	// test XMin
	minX := piece.XMin()
	if !testCase.expectedMinX.ignoreX {
		if minX.X != testCase.expectedMinX.x+offset.X {
			t.Errorf("Unexpected xMin in %s orientation (offset = %v)[expected = %d, actual = %d]", &orientation, offset, testCase.expectedMinX.x, minX.X+offset.X)
		}
	}

	if !testCase.expectedMinX.ignoreY {
		if minX.Y != testCase.expectedMinX.y+offset.Y {
			t.Errorf("Unexpected xMin.y in %s orientation (offset = %v) [expected = %d, actual = %d]", &orientation, offset, testCase.expectedMinX.y, minX.Y+offset.Y)
		}
	}

	// test YMax
	maxY := piece.YMax()
	if !testCase.expectedMaxY.ignoreX {
		if maxY.X != testCase.expectedMaxY.x+offset.X {
			t.Errorf("Unexpected yMax.x in %s orientation (offset = %v)[expected = %d, actual = %d]", &orientation, offset, testCase.expectedMaxY.x, maxY.X+offset.X)
		}
	}

	if !testCase.expectedMaxY.ignoreY {
		if maxY.Y != testCase.expectedMaxY.y+offset.Y {
			t.Errorf("Unexpected yMax in %s orientation (offset = %v) [expected = %d, actual = %d]", &orientation, offset, testCase.expectedMaxY.y, maxY.Y+offset.Y)
		}
	}

	// test YMin
	minY := piece.YMin()
	if !testCase.expectedMinY.ignoreX {
		if minY.X != testCase.expectedMinY.x+offset.X {
			t.Errorf("Unexpected yMin.x in %s orientation (offset = %v)[expected = %d, actual = %d]", &orientation, offset, testCase.expectedMinY.x, minY.X+offset.X)
		}
	}

	if !testCase.expectedMinY.ignoreY {
		if minY.Y != testCase.expectedMinY.y+offset.Y {
			t.Errorf("Unexpected yMin in %s orientation (offset = %v) [expected = %d, actual = %d]", &orientation, offset, testCase.expectedMinY.y, minY.Y+offset.Y)
		}
	}
}