package main

/*
 * @lc app=leetcode.cn id=44 lang=golang
 *
 * [44] 通配符匹配
 */

// @lc code=start

//通配符稍有不同的是
//?是单个任意字符
//*是任意字符串

//时间不够了，没时间详细的写了。写点伪代码
func isMatch(s string, p string) bool {
	//（1）二维数组伺候

	//（2）预处理dp数组，因为有可能s为空而p有*的情况

	//（3）状态方程式：
	//如果s[j] == p[i] || p[i] == ?  dp[i][j] = dp[i-1][j-1]
	//如果p[i] == * 那么dp[i][j] = dp[i-1][j] || dp[i][j-1]
	//dp[i-1][j]对应的是*匹配第i个字符 dp[i][j-1]对应的是*代表空字符
	//测试用例是1 abcd ab* 2 ab ab*

}

// @lc code=end
