package main

/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	wordDictMap := make(map[string]bool)
	for _, word := range wordDict {
		wordDictMap[word] = true
	}

	dp[0] = true
	//dp[i]的状态定义是，前i个字符是符合的
	//因为字符数组是前闭后开区间，因此i要从1开始，s[0:1]=s[0]
	//同时要注意，往往做字符串的题目，dp数组的索引定义和字符数组索引定义不一样
	//dp数组表示前i个，但是字符数组索引的前i个其实是i-1，因此字符是从0开始算的
	for i := 1; i < len(s)+1; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

// @lc code=end
