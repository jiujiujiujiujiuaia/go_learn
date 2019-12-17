package main

/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start

//因为题目要求了本地移动，因此只能采取了O(cnt * n)的时间复杂度
//思想就是把一个坑让出来，然后其他的数组数字就可以活动了。最后再把坑填上
func rotate1(nums []int, k int) {
	k = k % len(nums)
	for i := 0; i < k; i++ {
		temp := nums[len(nums)-1]
		for j := len(nums) - 2; j > -1; j-- {
			nums[j+1] = nums[j]
		}
		nums[0] = temp
	}
}

//解法二，如果对空间复杂度o(n)的话，那么就可以通过算法算出老位置对应的新位置的映射，就可以
//遍历一遍把数组算出来
func rotate(nums []int, k int) {
	k = k % len(nums)
	NewSlice := make([]int, len(nums))
	copy(NewSlice, nums)
	for i := 0; i < len(nums); i++ {
		index := (i + k) % len(nums)
		nums[index] = NewSlice[i]
	}
}

//解法三，经观察发现尾部的k%n个元素会被挪到头部，剩下的元素向后移动
//因此先全部反转，然后k%n个元素反转，后n-k%n进行反转（这个又是双指针法翻转数组的问题）

//解法四，依次遍历所有元素，找到元素的新位置，这样就会把老元素给挤掉，保存老元素。但是
//有可能会出现新老元素出现同一位置上，因此还要判断环状数组的问题

// @lc code=end
