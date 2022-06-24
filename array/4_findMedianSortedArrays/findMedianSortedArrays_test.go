package array

import (
	"fmt"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	num1 := []int{-1, 2}
	num2 := []int{-2, 0}
	result := FindMedianSortedArrays(num1, num2)
	fmt.Println("*********", result)

	if result+0.5 != 0 {
		t.Error("FindMedianSortedArrays test failed")
	}
	fmt.Println()
}
