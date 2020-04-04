package main

/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	//（1）n为0的时候任何数都为1
	if n == 0 {
		return 1
	}

	//（2）n小于0的时候转换为正数的求解
	if n < 0 {
		x = 1 / x
		n = -n
	}
	return fastPow(x, n)
}

//快速幂算法
func fastPow(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	res := fastPow(x, n/2)
	if n%2 == 0 {
		return res * res
	} else {
		return res * res * x
	}
}

// @lc code=end
