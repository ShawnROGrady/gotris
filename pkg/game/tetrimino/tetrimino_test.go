package tetrimino

import "testing"

func TestIPiece(t *testing.T) {
	var piece = &iPiece{
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
		t.Errorf("Unexpected yMax [expected = %d, actual = %d]", 3, maxY.Y)
	}
	if maxY.X != 2 {
		t.Errorf("Unexpected yMax.x [expected = %d, actual = %d]", 2, maxY.X)
	}

	minY := piece.YMin()
	if minY.Y != 0 {
		t.Errorf("Unexpected yMin [expected = %d, actual = %d]", 0, minY.Y)
	}
	if minY.X != 2 {
		t.Errorf("Unexpected yMin.x [expected = %d, actual = %d]", 2, minY.X)
	}

	maxX := piece.XMax()
	if maxX.X != 2 {
		t.Errorf("Unexpected xMax [expected = %d, actual = %d]", 2, maxX.X)
	}

	minX := piece.XMin()
	if minX.X != 2 {
		t.Errorf("Unexpected xMin [expected = %d, actual = %d]", 2, minX.X)
	}
}
