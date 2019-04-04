package game

import "testing"

func TestIPiece(t *testing.T) {
	var piece = &iPiece{
		box: box{
			topLeft: coordinates{
				x: 0,
				y: 3,
			},
			bottomRight: coordinates{
				x: 3,
				y: 0,
			},
		},
	}

	maxY := piece.yMax()
	if maxY.y != 3 {
		t.Errorf("Unexpected yMax [expected = %d, actual = %d]", 3, maxY.y)
	}
	if maxY.x != 2 {
		t.Errorf("Unexpected yMax.x [expected = %d, actual = %d]", 2, maxY.x)
	}

	minY := piece.yMin()
	if minY.y != 0 {
		t.Errorf("Unexpected yMin [expected = %d, actual = %d]", 0, minY.y)
	}
	if minY.x != 2 {
		t.Errorf("Unexpected yMin.x [expected = %d, actual = %d]", 2, minY.x)
	}

	maxX := piece.xMax()
	if maxX.x != 2 {
		t.Errorf("Unexpected xMax [expected = %d, actual = %d]", 2, maxX.x)
	}

	minX := piece.xMin()
	if minX.x != 2 {
		t.Errorf("Unexpected xMin [expected = %d, actual = %d]", 2, minX.x)
	}
}
