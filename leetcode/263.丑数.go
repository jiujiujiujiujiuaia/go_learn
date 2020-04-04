package main

/*
 * @lc app=leetcode.cn id=263 lang=golang
 *
 * [263] 丑数
 */

// @lc code=start
//题目：判断一个数是不是只有2 3 5 这个几个因子
//测试数据 负数 0 1 正常数字 大范围数字
//（1）数字范围（2）是否可以是负数

func isUgly(num int) bool {
	if num <= 0 {
		return false
	}

	if num%2 == 0 {
		return isUgly(num / 2)
	}

	if num%3 == 0 {
		return isUgly(num / 3)
	}

	if num%5 == 0 {
		return isUgly(num / 5)
	}

	return num == 1
}

// @lc code=end
