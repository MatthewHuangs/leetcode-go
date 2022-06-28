package string

import "testing"

func TestLongestPalindrome(t *testing.T) {
	s := "babad"
	res := LongestPalindrome(s)

	if res != "bab" && res != "aba" {
		t.Error("1.LongestPalindrome failed")
	}

	s = "cbbd"
	res = LongestPalindrome(s)

	if res != "bb" {
		t.Error("2.LongestPalindrome failed")
	}

	s = ""
	res = LongestPalindrome(s)

	if res != "" {
		t.Error("3.LongestPalindrome failed")
	}

	s = "a"
	res = LongestPalindrome(s)

	if res != "a" {
		t.Error("4.LongestPalindrome failed")
	}

	s = "bb"
	res = LongestPalindrome(s)

	if res != "bb" {
		t.Error("5.LongestPalindrome failed")
	}
}
