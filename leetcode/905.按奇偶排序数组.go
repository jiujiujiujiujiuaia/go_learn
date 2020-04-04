package main

/*
 * @lc app=leetcode.cn id=905 lang=golang
 *
 * [905] 按奇偶排序数组
 */

// @lc code=start

//解法一：额外空间换取时间
func sortArrayByParity1(A []int) []int {
	ans := make([]int, len(A))
	j := 0
	for i := 0; i < len(A); i++ {
		if A[i]&1 == 0 {
			ans[j] = A[i]
			j++
		}

	}

	for i := 0; i < len(A); i++ {
		if A[i]&1 == 1 {
			ans[j] = A[i]
			j++
		}
	}
	return ans
}

//解法二：原地的算法 双指针i和j一共有四种情况
//i j
//1 0 swap
//0 1 i++,j--
//0 0 i++
//1 1 j--
func sortArrayByParity(A []int) []int {
	i, j := 0, len(A)-1
	for i < j {
		if A[i]&1 == 1 && A[j]&1 == 0 {
			A[i], A[j] = A[j], A[i]
		}

		//(1)这两个if 包括了上面的情况二至四
		if A[i]&1 == 0 {
			i++
		}
		if A[j]&1 == 1 {
			j--
		}
	}
	return A
}

// @lc code=end
