package binarySearch

/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start

//2022/7/4
//这道题属于半有序，在basis基础上进行了修改。把一个有序数组通过循环的方式，变成了两段有序。
//比如{[5,6,7,8,9,10],[1,2,3,4]} 无论任意时刻，mid要么掉入上半有序区间，要么掉入下半有序区间。
//假设掉入上半有序区间（反之亦然），那么如果target在[l,mid)范围内，那就退化成二分搜索。
//如果不在范围内，那就放弃[l,mid) 就是[mid+1,r]范围，这个范围内，又会分为上半有序区间和下半有序区间。
func search(nums []int, target int) int {
	leftBoundary := 0
	rightBoundary := len(nums) - 1
	for leftBoundary <= rightBoundary {
		mid := leftBoundary + (rightBoundary-leftBoundary)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] >= nums[leftBoundary] {
			if target < nums[mid] && target >= nums[leftBoundary] {
				rightBoundary = mid - 1
			} else {
				leftBoundary = mid + 1
			}
		} else if nums[mid] < nums[rightBoundary] {
			if target > nums[mid] && target <= nums[rightBoundary] {
				leftBoundary = mid + 1
			} else {
				rightBoundary = mid - 1
			}
		}
	}
	return -1
}

//2020/3/10
//解法一：
//先找到整个有序数组旋转了多少，也就是向右移动了多少
//然后把数组当作是有序的，套用二分搜索的标准模板，只不过mid也要偏移，
//通过indexOld->indexNew的偏移找到原本mid的位置
//找到那个值
func search1(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	var b int
	//如果最后一个值大于第一个值，说明整体便宜了length，也就是没有偏移
	if nums[len(nums)-1] > nums[0] {
		b = 0
	} else {
		indexNew := findMax(nums)
		indexOld := len(nums) - 1
		//indexOld + b -len(nums) = indexNew
		//b代表偏移量
		b = indexNew + len(nums) - indexOld
		if b > len(nums) {
			b = b % len(nums)
		}
	}

	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		index := (mid + b) % len(nums)
		if nums[index] == target {
			return index
		} else if nums[index] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}

//由于旋转后，数组的最大值和最小值是紧挨着的，因此这个办法可以找到数组
//的最大和最小值（前提是数组中的元素不重复）
func findMax(nums []int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		//如果mid大于mid+1，说明mid就是分界
		if nums[mid] > nums[mid+1] {
			return mid
		}

		//如果mid-1大于mid，说明mid-1是分界
		if nums[mid-1] > nums[mid] {
			return mid - 1
		}

		if nums[mid] > nums[r] {
			l = mid
		} else {
			r = mid
		}
	}
	return -1
}

//解法二：不拐弯，直接二分
func search2(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}

		//旋转后数组变成了两段有序的，因此mid要么前半段有序，要么后半段有序
		if nums[l] <= nums[mid] {
			if target < nums[mid] && target >= nums[l] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

// @lc code=end
