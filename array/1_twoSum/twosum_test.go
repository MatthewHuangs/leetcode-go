package array

import "testing"

func TestTwoSum(t *testing.T) {
	var nums = []int{2, 7, 11, 15}
	target := 17
	result := TwoSum(nums, target)

	if len(result) == 0 || result[0] != 0 || result[1] != 3 {
		t.Error("TwoSum test failed")
	}

	t.Log("TwoSum test finished")
}
