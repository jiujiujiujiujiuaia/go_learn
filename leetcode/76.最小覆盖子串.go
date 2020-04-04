package main

import "math"

/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start

//这个题把我剥了层皮

func minWindow(s string, t string) string {
	//如果map已经为true 则count不++ 要不然会有一个字符算了多次cnt的情况
	var count int
	targetStringMap := make(map[byte]int)

	//1.预处理（最先我的思路是没有预处理，只是有个map记录某个字符访问的次数）
	for i := 0; i < len(t); i++ {
		targetStringMap[t[i]]++
	}

	leftBound := 0
	var minWindows = math.MaxInt32
	var minWindowString string
	var rightBound int
	for rightBound = 0; rightBound < len(s); rightBound++ {
		//如果这个字符是我想要的
		if val, ok := targetStringMap[s[rightBound]]; ok {
			//并且我任然想要这个，才对count++
			if val > 0 {
				count++
			}
			//然后把需求--
			targetStringMap[s[rightBound]]--

		}

		//如果某个窗口已经满足条件了，缩小这个窗口
		for count == len(t) {
			//如果这个字符是我想要的
			if val, ok := targetStringMap[s[leftBound]]; ok {
				//并且缩小窗口后（字符没了），我就需要他了
				if val > -1 {
					count--
				}
				targetStringMap[s[leftBound]]++
			}

			if count != len(s) && minWindows > rightBound-leftBound+1 {
				minWindows = rightBound - leftBound + 1
				minWindowString = s[leftBound : rightBound+1]
			}

			leftBound++
		}
	}
	return minWindowString
}

// @lc code=end
