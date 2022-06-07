package binarySearch

/*
 * @lc app=leetcode.cn id=74 lang=golang
 * [74] 搜索二维矩阵

 */

// @lc code=start

//2020/3/9
func searchMatrix1(matrix [][]int, target int) bool {

	if len(matrix) == 0 {
		return false
	}

	//1.行和列
	row := len(matrix)
	col := len(matrix[0])
	l, h := 0, row*col-1

	//2.按照一维数组二分的思想，0,1,2,3,4,5,6..n，只要把这个编号能够和二维数组
	//的坐标对应上，就可以迎刃而解了。
	for l <= h {
		mid := l + (h-l)/2
		var ok bool
		//3.这是重点，把二维当作一维来做
		i, j := mid/col, mid%col
		l, h, ok = help(matrix[i][j], target, mid, l, h)
		if ok {
			return true
		}

	}
	return false
}

func help(number, target, mid, l, h int) (int, int, bool) {
	if number > target {
		return l, mid - 1, false
	} else if number < target {
		return mid + 1, h, false
	} else {
		return 0, 0, true
	}
}

// @lc code=end
