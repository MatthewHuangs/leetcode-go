package string

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "pwwkew"
	res := LengthOfLongestSubstring(s)

	if res != 3 {
		t.Error("1.LengthOfLongestSubstring failed")
	}

	s = ""
	res = LengthOfLongestSubstring(s)
	if res != 0 {
		t.Error("2.LengthOfLongestSubstring failed")
	}
}
