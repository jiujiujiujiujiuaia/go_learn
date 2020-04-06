package main

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start

//滑动窗口的版本一,当出现了条件不满足时，要移动左边界了。这个时候让外层循环解决
func lengthOfLongestSubstring1(s string) int {
	stringMap := make(map[byte]bool)
	longestLength := 0
	startPoint := 0
	for i := 0; i < len(s); {
		if !stringMap[s[i]] {
			stringMap[s[i]] = true
			if longestLength < (i - startPoint + 1) {
				longestLength = (i - startPoint + 1)
			}
			i++
		} else {
			stringMap[s[startPoint]] = false
			startPoint++
		}
	}
	return longestLength
}

//滑动窗口的版本一,当出现了条件不满足时，要移动左边界了。这个时候加一个循环，不解决不走了。
func lengthOfLongestSubstring(s string) int {
	stringMap := make(map[byte]bool)
	longestLength := 0
	startPoint := 0
	for i := 0; i < len(s); {
		if !stringMap[s[i]] {
			stringMap[s[i]] = true
			i++
		} else {
			if longestLength < (i - startPoint) {
				longestLength = (i - startPoint)
			}
			for stringMap[s[i]] {
				stringMap[s[startPoint]] = false
				startPoint++
			}
		}
	}
	if len(s)-startPoint > longestLength {
		return len(s) - startPoint
	} else {
		return longestLength
	}
}

// @lc code=end
