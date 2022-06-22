package twosum

import "testing"

func TestTwoSum(t *testing.T) {
	var nums = []int{2, 7, 11, 15}
	target := 9
	result := TwoSum(nums, target)

	if len(result) == 0 || result[0] != 0 || result[1] != 1 {
		t.Error("TwoSum test failed")
	}

	t.Log("TwoSum test finished")
}
