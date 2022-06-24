package array

// MaxArea 盛最多水的容器 双指针，On
func MaxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		// 可简化：minInt可以和下面判断左右高度结合，短板效应
		area := minInt(height[left], height[right]) * (right - left)
		if area > maxArea {
			maxArea = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

func minInt(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// MaxArea 盛最多水的容器 暴力破解On^2，弟弟算法
// func MaxArea(height []int) int {
// 	leftIndex, rightIndex := 0, 0
// 	maxArea := 0
// 	for ; leftIndex < len(height); leftIndex++ {
// 		rightIndex = leftIndex
// 		for ; rightIndex < len(height); rightIndex++ {
// 			minHeight := height[rightIndex]
// 			if height[leftIndex] < height[rightIndex] {
// 				minHeight = height[leftIndex]
// 			}
// 			if minHeight*(rightIndex-leftIndex) > maxArea {
// 				maxArea = minHeight * (rightIndex - leftIndex)
// 			}
// 		}
// 	}

// 	return maxArea
// }
