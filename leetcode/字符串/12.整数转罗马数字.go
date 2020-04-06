package main

/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */

// @lc code=start

//解法一就是模拟
//解法二是列举，例如1 4 5 9这样列举
func intToRoman(num int) string {
	numberMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	toNumberMap := make(map[int]string)
	for key, val := range numberMap {
		toNumberMap[val] = key
	}

	multipy := 1000
	res := ""

	//1.每次都是这个地方弄得自己死循环
	//记住是系数大于0！！！ 因为任何数除1都是本身，所以用num>0会死循环
	for multipy > 0 {
		if num/multipy > 0 {

			res += calculate(multipy, num/multipy, toNumberMap)
			num %= multipy
		}
		multipy /= 10
	}
	return res
}

//核心代码，分情况讨论
func calculate(multipy int, n int, numberMap map[int]string) string {
	if n < 4 {
		return help(numberMap[multipy], n)
	} else if n == 4 {
		return help(numberMap[multipy], 1) + numberMap[multipy*5]
	} else if n <= 8 {
		return numberMap[multipy*5] + help(numberMap[multipy], n-5)
	} else {
		return help(numberMap[multipy], 1) + numberMap[multipy*10]
	}
}

func help(str string, n int) string {
	res := ""
	for n > 0 {
		res += str
		n--
	}
	return res
}

// @lc code=end
