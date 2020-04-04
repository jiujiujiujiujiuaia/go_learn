package main

/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个有序数组的中位数
 */

// @lc code=start

//TODO 二分的解法未完成啊

//解法一：把数组分成相等的两部分，并且要一直保持这个平衡
//知道找到前部分最大的值小于后部分最小值的，否则继续保持这个平衡
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	hanf := (len(nums1) + len(nums2)) / 2
	isOdd := (len(nums1)+len(nums2))%2 != 0
	//1.保证nums1是短数组，否则交换两个数组
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	//2.i是短数组索引，j是长数组索引
	var j int
	frontPartMax, backendPartMin := 0, 0

	//3.这里的i和j都是表示个数，i=0就是表示第一个数组0个
	for i := 0; i < len(nums1)+1; i++ {
		//4.假设总数是5-> 5/2=2 那么刚刚好表示的是3个数，即j是从数量上的第一个中位开始判断
		j = hanf - i

		if i == 0 {
			frontPartMax = nums2[j-1]
		} else if j == 0 {
			frontPartMax = nums1[i-1]
		} else{
			frontPartMax = max(nums1[i-1], nums2[j-1])
		}

		if i == len(nums1){
			backendPartMin = nums2[j]
		}else if j == len(nums2){
			backendPartMin = nums1[i]
		}else{
			backendPartMin = min(nums1[i], nums2[j])
		}
		
		
		//5.因为j表示的是第一个中位，因此奇数的时候就是backendPartMin
		if frontPartMax <= backendPartMin {
			if isOdd{
				return float64(backendPartMin)
			}else{
				return float64((frontPartMax + backendPartMin) / 2 )
			}
		}
	}
}

//解法二：解法一是对短的数组一个个的寻找，解法二的是在这个基础上
//把在短数组中的寻找过程变成二分的过程
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	hanf := (len(nums1) + len(nums2)) / 2
	isOdd := (len(nums1)+len(nums2))%2 != 0
	
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	
	var j,i int
	frontPartMax, backendPartMin := 0, 0

	l,r := 0,len(nums1)
	for l<=r {
		i = l + (r - l) / 2
		j = hanf - i

		if i == 0 {
			frontPartMax = nums2[j-1]
		} else if j == 0 {
			frontPartMax = nums1[i-1]
		} else{
			frontPartMax = max(nums1[i-1], nums2[j-1])
		}

		if i == len(nums1){
			backendPartMin = nums2[j]
		}else if j == len(nums2){
			backendPartMin = nums1[i]
		}else{
			backendPartMin = min(nums1[i], nums2[j])
		}
		
		
		
		if frontPartMax <= backendPartMin {
			if isOdd{
				return float64(backendPartMin)
			}else{
				return float64((frontPartMax + backendPartMin) / 2 )
			}
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// @lc code=end
