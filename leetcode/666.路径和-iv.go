package main

var res int
var numberMap map[int]int

//多次需要遍历的时候，思考是不是需要查找算法（最简单的hash表）
//然后就是第一步思考，先转换成树的结构，然后递归去遍历每条路径，注意回退
//然后一般套路是思考如何不转换，直接去寻找左右子节点

func pathSum(nums []int) int {
	preHandlePathSum(nums)
	if len(nums) == 0 {
		return 0
	}
	helpPathSum(nums[0]/10, 0)
	return res
}

func helpPathSum(node int, sum int) {
	number, ok := numberMap[node]
	if !ok {
		return
	}

	sum += number
	depth := node / 10
	pos := node % 10
	left := (depth+1)*10 + pos*2 - 1
	right := left + 1

	_, leftExist := numberMap[left]
	_, rightExsit := numberMap[right]
	if !leftExist && !rightExsit {
		res += sum
	} else {
		helpPathSum(left, sum)
		helpPathSum(right, sum)
	}

}

func preHandlePathSum(nums []int) {
	res = 0
	numberMap = make(map[int]int)
	for _, num := range nums {
		value := num % 10
		num = num / 10
		numberMap[num] = value
	}
}
