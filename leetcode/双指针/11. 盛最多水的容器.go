package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea1(a))
}

func maxArea(height []int) int {
	return twoPoint(height)
}

//复杂度为o(n^2)
func complexity_n2(height []int) int {
	best := 0
	for i, val := range height {
		witdh := 0
		for j, v := range height {
			if val <= v && math.Abs(float64(i)-float64(j)) > float64(witdh) {
				witdh = int(math.Abs(float64(i) - float64(j)))
			}
		}
		if witdh*val > best {
			best = witdh * val
		}
	}
	return best
}

//双指针 复杂度为o(n)
func twoPoint(height []int) int {
	best := 0
	i := 0
	for j := len(height) - 1; i < j; {
		if best < (j-i)*If(height[i], height[j]) {
			best = (j - i) * If(height[i], height[j])
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return best
}

func If(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

//2019/10/9重做这道题,复杂度为o(n)
//因为这个题目相当于是短板原理，依赖与短的一方高
//因此从两端开始计算，假设坐标为i与j（i<j），当前的面积就是i能组成的最大面积（因为最长），但是这个只是局部最大
//移动最短的一方前进或者后退（i++或者j--），才可以算出一组新的值

//2019/10/10 看完下面这个题解后，更深入的理解到，为什么o（n）可以做到和暴力法一样的效果，那是因为，o（n）法每一次消除掉了很多状态
//假设两个坐标是i和j（i<j）因为如果j--，肯定比i，j小，因此这个其实消除了(i,j-1),(i,j-2),...(i,i+1)这么多状态，而这些状态在暴力法中都算了一遍，而双指针法巧妙的避开了
//https://leetcode-cn.com/problems/container-with-most-water/solution/container-with-most-water-shuang-zhi-zhen-fa-yi-do/
func maxArea1(height []int) int {
	res := 0
	for i, j := 0, len(height)-1; i < j; {
		if length := j - i; height[i] < height[j] {
			res = max(res, height[i]*length)
			i++
		} else {
			res = max(res, height[j]*length)
			j--
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
