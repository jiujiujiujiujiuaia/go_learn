package main

/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
func generateMatrix(n int) [][]int {

	direction := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	res := make([][]int, n)
	route := make([][]bool, n)
	for i := 0; i < n; i++ {
		route[i] = make([]bool, n)
		res[i] = make([]int, n)
	}

	i, j := 0, 0
	d := 0
	row, col := n, n
	var iNext, jNext int
	//终止条件就是调转方向后还是走到已经走过的格子中去了
	for cnt := 0; cnt < n*n; cnt++ {
		res[i][j] = cnt + 1
		route[i][j] = true
		if iNext, jNext = i+direction[d][0], j+direction[d][1]; checkBorder(iNext, jNext, row, col) || route[iNext][jNext] {
			d = (d + 1) % 4
			//如果碰到了边界或者已经走过的地方，需要转换方向了
			iNext, jNext = i+direction[d][0], j+direction[d][1]
		}
		i, j = iNext, jNext
	}
	return res
}

func checkBorder(i, j, row, col int) bool {
	if i == row || j == col || i == -1 || j == -1 {
		return true
	}
	return false
}

// @lc code=end
