package array

import "testing"

func TestMaxArea(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	target := 49

	if MaxArea(height) != target {
		t.Error("MaxArea failed")
	}

	height1 := []int{1, 1}
	target1 := 1

	if MaxArea(height1) != target1 {
		t.Error("1.MaxArea failed")
	}

	height2 := []int{4, 1, 2, 3, 4}
	target2 := 16

	if MaxArea(height2) != target2 {
		t.Error("2.MaxArea failed")
	}
}
