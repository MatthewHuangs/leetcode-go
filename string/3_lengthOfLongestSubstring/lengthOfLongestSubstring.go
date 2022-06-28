package string

import "strings"

// LengthOfLongestSubstring 无重复字符的最长子串 滑动窗口
func LengthOfLongestSubstring(s string) int {
	res, next := 0, 0
	window := ""

	for next < len(s) {
		nextStr := string(s[next])
		next++

		if strings.Contains(window, nextStr) {
			index := strings.LastIndex(window, nextStr)
			window = window[index+1:]
		}
		window += nextStr

		if len(window) > res {
			res = len(window)
		}
	}

	return res
}

// LengthOfLongestSubstring 无重复字符的最长子串 部分匹配表（类似滑动窗口，但无序，只存储是否存在）
// func LengthOfLongestSubstring(s string) int {
// 	res := 0
// 	// 部分匹配表：Partial Match Table
// 	pmt := map[byte]int{}
// 	left, right := 0, 0

// 	for right < len(s) {
// 		nextByte := s[right]
// 		right++
// 		pmt[nextByte]++

// 		// 判断该字符是否存在于部分匹配表中，并依次删除左侧元素
// 		for pmt[nextByte] > 1 {
// 			pmt[s[left]]--
// 			left++
// 		}

// 		if right-left > res {
// 			res = right - left
// 		}
// 	}

// 	return res
// }

// LengthOfLongestSubstring 无重复字符的最长子串 优化的暴力模式
// func LengthOfLongestSubstring(s string) int {
// 	maxSubStrLen := 0
// 	sLen := len(s)

// 	for left := 0; left < sLen; left++ {
// 		if sLen-left <= maxSubStrLen {
// 			break
// 		}

// 		for right := left + 1; right <= sLen; right++ {
// 			// 左闭右开
// 			if strings.Contains(s[left:right-1], string(s[right-1])) {
// 				break
// 			}

// 			if maxSubStrLen < right-left {
// 				maxSubStrLen = right - left
// 			}
// 		}
// 	}

// 	return maxSubStrLen
// }
