package main

import "math"

//这道题可以简化为比较两个字符串的相似程度
func isOneEditDistance(s string, t string) bool {
	//1.如果两个字符长度的绝对值大于1，则fasle
	firstStringLength := len(s)
	secondStringLength := len(t)
	sub := float64(firstStringLength - secondStringLength)
	if math.Abs(sub) > 1 {
		return false
	}

	//2.保证s永远比t长度大1
	if firstStringLength < secondStringLength {
		s, t = t, s
	}

	//3.以短字符为遍历，防止数组越界，如果出现某个字符不同，截断并比较后面的字符
	for i := 0; i < len(t); i++ {
		if s[i] != t[i] {
			//替换的case
			if firstStringLength == secondStringLength {
				return helpCompareTwoString(s[i+1:], t[i+1:])
			} else {
				//由于前面保证了s字符比t字符长，
				//因此把删除和插入的情况合并成了只有删除的情况
				return helpCompareTwoString(s[i+1:], t[i:])
			}
		}
	}

	//确保两个字符均为空的异常情况
	return len(t)+1 == len(s)
}

//两个字符长度需要完全一样
//比较两个字符是否完全相同
func helpCompareTwoString(a string, b string) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
