package string

// LongestPalindrome 最长回文子串 动态规划
func LongestPalindrome(s string) string {
	maxLen, begin := 1, 0
	df := make([][]bool, len(s))

	if len(s) < 2 {
		return s
	}

	// 状态转移方程：P(i,j)=P(i+1,j−1)∧(Si==Sj)
	// 边界条件：P(i,i)=true 和 P(i,i+1)=(Si==Si+1)

	// 初始化，P(i,i)=true
	for i := 0; i < len(s); i++ {
		df[i] = make([]bool, len(s))
		df[i][i] = true
	}

	data := []byte(s)

	// 子串长度遍历
	for l := 2; l <= len(s); l++ {
		for i := 0; i < len(s); i++ {
			// j为右边界，控制出界
			j := l + i - 1
			if j >= len(s) {
				break
			}

			// 匹配状态转移方程
			if data[i] != data[j] {
				df[i][j] = false
			} else {
				if l < 3 {
					df[i][j] = true
				} else {
					df[i][j] = df[i+1][j-1]
				}
			}

			if df[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
	}

	return s[begin:(begin + maxLen)]
}

// LongestPalindrome 最长回文子串 暴力求解
// func LongestPalindrome(s string) string {
// 	res := ""
// 	for left := 0; left < len(s); left++ {
// 		if len(res) > len(s)-left {
// 			break
// 		}

// 		for right := left + 1; right <= len(s); right++ {
// 			currStr := s[left:right]
// 			if isPalindrome(currStr) && right-left > len(res) {
// 				res = currStr
// 			}
// 		}
// 	}

// 	return res
// }

// 判断是否为回文串
// func isPalindrome(s string) bool {
// 	for i := 0; i < len(s)/2; i++ {
// 		if s[i] != s[len(s)-i-1] {
// 			return false
// 		}
// 	}
// 	return true
// }
