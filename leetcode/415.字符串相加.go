package main

import "strconv"

/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 */

// @lc code=start

//边界情况：1.是否含有除0-9外的字符2.是否含有前导0

//1.分case （1）字符串长度相同（2）字符串长度不同
//2.存在进位的情况要处理
//3.如果一方长度耗尽，另一方应该可以直接加

//测试用例：
//123 456
//44 456 （长度不同发生进位）
//1 9 （长度相同发生进位）
//0 0
//"" ""
//"" "123"
func addStrings(num1 string, num2 string) string {
	res := ""
	i, j := len(num1)-1, len(num2)-1
	carry := 0
	for i >= 0 || j >= 0 {
		number1, number2 := 0, 0
		if num1 != "" {
			number1, _ = strconv.Atoi(num1[i:])
			num1 = num1[0:i]
			i--
		}

		if num2 != "" {
			number2, _ = strconv.Atoi(num2[j:])
			num2 = num2[0:j]
			j--
		}

		if (number1+number2+carry)/10 == 1 {
			sub := (number1 + number2 + carry) % 10
			res = strconv.Itoa(sub) + res
			carry = 1
		} else {
			sub := number1 + number2 + carry
			res = strconv.Itoa(sub) + res
			carry = 0
		}
	}

	//重点！注意num1 num2长度相同但是发生了进位的情况
	if carry != 0 {
		res = "1" + res
	}
	return res
}

// @lc code=end
