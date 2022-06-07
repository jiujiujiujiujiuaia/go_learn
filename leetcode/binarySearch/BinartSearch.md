## 1.循环终止条件

下面是一个basic的二分搜索模板

```go
package main

func search(nums []int, target int) int {
	leftBoundary := 0
	rightBoundary := len(nums) - 1

	for leftBoundary <= rightBoundary {
		//防止溢出的写法
		mid := leftBoundary + (rightBoundary-leftBoundary)/2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			leftBoundary = mid + 1
		} else if target < nums[mid] {
			rightBoundary = mid - 1
		}
	}
	return -1
}
```

如果target存在nums数组中，那么循环终止条件就是
```
if target == nums[mid] {
	return mid
}
```

如果target不存在nums数组中，循环会如何终止呢？ 举两个极端例子

### case1：target比nums数组所有值都大

由于targer最大，所以，r = nums.length-1且不变，l不断向右靠拢。

直到搜索到最后一个值是否等于target，l=r=nums.length-1，由于target>nums[nums.length-1]，因此
newLeft = mid + 1 = r + 1。（由于r=nums.length,newLeft会超过数组索引）

### case2: target比nums数组所有值都小
由于targer最小，所以，l = 0且不变，r不断向左靠拢。

直到搜索到第一个值是否等于target，l=r=0，由于target<nums[0]，因此
newRight = mid - 1 = l - 1。（由于l=0,newRight等于-1，超过数组索引）

### case3: target比数组部分元素大，部分元素小
例如:nums = [1,2,3,4,5],target=3.5. 最后一次[l,r] = [3,2]
例如:nums = [1,2,3,4],target=3.5. 最后一次[l,r] = [3,2]

循环终止条件是l = r + 1

### 结论

所以循环终止条件最终都是l = r + 1，二分搜索完数组内潜在的值。

## 2.搜索边界
在basic的二分搜索进行升级，例如34th题，允许数组中有多个元素，返回range。
```
func searchLeftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1 
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			right = mid - 1
		}
	}

	if left == len(nums) || nums[left] != target {
		return -1
	}
	return left
}
```
为什么right = mid - 1？为什么left就是最后的边界呢？为什么最后要额外判断一下？

### 2.0 额外判断
如果target不在nums数组中，那么根据case1和case2，要么left=len(nums)，要么left=[0,len(nums))。

### 2.1 为什么right = mid - 1
和传统的二分搜索不同的是，当搜索到target的时候，没有停止搜索，而是缩小右区间。

### 2.2 为什么left是最后的边界呢？

#### (1)假设数组只有一个target
r=targetIndex-1, 搜索范围是[l,targetIndex-1],因此这个case就是上面的例子1。target比搜索范围内所有元素都大。

因此newLeft = mid + 1 = r + 1 = targetIndex - 1 + 1 = targetIndex

#### (2)假设数组有多个target 

只要不是最左边的target,循环都不会停下来，不断的缩小右边界缩小范围寻找target，直到命中最后也是最左边的的target。

然后由于最左边的target是最后一个target，因此r=targetIndex-1, 搜索范围是[l,targetIndex-1]，退化到**假设数组只有一个target**


