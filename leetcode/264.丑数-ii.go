package main

/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] 丑数 II
 */

// @lc code=start
func nthUglyNumber(n int) int {
	num := make([]int, 1)
	num[0] = 1
	i, j, k := 0, 0, 0
	for n := 1; n < 1690; n++ {
		//1.可以看到某一个丑数一定是由比他小的另外一个丑数*2 or *3 or *5得来的
		//因此可以用dp来实现这种转换
		ugly := min(min(num[i]*2, num[j]*3), num[k]*5)
		num = append(num, ugly)

		//2.以6举例 6可以由2*3 也可以由3*2 因此i和j都要++
		if ugly == num[i]*2 {
			i++
		}
		if ugly == num[j]*3 {
			j++
		}
		if ugly == num[k]*5 {
			k++
		}
	}
	return num[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end
