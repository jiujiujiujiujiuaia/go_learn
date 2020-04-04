package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start

//厘清题意：1.有空格吗2.数据范围最大多大3.有正负号吗？4.会出现其他字符吗？出现了怎么办

//测试用例:
//"+123" 正常数据
//"-123" 正常数据
//"123"无符号
//"+111111111111"超出范围
//"-111111111111"超出范围
//"   123" 有空格
//"   12 3"
//"123A323"异常字符

//思路：按照每一部分可能出现的情况，一点一点把str缩小
func myAtoi(str string) int {
	if str == "" {
		return 0
	}
	isNegative := false
	res := 0

	i := 0
	//1.处理空格
	for i < len(str) && str[i] == ' ' {
		i++
	}
	if i == len(str) {
		return 0
	}

	//2.处理符号情况
	if str[i] == '-' {
		isNegative = true
		i++
	} else if str[i] == '+' {
		i++
	}

	max := math.MaxInt32
	min := math.MinInt32

	//3.处理数字部分，如果超出范围，则要直接返回
	for i < len(str) && str[i] >= '0' && str[i] <= '9' {

		number, _ := strconv.Atoi(str[i : i+1])
		//4.如果超过范围了直接返回：（1）大于max/10 （2）==max/10但是个位数超过的了max的个位数
		fmt.Println(isNegative)
		if !isNegative && (res > max/10 || (res == max/10 && number > max%10)) {
			return max
		}
		if isNegative && (-res < min/10 || (-res == min/10 && number > -min%10)) {
			return min
		}

		res *= 10
		res += number
		i++
	}
	return generateResult(res, isNegative)
}

func generateResult(number int, isNegative bool) int {
	if isNegative {
		return -number
	} else {
		return number
	}
}

// @lc code=end
