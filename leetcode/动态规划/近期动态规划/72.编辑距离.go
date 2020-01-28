package main

/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */

// @lc code=start

//这个题目很复杂，没有见过，先去做了相关题目，了解了两个字符串之间的比较
//1.发现，增加和删除其实是一种操作，替换是另一种操作
func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(word1)+1; i++ {
		dp[i] = make([]int, len(word2)+1)
	}

	//1.预处理
	for i := 0; i < len(word1)+1; i++ {
		dp[i][0] = i
	}

	for i := 0; i < len(word2)+1; i++ {
		dp[0][i] = i
	}

	//2.开始自底向上的递推
	for i := 1; i < len(word1)+1; i++ {
		for j := 1; j < len(word2)+1; j++ {
			left := dp[i-1][j] + 1
			up := dp[i][j-1] + 1
			left_up := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				left_up++
			}
			dp[i][j] = min(left, up, left_up)
		}
	}
	return dp[len(word1)][len(word2)]
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		} else {
			return c
		}
	} else {
		if b < c {
			return b
		} else {
			return c
		}
	}
}

// @lc code=end
