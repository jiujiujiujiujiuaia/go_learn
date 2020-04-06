package main

import "math"

/*
 * @lc app=leetcode.cn id=179 lang=golang
 *
 * [179] 最大数
 */

// @lc code=start
func largestNumber(nums []int) string {
	sort(nums,0,len(nums)-1)
	//(1)边界条件，如果数组最大的值是0，那么就返回‘0’
	if nums[len(nums)-1] == 0{
		return "0"
	}
	fmt.Println(nums)
	//(2)升序数组
	res :=""
	for i:=len(nums)-1;i >= 0;i--{
		res += strconv.Itoa(nums[i])
	}
	return res
}

//手写快排
func sort(nums []int,start,end int){
	if start > end {
		return
	}
	index := patipation(nums,start,end)
	sort(nums,start,index-1)
	sort(nums,index+1,end)

}

//这里的排序是升序
func patipation(nums []int, start, end int) int {
	tempNumber := nums[end]
	tempNumberIndex := start
	for i:=start;i<=end;i++{
		if compare(tempNumber,nums[i]){
			nums[i], nums[tempNumberIndex] = nums[tempNumberIndex], nums[i]
			tempNumberIndex++
		}
	}
	nums[tempNumberIndex], nums[end] = nums[end], nums[tempNumberIndex]
	return tempNumberIndex
}

//这里的排序规则不是单纯的比较大小，而是看哪个数在前更好
func compare(a, b int) bool {
	//(1)计算出a，b两者的位数,这里有一个特殊用例0
	tempA,tempB := a,b
	lengthA := 0
	lengthB := 0

	if a != 0{
		for tempA > 0 {
			lengthA++
			tempA /= 10
		}
	}else{
		lengthA=1
	}
	
	if b != 0{
		for tempB > 0 {
			lengthB++
			tempB /= 10
		}
	}else{
		lengthB=1
	}
	
	

	//(2)长度相同
	if lengthA == lengthB {
		if a > b {
			return true
		} else {
			return false
		}
	}

	//(3)长度不同
	combination1 := float64(a)*math.Pow10(lengthB) + float64(b)
	combination2 := float64(b)*math.Pow10(lengthA) + float64(a)
	if combination1 > combination2 {
		return true
	} else {
		return false
	}
}

// @lc code=end
