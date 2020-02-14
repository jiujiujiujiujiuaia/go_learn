package main /*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */

import "fmt"

// @lc code=start
func isMatch(s string, p string) bool {
	col := len(s) + 1
	row := len(p) + 1
	dp := make([][]bool, col)
	for i := 0; i < col; i++ {
		dp[i] = make([]bool, row)
	}

	//1.dp的预处理，如果s为空而p不为空，是有可能true
	//eg.s="" p="a*" or p="a*b*"
	dp[0][0] = true
	for j := 1; j < row; j++ {
		if j == 1 {
			dp[0][j] = false
		} else if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	for i := 1; i < col; i++ {
		for j := 1; j < row; j++ {
			//2.在dp数组中，dp[i][j]表示前i个字符和前j个模式字符的状态
			//因此dp数组多了一行一列表示s为空or p为空的状态，但是字符串是索引从0
			//开始计算，因此这里dp数组和字符串索引表示不一样

			//3.1 如果s当前字符相同或者p当前字符是万能字符，则看前面
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			}

			//3.2 如果p当前字符是*，则看前面一个字符，如果前面一个字符与s当前字符不相同，可以抹去
			if p[j-1] == '*' {
				if p[j-2] == '.' || p[j-2] == s[i-1] {
					fmt.Println(i, j, dp[i][j-1], dp[i][j-2], dp[i-1][j])
					//eg dp[i][j-2]解决的是s="a" p="aa*"  把前一个字符重复0次
					//eg dp[i-1][j] 解决的是s="aa" p="a*"  把前一个字符重复1-n次
					dp[i][j] = dp[i][j-2] || dp[i-1][j]
					fmt.Println(dp[i][j])
				} else if p[j-2] != s[i-1] {
					//eg.s="a" p="ab*"
					dp[i][j] = dp[i][j-2]
				}
			}
		}
	}
	fmt.Println(dp)
	return dp[col-1][row-1]
}

// @lc code=end
