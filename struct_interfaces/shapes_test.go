package structinterfaces

import "testing"

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("shape: %v, wanted: %f, got: %f", shape, want, got)
		}
	}

	testCases := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10, 2.5}, 25.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{6, 4}, 12.0},
	}

	for _, testCase := range testCases {
		checkArea(t, testCase.shape, testCase.want)
	}
}
