package array

// TwoSum 两数之和，小于On^2
func TwoSum(nums []int, target int) []int {
	set := make(map[int]int, len(nums))
	for i, v := range nums {
		if index, exist := set[v]; exist {
			return []int{index, i}
		} else {
			set[target-v] = i
		}
	}
	return nil
}

// TwoSum 两数之和，On^2
// func TwoSum(nums []int, target int) []int {
// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i]+nums[j] == target {
// 				var result = []int{i, j}
// 				return result
// 			}
// 		}
// 	}
// 	return nil
// }
