package main

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 */

// @lc code=start

//测试用例：
// 123 456
// 0 0
// 0 1
// 1 0
// "" ""
//123 * 99有进位的情况
//123 * 2没有进位的情况

//解法一：普通竖式思维
func multiply1(num1 string, num2 string) string {
	if num1 == "" || num2 == "" {
		return ""
	}

	if num1 == "0" || num2 == "0" {
		return "0"
	}

	res := ""
	for i := len(num2) - 1; i >= 0; i-- {
		//1.用num2的第i位乘num1的所有位
		number1, _ := strconv.Atoi(num2[i : i+1])
		multiplyRes := ""
		//2.在每次乘法之前这里要清零
		carry := 0
		for j := len(num1) - 1; j >= 0; j-- {
			number2, _ := strconv.Atoi(num1[j : j+1])
			//3.进行计算
			current := (number1*number2 + carry) % 10
			carry = (number1*number2 + carry) / 10
			multiplyRes = strconv.Itoa(current) + multiplyRes
		}

		//4.如果还有进位则加上
		if carry != 0 {
			multiplyRes = strconv.Itoa(carry) + multiplyRes
		}

		//5.补0
		for k := i; k < len(num2)-1; k++ {
			multiplyRes = multiplyRes + "0"
		}

		fmt.Println(multiplyRes)
		res = addStrings(res, multiplyRes)
	}

	return res
}

//解法二：优化解法一
func multiply2(num1 string, num2 string) string {
	//1.throw error
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	//2.pre processing
	res := make([]int, len(num1)+len(num2))
	resString := ""

	//3.n位的num1 和m位num2 最大的乘积也就是n+m位 最小的成绩是n+m-1位
	//那么num1[j] * num2[i] 表示的最大也就是一个两位数 -> 个位数放在了res[i+j+1]，十位数放在了res[i+j]
	//再去考虑有进位的情况
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			n1, _ := strconv.Atoi(num1[i : i+1])
			n2, _ := strconv.Atoi(num2[j : j+1])
			sum := n1*n2 + res[i+j+1]
			res[i+j+1] = sum % 10
			res[i+j] += sum / 10
		}
	}

	//3.由于最小的位数是n+m-1 因此索引0可能没有这个数
	for i := 0; i < len(res); i++ {
		if i == 0 && res[i] == 0 {
			continue
		} else {
			resString += strconv.Itoa(res[i])
		}
	}
	return resString
}

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
