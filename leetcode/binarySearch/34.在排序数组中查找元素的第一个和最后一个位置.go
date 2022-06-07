package binarySearch

/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置

 */

//2020/3/10

// @lc code=start

//解法一：找到目标值后，然后向左向右进行搜索，找到起点和终点
//二分查找可以分为三种解法，也就是找到值，找到值得左边界，找到值得右边界，详情可见labuladong
func searchRange1(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return handleEqualCondition(nums, target, mid)
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return []int{-1, -1}
}

//最坏的情况全部都是target，然后算法退化成了o(n)
func handleEqualCondition(nums []int, target, index int) []int {
	l, h := index, index
	for h < len(nums)-1 && nums[h+1] == target {
		h++
	}

	for l > 0 && nums[l-1] == target {
		l--
	}
	return []int{l, h}
}

//解法二：这里运用一下二分搜索找到左右边界的方法
//假设数组中没有target，即最终的情况是mid=left=right
//如果target>nums[mid] -> left = mid + 1 = right + 1
//如果target<nums[mid] -> right = mid - 1= left - 1 (这个式子可以转换到上面)
//循环的终止条件就是left = right + 1

func searchRange(nums []int, target int) []int {
	left := searchLeftBound(nums, target)
	right := searchRightBound(nums, target)
	return []int{left, right}
}

//这里返回的是target的左边界
func searchLeftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1 //1.这里表示的搜索区间为[l,r]双闭区间
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			//3.和传统的二分搜索不同的是，当搜索到target的时候，没有停止搜索，而是缩小右区间。那么就会有以下集中可能:
			//3.1 如果数组只有一个target, 然后l和r就会在target的左边一直搜索，r=targetIndex-1不变，l不断的向r靠拢，直到l=r=targetIndex-1
			//那么最后一次肯定进nums[mid] < target这个条件，最后一次循环为mid=l=r，且target>nums[mid]，因此l=r+1,而r=targetIndex -1, l = targetIndex
			//3.2 如果数组有多个target,只要不是最左边的target,循环都不会停下来，不断的缩小右边界缩小范围寻找target，直到命中最后也是最左边的的target, 然后就退化到3.1的case
			//l=targetIndex
			right = mid - 1
		}
	}

	//4.因为left=right+1，那么如果target不在nums数组中
	//（1）target>nums 则left=right+1=数组长度会越界
	//（2）target<nums 则left=right+1=1
	//（3）target between nums 则left=right+1
	if left == len(nums) || nums[left] != target {
		return -1
	}
	return left
}

//这里返回的是target的右边界
func searchRightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1 //1.这里表示的搜索区间为[l,r]双闭区间
	for left <= right {
		mid := left + (right-left)/2
		//2.因此每一次搜索完后，区间都要是双闭[left,mid-1],[mid+1,right]
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			//3.当搜索到这个值后，不是着急把他返回，而是增加左区间
			//因此循环终止条件是left = right + 1 -> right=left-1=mid
			//如果数组中只有一个目标值的话，left还是能==mid
			//如果数组中有2个或以上目标值，那么right就会最后一个值
			left = mid + 1
		}
	}

	//4.因为left=right+1，那么如果target不在nums数组中
	//（1）target>nums 则right=left-1=len(nums)-1
	//（2）target<nums 则right=left-1=-1
	//（3）target between nums 则right=left-1
	if right == -1 || nums[right] != target {
		return -1
	}
	return right
}

// @lc code=end
