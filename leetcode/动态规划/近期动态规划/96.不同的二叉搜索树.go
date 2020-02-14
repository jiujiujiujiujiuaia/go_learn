package main

/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] 不同的二叉搜索树
 */

// @lc code=start

//思想：值得一提的是，无论是1，2，3还是4，5，6 能组成的二叉树序列都是5
//因此只要是递增的一组数，其长度决定了二叉树的个数
//也就是说序列的长度决定了二叉树的个数
//然后以其中的每个元素为根进行排列组合。

func numTrees(n int) int {
	//dp(n)表示序列长为n的不同二叉树搜索树的个数
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		//第二层循环表示，以j为根的二叉搜索树
		for j := 1; j <= i; j++ {
			//当j为根时，左子树是j-1序列长度的个数，右边是i-j序列长度的个数
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

// @lc code=end
