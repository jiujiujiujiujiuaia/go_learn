package 贪心

func Jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	jumpCnt := 0
	for i := 0; i < len(nums); {
		mostFar := i + nums[i]
		if mostFar >= len(nums)-1 {
			jumpCnt++
			return jumpCnt
		}
		nextStep := 0
		nextIndex := 0
		for j := i + 1; j < len(nums) && j <= mostFar; j++ {
			if nums[j] >= nextStep {
				nextIndex = j
				nextStep = nums[j]
			}
		}
		i = nextIndex
		jumpCnt++
		if i >= len(nums)-1 {
			break
		}
	}
	return jumpCnt
}
