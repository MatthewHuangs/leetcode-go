package main

import (
	"fmt"
	twosum "leetcode-go/1_twosum"
)

func main() {
	var nums = []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twosum.TwoSum(nums, target))
}
