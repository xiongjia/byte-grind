package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Leetcode problem #5. Longest Palindromic Substring
// https://leetcode.com/problems/longest-palindromic-substring/
//
//
// Given a string s, return the longest palindromic substring in s.
//
// Example 1:
// Input: s = "babad"
// Output: "bab"
// Explanation: "aba" is also a valid answer.
//
// Example 2:
// Input: s = "cbbd"
// Output: "bb"
//
// Constraints:
// 1 <= s.length <= 1000
// s consist of only digits and English letters.

// Manacher algorithm
func longestPalindrome001(s string) string {
	if len(s) < 2 {
		return s
	}
	newS := make([]rune, 0)
	newS = append(newS, '#')
	for _, c := range s {
		newS = append(newS, c)
		newS = append(newS, '#')
	}

	maxLen, begin := 1, 0
	dp := make([]int, len(newS))
	maxRight, center := 0, 0
	for i := 0; i < len(newS); i++ {
		if i < maxRight {
			dp[i] = min(maxRight-i, dp[2*center-i])
		}
		left, right := i-(1+dp[i]), i+(1+dp[i])
		for left >= 0 && right < len(newS) && newS[left] == newS[right] {
			dp[i]++
			left--
			right++
		}
		if i+dp[i] > maxRight {
			maxRight = i + dp[i]
			center = i
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
			begin = (i - maxLen) / 2
		}
	}
	return s[begin : begin+maxLen]
}

func longestPalindrome002(s string) string {
	if len(s) == 0 {
		return ""
	}
	left, right, pl, pr := 0, -1, 0, 0
	for left < len(s) {

		for right+1 < len(s) && s[left] == s[right+1] {
			right++
		}
		for left-1 >= 0 && right+1 < len(s) && s[left-1] == s[right+1] {
			left--
			right++
		}
		if right-left > pr-pl {
			pl, pr = left, right
		}
		left = (left+right)/2 + 1
		right = left
	}
	return s[pl : pr+1]
}

func TestP00005(t *testing.T) {
	type sample struct {
		input  string
		output string
	}

	samples := []sample{
		{
			input:  "babad",
			output: "bab",
		},
	}

	// Solution001
	for _, s := range samples {
		result := longestPalindrome001(s.input)
		t.Logf("Input: %s, Result: %s (%s)", s.input, result, s.output)
		assert.Equal(t, s.output, result)
	}

	// Solution002
	for _, s := range samples {
		result := longestPalindrome002(s.input)
		t.Logf("Input: %s, Result: %s (%s)", s.input, result, s.output)
		assert.Equal(t, s.output, result)
	}
}
