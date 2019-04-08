package tetrimino

import "testing"

func TestIPiece(t *testing.T) {
	spawnOrientation := spawn
	var piece = &iPiece{
		orientation: &spawnOrientation,
		box: Box{
			TopLeft: Coordinates{
				X: 0,
				Y: 3,
			},
			BottomRight: Coordinates{
				X: 3,
				Y: 0,
			},
		},
	}

	maxY := piece.YMax()
	if maxY.Y != 3 {
		t.Errorf("Unexpected yMax in spawn orientation [expected = %d, actual = %d]", 3, maxY.Y)
	}
	if maxY.X != 2 {
		t.Errorf("Unexpected yMax.x in spawn orientation [expected = %d, actual = %d]", 2, maxY.X)
	}

	minY := piece.YMin()
	if minY.Y != 0 {
		t.Errorf("Unexpected yMin in spawn orientation [expected = %d, actual = %d]", 0, minY.Y)
	}
	if minY.X != 2 {
		t.Errorf("Unexpected yMin.x in spawn orientation [expected = %d, actual = %d]", 2, minY.X)
	}

	maxX := piece.XMax()
	if maxX.X != 2 {
		t.Errorf("Unexpected xMax in spawn orientation [expected = %d, actual = %d]", 2, maxX.X)
	}

	minX := piece.XMin()
	if minX.X != 2 {
		t.Errorf("Unexpected xMin in spawn orientation [expected = %d, actual = %d]", 2, minX.X)
	}

	piece.RotateClockwise()
	if *piece.orientation != clockwise {
		t.Errorf("Unexpected orientation after rotating clockwise: %d", *piece.orientation)
	}

	maxY = piece.YMax()
	if maxY.Y != 1 {
		t.Errorf("Unexpected yMax in clockwise orientation [expected = %d, actual = %d]", 1, maxY.Y)
	}

	minY = piece.YMin()
	if minY.Y != 1 {
		t.Errorf("Unexpected yMin in clockwise orientation [expected = %d, actual = %d]", 1, minY.Y)
	}

	maxX = piece.XMax()
	if maxX.X != 3 {
		t.Errorf("Unexpected xMax in clockwise orientation [expected = %d, actual = %d]", 3, maxX.X)
	}
	if maxX.Y != 1 {
		t.Errorf("Unexpected xMax.y in clockwise orientation [expected = %d, actual = %d]", 1, maxX.Y)
	}

	minX = piece.XMin()
	if minX.X != 0 {
		t.Errorf("Unexpected xMin in clockwise orientation [expected = %d, actual = %d]", 0, minX.X)
	}
	if minX.Y != 1 {
		t.Errorf("Unexpected xMin.y in clockwise orientation [expected = %d, actual = %d]", 1, minX.Y)
	}
}
