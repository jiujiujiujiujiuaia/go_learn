package binarySearch

/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start

//2022/7/28
//还是比两年前聪明了，秒解
func mySqrt(x int) int {
	//特殊情况
	if x == 1{
		return 1
	}
	leftBoundary, rightBoundary := 1, x/2
	for leftBoundary <= rightBoundary {
		mid := leftBoundary + (rightBoundary-leftBoundary)/2
		midSquare := mid * mid
		if midSquare == x {
			return mid
		} else if midSquare > x {
			rightBoundary = mid - 1
		} else {
			leftBoundary = mid + 1
		}
	}
	return rightBoundary
}

//2020/3/10 缺一个牛顿法
//解法一：二分法
func mySqrt1(x int) int {
	if x == 1 {
		return 1
	}
	l, r := 0, x/2
	for l <= r {
		mid := l + (r-l)/2
		number := mid * mid
		//这个是如4，9，16等等
		if number == x {
			return mid
		} else if number > x {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	//为什么是r？
	//当x不是4，9，16等等数时，如5   2<根号5<3
	//即最终l=3,r=2 因为r=l+1嘛，因此永远都是l，r中那个小一点的数，因此返回r
	return r
}

// @lc code=end
