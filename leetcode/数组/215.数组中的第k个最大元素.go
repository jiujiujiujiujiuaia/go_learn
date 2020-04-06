package main

/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start

//快速排序！！！
func findKthLargest(nums []int, k int) int {
	return sort(nums, 0, len(nums)-1, len(nums)-k)
}

//快速排序中这个函数的作用是把start-end间的数据进行排序
//这个算法的时间复杂度是o(n) n+1/2n+1/4n 收敛后还是n
func sort(nums []int, start, end, targetIndex int) int {
	index := patipation(nums, start, end)
	if targetIndex == index {
		return nums[targetIndex]
	} else if targetIndex > index {
		return sort(nums, index+1, end, targetIndex)
	} else {
		return sort(nums, start, index-1, targetIndex)
	}
}

//保证start-end中的任意一个数能够找到他在原数组正确的位置，并返回其索引
func patipation(nums []int, start, end int) int {
	//1.目标是要找到tempnumber这个数排序后的索引

	//tempNumberIndex := random.() 更标准的做法应该是随机
	//swap(end,tempNumberIndex)
	tempNumber := nums[end]

	//2.[tempNumberIndex,end]表示大于等于tempNumber
	//如果有小于tempNumber的j
	//那就j和tempNumberIndex交换，这样就需要tempNumberIndex++
	tempNumberIndex := start
	for i := start; i <= end; i++ {
		if nums[i] < tempNumber {
			nums[i], nums[tempNumberIndex] = nums[tempNumberIndex], nums[i]
			tempNumberIndex++
		}
	}

	//nums[end]符合大于等于的条件，所以不会动，因此交换
	nums[tempNumberIndex], nums[end] = nums[end], nums[tempNumberIndex]
	return tempNumberIndex
}

//解法二就是利用堆去建堆，建好之后poll出第k个值
// @lc code=end
