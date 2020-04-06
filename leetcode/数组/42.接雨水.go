package main

/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start

//暴力解法，每遍历到一个点位，不要管是谁把他围起来接雨水了
//而是去找以他为中心左边最大的和右边最大的柱子，然后这两个之中较小的柱子减去他就是了
//是从全局上来考虑的而不是从模拟能否把它围起来
func trap1(height []int) int {
	res := 0
	for i:=0;i<len(height);i++{
		maxLeft,maxRight := height[i],height[i]
		for j:=i;j>=0;j--{
			if height[j] >maxLeft{
				maxLeft = height[j]
			}
		}
		for j:=i;j<len(height);j++{
			if height[j] > maxRight{
				maxRight=height[j]
			}
		}
		
		if maxLeft > maxRight{
			res += maxRight - height[i]
		}else{
			res += maxLeft - height[i]
		}
	}

	return res
}

//解法二：从o(n^2)时间复杂度进一步优化，那就说明要么是动态规划类型，要么题目有一些条件没用（如有序），
//那么要么就是利用空间换时间，诸如查找之类的法子。

//那么这道题暴力是每一个点位都去暴力的查找左右最大的值
//那么我们可以把这个最大值储存起来，然后下一个最大值只需要和前面比就可以了，类似于记忆化搜索
func trap2(height []int) int {
	if len(height) == 0{
		return 0
	}

	ans := 0
	//每一位但是比height[i]大得值
	leftMaxValArray,rightMaxValArray := make([]int,len(height)),make([]int,len(height))
	leftMaxValArray[0] = height[0]
	for i:=1;i<len(height);i++{
		leftMaxValArray[i] = max(leftMaxValArray[i-1],height[i])
	}

	rightMaxValArray[len(height) - 1] = height[len(height) - 1]
	for i:=len(height) - 2;i>=0;i--{
		rightMaxValArray[i] = max(rightMaxValArray[i+1],height[i])
	}

	for i:=0;i<len(height);i++{
		if leftMaxValArray[i] > rightMaxValArray[i]{
			ans +=  rightMaxValArray[i] - height[i]
		}else{
			ans +=  leftMaxValArray[i] - height[i]
		}
	}
	return ans
}

func max(a,b int)int{
	if a > b {
		return a
	}else{
		return b
	}
}

// @lc code=end
