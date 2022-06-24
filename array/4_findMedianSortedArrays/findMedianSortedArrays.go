package array

import "fmt"

// FindMedianSortedArrays 寻找两个正序数组的中位数
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	leftValue, rightValue := 0, 0
	index1, index2 := 0, 0
	l := len(nums1) + len(nums2)

	// sort
	for i := 0; i <= l/2; i++ {
		leftValue = rightValue
		if index1 < len(nums1) && (index2 >= len(nums2) || nums1[index1] < nums2[index2]) {
			rightValue = nums1[index1]
			index1++
		} else {
			rightValue = nums2[index2]
			index2++
		}
	}
	fmt.Println(leftValue, rightValue)

	// return
	if l%2 == 0 {
		return (float64(leftValue) + float64(rightValue)) / 2
	}
	return float64(rightValue)
}

// func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	isEven := (len(nums1)+len(nums2))%2 == 0
// 	midIndex := (len(nums1) + len(nums2)) / 2
// 	sortNums := make([]int, len(nums1)+len(nums2))

// 	// 排序
// 	if len(nums1) == 0 {
// 		sortNums = nums2
// 	} else if len(nums2) == 0 {
// 		sortNums = nums1
// 	} else {
// 		var i, j, k int = 0, 0, 0
// 		for k < len(nums1)+len(nums2) {
// 			if nums1[i] < nums2[j] {
// 				sortNums[k] = nums1[i]
// 				i++
// 			} else {
// 				sortNums[k] = nums2[j]
// 				j++
// 			}

// 			k++

// 			if i == len(nums1) {
// 				for ; j < len(nums2); j++ {
// 					sortNums[k] = nums2[j]
// 					k++
// 				}
// 				break
// 			}

// 			if j == len(nums2) {
// 				for ; i < len(nums1); i++ {
// 					sortNums[k] = nums1[i]
// 					k++
// 				}
// 				break
// 			}
// 		}
// 	}

// 	// even
// 	if isEven {
// 		return (float64(sortNums[midIndex-1]) + float64(sortNums[midIndex])) / 2
// 	}

// 	// odd
// 	return float64(sortNums[midIndex])
// }
