package main

/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */

// @lc code=start

//测试用例:110->1 11->2
//这个测试用例最坑的点在于，由于0在最末尾，而0一旦加入了最末尾
//就会导致之前的值不准
func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}

	dp := make([]int, len(s)+1)
	//预处理
	dp[0] = 1
	dp[1] = 1
	for i := 1; i < len(s); i++ {
		//（1）如果该字符是0并且前面是1或2则，则和1或2组成字母，并且dp等于前2个的字符合
		//否则直接返回0，因为0不对应有字符
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				dp[i+1] = dp[i-1]
			} else {
				return 0
			}
			//（2）如果i指向非0并且前面是1或者2，那么就有两种情况，一种和前一个字符组合一种是不组合
		} else if s[i-1] == '1' || (s[i-1] == '2' && (s[i] >= '1' && s[i] <= '6')) {
			dp[i+1] = dp[i-1] + dp[i]
			//（3）前面的字符是3-9，和i指向的字符无法组合，因此直接复用前面的
		} else {
			dp[i+1] = dp[i]
		}
	}
	return dp[len(s)]
}

//想要直接递归+回溯解，超时。这道题有最优子结构

// func help(s string, start int) int {
// 	if start == len(s) {
// 		return 1
// 	}
// 	cnt := 0
// 	for i := start; i <= len(s); i++ {
// 		if _, ok := chMap[s[start:i]]; ok {
// 			cnt += help(s, i)
// 		}
// 	}
// 	return cnt
// }

// @lc code=end
