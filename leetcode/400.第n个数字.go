package main

/*
 * @lc app=leetcode.cn id=400 lang=golang
 *
 * [400] 第N个数字
 */

// @lc code=start

//1位数 1*1*9 2位数10*2*9 3位数100*3*9
//规律base * 位数 * 9
func findNthDigit(n int) int {
	base := 1
	//位数
	digis := 1

	//(1)知道第n这个数是第digis位
	for n > 0 {
		n -= base * digis * 9
		base *= 10
		digis++
	}
	//n表示从第digis位第一个开始数，第n位数
	n += base * digis * 9
	digis--

	//(2)index表示第digis位第一个开始数，第index个数
	index := n % digis

}

// @lc code=end
