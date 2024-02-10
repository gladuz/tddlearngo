package arraysslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("wanted %d, got %d", want, got)
	}
}

func TestSumAll(t *testing.T) {
	t.Run("Run with 3", func(t *testing.T) {
		sl1 := []int{2, 6, 8}
		sl2 := []int{5, 7, 4, 12}
		sl3 := []int{5, 4}
		got := SumZip(sl1, sl2, sl3)
		want := []int{12, 17}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted %v, got %v", want, got)
		}
	})
	t.Run("With one off slice", func(t *testing.T) {
		sl := []int{1, 2, 3, 4, 5}
		got := SumZip(sl, sl[1:])
		want := []int{3, 5, 7, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted %v, got %v", want, got)
		}
	})
}
